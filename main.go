package main

import "go-bookshelf/database"

func main() {
	database.ConnectDB()
}

// type Author struct {
//   Id   uint64 `json:"id"`
//   Name string `json:"name"`
// }

// type Authors struct {
//   Authors []Author `json:"authors"`
// }

// type Book struct {
//   Id       uint64 `json:"id"`
//   Title    string `json:"title"`
//   AuthorId Author `json:"author"`
// }

// type Books struct {
//   Books []Book `json:"books"`
// }
