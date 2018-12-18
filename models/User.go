package models

/*
 *  @Author Muhammad Edwin (muhammad.edwin@ruma.co.id)
 *
 *
 */
type User struct {
	Id       string
	Username string
	Password string
}

func Login(username, password string) bool {
	if username == "edwin" && password == "password" {
		return true
	}
	return false
}
