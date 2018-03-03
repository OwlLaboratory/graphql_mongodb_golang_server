package channels

import (
	"github.com/OwlLaboratory/mongodm"
)

type Channel struct {
	mongodm.DocumentBase 	`json:",inline" bson:",inline"`

	Name     *string 		`json:"name" bson:"name"`
	Platform Platform		`json:"platform" bson:"platform"`
}

type ChannelResolver struct {
	c *Channel
}

func (r *ChannelResolver) ID() string {
	return r.c.Id.Hex()
}

func (r *ChannelResolver) Name() string {
	return *r.c.Name
}

func (r *ChannelResolver) Created() string {
	return r.c.Created.String()
}

func (r *ChannelResolver) Updated() string {
	return r.c.Updated.String()
}