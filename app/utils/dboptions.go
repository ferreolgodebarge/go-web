package utils

import "fmt"

// CreateOptions generates connection configurations
func CreateOptions(
	host string,
	port string,
	user string,
	password string,
	dbname string,
	sslmode string,
) (opts string) {
	opts = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, user, dbname, password, sslmode)
	return
}
