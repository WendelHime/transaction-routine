// Package postgres contain implementations that integrates with PostgreSQL
// database
package postgres

import (
	"database/sql"
	"fmt"

	// registering postgresql drivers
	_ "github.com/lib/pq"
)

// Connection establish informations to access
// the selected database
type Connection struct {
	// Host specify the database hostname like
	// localhost or 127.0.0.1 or a real IP address
	host string
	// Port specify the database port, by default
	// postgresql use 5432
	port int
	// User specify the user that can access the
	// postgresql database
	user string
	// Password specify the user password that can
	// access the postgresql database
	password string
	// DBname contains the database name
	dbname   string
	psqlinfo string
}

// NewConnection creates a new connection with the
// specified parameters
func NewConnection(host string, port int, user, password, dbname string) *Connection {
	return &Connection{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbname:   dbname,
	}
}

// Open connection to database
func (c *Connection) Open() (*sql.DB, error) {
	if c.psqlinfo == "" {
		c.psqlinfo = fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable", c.host, c.port,
			c.user, c.password, c.dbname)
	}
	return sql.Open("postgres", c.psqlinfo)
}
