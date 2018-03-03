package graph

import "github.com/OwlLaboratory/go_api/channels"

var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}

	# The query type, represents all of the entry points into our object graph
	type Query {
		` + channels.Queries + `
	}

	# The mutation type, represents all updates we can make to our data
	type Mutation {
		` + channels.Mutations + `
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

	` + channels.Interfaces + `
	` + channels.Inputs + `
	` + channels.Types + `

	union SearchResult = Channel
`