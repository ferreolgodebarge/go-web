package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"./apis"
	"./front"
	"./models/users"
	"./utils"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

// CLI parses arguments to configure server
func CLI() (opts string, bindAddress string) {
	host := flag.String("dbhost", "127.0.0.1", "Database hostname")
	port := flag.String("dbport", "5432", "Database connection port")
	user := flag.String("dbuser", "user", "Database connection username")
	password := flag.String("password", "pass", "Database connection password")
	dbname := flag.String("dbname", "database", "Database name")
	ba := flag.String("bind", "127.0.0.1:80", "Bind address for server")
	sslmode := flag.String("sslmode", "disable", "Activate ssl connection or not")
	flag.Parse()
	opts = utils.CreateOptions(*host, *port, *user, *password, *dbname, *sslmode)
	bindAddress = *ba
	return
}

// InitDB opens connection to database
func InitDB(opts string) {
	db, err = gorm.Open("postgres", opts)
	if err != nil {
		panic(err)
	}
}

func main() {

	opts, bindAddress := CLI()
	InitDB(opts)
	users.Init(db)
	r := mux.NewRouter()
	r.HandleFunc("/home", front.HomeHandler)
	r.HandleFunc("/api/users", apis.ListUsersHandlers).Methods("GET")
	r.HandleFunc("/api/users", apis.CreateUserHandlers).Methods("POST")
	http.Handle("/", r)

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Server is starting...")

	nextRequestID := func() string {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}

	server := &http.Server{
		Addr:         bindAddress,
		Handler:      utils.Tracing(nextRequestID)(utils.Logging(logger)(r)),
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	server.ListenAndServe()
}
