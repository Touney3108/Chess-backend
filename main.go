package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	_ "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections (adjust for security in production)
	},
}
//ctrl shift p tasksa nodemon golang to run
func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error while upgrading connection:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected")

	for {
		// Read message from client
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error while reading message:", err)
			break
		}

		fmt.Printf("Received message: %s\n", msg)

		// Echo the message back to the client
		if err := conn.WriteMessage(messageType, msg); err != nil {
			fmt.Println("Error while writing message:", err)
			break
		}
	}
}

func main() {
    username := "admin"  // Replace with your MySQL username
	password := "admin"  // Replace with your MySQL password
	ip := "127.0.0.1"            // MySQL server IP
	port := "3306"                // MySQL port
	database := "sah_database"  // Replace with your database name

	// Step 2: Combine them into a connection string
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, ip, port, database)
	db, err := sql.Open("mysql", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
	// Optionally, check if the connection is alive
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Welcome to Go + MySQL API apps")
    })
	http.HandleFunc("/ws", handleConnection)

    log.Println("Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
