package entity

import (
	
	"gorm.io/gorm"

	"time"

)

type Profile struct {
	gorm.Model
	ProfileID uint              
	Firstname string    
	Lastname  string    
	Birthdate time.Time 
	WriterID  string    

	// Relationships
	UserID    *uint 
	User User `gorm:"foreignKey:UserID"` 
}
