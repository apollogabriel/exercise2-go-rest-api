package main

import (
	"crypto/tls"
	"fmt"
	"go-rest-api/internal/api/handlers"
	mw "go-rest-api/internal/api/middleware"
	"go-rest-api/internal/repository/sqlconnect"
	"log"
	"net/http"
)

func main() {
	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	_, err := sqlconnect.ConnectDb("school")
	if err != nil {
		fmt.Println(err)
		return
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.RootHandler)

	mux.HandleFunc("/teachers/", handlers.TeachersHandler)

	mux.HandleFunc("/students/", handlers.StudentsHandler)

	mux.HandleFunc("/execs/", handlers.ExecsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	//rl := mw.NewRateLimiter(5, time.Minute)

	//hppOptions := mw.HPPOptions{
	//	CheckQuery:                  true,
	//	CheckBody:                   true,
	//	CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
	//	Whitelist:                   []string{"sortBy", "sortOrder", "name", "age", "class"},
	//}

	server := &http.Server{
		Addr: port,
		//Handler:   mw.Hpp(hppOptions)(rl.Middleware(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Cors(mux))))),
		//secureMux := utils.ApplyMiddlewares(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTimeMiddleware, rl.Middleware, mw.Cors)
		Handler:   mw.SecurityHeaders(mux),
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port:", port)
	err = server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}
