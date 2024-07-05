package models

// User schema of the user table
type User struct {
	ID             string `json:"id"`
	First_Name     string `json:"first_name"`
	Last_Name      string `json:"last_name"`
	Preffered_Name string `json:"nick_name"`
	Email_Address  string `json:"email_address"`
}