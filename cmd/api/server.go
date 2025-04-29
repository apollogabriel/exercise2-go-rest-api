package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":3000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Root Route"))
		fmt.Println("Hello Root Route")
	})

	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
		if r.Method == http.MethodGet {
			w.Write([]byte("Hello GET teachers Route"))
			fmt.Println("Hello GET teachers Route")
			return
		}
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("Hello GET teachers Route"))
			fmt.Println("Hello GET teachers Route")
		case http.MethodPost:
			w.Write([]byte("Hello POST teachers Route"))
			fmt.Println("Hello POST teachers Route")
		case http.MethodPut:
			w.Write([]byte("Hello PUT teachers Route"))
			fmt.Println("Hello PUT teachers Route")
		case http.MethodPatch:
			w.Write([]byte("Hello PATCH teachers Route"))
			fmt.Println("Hello PATCH teachers Route")
		case http.MethodDelete:
			w.Write([]byte("Hello DELETE teachers Route"))
			fmt.Println("Hello DELETE teachers Route")
		}
	})

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello students Route"))
		fmt.Println("Hello students Route")
	})

	http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello execs Route"))
		fmt.Println("Hello execs Route")
	})

	fmt.Println("Server is running on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}
