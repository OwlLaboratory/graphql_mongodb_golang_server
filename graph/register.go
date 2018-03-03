package graph

import (
	"../channels"
	"../DB"
)

func RegisterNodes() {
	s, _ := DB.Session.GetSession()

	// Register each module (collection) here
	s.Register(&channels.Channel{}, "channels")
}