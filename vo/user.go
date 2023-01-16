package vo

type AddUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
