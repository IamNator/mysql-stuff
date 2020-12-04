package user

import "database/sql"

type RegisterUser struct {
	UserName string  `json:"username"`
	//Email string	 `json:"email"`
	Password string   `json:"password"`
}

type LoginUser struct {
	UserName string  `json:"username"`
	//Email string	 `json:"email"`
	Password string   `json:"password"`
}

type DBData struct {
	DBType, User, Password, Host, DBName string
	Session                              *sql.DB
	SessionIDs							 map[string]string
	SessionUsers						 map[string]string
}