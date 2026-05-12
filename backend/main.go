package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"

    _ "github.com/go-sql-driver/mysql"
)

type Course struct {
    ID int `json:"id"`
    Name string `json:"name"`
}

func main() {

    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
        dbUser,
        dbPassword,
        dbHost,
        dbPort,
        dbName,
    )

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }

    http.HandleFunc("/api/courses", func(w http.ResponseWriter, r *http.Request) {

        if r.Method == "GET" {

            rows, err := db.Query("SELECT id, name FROM courses")

            if err != nil {
                http.Error(w, err.Error(), 500)
                return
            }

            defer rows.Close()

            var courses []Course

            for rows.Next() {
                var c Course
                rows.Scan(&c.ID, &c.Name)
                courses = append(courses, c)
            }

            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(courses)

        } else if r.Method == "POST" {

            var course Course

            err := json.NewDecoder(r.Body).Decode(&course)

            if err != nil {
                http.Error(w, err.Error(), 400)
                return
            }

            _, err = db.Exec(
                "INSERT INTO courses(name) VALUES(?)",
                course.Name,
            )

            if err != nil {
                http.Error(w, err.Error(), 500)
                return
            }

            w.WriteHeader(http.StatusCreated)
            w.Write([]byte("Course Added Successfully"))
        }
    })

    fmt.Println("Backend running on port 8080")

    log.Fatal(http.ListenAndServe(":8080", nil))
}