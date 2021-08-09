package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.Error(writer, "404 page not found", http.StatusNotFound)
		return
	}

	switch request.Method {
	case http.MethodGet:
		http.ServeFile(writer, request, "index.html")
	case http.MethodPost:
		err := request.ParseForm()
		if err != nil {
			fmt.Fprintf(writer, "ParseForm() error := %v", err)
			return
		}
		name := request.FormValue("name")
		address := request.FormValue("address")
		token := http.Cookie{
			Name: "token",
			Value: fmt.Sprintf("%s%s", name, address),
		}
		http.SetCookie(writer, &token)
		http.Redirect(writer, request, "/", http.StatusFound)
	default:
		fmt.Fprintf(writer, "Only GET and POST requests are available")
	}
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
