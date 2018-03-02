package channels

import (
	"time"
	"github.com/zebresel-com/mongodm"
	"github.com/OwlLaboratory/go_api/DB"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

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

func ChannelResolve(args struct{ ID string }) (*ChannelResolver, error) {
	c, _ := DB.Session.GetSession()
	collection := c.Model("Channel")

	channel := &Channel{}
	if !bson.IsObjectIdHex(args.ID) {
		return nil,  errors.New("invalid ObjectID")
	}

	collection.FindOne(bson.M{"_id" : bson.ObjectIdHex(args.ID)}).Exec(channel)

	return &ChannelResolver{c: channel}, nil
}