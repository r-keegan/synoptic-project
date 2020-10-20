package Models

type User struct {
	ID         uint
	EmployeeID int    `gorm:"not null;unique" json:"employeeID"`
	CardID     string `gorm:"length:16;not null;unique" json:"cardID"`
	Name       string
	Email      string
	Phone      string
	Pin        string
	Balance    int
}

func (b *User) TableName() string {
	return "user"
}
