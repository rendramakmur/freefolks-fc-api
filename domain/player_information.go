package domain

import "time"

type PlayerInformation struct {
	Id             int       `gorm:"column:pi_id;primaryKey"`
	UserType       int       `gorm:"column:pi_user_type"`
	CustomerNumber string    `gorm:"column:pi_customer_number"`
	FirstName      string    `gorm:"column:pi_first_name"`
	LastName       string    `gorm:"column:pi_last_name"`
	Email          string    `gorm:"column:pi_email"`
	Password       string    `gorm:"column:pi_password"`
	MobileNumber   string    `gorm:"column:pi_mobile_number"`
	Occupation     int       `gorm:"column:pi_occupation"`
	DateOfBirth    time.Time `gorm:"column:pi_date_of_birth"`
	Gender         string    `gorm:"column:pi_gender"`
	PhotoProfile   string    `gorm:"column:pi_photo_profile"`
	Address        string    `gorm:"column:pi_address"`
	City           string    `gorm:"column:pi_city"`
	PostalCode     string    `gorm:"column:pi_postal_code"`
	BodySize       string    `gorm:"column:pi_body_size"`
	ActivationCode string    `gorm:"column:pi_activation_code"`
	EmailStatus    string    `gorm:"column:pi_email_status"`
	VerifiedAt     time.Time `gorm:"column:pi_verified_at"`
	CreatedAt      time.Time `gorm:"column:pi_created_at"`
	CreatedBy      int       `gorm:"column:pi_created_by"`
	UpdatedAt      time.Time `gorm:"column:pi_updated_at"`
	UpdatedBy      int       `gorm:"column:pi_updated_by"`
}
