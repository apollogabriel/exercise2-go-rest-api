package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	mathrand "math/rand/v2"

	_ "github.com/go-sql-driver/mysql"
)

// Subject options for teachers
var subjects = []string{
	"Mathematics", "Physics", "Chemistry", "Biology", "History",
	"Geography", "English", "Literature", "Computer Science", "Art",
	"Music", "Physical Education", "Economics", "Psychology", "Philosophy",
}

// Class options
var classes = []string{
	"1A", "1B", "1C", "2A", "2B", "2C", "3A", "3B", "3C",
	"4A", "4B", "4C", "5A", "5B", "5C", "6A", "6B", "6C",
}

// Role options for executives
var roles = []string{
	"Principal", "Vice Principal", "Head of Department", "Administrator",
	"Counselor", "IT Manager", "Financial Officer", "HR Manager",
}

// Sample first names
var firstNames = []string{
	"Emma", "Liam", "Olivia", "Noah", "Ava", "Ethan", "Sophia", "Mason",
	"Isabella", "Jacob", "Mia", "William", "Charlotte", "James", "Amelia",
	"Alexander", "Harper", "Michael", "Evelyn", "Benjamin", "Abigail", "Daniel",
	"Emily", "Matthew", "Elizabeth", "Henry", "Sofia", "Jackson", "Avery", "Sebastian",
	"Ella", "Aiden", "Scarlett", "Lucas", "Grace", "Jack", "Chloe", "Owen",
	"Victoria", "Gabriel", "Riley", "Carter", "Aria", "Jayden", "Lily", "John",
	"Aubrey", "Luke", "Zoey", "Anthony", "Hannah", "Isaac", "Layla", "Dylan",
	"Ellie", "Wyatt", "Zoe", "Andrew", "Penelope", "Joshua", "Audrey", "Christopher",
	"Natalie", "Grayson", "Leah", "Julian", "Savannah", "Leo", "Aaliyah", "David",
}

// Sample last names
var lastNames = []string{
	"Smith", "Johnson", "Williams", "Jones", "Brown", "Davis", "Miller", "Wilson",
	"Moore", "Taylor", "Anderson", "Thomas", "Jackson", "White", "Harris", "Martin",
	"Thompson", "Garcia", "Martinez", "Robinson", "Clark", "Rodriguez", "Lewis", "Lee",
	"Walker", "Hall", "Allen", "Young", "Hernandez", "King", "Wright", "Lopez",
	"Hill", "Scott", "Green", "Adams", "Baker", "Gonzalez", "Nelson", "Carter",
	"Mitchell", "Perez", "Roberts", "Turner", "Phillips", "Campbell", "Parker", "Evans",
	"Edwards", "Collins", "Stewart", "Sanchez", "Morris", "Rogers", "Reed", "Cook",
	"Morgan", "Bell", "Murphy", "Bailey", "Rivera", "Cooper", "Richardson", "Cox",
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Connect to the database
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")
	host := os.Getenv("HOST")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, dbport, dbname)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Ping the database to verify connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Generate and insert data
	generateTeachers(db, 1000)
	generateStudents(db, 1000)
	generateExecutives(db, 1000)

	fmt.Println("Seed data generated successfully!")
}

// generateTeachers creates and inserts teacher records
func generateTeachers(db *sql.DB, count int) {
	stmt, err := db.Prepare(`
		INSERT INTO teachers (id, first_name, last_name, subject, class, email)
		VALUES (?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		log.Fatalf("Failed to prepare teachers statement: %v", err)
	}
	defer stmt.Close()

	fmt.Printf("Generating %d teacher records...\n", count)

	for i := 0; i < count; i++ {
		firstName := firstNames[mathrand.IntN(len(firstNames))]
		lastName := lastNames[mathrand.IntN(len(lastNames))]
		subject := subjects[mathrand.IntN(len(subjects))]
		class := classes[mathrand.IntN(len(classes))]

		// Creating a unique email with random suffix
		randomSuffix := makeRandomString(4)
		email := fmt.Sprintf("%s.%s.%s@school.edu",
			strings.ToLower(firstName),
			strings.ToLower(lastName),
			randomSuffix)

		// Create UUID for ID
		id := uuid.New().String()

		_, err := stmt.Exec(id, firstName, lastName, subject, class, email)
		if err != nil {
			// Just log and continue if there's an error (likely duplicate email)
			log.Printf("Error inserting teacher #%d: %v", i+1, err)
			i-- // retry
			continue
		}

		if (i+1)%100 == 0 {
			fmt.Printf("  %d teacher records generated\n", i+1)
		}
	}
}

// generateStudents creates and inserts student records
func generateStudents(db *sql.DB, count int) {
	stmt, err := db.Prepare(`
		INSERT INTO students (id, first_name, last_name, class, email)
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		log.Fatalf("Failed to prepare students statement: %v", err)
	}
	defer stmt.Close()

	fmt.Printf("Generating %d student records...\n", count)

	for i := 0; i < count; i++ {
		firstName := firstNames[mathrand.IntN(len(firstNames))]
		lastName := lastNames[mathrand.IntN(len(lastNames))]
		class := classes[mathrand.IntN(len(classes))]

		// Creating a unique email with random suffix
		randomSuffix := makeRandomString(4)
		email := fmt.Sprintf("student.%s.%s.%s@school.edu",
			strings.ToLower(firstName),
			strings.ToLower(lastName),
			randomSuffix)

		// Create UUID for ID
		id := uuid.New().String()

		_, err := stmt.Exec(id, firstName, lastName, class, email)
		if err != nil {
			// Just log and continue if there's an error (likely duplicate email)
			log.Printf("Error inserting student #%d: %v", i+1, err)
			i-- // retry
			continue
		}

		if (i+1)%100 == 0 {
			fmt.Printf("  %d student records generated\n", i+1)
		}
	}
}

// generateExecutives creates and inserts executive records
func generateExecutives(db *sql.DB, count int) {
	stmt, err := db.Prepare(`
		INSERT INTO executives (
			id, first_name, last_name, email, username, password, 
			last_password_change, user_creation_time, role, user_inactive
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		log.Fatalf("Failed to prepare executives statement: %v", err)
	}
	defer stmt.Close()

	fmt.Printf("Generating %d executive records...\n", count)

	for i := 0; i < count; i++ {
		firstName := firstNames[mathrand.IntN(len(firstNames))]
		lastName := lastNames[mathrand.IntN(len(lastNames))]
		role := roles[mathrand.IntN(len(roles))]

		// Creating a unique username and email with random suffix
		randomSuffix := makeRandomString(4)
		username := fmt.Sprintf("%s_%s_%s",
			strings.ToLower(firstName),
			strings.ToLower(lastName),
			randomSuffix)

		email := fmt.Sprintf("exec.%s.%s.%s@school.edu",
			strings.ToLower(firstName),
			strings.ToLower(lastName),
			randomSuffix)

		// Create UUID for ID
		id := uuid.New().String()

		// Create a password and hash it with bcrypt
		rawPassword := fmt.Sprintf("Pass%s!", makeRandomString(8))
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Failed to hash password: %v", err)
		}

		// Create random timestamps for last year
		now := time.Now()
		creationTime := now.Add(-time.Duration(mathrand.IntN(365)) * 24 * time.Hour)
		passwordChangeTime := creationTime.Add(time.Duration(mathrand.IntN(int(now.Sub(creationTime).Hours()/24))) * 24 * time.Hour)

		// Random boolean for user_inactive (mostly active)
		inactive := mathrand.IntN(10) == 0 // 10% chance of being inactive

		_, err = stmt.Exec(
			id, firstName, lastName, email, username, hashedPassword,
			passwordChangeTime, creationTime, role, inactive,
		)
		if err != nil {
			// Just log and continue if there's an error (likely duplicate email/username)
			log.Printf("Error inserting executive #%d: %v", i+1, err)
			i-- // retry
			continue
		}

		if (i+1)%100 == 0 {
			fmt.Printf("  %d executive records generated\n", i+1)
		}
	}
}

// makeRandomString generates a random string of specified length
func makeRandomString(length int) string {
	randomBytes := make([]byte, length/2) // Each byte becomes 2 hex chars
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatalf("Failed to generate random string: %v", err)
	}
	return hex.EncodeToString(randomBytes)
}
