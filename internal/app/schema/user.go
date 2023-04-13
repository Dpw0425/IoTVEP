package schema

type User struct {
	UID        uint   `json:"uid"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	VerifyCode string `json:"verify_code"`
}
