package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
)
type Comment struct {
	ID        int
	Username  string
	Content   string
	CreatedAt time.Time
}


var db *sql.DB
var err error

func main() {
	db, err := sql.Open("mysql", "root: @tcp(127.0.0.1:3306)/database")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)

	log.Println("Server started on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		t, _ := template.ParseFiles("static/register.html")
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username=?", username).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count > 0 {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}

	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE email=?", email).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count > 0 {
		http.Error(w, "Email already exists", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", username, email, string(hashedPassword))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "User %s has been registered!", username)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "static/login.html")
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	var dbUsername, dbPassword string
	err := db.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&dbUsername, &dbPassword)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Welcome, %s!", username)
}
func createComment(db *sql.DB, comment Comment) {
	insertComment := `
		INSERT INTO comments (username, content, created_at) VALUES (?, ?, ?)
	`
	_, err := db.Exec(insertComment, comment.Username, comment.Content, comment.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Комментарий создан")
}

func readComments(db *sql.DB) []Comment {
	getComments := `
		SELECT id, username, content, created_at FROM comments
	`
	rows, err := db.Query(getComments)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.ID, &comment.Username, &comment.Content, &comment.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return comments
}

func readComment(db *sql.DB, id int) Comment {
	getComment := `
		SELECT id, username, content, created_at FROM comments WHERE id=?
	`
	row := db.QueryRow(getComment, id)

	var comment Comment
	err := row.Scan(&comment.ID, &comment.Username, &comment.Content, &comment.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("Комментарий с ID=%d не найден\n", id)
		} else {
			log.Fatal(err)
		}
	}

	return comment
}

func updateComment(db *sql.DB, id int, content string) {
	updateComment := `
		UPDATE comments SET content=? WHERE id=?
	`
	result, err := db.Exec(updateComment, content, id)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		log.Fatalf("Комментарий с ID=%d не найден\n", id)
	}
	fmt.Printf("Комментарий с ID=%d изменен\n", id)
}

func deleteComment(db *sql.DB, id int) {
	deleteComment := `
		DELETE FROM comments WHERE id=?
	`
	result, err := db.Exec(deleteComment, id) 
	
		defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	
	comment1 := Comment{
		Username:  "user1",
		Content:   "Привет, мир!",
		CreatedAt: time.Now(),
	}
	createComment(db, comment1)
	
	comments := readComments(db)
	fmt.Println("Все комментарии:")
	for _, comment := range comments {
		fmt.Printf("ID: %d\n", comment.ID)
		fmt.Printf("Username: %s\n", comment.Username)
		fmt.Printf("Content: %s\n", comment.Content)
		fmt.Printf("CreatedAt: %s\n", comment.CreatedAt)
	}
	
	comment2 := readComment(db, 2)
	fmt.Printf("Комментарий с ID=2: %v\n", comment2)
	
	updateComment(db, 1, "Новое содержание комментария")
	fmt.Printf("Измененный комментарий с ID=1: %v\n", readComment(db, 1))
	
	deleteComment(db, 2)
	
	comments = readComments(db)
	fmt.Println("Все комментарии:")
	for _, comment := range comments {
		fmt.Printf("ID: %d\n", comment.ID)
		fmt.Printf("Username: %s\n", comment.Username)
		fmt.Printf("Content: %s\n", comment.Content)
		fmt.Printf("CreatedAt: %s\n", comment.CreatedAt)
	}
}
defer db.Close()

createTable := `
	CREATE TABLE IF NOT EXISTS comments (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(255) NOT NULL,
		content VARCHAR(255) NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)
`
_, err = db.Exec(createTable)
if err != nil {
	log.Fatal(err)
}

comment1 := Comment{
	Username:  "user1",
	Content:   "Комментарий от user1",
	CreatedAt: time.Now(),
}
createComment(db, comment1)

comment2 := Comment{
	Username:  "user2",
	Content:   "Комментарий от user2",
	CreatedAt: time.Now(),
}
createComment(db, comment2)

comments := readComments(db)
fmt.Println("Все комментарии:")
for _, comment := range comments {
	fmt.Printf("%d\t%s\t%s\t%s\n", comment.ID, comment.Username, comment.Content, comment.CreatedAt)
}

comment := readComment(db, 1)
fmt.Printf("Комментарий с ID=%d:\n%s\t%s\t%s\t%s\n", comment.ID, comment.Username, comment.Content, comment.CreatedAt)

updateComment(db, 2, "Новый комментарий от user2")

deleteComment(db, 1)

func main() {
db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/database")
if err != nil {
log.Fatal(err)
}
defer db.Close()

newComment := Comment{
	Username:  "JohnDoe",
	Content:   "Привет, мир!",
	CreatedAt: time.Now(),
}
createComment(db, newComment)

fmt.Println(readComments(db))

fmt.Println(readComment(db, 1))

updateComment(db, 1, "Новое содержание комментария")
fmt.Println(readComment(db, 1))

deleteComment(db, 1)
fmt.Println(readComments(db))
}