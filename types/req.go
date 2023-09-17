package types

type LoginUser struct {
	En       string `json:"en"`
	Password string `json:"password"`
}

type LdapUser struct {
	Cn       string `json:"cn" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LdapDeleteUser struct {
	Id string `json:"id" binding:"required"`
	Cn string `json:"cn" binding:"required"`
}

type LdapResetPass struct {
	Cn       string `json:"cn" binding:"required"`
	Password string `json:"password" binding:"required"`
}
