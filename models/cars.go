package models

type Cars struct {
	ID           uint   `json:"id" gorm:"primarykey"`
	Categories   string `json:"categories"`
	Brand        string `json:"brand"`
	Model        string `json:"model"`
	Availability string `json:"availability"`
}
