package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lib/pq"
)

//Dbtype :will contain all variables required for db
type Dbtype struct {
	db *sql.DB
}

var blacklist map[string]bool

// SecretKey used for jwt token verification and generation
var SecretKey string = "mykey"

// Dbhandler database struct global variable
var Dbhandler Dbtype

func main() {
	blacklist = make(map[string]bool)
	pqurl, err := pq.ParseURL("postgres://ccmnzryy:12OjOODSZeS_yTLUB-sDdJ3sU7swHDuz@arjuna.db.elephantsql.com:5432/ccmnzryy")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(pqurl)
	Dbhandler.db, err = sql.Open("postgres", pqurl)
	defer Dbhandler.db.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = Dbhandler.db.Ping()

	if err != nil {
		log.Fatal(err)
	}
	EndpointsInit()
}

// Homepage used for homepage routing
func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the home page")
}
