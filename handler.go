package handler

import (
        "fmt"
        "log"
        "net/http"
        "github.com/gin-gonic/gin"
//        "strings"
)

type album struct {
        ID     string `json:"id"`
        Title  string `json:"title"`
        Artist string `json:"artist"`
        Price  string `json:"price"`
}

var albums = []album{
        {ID: "1", Title: "abcc", Artist: "Rajat", Price: "64"},
}

func getAlbums(c *gin.Context) {
        c.IndentedJSON(http.StatusOK, albums)
}

func homePage(w http.ResponseWriter, r *http.Request) {

        fmt.Fprintf(w, "endpoint hit")

}

func Handler() {
        http.HandleFunc("/", homePage)
        log.Fatal(http.ListenAndServe(":8081", nil))
}

