package models

type UserLoginData struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

//type EmailCheck struct {
//	ID   uint   `json:"id"`
//	UUID string `json:"uuid"`
//}
