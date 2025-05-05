package main

import (
	
	"crypto/tls"
	"fmt"
	"github.com/joho/godotenv"
	"go-rest-api/internal/api/handlers"
	mw "go-rest-api/internal/api/middleware"
	"go-rest-api/internal/repository/sqlconnect"
	"log"
	"net/http"
	"os"
	"github.com/swaggo/http-swagger"
    
	
)

// @title Sample API
// @version 1.0
// @description A simple API with Swagger in Go
// @host localhost:3000
// @BasePath /api

// @Summary Returns pong
// @Description A basic ping endpoint
// @Success 200 {string} string "pong"
// @Router /api/ping [get]
func pingHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, `"pong"`)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
	}

	port := os.Getenv("API_PORT")

	cert := "cert.pem"
	key := "key.pem"

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.RootHandler)

	mux.HandleFunc("/login/", handlers.LoginHandler)

	mux.HandleFunc("/adduser/", handlers.AddUserHandler)

	//http.Handle("/swagger/", httpSwagger.WrapHandler)

	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)

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
