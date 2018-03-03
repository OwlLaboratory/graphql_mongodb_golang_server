package graph

var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}

	# The query type, represents all of the entry points into our object graph
	type Query {
		channel(id: String!): Channel
		channels(first: Int, offset: Int): [Channel]
	}

	# The mutation type, represents all updates we can make to our data
	type Mutation {
		createChannel(input: CreateChannelInput!): Channel
		updateChannel(input: UpdateChannelInput!): Channel
	}

	# A character from the Star Wars universe
	interface Entity {
		# The ID of the entity
		id: String!
		# The created time of the entity
		created: String!
		# The updated time of the entity
		updated: String!
	}

	interface IChannel {
		# The platform connected to the channel
		platform(): Platform!
	}

	input CreateChannelInput {
		channel: ChannelInput!
	}

	input UpdateChannelInput {
		id: String!
		patch: ChannelInput
	}

	# The input object sent when someone is creating a new review
	input ChannelInput {
		# Name of the channel
		name: String!
		# Platform of the channel
		platform: PlatformInput
	}

	# The input object sent when someone is creating a new review
	input PlatformInput {
		# Name of the platform
		name: String!
	}

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
		# The created time of the entity
		created: String!
		# The updated time of the entity
		updated: String!
		# The updated time of the entity
		platform: Platform!
	}
	union SearchResult = Channel
`