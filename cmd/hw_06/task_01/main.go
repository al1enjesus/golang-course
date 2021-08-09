package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Host       string              `json:"host: "`
	UserAgent  string              `json:"user_agent: "`
	RequestUri string              `json:"request_uri: "`
	Headers    http.Header 		   `json:"headers: "`
}

func (response Response) String() string {
	j, _ := json.MarshalIndent(response, "", "  ")
	return string(j)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		response := Response{
			Host: r.Host,
			UserAgent: r.UserAgent(),
			RequestUri: r.RequestURI,
			Headers: r.Header,
		}
		fmt.Fprintf(w, response.String())
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
