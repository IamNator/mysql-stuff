package user

type RegisterUser struct {
	userName string  `json:"username"`
	//Email string	 `json:"email"`
	Password string   `json:"password"`
}

type LoginUser struct {
	userName string  `json:"username"`
	//Email string	 `json:"email"`
	Password string   `json:"password"`
}