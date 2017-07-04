package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port := "3000"
	if len(os.Getenv("PORT")) != 0 {
		port = os.Getenv("PORT")
	}

	http.HandleFunc("/cli", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello PaaS meetup !")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, _ := os.Open("./templates/index.html")
		io.Copy(w, f)
		f.Close()
	})

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))

	log.Println("Listening on: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
