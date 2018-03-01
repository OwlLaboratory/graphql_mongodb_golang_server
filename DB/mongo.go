/*
	Singleton for MongoDB
 */

package DB

import (
	"gopkg.in/mgo.v2"
)

type session struct {
	session *mgo.Session
	err		error
}

func (s *session) GetSession() (*mgo.Session, error) {
	// connect to the database
	if s.session == nil {
		s.session, s.err = mgo.Dial("localhost")
	}
	return s.session, s.err
}

func (s *session) GetCollection(collection string) (*mgo.Collection) {
	session, _ := s.GetSession()
	c := session.DB("fameus").C(collection)

	return c
}

func (s *session) CloseSession() () {
	if s.session != nil {
		s.session.Close()
	}
}

var Session = session{}