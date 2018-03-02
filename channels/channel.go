package channels

import (
	"time"
	"github.com/zebresel-com/mongodm"
	"github.com/OwlLaboratory/go_api/DB"
)

type iChannel interface {
	Platform() (*platformResolver, error)
}

type Channel struct {
	mongodm.DocumentBase 	`json:",inline" bson:",inline"`

	Name     string 		`json:"name" bson:"name"`
	Platform Platform		`json:"platform" bson:"platform"`
	Created  *time.Time		`json:"created" bson:"created"`
	Updated  *time.Time		`json:"updated" bson:"updated"`
}


type ChannelResolver struct {
	c *Channel
}

func (r *ChannelResolver) ID() string {
	return r.c.Id.Hex()
}

func (r *ChannelResolver) Name() string {
	return r.c.Name
}

func (r *ChannelResolver) Created() string {
	if r.c.Created == nil {
		return ""
	}
	return r.c.Created.String()
}

func (r *ChannelResolver) Updated() string {
	if r.c.Updated == nil {
		return ""
	}
	return r.c.Updated.String()
}

func ChannelResolverId(args struct{ ID string }) *ChannelResolver {
	c, _ := DB.Session.GetSession()
	collection := c.Model("Channel")


	channel := &Channel{}
	collection.FindOne().Exec(channel)

	return &ChannelResolver{c: channel}
}