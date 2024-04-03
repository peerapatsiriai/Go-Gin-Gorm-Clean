package domain

type Book struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Title     string `json:"title" required:"true"`
	Author    string `json:"author" required:"true"`
	Pages     int    `json:"pages" required:"true"`
	Published bool   `json:"published" required:"true"`
}
