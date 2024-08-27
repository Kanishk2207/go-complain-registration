package models

type User struct {
    ID        uint   `gorm:"primaryKey"`
    FirstName string
    LastName  string
    Email     string
    Password  string
    Status    string
    Complaints []Complaint
}