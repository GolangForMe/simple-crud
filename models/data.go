package models

type Author struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

type Book struct {
	Id            int     `json:"id"`
	Title         string  `json:"title"`
	Price         float64 `json:"price"`
	YearPublished int     `json:"year_published"`
	Author        Author
}

var Db []Book

func init() {
	Db = append(Db, Book{
		Id:            1,
		Title:         "War and Peace",
		Price:         123.32,
		YearPublished: 1868,
		Author: Author{
			Name:    "Lev",
			Surname: "Tolstoy",
			Age:     41,
		},
	}, Book{
		Id:            2,
		Title:         "The war of the worlds",
		Price:         133.32,
		YearPublished: 1897,
		Author: Author{
			Name:    "Gerbert",
			Surname: "Wells",
			Age:     31,
		},
	}, Book{
		Id:            3,
		Title:         "Robinson Crusoe",
		Price:         143.32,
		YearPublished: 1719,
		Author: Author{
			Name:    "Daniel",
			Surname: "Defaus",
			Age:     59,
		},
	})
}

func FindBookById(id int) (Book, bool) {
	for _, b := range Db {
		if b.Id == id {
			return b, true
		}
	}
	return Book{}, false
}
