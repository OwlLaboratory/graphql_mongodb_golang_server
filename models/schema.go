package models
/*
import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	graphql "github.com/neelance/graphql-go"
)*/
/*

type Resolver struct{}

var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}
	# The query type, represents all of the entry points into our object graph
	type Query {
		search(text: String!): [SearchResult]!
		character(id: ID!): Character
		human(id: ID!): Human
	}
	# The mutation type, represents all updates we can make to our data
	type Mutation {
	}
	# The episodes in the Star Wars trilogy
	enum Episode {
		# Star Wars Episode IV: A New Hope, released in 1977.
		NEWHOPE
		# Star Wars Episode V: The Empire Strikes Back, released in 1980.
		EMPIRE
		# Star Wars Episode VI: Return of the Jedi, released in 1983.
		JEDI
	}
	# A character from the Star Wars universe
	interface Character {
		# The ID of the character
		id: ID!
		# The name of the character
		name: String!
		# The friends of the character, or an empty list if they have none
		friends: [Character]
		# The movies this character appears in
		appearsIn: [Episode!]!
	}
	# Units of height
	enum LengthUnit {
		# The standard unit around the world
		METER
		# Primarily used in the United States
		FOOT
	}
	# A humanoid creature from the Star Wars universe
	type Human implements Character {
		# The ID of the human
		id: ID!
		# What this human calls themselves
		name: String!
		# Height in the preferred unit, default is meters
		height(unit: LengthUnit = METER): Float!
		# Mass in kilograms, or null if unknown
		mass: Float
		# This human's friends, or an empty list if they have none
		friends: [Character]
		# The movies this human appears in
		appearsIn: [Episode!]!
	}
	union SearchResult = Human
`
*/
/*

type human struct {
	ID        graphql.ID
	Name      string
	Friends   []graphql.ID
	AppearsIn []string
	Height    float64
	Mass      int
	Starships []graphql.ID
}

var humans = []*human{
	{
		ID:        "1000",
		Name:      "Luke Skywalker",
		Friends:   []graphql.ID{"1002", "1003", "2000", "2001"},
		AppearsIn: []string{"NEWHOPE", "EMPIRE", "JEDI"},
		Height:    1.72,
		Mass:      77,
		Starships: []graphql.ID{"3001", "3003"},
	},
	{
		ID:        "1001",
		Name:      "Darth Vader",
		Friends:   []graphql.ID{"1004"},
		AppearsIn: []string{"NEWHOPE", "EMPIRE", "JEDI"},
		Height:    2.02,
		Mass:      136,
		Starships: []graphql.ID{"3002"},
	},
	{
		ID:        "1002",
		Name:      "Han Solo",
		Friends:   []graphql.ID{"1000", "1003", "2001"},
		AppearsIn: []string{"NEWHOPE", "EMPIRE", "JEDI"},
		Height:    1.8,
		Mass:      80,
		Starships: []graphql.ID{"3000", "3003"},
	},
	{
		ID:        "1003",
		Name:      "Leia Organa",
		Friends:   []graphql.ID{"1000", "1002", "2000", "2001"},
		AppearsIn: []string{"NEWHOPE", "EMPIRE", "JEDI"},
		Height:    1.5,
		Mass:      49,
	},
	{
		ID:        "1004",
		Name:      "Wilhuff Tarkin",
		Friends:   []graphql.ID{"1001"},
		AppearsIn: []string{"NEWHOPE"},
		Height:    1.8,
		Mass:      0,
	},
}

var humanData = make(map[graphql.ID]*human)

func init() {
	for _, h := range humans {
		humanData[h.ID] = h
	}
}
func (r *Resolver) Search(args struct{ Text string }) []*searchResultResolver {
	var l []*searchResultResolver
	for _, h := range humans {
		if strings.Contains(h.Name, args.Text) {
			l = append(l, &searchResultResolver{&humanResolver{h}})
		}
	}
	return l
}

func (r *Resolver) Character(args struct{ ID graphql.ID }) *characterResolver {
	if h := humanData[args.ID]; h != nil {
		return &characterResolver{&humanResolver{h}}
	}
	return nil
}

func (r *Resolver) Human(args struct{ ID graphql.ID }) *humanResolver {
	if h := humanData[args.ID]; h != nil {
		return &humanResolver{h}
	}
	return nil
}


type friendsConnectionArgs struct {
	First *int32
	After *graphql.ID
}

type character interface {
	ID() graphql.ID
	Name() string
	Friends() *[]*characterResolver
	FriendsConnection(friendsConnectionArgs) (*friendsConnectionResolver, error)
	AppearsIn() []string
}

type characterResolver struct {
	character
}

func (r *characterResolver) ToHuman() (*humanResolver, bool) {
	c, ok := r.character.(*humanResolver)
	return c, ok
}

type humanResolver struct {
	h *human
}

func (r *humanResolver) ID() graphql.ID {
	return r.h.ID
}

func (r *humanResolver) Name() string {
	return r.h.Name
}

func (r *humanResolver) Height(args struct{ Unit string }) float64 {
	return convertLength(r.h.Height, args.Unit)
}

func (r *humanResolver) Mass() *float64 {
	if r.h.Mass == 0 {
		return nil
	}
	f := float64(r.h.Mass)
	return &f
}

func (r *humanResolver) Friends() *[]*characterResolver {
	return resolveCharacters(r.h.Friends)
}

func (r *humanResolver) FriendsConnection(args friendsConnectionArgs) (*friendsConnectionResolver, error) {
	return newFriendsConnectionResolver(r.h.Friends, args)
}

func (r *humanResolver) AppearsIn() []string {
	return r.h.AppearsIn
}


type searchResultResolver struct {
	result interface{}
}

func (r *searchResultResolver) ToHuman() (*humanResolver, bool) {
	res, ok := r.result.(*humanResolver)
	return res, ok
}

func convertLength(meters float64, unit string) float64 {
	switch unit {
	case "METER":
		return meters
	case "FOOT":
		return meters * 3.28084
	default:
		panic("invalid unit")
	}
}

func resolveCharacters(ids []graphql.ID) *[]*characterResolver {
	var characters []*characterResolver
	for _, id := range ids {
		if c := resolveCharacter(id); c != nil {
			characters = append(characters, c)
		}
	}
	return &characters
}

func resolveCharacter(id graphql.ID) *characterResolver {
	if h, ok := humanData[id]; ok {
		return &characterResolver{&humanResolver{h}}
	}
	return nil
}

type friendsConnectionResolver struct {
	ids  []graphql.ID
	from int
	to   int
}

func newFriendsConnectionResolver(ids []graphql.ID, args friendsConnectionArgs) (*friendsConnectionResolver, error) {
	from := 0
	if args.After != nil {
		b, err := base64.StdEncoding.DecodeString(string(*args.After))
		if err != nil {
			return nil, err
		}
		i, err := strconv.Atoi(strings.TrimPrefix(string(b), "cursor"))
		if err != nil {
			return nil, err
		}
		from = i
	}

	to := len(ids)
	if args.First != nil {
		to = from + int(*args.First)
		if to > len(ids) {
			to = len(ids)
		}
	}

	return &friendsConnectionResolver{
		ids:  ids,
		from: from,
		to:   to,
	}, nil
}

func (r *friendsConnectionResolver) TotalCount() int32 {
	return int32(len(r.ids))
}

func (r *friendsConnectionResolver) Edges() *[]*friendsEdgeResolver {
	l := make([]*friendsEdgeResolver, r.to-r.from)
	for i := range l {
		l[i] = &friendsEdgeResolver{
			cursor: encodeCursor(r.from + i),
			id:     r.ids[r.from+i],
		}
	}
	return &l
}

func (r *friendsConnectionResolver) Friends() *[]*characterResolver {
	return resolveCharacters(r.ids[r.from:r.to])
}

func (r *friendsConnectionResolver) PageInfo() *pageInfoResolver {
	return &pageInfoResolver{
		startCursor: encodeCursor(r.from),
		endCursor:   encodeCursor(r.to - 1),
		hasNextPage: r.to < len(r.ids),
	}
}

func encodeCursor(i int) graphql.ID {
	return graphql.ID(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("cursor%d", i+1))))
}

type friendsEdgeResolver struct {
	cursor graphql.ID
	id     graphql.ID
}

func (r *friendsEdgeResolver) Cursor() graphql.ID {
	return r.cursor
}

func (r *friendsEdgeResolver) Node() *characterResolver {
	return resolveCharacter(r.id)
}

type pageInfoResolver struct {
	startCursor graphql.ID
	endCursor   graphql.ID
	hasNextPage bool
}

func (r *pageInfoResolver) StartCursor() *graphql.ID {
	return &r.startCursor
}

func (r *pageInfoResolver) EndCursor() *graphql.ID {
	return &r.endCursor
}

func (r *pageInfoResolver) HasNextPage() bool {
	return r.hasNextPage
}*/
