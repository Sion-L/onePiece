package dao

import (
	"fmt"
	"github.com/Sion-L/onePiece/db"
	ldapv3 "github.com/go-ldap/ldap/v3"
	"github.com/mozillazg/go-pinyin"
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
	conn, err := db.NewClientLdap()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	dn := fmt.Sprintf("uid=%s,ou=%s,dc=lang,dc=com", sn, ou) // uid跟sn一样
	addResponse := ldapv3.NewAddRequest(dn, []ldapv3.Control{})
	addResponse.Attribute("objectClass", []string{"top", "organizationalPerson", "inetOrgPerson", "person"}) // 必填字段 否则报错 LDAP Result Code 65 "Object Class Violation"
	addResponse.Attribute("employeeType", []string{"1"})                                                     // 工号 暂时没用到
	addResponse.Attribute("sn", []string{sn})
	addResponse.Attribute("cn", []string{cn})
	addResponse.Attribute("mail", []string{fmt.Sprintf("%s@lang.com", sn)})
	addResponse.Attribute("uid", []string{sn})
	addResponse.Attribute("userPassword", []string{pass})
	if err := conn.Add(addResponse); err != nil {
		if ldapv3.IsErrorWithCode(err, 68) {
			return fmt.Errorf("user %s already exist", sn)
		}
		return err
	}
	return nil

}

// 检查用户是否已经存在,存在就返回true
func LdapSearch(en string) (string, bool, error) {
	conn, err := db.NewClientLdap()
	defer conn.Close()

	if err != nil {
		fmt.Println(err)
		return "", false, err
	}

	filter := fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", ldapv3.EscapeFilter(en))
	searchRequest := ldapv3.NewSearchRequest(
		"dc=lang,dc=com",
		ldapv3.ScopeWholeSubtree, ldapv3.NeverDerefAliases, 0, 0, false,
		//fmt.Sprintf("(&(objectClass=organizationalUnit))"),
		filter,
		[]string{"cn"},
		nil,
	)

	sr, err := conn.Search(searchRequest)
	if err != nil {
		fmt.Println(err)
		return "", false, err
	}

	if len(sr.Entries) == 0 {
		return en, true, nil
	}

	return en, false, nil
}
