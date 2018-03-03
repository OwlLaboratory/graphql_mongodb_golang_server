package channels

type CreateChannelInput struct {
	Channel 	*ChannelInput
}

type UpdateChannelInput struct {
	ID    		string
	Patch 		*ChannelInput
}

type DeleteChannelInput struct {
	ID    		string
}

type ChannelInput struct {
	Name     	*string
	Platform 	*PlatformInput
}

type PlatformInput struct {
	Name     	*string
}

type PaginationInput struct {
	First 		*int32
	Offset 		*int32
}