// migrations/migrations.go
package migrations

import (
	"gorm.io/gorm"
)

type User struct {
    ID        uint   `gorm:"primaryKey"`
    FirstName string
    LastName  string
    Email     string
    Password  string
    Status    string
}

type Complaint struct {
    ID       uint   `gorm:"primaryKey"`
    UserID   uint   
    Title    string
    Summary  string
    Severity string
    Resolved bool
}
func RunMigrations(db *gorm.DB) {
    // AutoMigrate will create the table if it doesn't exist and add any missing fields
    db.AutoMigrate(&User{})
    db.AutoMigrate(&Complaint{})

    // Check if the foreign key constraint exists
    var fkExists bool
    if err := db.Raw("SELECT COUNT(*) FROM information_schema.TABLE_CONSTRAINTS WHERE CONSTRAINT_NAME = 'fk_user_id' AND TABLE_NAME = 'complaints'").Scan(&fkExists).Error; err != nil {
        panic("failed to check foreign key constraint: " + err.Error())
    }

    // Drop existing foreign key constraint (if it exists)
    if fkExists {
        if err := db.Exec("ALTER TABLE complaints DROP FOREIGN KEY fk_user_id").Error; err != nil {
            panic("failed to drop foreign key constraint: " + err.Error())
        }
    }

    // Add foreign key constraint
    if err := db.Exec("ALTER TABLE complaints ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE").Error; err != nil {
        panic("failed to add foreign key constraint: " + err.Error())
    }
}
