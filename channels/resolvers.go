package channels

import (
	"../DB"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

func GetChannel(args struct{ ID string }) (*ChannelResolver, error) {
	c, _ := DB.Session.GetSession()
	collection := c.Model("Channel")

	channel := &Channel{}
	if !bson.IsObjectIdHex(args.ID) {
		return nil,  errors.New("invalid ObjectID")
	}

	collection.FindOne(bson.M{"_id" : bson.ObjectIdHex(args.ID)}).Exec(channel)

	return &ChannelResolver{c: channel}, nil
}

func GetChannels(args PaginationInput) (*[]*ChannelResolver, error) {
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