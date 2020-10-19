package Models

type User struct {
	ID         uint32 `gorm:"primary_key;auto_increment" json:"id"`
	EmployeeID int    `gorm:"size:10;not null;unique" json:"employeeID"`
	Name       string `gorm:"size:255;not null" json:"name"`
	Email      string `gorm:"size:100;not null" json:"email"`
	Phone      string `gorm:"size:100;not null" json:"phone"`
	Pin        int    `gorm:"size:4;not null" json:"pin"`
}

func (b *User) TableName() string {
	return "user"
}
