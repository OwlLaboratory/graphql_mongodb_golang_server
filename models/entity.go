package models

import (
	"github.com/neelance/graphql-go"
)


type entity interface {
	ID() graphql.ID
	Name() string
}

type entityResolver struct {
	entity
}

func (r *entityResolver) ToChannel() (*channelResolver, bool) {
	c, ok := r.entity.(*channelResolver)
	return c, ok
}