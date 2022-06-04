package models

type GroundAssociation struct {
	Name    string `json:"name" gorm:"not null; size:256"`
	Address string `json:"address" gorm:"not null; size:256"`
	Tell    string `json:"tell" gorm:"not null; size:256"`
	Email   string `json:"email" gorm:"not null; size:256"`
	Break   string `json:"break" gorm:"not null; size:256"`
	Styles  []int  `json:"styles" gorm:"many2many:ground_styles"`
	Price   string `json:"price" gorm:"not null; size:256"`
	Url     string `json:"url" gorm:"not null; size:256"`
	Feature string `json:"feature" gorm:"not null; size:256"`
	Rule    string `json:"rule" gorm:"not null; size:256"`
	Other   string `json:"other" gorm:"not null; size:256"`
}
