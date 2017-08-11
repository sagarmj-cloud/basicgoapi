package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sagarmj76/basicgoapi/controllers"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home...Sweet home!")
	fmt.Println("Enpoint Hit : Home")
}

func handleRequests() {
	http.HandleFunc("/", home)
	http.HandleFunc("/all", controllers.ReturnAllArticles)
	log.Fatal(http.ListenAndServe(":9081", nil))
}

func main() {
	handleRequests()
}
