package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Драйвер PostgreSQL
	"log"
	"net/http"
)

func main() {
	// Параметры подключения
	connStr := "host=192.168.145.185 port=5433 user=manu password=password dbname=lesson15 sslmode=disable"

	// Открываем соединение
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Проверяем подключение
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ Успешно подключились к базе данных!")
	//users, err := getAllUsers(db)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for i := range users {
	//	log.Println("user:", users[i])
	//}

	if err := deleteUser(db, 3); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/login", LoginHandler)
	mux.Handle("/protected", AuthMiddleware(http.HandlerFunc(ProtectedHandler)))

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", mux)
}
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("userClaims").(*Claims)
	fmt.Fprintf(w, "Welcome, %s (user ID: %d)\n", claims.Email, claims.UserID)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	token, _ := GenerateToken(1, "user@example.com")
	fmt.Fprintf(w, "Your token:\n%s", token)
}

func deleteUser(db *sql.DB, id int) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}
	log.Println("deleted user by id:", id)
	return nil
}

func updateUser(db *sql.DB, user *User) error {
	query := `UPDATE users SET name = $1, email = $2`

	result, err := db.Exec(query, user.Name, user.Email)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	log.Println("RowsAffected:", rowsAffected)
	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}
func getAllUsers(db *sql.DB) ([]User, error) {
	var users []User
	query := `SELECT id, name, email FROM users`

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func getUserByID(db *sql.DB, id int) (*User, error) {
	var user User
	query := `SELECT id, name, email FROM users WHERE id = $1`

	err := db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &user, nil
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func createUser(db *sql.DB, user *User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`

	err := db.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		return err
	}
	log.Println("User created successfully")
	return nil
}
func createTables(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(100) UNIQUE NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    
    CREATE TABLE IF NOT EXISTS products (
        id SERIAL PRIMARY KEY,
        name VARCHAR(200) NOT NULL,
        price DECIMAL(10,2) NOT NULL,
        stock INTEGER DEFAULT 0,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`

	_, err := db.Exec(query)
	return err
}
