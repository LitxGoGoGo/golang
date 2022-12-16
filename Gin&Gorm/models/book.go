package models

type Book struct {
	Id     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type UploadFile struct {
	Id   uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name string `json:"name" gorm:"type:varchar(100)"`
	Path string `json:"path" gorm:"type:varchar(100)"`
}
