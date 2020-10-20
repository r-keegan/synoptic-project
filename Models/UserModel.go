package Models

// TODO generate 16 char length Card ID alphanumeric
type User struct {
	ID         uint
	EmployeeID int    `gorm:"size:10;not null;unique" json:"employeeID"`
	CardID     string `gorm:"length:16;type:uuid;not null;unique" json:"cardID"`
	Name       string
	Email      string
	Phone      string
	Pin        string
	ConfirmPin string
	Balance    int
}

func (b *User) TableName() string {
	return "user"
}
