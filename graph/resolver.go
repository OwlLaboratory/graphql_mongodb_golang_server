package graph

import "github.com/OwlLaboratory/go_api/channels"

type Resolver struct{}

func (r *Resolver) Channel(args struct{ ID string }) *channels.ChannelResolver {
	return channels.ChannelResolverId(args)
}