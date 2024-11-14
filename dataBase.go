package main

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Titles struct {
	Title1 string `json:"title1"`
	Title2 string `json:"title2"`
	Title3 string `json:"title3"`
}

func connectDB() {
	// JSON data
	jsonData := []byte(`{
		"title1": "Welcome to MyWeihnachtsMarkt!",
		"title2": "Eat, Trink, Rate, Repeat",
		"title3": "From Hamburg with love"
	}`)

	var titles Titles
	if err := json.Unmarshal(jsonData, &titles); err != nil {
		log.Fatal(err)
	}

	// Connect to SQLite database
	db, err := sql.Open("sqlite3", "./mydatabase.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the table if it doesn't exist
	sqlStmt := `CREATE TABLE IF NOT EXISTS market (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT
	);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}

	// Insert data into the database
	insertStmt := `INSERT INTO market (title) VALUES (?)`
	_, err = db.Exec(insertStmt, titles.Title1)
	if err != nil {
		log.Fatalf("%q: %s\n", err, insertStmt)
	}
	_, err = db.Exec(insertStmt, titles.Title2)
	if err != nil {
		log.Fatalf("%q: %s\n", err, insertStmt)
	}
	_, err = db.Exec(insertStmt, titles.Title3)
	if err != nil {
		log.Fatalf("%q: %s\n", err, insertStmt)
	}

	log.Println("Data inserted successfully to the SQL_Lite_3!")

	// Get all titles from the database
	getAllTitles(db)
}

func getAllTitles(db *sql.DB) {
	rows, err := db.Query("SELECT title FROM market")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var titles []string
	for rows.Next() {
		var title string
		if err := rows.Scan(&title); err != nil {
			log.Fatal(err)
		}
		titles = append(titles, title)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println("Titles from the database:")
	for _, title := range titles {
		log.Println(title)
	}
}
