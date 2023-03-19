// This is a golang program for learning go and has a few assorted programs!
package main

import (
	//"database/sql"
	//"log"
	//"time"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sid-008/Learning-Golang/mux"
)

func main() {
	var i int
	fmt.Println("Enter which code you want to run: ")
	fmt.Println("1. http server with mux thing(yes this is vague look at the package)")
	fmt.Scanln(&i)

	switch i {
	case 1:
		mux.Server()

	}
}

// http server,
/*func main() {
	r := mux.NewRouter()

	r.HandleFunc(
		"/books/{title}/page/{page}/car/{car}",
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			title := vars["title"]
			page := vars["page"]
			car := vars["car"]

			fmt.Fprintf(w, "you've requested %s page num %s\n", title, page)
			fmt.Fprintf(w, "car: %s", car)
		},
	)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})

	http.ListenAndServe(":3000", r)
}*/

/*func main() {
	db, err := sql.Open("mysql", "root:admin@(localhost)/db_1?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	if err == nil {
		log.Println("db connected!")
	}

	/*query := `
	    CREATE TABLE users (
	        id INT AUTO_INCREMENT,
	        username TEXT NOT NULL,
	        password TEXT NOT NULL,
	        created_at DATETIME,
	        PRIMARY KEY (id)
	    );`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}

	{ //insert a new user
		username := "malice"
		password := "bobbie"
		createdAt := time.Now()

		result, err := db.Exec(
			`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`,
			username,
			password,
			createdAt,
		)
		if err != nil {
			log.Fatal(err)
		}
		id, err := result.LastInsertId()
		fmt.Println(id)
	}

	{ //Query a single user
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)

		query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, username, password, createdAt)
	}

	{ // Query all users
		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}

		rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []user
		for rows.Next() {
			var u user

			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v", users)
	}

	{ // Delete user by id
		_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 5)
		if err != nil {
			log.Fatal(err)
		}
	}
}*/
