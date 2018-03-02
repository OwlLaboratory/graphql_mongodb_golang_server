package models

type entity interface {
	ID() string
	Created() string
	Updated() string
}

type entityResolver struct {
	entity
}

func (r *entityResolver) ToChannel() (*channelResolver, bool) {
	c, ok := r.entity.(*channelResolver)
	return c, ok
}