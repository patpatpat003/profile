package	entity

import(
	"time"

	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	User_ID string
	Username string
	Password string
	Email string
	ID_Type string
	Firstname string
	Lastname string
	BirthDate time.Time


}