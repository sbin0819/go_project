package main

import (
	"fmt"
	"log"
	"net/http"
	"src/blockchain"
	"text/template"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	tmpl.Execute(rw, data)
}

func user(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello Users!")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/users", user)

	fmt.Printf("Listening on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
