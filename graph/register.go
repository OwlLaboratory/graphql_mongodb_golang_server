package graph

import (
	"github.com/OwlLaboratory/go_api/channels"
	"github.com/OwlLaboratory/go_api/DB"
)

func RegisterNodes() {
	s, _ := DB.Session.GetSession()

	// Register each module (collection) here
	s.Register(&channels.Channel{}, "channels")
}