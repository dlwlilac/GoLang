package models

// type Person struct {
// 	Name string `json:"name"`
// 	Pass string `json:"pass"`
// }

// type User struct {
// 	Name     string `json:"name" validate:"required,min=3,max=32"`
// 	IsActive *bool  `json:"isactive" validate:"required"`
// 	Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
// }

type User struct {
	Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
	Username string `json:"username" validate:"username"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	LineID   string `json:"lineid"`
	Tel      string `json:"tel" validate:"required,min=10,max=10"`
	Type     string `json:"type"`
	Url      string `json:"url" validate:"required,min=3,max=30,url"`
}
