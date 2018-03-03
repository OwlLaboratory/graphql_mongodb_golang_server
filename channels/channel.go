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

type ChannelArgs struct {
	Name     string
	Platform PlatformArgs
}

type PlatformArgs struct {
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

func CreateChannelMutation(entity *ChannelArgs) (*ChannelResolver, error) {
	c, _ := DB.Session.GetSession()
	collection := c.Model("Channel")


	channel := &Channel{}
	today := time.Now()
	collection.New(channel) //this sets the connection/collection for this type and is strongly necessary(!) (otherwise panic)

	channel.Name = entity.Name
	channel.Platform = Platform(entity.Platform)
	channel.Created = &today

	err := channel.Save()

	return &ChannelResolver{channel}, err
}