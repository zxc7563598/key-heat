package model

type KeyEvent struct {
	ID      int    `gorm:"primaryKey;autoIncrement"`
	Date    string `gorm:"type:text;not null;uniqueIndex:idx_date_key"`
	KeyName string `gorm:"type:text;not null;uniqueIndex:idx_date_key"`
	Count   int    `gorm:"not null;default:0"`
}

func (KeyEvent) TableName() string {
	return "key_events"
}
