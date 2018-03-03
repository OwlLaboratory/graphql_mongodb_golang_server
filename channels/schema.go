package channels

var Queries = `
		channel(id: String!): Channel
		channels(first: Int, offset: Int): [Channel]
`

var Mutations = `
		createChannel(input: CreateChannelInput!): Channel
		updateChannel(input: UpdateChannelInput!): Channel
		deleteChannel(input: DeleteChannelInput!): Channel
`

var Interfaces = `
	interface IChannel {
		# The platform connected to the channel
		platform(): Platform!
	}
`

var Inputs = `
	input CreateChannelInput {
		channel: ChannelInput!
	}

	input UpdateChannelInput {
		id: String!
		patch: ChannelInput
	}

	input DeleteChannelInput {
		id: String!
	}

	# The input object sent when someone is creating/updating a new channel
	input ChannelInput {
		# Name of the channel
		name: String
		# Platform of the channel
		platform: PlatformInput
	}

	# The input platform object sent when someone is creating/updating a new channel
	input PlatformInput {
		# Name of the platform
		name: String
	}
`

var Types = `
	type Platform {
		# The platform name
		name: String!
	}

	# A channel entity
	type Channel implements Entity, IChannel {
		# The ID of the channel
		id: String!
		# What this human calls themselves
		name: String!
		# The updated time of the entity
		platform: Platform!
		# The created time of the entity
		created: String!
		# The updated time of the entity
		updated: String!
	}
`