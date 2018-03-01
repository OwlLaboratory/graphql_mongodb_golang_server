package models

import (
	"github.com/neelance/graphql-go"
	"github.com/OwlLaboratory/go_api/DB"
)

var collectionName = "channels"

type channelResolver struct {
	c *Channel
}

func (r *channelResolver) ID() graphql.ID {
	return r.c.ID
}

func (r *channelResolver) Name() string {
	return r.c.Name
}

func (r *channelResolver) Platform() Platform {
	return r.c.Platform
}

type Platform struct {
	Name	string
}

type Channel struct {
	ID        	graphql.ID
	Name      	string
	Platform	Platform
}

func (r *Resolver) Channel(args struct{ ID graphql.ID }) *channelResolver {
	s := DB.Session.GetCollection(collectionName)

	channel := Channel{}

	s.Find(nil).One(&channel)
	return &channelResolver{c: &channel}
}
