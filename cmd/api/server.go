package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Root Route"))
	fmt.Println("Hello Root Route")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Println(r.URL.Path)
		path := strings.TrimPrefix(r.URL.Path, "/teachers/")
		userID := strings.TrimSuffix(path, "/")
		fmt.Println(path)
		fmt.Println("The ID is:", userID)

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
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET students Route"))
		fmt.Println("Hello GET students Route")
	case http.MethodPost:
		w.Write([]byte("Hello POST students Route"))
		fmt.Println("Hello POST students Route")
	case http.MethodPut:
		w.Write([]byte("Hello PUT students Route"))
		fmt.Println("Hello PUT students Route")
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH students Route"))
		fmt.Println("Hello PATCH students Route")
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE students Route"))
		fmt.Println("Hello DELETE students Route")
	}
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET execs Route"))
		fmt.Println("Hello GET execs Route")
	case http.MethodPost:
		w.Write([]byte("Hello POST execs Route"))
		fmt.Println("Hello POST execs Route")
	case http.MethodPut:
		w.Write([]byte("Hello PUT execs Route"))
		fmt.Println("Hello PUT execs Route")
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH execs Route"))
		fmt.Println("Hello PATCH execs Route")
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE execs Route"))
		fmt.Println("Hello DELETE execs Route")
	}
}

func main() {
	port := ":3000"

	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/teachers", teachersHandler)

	http.HandleFunc("/students", studentsHandler)

	http.HandleFunc("/execs", execsHandler)

	fmt.Println("Server is running on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}
