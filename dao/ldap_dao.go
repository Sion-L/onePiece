package dao

import (
	"fmt"
	"github.com/Sion-L/onePiece/db"
	ldapv3 "github.com/go-ldap/ldap/v3"
	"github.com/mozillazg/go-pinyin"
	"log"
	"strings"
)

func removeDuplication(arr []string) []string {
	set := make(map[string]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}

	return arr[:j]
}

func Product(sets ...[]string) [][]string {
	lens := func(i int) int { return len(sets[i]) }
	product := [][]string{}
	for ix := make([]int, len(sets)); ix[0] < lens(0); nextIndex(ix, lens) {
		var r []string
		for j, k := range ix {
			r = append(r, sets[j][k])
		}
		product = append(product, r)
	}
	return product
}

func nextIndex(ix []int, lens func(i int) int) {
	for j := len(ix) - 1; j >= 0; j-- {
		ix[j]++
		if j == 0 || ix[j] < lens(j) {
			return
		}
		ix[j] = 0
	}
}

// 输入中文转换成英文,再判断此用户是否已经存在,存在就在pinyin后加一个数字返回
func FilterUser(cn string) ([]string, error) {
	var (
		enList   [][]string
		toList   []string
		nameList []string
	)
	en := pinyin.Pinyin(cn, pinyin.Args{
		Style:     pinyin.Normal,
		Heteronym: true,
	})

	for _, v := range en {
		duplication := removeDuplication(v)
		enList = append(enList, duplication)
	}

	sets := Product(enList...)
	for _, set := range sets {
		toList = append(toList, strings.Join(set, ""))
	}

	for _, v := range toList {
		name, ok, err := LdapSearch(v) // ok则为不存在 ，!ok 则为存在 需要处理
		if err != nil {
			return nil, err
		}
		if ok {
			nameList = append(nameList, name) // 不存在的部分
		} else {
			//name01 := fmt.Sprintf("%s01", name) // 存在的部分做处理返回
			//name02 := fmt.Sprintf("%s02", name)
			nameList = append(nameList, name)
			return nil, fmt.Errorf("%s 用户已存在,建议取个别名", nameList)
		}
	}
	return nameList, nil
}

// AddUser add user to ldap
func AddUser(ou, sn, cn, pass string) error {

	dn := fmt.Sprintf("uid=%s,ou=%s,dc=lang,dc=com", sn, ou) // uid跟sn一样
	addResponse := ldapv3.NewAddRequest(dn, []ldapv3.Control{})
	addResponse.Attribute("objectClass", []string{"top", "organizationalPerson", "inetOrgPerson", "person"}) // 必填字段 否则报错 LDAP Result Code 65 "Object Class Violation"
	addResponse.Attribute("employeeType", []string{"1"})                                                     // 工号 暂时没用到
	addResponse.Attribute("sn", []string{sn})
	addResponse.Attribute("cn", []string{cn})
	addResponse.Attribute("mail", []string{fmt.Sprintf("%s@lang.com", sn)})
	addResponse.Attribute("uid", []string{sn})
	addResponse.Attribute("userPassword", []string{pass})

	if err := db.LdapConn.Add(addResponse); err != nil {
		if ldapv3.IsErrorWithCode(err, 68) {
			return fmt.Errorf("user %s already exist", sn)
		}
		return err
	}
	return nil

}

// 检查用户是否已经存在,存在就返回true
func LdapSearch(sn string) (string, bool, error) {

	filter := fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", ldapv3.EscapeFilter(sn))
	searchRequest := ldapv3.NewSearchRequest(
		"dc=lang,dc=com",
		ldapv3.ScopeWholeSubtree, ldapv3.NeverDerefAliases, 0, 0, false,
		//fmt.Sprintf("(&(objectClass=organizationalUnit))"),
		filter,
		[]string{"cn"},
		nil,
	)

	sr, err := db.LdapConn.Search(searchRequest)

	if err != nil {
		fmt.Println(err)
		return "", false, err
	}

	if len(sr.Entries) == 0 {
		return sn, true, nil
	}

	return sn, false, nil
}

// 删除用户同时删除数据
func LdapDeleteUser(cn string) error {

	sn, err := FindUserByLdap(cn)
	if err != nil {
		return err
	}
	dn := fmt.Sprintf("uid=%s,ou=employee,dc=lang,dc=com", sn)
	delReq := ldapv3.NewDelRequest(dn, []ldapv3.Control{})
	if err := db.LdapConn.Del(delReq); err != nil {
		return fmt.Errorf("[DeleteUser] Ldap failed to delete user %s: %s", sn, err.Error())
	}

	err = DeleteUserByName(cn)
	if err != nil {
		return fmt.Errorf("[DeleteUser] database delete user %s failed: %s", sn, err.Error())
	}
	return nil
}

// 汉字转拼音
func FindUserByLdap(cn string) (string, error) {
	var (
		enList   [][]string
		toList   []string
		nameList []string
	)
	en := pinyin.Pinyin(cn, pinyin.Args{
		Style:     pinyin.Normal,
		Heteronym: true,
	})

	for _, v := range en {
		duplication := removeDuplication(v)
		enList = append(enList, duplication)
	}

	sets := Product(enList...)
	for _, set := range sets {
		toList = append(toList, strings.Join(set, ""))
	}

	for _, v := range toList {
		name, ok, err := LdapSearch(v) // ok则为不存在 ，!ok 则为存在
		if err != nil {
			return "", err
		}
		if !ok { // 存在就返回
			nameList = append(nameList, name) // 不存在的部分
		}
	}

	if nameList == nil {
		return "", fmt.Errorf("要删除的用户不存在: %s")
	}
	return nameList[0], nil
}

// 修改密码
func LdapResetPassword(sn, password string) error {

	dn := fmt.Sprintf("uid=%s,ou=employee,dc=lang,dc=com", sn)
	modReq := ldapv3.NewModifyRequest(dn, []ldapv3.Control{})
	modReq.Replace("userPassword", []string{password})
	if err := db.LdapConn.Modify(modReq); err != nil {
		return fmt.Errorf("[ResetPassword] %s reset password Failed: %s", sn, err.Error())
	}

	return nil
}

// 用户登陆验证
func LoginForLdap(sn, password string) bool {

	// employeeType再加一层过滤。只有类型为1的才能登陆
	filter := fmt.Sprintf("(&(uid=%s)(employeeType=1)(memberof=%s))", sn, "cn=Ops,cn=group,dc=lang,dc=com")

	//filter := fmt.Sprintf("(uid=%s)", sn)
	sql := ldapv3.NewSearchRequest(
		"dc=lang,dc=com",
		ldapv3.ScopeWholeSubtree,
		ldapv3.NeverDerefAliases,
		0,
		0,
		false,
		filter,
		nil, // 为nil表示返回所有属性
		nil)
	fmt.Println(sql)
	cur, err := db.LdapConn.Search(sql)
	if err != nil {
		log.Fatal(err)
	}
	if len(cur.Entries) == 0 {
		fmt.Println(cur.Entries)
		return false
	}

	entry := cur.Entries[0]
	fmt.Println(entry.GetAttributeValue("userPassword"))
	err = db.LdapConn.Bind(entry.DN, password)
	if err != nil {
		fmt.Println(entry.DN)
		fmt.Println(err)
		return false
	}
	return true
}
