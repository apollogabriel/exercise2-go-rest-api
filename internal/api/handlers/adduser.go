package handlers

import (
	"fmt"
	"net/http"

	"encoding/json"
	
	"github.com/google/uuid"
	"go-rest-api/internal/models"
	"go-rest-api/internal/repository/sqlconnect"
	
	
	
)



// @Summary Add User
// @Description A basic ping endpoint
// @Success 200 {string} string "success added user"
// @Router /adduser/ [post]
func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Add User Route"))
	fmt.Println("Hello Add User Route")

	// if r.Method == http.MethodOptions {
	// 	handlePreflight(w, r)
	// 	return
	// }

	// // Set CORS headers for the main request
	// w.Header().Set("Access-Control-Allow-Origin", "https://localhost:5173")
	// w.Header().Set("Access-Control-Allow-Credentials", "true")
	// w.Header().Set("Content-Type", "application/json")

	// // Example login response
	// fmt.Fprintf(w, `{"message":"Connected successful"}`)

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	fmt.Println("AA")

	var newLogins []models.Login
	err = json.NewDecoder(r.Body).Decode(&newLogins)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	fmt.Println("AB")
	stmt, err := db.Prepare(`
		INSERT INTO login (id2, username, password, account_status, account_group, email)
		VALUES (?,?, ?, ?, ?, ?)
	`)
	if err != nil {
		http.Error(w, "Error preparing statement", http.StatusInternalServerError)
	}
	defer stmt.Close()
	fmt.Println("AC")
	//addedLogins := make([]models.Login, 0, len(newLogins))
	for _, newLogin := range newLogins {
		//check for duplicate email
		var existingID string
		err = db.QueryRow("SELECT id FROM login WHERE username = ? OR email = ?", newLogin.USERNAME,newLogin.EMAIL).Scan(&existingID)
		if err == nil {
			//Email already exists
			http.Error(w, "Username or email already exists", http.StatusBadRequest)
			return
		}
		fmt.Println("AD")
		

		id := uuid.New().String()
		_, err := stmt.Exec(id, newLogin.USERNAME, newLogin.PASSWORD, newLogin.ACCOUNT_STATUS, newLogin.ACCOUNT_GROUP, newLogin.EMAIL)
		if err != nil {
			http.Error(w, "Error inserting login", http.StatusInternalServerError)
			return
		}

		fmt.Println("AE")
		w.Header().Set("Content-Type", "application/json")
		//w.WriteHeader(http.StatusCreated)
		response := struct {
			Status string           `json:"status"`
			Count  int              `json:"count"`
			Data   string `json:"data"`
		}{
			Status: "success",
			Count:  1,
			Data:   "Added User",
		}
	
		json.NewEncoder(w).Encode(response)
		
		//newTeacher.ID = id
		//addedTeachers = append(addedTeachers, newTeacher)
	}

	
}

