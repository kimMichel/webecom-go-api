package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"img_url"`
	Rating      float64 `json:"rating"`
	Quantity    int     `json:"quantity"`
	Gender      string  `json:"gender"`
}
