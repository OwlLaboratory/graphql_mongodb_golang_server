package models

import (
	"github.com/neelance/graphql-go"
)

type channelResolver struct {
	c *channel
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

type channel struct {
	ID        	graphql.ID
	Name      	string
	Platform	Platform
}

var channels = []*channel{
	{ ID: "1000", Name: "Luke Skywalker", Platform: Platform{"platform"}},
}

var channelData = make(map[graphql.ID]*channel)

func init() {
	for _, c := range channels {
		channelData[c.ID] = c
	}
}

func (r *Resolver) Channel(args struct{ ID graphql.ID }) *channelResolver {
	if c := channelData[args.ID]; c != nil {
		return &channelResolver{c}
	}
	return nil
}
