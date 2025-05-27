package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Specify what a item record would look like
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	var err error
	conn_retries := 10

	// Retrieve the data source name from the environment. Docker would have injected the variable into the environment to be retrieved
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		log.Fatal("MYSQL_DSN env variable could not be found. Please ensure it is available before running again")
	}

	// Set a loop that retries connection to the DB up to X number of times with x-second delay. Set the environment variables as you see fit
	for i := 0; i < conn_retries; i++ {
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Printf("Attempt number %d: DB connection error: %v", i+1, err)
		} else if err = db.Ping(); err != nil {
			log.Printf("Attempt number %d: DB ping failed: %v", i+1, err)
		} else {
			log.Println("Successfully connected to DB")
			break
		}

		// After x number of tries, accept defeat and abort the mission.
		if i == conn_retries-1 {
			log.Fatalf("Could not connect to DB after %v attempts. Aborting.", 15)
		}

		// Delay. Perhaps exponential backoff would be good too.
		time.Sleep(2 * time.Second)
	}
	// If all good, close the db connection
	defer db.Close()

	// Create a table if it doesn't already exist. In this case we will be doing an items table.
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS items (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL
	)`)
	if err != nil {
		log.Fatal("Table creation failed with the error:", err)
	}

	http.HandleFunc("/items", itemsHandler)
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// This will handle the different requests types sent in. For now it will handle for POST and GET and throw a warning if some other method was sent in.
func itemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var item Item
		// First attempt to parse the JSON body received from the request into an Item struct (the one specified above)
		// Should the JSON structure fail to be parsed, HTTP 400 will be returned
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// Executes the insertion query. ? serves as a placeholder to prevent SQL injections, and item.Name will be passed in as an argument
		res, err := db.Exec("INSERT INTO items (name) VALUES (?)", item.Name)
		if err != nil {
			http.Error(w, "DB insertion failed", http.StatusInternalServerError)
			return
		}

		// The id of the newly inserted row is then retrieved and is assigned to the Item struct
		id, _ := res.LastInsertId()
		item.ID = int(id)

		// Response body is then crafted and the newly created item (record) is returned to the requester
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(item)

	case "GET":
		// Executes a select query
		rows, err := db.Query("SELECT id, name FROM items")
		if err != nil {
			http.Error(w, "DB query failed", http.StatusInternalServerError)
			return
		}

		// This helps ensure that the rows result is properly closed when the function exits (NOTE: this be important as it frees up DB resources)
		defer rows.Close()

		// Each record in rows is iterated through, created an Item struct for each and appends it to the items slice if Item struct was successfully created
		var items []Item
		for rows.Next() {
			var it Item
			if err := rows.Scan(&it.ID, &it.Name); err != nil {
				http.Error(w, "DB scan failed", http.StatusInternalServerError)
				return
			}
			items = append(items, it)
		}

		// Similarly, crafts and sends the items back to the requester in JSON format
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)

	default:
		// Handles for any methods that aren't part of the handler
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
