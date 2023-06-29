// package main

// import (
//     "net/http"
// 		// "log"
// 		// "encoding/json"
// 		// "github.com/gorilla/mux"
//     "github.com/gin-gonic/gin"
// )

// // album represents data about a record album.
// type album struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

// type song struct {
// 	ID string `json:"id"`
// 	Name string `json:"name"`
// }

// // albums slice to seed record album data.
// var albums = []album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

// var songs = []song{
// 	{ID: "1", Name: "Shape of You"},
// 	{ID: "2", Name: "Lạ Lùng"},
// }

// // getAlbums responds with the list of all albums as JSON.
// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }

// func getSongs(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, songs)
// }


// func main() {
// 	router := gin.Default()
// 	router.GET("/albums", getAlbums)
// 	router.GET("/songs", getSongs)

// 	router.Run("localhost:8080")
// }
// // func main() {
// // 	r := gin.Default()

// // 	r.GET("/hello", func(c *gin.Context) {
// // 		c.JSON(http.StatusOK, gin.H{
// // 			"message": "Hello from Gin!",
// // 		})
// // 	})

// // 	r.Run()
// // }

package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type Product struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Price int    `json:"price"`
}

type Song struct {
	ID    int    `json:"id"`
  Name  string `json:"name"`
  Single string    `json:"single"`
}

var products []Product
var song []Song

func main() {
    // Set up some sample products
    products = []Product{
        {ID: 1, Name: "Product 1", Price: 10},
        {ID: 2, Name: "Product 2", Price: 20},
        {ID: 3, Name: "Product 3", Price: 30},
    }

		song = []Song{
			{ID: 1, Name: "Em Của Ngày Xưa Chết Rồi ", Single: "Hiền Hồ"},
			{ID: 2, Name: "Anh Thật Sự Ngu Ngốc ", Single: "Sơn Tùng"},
		}

    http.HandleFunc("/api/products", getProducts)
		http.HandleFunc("/api/songs", getSongs)

    log.Println("Server started on http://localhost:8000")

    err := http.ListenAndServe(":8000", nil)
    if err != nil {
        log.Fatal(err)
    }
}

func getProducts(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
}

func getSongs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(song)
}
