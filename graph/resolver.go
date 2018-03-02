package graph

import (
	"github.com/OwlLaboratory/go_api/channels"
)

type Resolver struct{}

func (r *Resolver) Channel(args struct{ ID string }) (*channels.ChannelResolver, error) {
	return channels.ChannelResolve(args)
}

func (r *Resolver) Channels() (*[]*channels.ChannelResolver, error) {
	return channels.ChannelListResolve()
}