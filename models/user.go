package models

type User struct{
	ID 			uint 	`gorm:"primarykey"`
	Name 		string	`json:"name"`
	Email 		string	`json:"email" gorm:"unique"`
	Password 	string 	`json:"password"`
}