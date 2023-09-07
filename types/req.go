package types

type LoginUser struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LdapUser struct {
	OU   string `json:"ou" binding:"required"`
	CN   string `json:"cn" binding:"required"`
	Pass string `json:"pass" binding:"required"`
}

type LdapDeleteUser struct {
	CN string `json:"cn" binding:"required"`
}

type LdapResetPass struct {
	CN       string `json:"cn" binding:"required"`
	Password string `json:"password" binding:"required"`
}
