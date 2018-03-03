package channels

import (
	"github.com/OwlLaboratory/mongodm"
	"github.com/OwlLaboratory/go_api/DB"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

type Channel struct {
	mongodm.DocumentBase 	`json:",inline" bson:",inline"`

	Name     string 		`json:"name" bson:"name"`
	Platform Platform		`json:"platform" bson:"platform"`
}

type CreateChannelInput struct {
	Channel *ChannelInput
}

type UpdateChannelInput struct {
	ID    string
	Patch *ChannelInput
}

type ChannelInput struct {
	Name     string
	Platform *PlatformInput
}

type PlatformInput struct {
	Name     string
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
	return r.c.Created.String()
}

func (r *ChannelResolver) Updated() string {
	return r.c.Updated.String()
}

type PaginationArgs struct {
	First *int32
	Offset *int32
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

func ChannelListResolve(args PaginationArgs) (*[]*ChannelResolver, error) {
	c, _ := DB.Session.GetSession()
	collection := c.Model("Channel")

	channels := []*Channel{}

	limit := 0
	offset := 0

	if args.First != nil {
		limit = int(*args.First)
	}

	if args.Offset != nil {
		offset = int(*args.Offset)
	}

	collection.Find().Limit(int(limit)).Skip(int(offset)).Exec(&channels)

	var ret []*ChannelResolver

	for _, channel := range channels {
		ret = append(ret, &ChannelResolver{c: channel})
	}

	return &ret, nil
}

func CreateChannelMutation(input *CreateChannelInput) (*ChannelResolver, error) {
	c, _ := DB.Session.GetSession()
	collection := c.Model("Channel")

	channel := &Channel{}
	collection.New(channel) //this sets the connection/collection for this type and is strongly necessary(!) (otherwise panic)

	channel.Name = input.Channel.Name
	channel.Platform = Platform(*input.Channel.Platform)

	err := channel.Save()

	return &ChannelResolver{channel}, err
}

func UpdateChannelMutation(input *UpdateChannelInput) (*ChannelResolver, error) {
	c, _ := DB.Session.GetSession()
	collection := c.Model("Channel")

	channel := &Channel{}

	if !bson.IsObjectIdHex(input.ID) {
		return nil,  errors.New("invalid ObjectID")
	}

	err := collection.FindOne(bson.M{"_id" : bson.ObjectIdHex(input.ID)}).Exec(channel)

	if _, ok := err.(*mongodm.NotFoundError); ok {
		return nil, err
	}

	collection.New(channel) //this sets the connection/collection for this type and is strongly necessary(!) (otherwise panic)

	channel.Name = input.Patch.Name
	channel.Platform = Platform(*input.Patch.Platform)

	err = channel.Save()

	return &ChannelResolver{channel}, err
}