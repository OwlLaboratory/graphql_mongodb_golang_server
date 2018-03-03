package graph

import (
	"github.com/OwlLaboratory/go_api/channels"
)

type Resolver struct{}

func (r *Resolver) Channel(args struct{ ID string }) (*channels.ChannelResolver, error) {
	return channels.ChannelResolve(args)
}

func (r *Resolver) Channels(args channels.PaginationArgs) (*[]*channels.ChannelResolver, error) {
	return channels.ChannelListResolve(args)
}

func (r *Resolver) CreateChannel(args struct{Input *channels.CreateChannelInput }) (*channels.ChannelResolver, error) {
	return channels.CreateChannelMutation(args.Input)
}

func (r *Resolver) UpdateChannel(args struct{Input *channels.UpdateChannelInput }) (*channels.ChannelResolver, error) {
	return channels.UpdateChannelMutation(args.Input)
}