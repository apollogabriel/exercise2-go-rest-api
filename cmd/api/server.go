package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	mw "go-rest-api/internal/api/middleware"
	"log"
	"net/http"
	"time"
)

type Teacher struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Class     string `json:"class,omitempty"`
	Subject   string `json:"subject,omitempty"`
}

var (
	teachers = make(map[int]Teacher)
	nextID   = 1
)

func init() {
	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "Juan",
		LastName:  "Perez",
		Class:     "Apo",
		Subject:   "Math",
	}
	nextID++
	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "Pedro",
		LastName:  "Castro",
		Class:     "Rizal",
		Subject:   "Science",
	}
	nextID++
}

func getTeachersHandler(w http.ResponseWriter, r *http.Request) {
	teacherList := make([]Teacher, len(teachers))

	for _, value := range teachers {
		teacherList = append(teacherList, value)
	}

	response := struct {
		Status string    `json:"status"`
		Count  int       `json:"count"`
		Data   []Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(teachers),
		Data:   teacherList,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Root Route"))
	fmt.Println("Hello Root Route")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTeachersHandler(w, r)
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

	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	mux.HandleFunc("/teachers/", teachersHandler)

	mux.HandleFunc("/students/", studentsHandler)

	mux.HandleFunc("/execs/", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	rl := mw.NewRateLimiter(5, time.Minute)

	hppOptions := mw.HPPOptions{
		CheckQuery:                  true,
		CheckBody:                   true,
		CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
		Whitelist:                   []string{"sortBy", "sortOrder", "name", "age", "class"},
	}

	server := &http.Server{
		Addr:      port,
		Handler:   mw.Hpp(hppOptions)(rl.Middleware(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Cors(mux))))),
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port:", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}
