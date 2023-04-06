package pkg

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

func main() {
    connStr := "user=<user> password=<password> dbname=neondb host=<hostname>"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    rows, err := db.Query("select version()")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var version string
    for rows.Next() {
        err := rows.Scan(&version)
        if err != nil {
            log.Fatal(err)
        }
    }
    fmt.Printf("version=%s\n", version)
}