package channels

type Platform struct {
	Name	string
}

type platformResolver struct {
	p *Platform
}

func (r *ChannelResolver) Platform() (*platformResolver, error) {
	return &platformResolver{p: &r.c.Platform}, nil
}

func (r *platformResolver) Name() string {
	return r.p.Name
}