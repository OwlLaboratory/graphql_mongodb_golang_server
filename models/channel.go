package models

import (
	"time"
	"github.com/OwlLaboratory/go_api/DB"
	"github.com/zebresel-com/mongodm"
)

var collectionName = "channels"

type channelResolver struct {
	c *Channel
}

type platformResolver struct {
	p *Platform
}


func (r *channelResolver) ID() string {
	return r.c.Id.Hex()
}

func (r *channelResolver) Name() string {
	return r.c.Name
}

func (r *channelResolver) Platform() (*platformResolver, error) {
	return &platformResolver{p: &r.c.Platform}, nil
}

func (r *platformResolver) Name() string {
	return r.p.Name
}

func (r *channelResolver) Created() string {
	if r.c.Created == nil {
		return ""
	}
	return r.c.Created.String()
}

func (r *channelResolver) Updated() string {
	if r.c.Updated == nil {
		return ""
	}
	return r.c.Updated.String()
}

type Platform struct {
	Name	string
}

type Channel struct {
	mongodm.DocumentBase 	`json:",inline" bson:",inline"`

	Name     string 		`json:"name" bson:"name"`
	Platform Platform		`json:"platform" bson:"platform"`
	Created  *time.Time		`json:"created" bson:"created"`
	Updated  *time.Time		`json:"updated" bson:"updated"`
}

func (r *Resolver) Channel(args struct{ ID string }) *channelResolver {
	c, _ := DB.Session.GetSession()
	collection := c.Model("Channel")


	channel := &Channel{}
	collection.FindOne().Exec(channel)

	return &channelResolver{c: channel}
}
