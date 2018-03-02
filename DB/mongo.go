/*
	Singleton for MongoDB
 */

package DB

import (
	"fmt"
	"github.com/zebresel-com/mongodm"
)

type session struct {
	connection *mongodm.Connection
	err		error
}

func (s *session) GetSession() (*mongodm.Connection, error) {
	// connect to the database
	if s.connection == nil {

		dbConfig := &mongodm.Config{
			DatabaseHosts: []string{"127.0.0.1"},
			DatabaseName: "fameus",
		}

		s.connection, s.err = mongodm.Connect(dbConfig)

		if s.err != nil {
			fmt.Println("Database connection error: %v", s.err)
		}
	}
	return s.connection, s.err
}

func (s *session) CloseSession() () {
	if s.connection != nil {
		s.connection.Close()
	}
}

var Session = session{}