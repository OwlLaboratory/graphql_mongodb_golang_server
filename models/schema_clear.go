package models

type Resolver struct{}

var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}

	# The query type, represents all of the entry points into our object graph
	type Query {
		channel(id: ID!): Channel
	}

	# A character from the Star Wars universe
	interface Entity {
		# The ID of the entity
		id: ID!
		# The created time of the entity
		created: String!
		# The updated time of the entity
		updated: String!
	}

	# The mutation type, represents all updates we can make to our data
	type Mutation {
	}

	# A channel entity
	type Channel implements Entity {
		# The ID of the channel
		id: ID!
		# What this human calls themselves
		name: String!
	}
	union SearchResult = Channel
`