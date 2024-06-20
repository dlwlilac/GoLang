package models

import (
	"gorm.io/gorm"
)

type User struct {
	Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
	Username string `json:"username" validate:"username"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	LineID   string `json:"lineid"`
	Tel      string `json:"tel" validate:"required,min=10,max=10"`
	Type     string `json:"type"`
	Url      string `json:"url" validate:"required,min=3,max=30,url"`
}

type Dogs struct {
	gorm.Model
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
}

type DogsRes struct {
	Name  string `json:"name" `
	DogID int    `json:"dog_id" `
	Type  string `json:"type" `
}

type ResultData struct {
	Data        []DogsRes   `json:"data"`
	Name        string      `json:"name"`
	Count       int         `json:"count"`
	ColorCounts ColorCounts `json:"color_counts"`
}

type ColorCounts struct {
	Pink    int `json:"pink"`
	Green   int `json:"green"`
	Red     int `json:"red"`
	NoColor int `json:"NoColor"`
}

type Company struct { //7.0.1
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone" validate:"required,min=10,max=10"`
	Email   string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
}
