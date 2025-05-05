package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"go-rest-api/internal/models"
	"go-rest-api/internal/repository/sqlconnect"
	"net/http"
	
	"strings"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		//fmt.Println("A")
		getLoginHandler2(w, r)
	case http.MethodPost:
		fmt.Println("B")
		getLoginHandler(w, r)
		//addLoginHandler(w, r)
		//getLoginHandler(w, r)
	case http.MethodPut:
		updateUserHander(w,r)
		//w.Write([]byte("Hello PUT teachers Route"))
		//fmt.Println("Hello PUT teachers Route")
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH teachers Route"))
		fmt.Println("Hello PATCH teachers Route")
	case http.MethodDelete:
		deleteUserHandler(w,r)
		//w.Write([]byte("Hello DELETE teachers Route"))
		//fmt.Println("Hello DELETE teachers Route")
	}
}


// Task godoc
// @Summary Login
// @Tags V1
// @Description Get tasks
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Param X-API-Key header string true "API Key"
// @Success 200 {object}  models.Login2
// @Router /login/ [post]
func getLoginHandler(w http.ResponseWriter, r *http.Request) {
	

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	path := strings.TrimPrefix(r.URL.Path, "/login/")
	usr := strings.TrimSuffix(path, "/")
	


	//fmt.Println("APOLLO PARA ",usr)
	
	if usr == "" {
		fmt.Println("A123")
		//GET PARA USSERNAME and PASSWORD
		var newLogins []models.Login2
		err = json.NewDecoder(r.Body).Decode(&newLogins)
		fmt.Println("A1234")
		fmt.Println("PARA ",newLogins)
		if err != nil {
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}
		fmt.Println("A12345")
		fmt.Println("PARA ",newLogins)
		//addedLogins := make([]models.Login, 0, len(newLogins))
		for _, newLogin := range newLogins {
			query := "SELECT id, id2, username, password FROM login where username='"+newLogin.USERNAME+"' AND password='"+newLogin.PASSWORD+"'"
			//queryTotal := "SELECT COUNT(id) FROM login"
			

			loginList := make([]models.Login, 0)

			fmt.Println("SQL QUERY : ",query)

			rows, err := db.Query(query)
			if err != nil {
				http.Error(w, "Error querying teachers", http.StatusInternalServerError)
				return
			}
			defer rows.Close()

			for rows.Next() {
				var login models.Login
				err = rows.Scan(&login.ID, &login.ID2, &login.USERNAME, &login.PASSWORD)
				if err != nil {
					http.Error(w, "Error scanning teachers", http.StatusInternalServerError)
					return
				}
				loginList = append(loginList, login)
			}

			response := struct {
				Status string           `json:"status"`
				Data   []models.Login `json:"data"`
			}{
				Status: "success",
				Data:   loginList,
			}
	
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}

		

		

		

		
	}

	
}

// Task godoc
// @Summary Get all saved tasks
// @Tags V1
// @Description Get tasks
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Param X-API-Key header string true "API Key"
// @Success 200 {object}  string "is online"
// @Router /login [get]
func getLoginHandler2(w http.ResponseWriter, r *http.Request) {
	type User struct {
		ID    int    `json:"id"`
		IS_ONLINE  string `json:"is_online"`
		ACCOUNT_GROUP string `json:"account_group"`
	}

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	name := r.URL.Query().Get("name") // Get ?name=John

	if name == "" {
        http.Error(w, "Missing id parameter", http.StatusBadRequest)
        return
    }


	var user User
	err = db.QueryRow("SELECT id,is_online,account_group FROM login WHERE account_status='Y' AND username = ?", name).Scan(&user.ID, &user.IS_ONLINE, &user.ACCOUNT_GROUP)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
    
	
}

func updateUserHander(w http.ResponseWriter, r *http.Request){
	fmt.Println("UPDATE PUT")
	type User2 struct {
		ID    int    `json:"id"`
		USERNAME  string `json:"username"`
		PASSWORD string `json:"password"`
		ACCOUNT_STATUS string `json:"account_status"`
		ACCOUNT_GROUP string `json:"account_group"`
		EMAIL string `json:"email"`
	}

	if r.Method != http.MethodPut {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	fmt.Println(r.Body)
    var user User2
    err = json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

	

    res, err := db.Exec("UPDATE login SET username = ?, password = ? , account_status = ? , account_group = ? ,email = ? WHERE id = ?", user.USERNAME, user.PASSWORD, user.ACCOUNT_STATUS, user.ACCOUNT_GROUP, user.EMAIL, user.ID)
    if err != nil {
        http.Error(w, "Failed to update user", http.StatusInternalServerError)
        return
    }

    rowsAffected, _ := res.RowsAffected()
    if rowsAffected == 0 {
        http.Error(w, "Zero affected rows", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
        return
    }

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

    res, err := db.Exec("DELETE FROM login WHERE id = ?", id)
    if err != nil {
        http.Error(w, "Failed to delete user", http.StatusInternalServerError)
        return
    }

    rowsAffected, _ := res.RowsAffected()
    if rowsAffected == 0 {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, `{"message": "User deleted successfully"}`)
}

func addLoginHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sqlconnect.ConnectDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	var newLogins []models.Login
	err = json.NewDecoder(r.Body).Decode(&newLogins)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare(`
		INSERT INTO login (id2, username, password, account_status, account_group, email)
		VALUES (?,?, ?, ?, ?, ?)
	`)
	if err != nil {
		http.Error(w, "Error preparing statement", http.StatusInternalServerError)
	}
	defer stmt.Close()

	//addedLogins := make([]models.Login, 0, len(newLogins))
	for _, newLogin := range newLogins {
		//check for duplicate email
		var existingID string
		err = db.QueryRow("SELECT id FROM login WHERE username = ?", newLogin.USERNAME).Scan(&existingID)
		if err == nil {
			//Email already exists
			http.Error(w, "Username already exists", http.StatusBadRequest)
			return
		}
		var existingID2 string
		err = db.QueryRow("SELECT id FROM login WHERE email = ?", newLogin.EMAIL).Scan(&existingID2)
		if err == nil {
			//Email already exists
			http.Error(w, "Email already exists", http.StatusBadRequest)
			return
		}

		id := uuid.New().String()
		_, err := stmt.Exec(id, newLogin.USERNAME, newLogin.PASSWORD, newLogin.ACCOUNT_STATUS, newLogin.ACCOUNT_GROUP, newLogin.EMAIL)
		if err != nil {
			http.Error(w, "Error inserting login", http.StatusInternalServerError)
			return
		}else{
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			response := struct {
				Status string           `json:"status"`
				Count  int              `json:"count"`
				Data   string `json:"data"`
			}{
				Status: "success",
				Count:  1,
				Data:   "Added login",
			}
		
			json.NewEncoder(w).Encode(response)
		}
		//newTeacher.ID = id
		//addedTeachers = append(addedTeachers, newTeacher)
	}

	
}
