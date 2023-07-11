package entities

type Counter struct {
	WinCode string `gorm:"column:wincode"`
	Count   int64  `gorm:"column:count"`
}
