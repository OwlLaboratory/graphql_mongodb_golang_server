package channels

import (
	"../DB"
	"gopkg.in/mgo.v2/bson"
	"errors"
	"github.com/OwlLaboratory/mongodm"
)

func CreateChannel(input *CreateChannelInput) (*ChannelResolver, error) {
	c, _ := DB.Session.GetSession()
	collection := c.Model("Channel")

	channel := &Channel{}
	collection.New(channel) //this sets the connection/collection for this type and is strongly necessary(!) (otherwise panic)

	channel.Name = input.Channel.Name
	channel.Platform = Platform(*input.Channel.Platform)

	err := channel.Save()

	return &ChannelResolver{channel}, err
}

func UpdateChannel(input *UpdateChannelInput) (*ChannelResolver, error) {
	c, _ := DB.Session.GetSession()
	collection := c.Model("Channel")

	channel := &Channel{}

	if !bson.IsObjectIdHex(input.ID) {
		return nil, errors.New("invalid ObjectID")
	}

	err := collection.FindOne(bson.M{"_id" : bson.ObjectIdHex(input.ID)}).Exec(channel)

	if _, ok := err.(*mongodm.NotFoundError); ok {
		return nil, err
	}

	if input.Patch.Name != nil {
		channel.Name = input.Patch.Name
	}

	if input.Patch.Platform != nil {
		channel.Platform = Platform(*input.Patch.Platform)
	}

	err = channel.Save()

	return &ChannelResolver{channel}, err
}

func DeleteChannel(input *DeleteChannelInput) (*ChannelResolver, error) {
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

	err = channel.FullDelete()

	return &ChannelResolver{channel}, err
}