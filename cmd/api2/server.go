package main

import (
	"fmt"
	"go-rest-api/internal/api/handlers"

	"go-rest-api/internal/repository/sqlconnect"

	"net/http"

	"github.com/joho/godotenv"

	httpSwagger "github.com/swaggo/http-swagger"

	_ "go-rest-api/docs"
)

// @title Inventory API
// @version 1.0
// @description Login API with Swagger in Go
// @host localhost:3000
// @BasePath /

// @Summary Delete User
// @Description Delete credentials  http://localhost:3000/login/1
// @Tags delete
// @Accept  json
// @Produce  json
// @Param   id     path     int     true  "ID"
// @Success 200 {string} string "delete status"
// @Router /login/ [delete]
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
	}

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	http.HandleFunc("/login/", loginHandler)
	http.HandleFunc("/adduser/", handlers.AddUserHandler)
	http.Handle("/swagger/", httpSwagger.WrapHandler)
	fmt.Println("Server running on http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}

// @Summary Returns login username and password 
// @Description Login endpoint  http://localhost:3000/login/     [{"username":"admin","password","iloveadmin"}]
// @Tags login
// @Accept  json
// @Produce  json
// @Param   username     path     string     true  "Username"
// @Param   password     path     string     true  "Password"
// @Success 200 {string} string "login status"
// @Router /login/ [post]
func loginHandler(w http.ResponseWriter, r *http.Request) {
	
	// Handle CORS preflight request
	if r.Method == http.MethodOptions {
		handlePreflight(w, r)
		return
	}

	// Set CORS headers for the main request
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")

	// Example login response
	//fmt.Fprintf(w, `{"message":"Login successful"}`)

	//loginHandler(w,r)
	handlers.LoginHandler(w, r)
}


// @Summary Verify User status
// @Description Login Verify status endpoint http://localhost:3000/login/?name=admin
// @Tags users
// @Accept  json
// @Produce  json
// @Param   name     path     string     true  "Username"
// @Success 200 {string} string "login status"
// @Router /login/ [get]
func handlePreflight(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.WriteHeader(http.StatusOK)
}

// @Summary Add User
// @Description Add login credentials http://localhost:3000/login/     {"username":"admin","password","iloveadmin","account_status":"Y","account_group":"A","email":your@email.com}
// @Tags adduser
// @Accept  json
// @Produce  json
// @Param   username     path     string     true  "Username"
// @Param   password     path     string     true  "Password"
// @Param   account_status     path     string     true  "Y"
// @Param   account_group     path     string     true  "A"
// @Param   email     path     string     true  "Email Add"
// @Success 200 {string} string "add status"
// @Router /adduser/ [post]
func is_onlineHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // Get ?name=John
	if name == "" {
		name = "World"
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Send JSON response
	fmt.Fprintf(w, `{"message": "Hello, %s!"}`, name)
}


// @Summary Update User
// @Description Update login credentials http://localhost:3000/login/     {"username":"admin","password","iloveadmin","account_status":"Y","account_group":"A","email":your@email.com}
// @Tags adduser
// @Accept  json
// @Produce  json
// @Param   username     path     string     true  "Username"
// @Param   password     path     string     true  "Password"
// @Param   account_status     path     string     true  "Y"
// @Param   account_group     path     string     true  "A"
// @Param   email     path     string     true  "Email Add"
// @Success 200 {string} string "update status"
// @Router /login/ [PUT]
func updateuserHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // Get ?name=John
	if name == "" {
		name = "World"
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Send JSON response
	fmt.Fprintf(w, `{"message": "Hello, %s!"}`, name)
}


