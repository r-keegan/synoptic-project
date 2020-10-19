package Models

type Card struct {
	CardID  string `gorm:"hex(randomblob(2));not null;unique" json:"cardID"`
	Balance int    `gorm:"not null" json:"balance"`
}

func (b *Card) TableName() string {
	return "card"
}
