package entity


import (
	
	"gorm.io/gorm"

	"time"

)

type EditProfile struct {
	gorm.Model
	EditProfileID uint                           
	TempName      string 
	TempEmail     string 
	Birthdate time.Time 

	// Relationships
	ProfileID     *uint  
	Profile Profile `gorm:"foreignKey:ProfileID"` // Belongs to Profile
}