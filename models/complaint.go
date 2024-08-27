package models

type Complaint struct {
    ID       uint   `gorm:"primaryKey"`
    UserID   uint   `gorm:"unique"` // Define foreign key with unique constraint
    Title    string
    Summary  string
    Severity string
    Resolved   bool
}
