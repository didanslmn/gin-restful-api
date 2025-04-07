package model

type Book struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" gorm:"type:varchar(255)"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
}
