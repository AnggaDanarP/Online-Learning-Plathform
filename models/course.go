package models

type Course struct {
	Id        int    `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Describe  string `json:"describe"`
	Category  string `json:"category"`
	Price     string `json:"price"`
	Favorites string `json:"favorites"`
}

type RequestUpdateCourse struct {
	Name     string `json:"name"`
	Describe string `json:"describe"`
	Category string `json:"category"`
	Price    string `json:"price"`
}

type GetCourse struct {
	Name	 string `json:"name"`
}