package graph

import (
	"github.com/OwlLaboratory/graphql_mongodb_golang_server/channels"
	"github.com/OwlLaboratory/graphql_mongodb_golang_server/DB"
)

func RegisterNodes() {
	s, _ := DB.Session.GetSession()

	// Register each module (collection) here
	s.Register(&channels.Channel{}, "channels")
}