package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
}

func handleLogin(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	password := r.FormValue("password")
	query := fmt.Sprintf("SELECT * FROM users WHERE username='%s' AND password='%s'", username, password)
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	if rows.Next() {
		http.ServeFile(w, r, "welcome.html")
	} else {
		http.ServeFile(w, r, "login_failed.html")
	}
}

func handleSensitiveData(w http.ResponseWriter, r *http.Request) {
	var username, password string
	err := db.QueryRow("SELECT username, password FROM users WHERE id=1").Scan(&username, &password)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("User: %s, Password: %s", username, password)))
}

func handleXSS(w http.ResponseWriter, r *http.Request) {

	userInput := r.URL.Query().Get("user_input")
	htmlResponse := fmt.Sprintf("<html><body><h1>Hello, %s!</h1></body></html>", userInput)
	w.Write([]byte(htmlResponse))
}

func main() {
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/sensitive", handleSensitiveData)
	http.HandleFunc("/xss", handleXSS)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
