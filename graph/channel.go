package graph

import "github.com/OwlLaboratory/go_api/channels"

/*
	Resolver binding for channel module.
*/

func (r *Resolver) Channel(args struct{ ID string }) (*channels.ChannelResolver, error) {
	return channels.GetChannel(args)
}

func (r *Resolver) Channels(args channels.PaginationInput) (*[]*channels.ChannelResolver, error) {
	return channels.GetChannels(args)
}

func (r *Resolver) CreateChannel(args struct{Input *channels.CreateChannelInput }) (*channels.ChannelResolver, error) {
	return channels.CreateChannel(args.Input)
}

func (r *Resolver) UpdateChannel(args struct{Input *channels.UpdateChannelInput }) (*channels.ChannelResolver, error) {
	return channels.UpdateChannel(args.Input)
}

func (r *Resolver) DeleteChannel(args struct{Input *channels.DeleteChannelInput }) (*channels.ChannelResolver, error) {
	return channels.DeleteChannel(args.Input)
}