package dao

import (
	"fmt"
	"github.com/Sion-L/onePiece/db"
	ldapv3 "github.com/go-ldap/ldap/v3"
)

func AddUser(ou, cn, sn, password string) error {
	conn, err := db.NewClientLdap()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	dn := fmt.Sprintf("uid=%s,ou=%s,dc=lang,dc=com", sn, ou)
	addResponse := ldapv3.NewAddRequest(dn, []ldapv3.Control{})
	addResponse.Attribute("objectClass", []string{"top", "organizationalPerson", "inetOrgPerson", "person"}) // 必填字段 否则报错 LDAP Result Code 65 "Object Class Violation"
	addResponse.Attribute("employeeType", []string{"1"})                                                     // 工号 暂时没用到
	addResponse.Attribute("sn", []string{sn})
	addResponse.Attribute("cn", []string{cn})
	addResponse.Attribute("mail", []string{fmt.Sprintf("%s@lang.com", sn)})
	addResponse.Attribute("uid", []string{sn})
	addResponse.Attribute("userPassword", []string{password})
	if err := conn.Add(addResponse); err != nil {
		if ldapv3.IsErrorWithCode(err, 68) {
			return fmt.Errorf("user %s already exist", sn)
		}
		return err
	}
	return nil

}
