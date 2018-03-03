# Simple GraphQL + MongoDB Golang server

## Why this repo ?

Since a few months I was curious to discover two new concepts: Golang and GraphQL
A few days ago, I started to create my own GraphQL Golang server with MongoDB as a database.


I'm pretty sure it will help some people understand how a Golang GraphQL server works.

Of course I have very little experience under these two technologies, so I do not guarantee anything.

## Run server

```bash
    git clone https://github.com/OwlLaboratory/graphql_mongodb_golang_server $HOME/go/src/github.com/OwlLaboratory/graphql_mongodb_golang_server
    cd $HOME/go/src/github.com/OwlLaboratory/graphql_mongodb_golang_server/
    go get ./...
    go build server.go
    ./server
```

## Server + ODM

For the server part I based on the [neelance server](https://github.com/graph-gophers/graphql-go) (he is the main creator of GopherJS)

For the ODM I based on [mongodm](https://github.com/zebresel-com/mongodm) (which I slightly modified)

MongoDB is configured on localhost:27017 by default (you can change config in mongo.go)

I try to follow [some best practices](https://dev-blog.apollodata.com/designing-graphql-mutations-e09de826ed97) and implement a CLEAR architecture for a module

I implemented a simple CRUD on a Mongo collection called "Channel" and defined like this:

```
{
  id: ObjectID,
  name: string,
  platform {
    name: string
  },
  created: MongoDate
  updated: MongoDate
}
```

### Available queries:

```GraphQL
		channel(id: String!): Channel
		channels(first: Int, offset: Int): [Channel]
```  
### Available mutations:

```GraphQL
		createChannel(input: CreateChannelInput!): Channel
		updateChannel(input: UpdateChannelInput!): Channel
		deleteChannel(input: DeleteChannelInput!): Channel
```

### Example of GraphQL requests:

```GraphQL
mutation createChannel {
	createChannel(input: {channel: {name: "Channel 1", platform: {name: "Platform for channel 1"}}}) {
    id
 		name
    platform {
      name
    }
	}
}

mutation updateChannel {
	updateChannel(input: {id: "5a9b13a7e1382307023003d7", patch: {platform: {name: "Yolo"}}}) {
 		name
    platform {
      name
    }
	}
}

mutation deleteChannel {
	deleteChannel(input: {id: "5a9b13a7e1382307023003d7"}) {
 		name
	}
}

{
  channels(first: 1, offset: 1) {
    id
    name
  }
}


{
  channel(id: "5a9b13a7e1382307023003d7") {
    name
    platform {
      name
    }
  }
}
```

## Architecture 

I tried to implement a clear architecture rather than storing everything in a "models" package

Each collection is represented as an independent package containing all the resources needed for GraphQL Server.

### Example for channel collection

![module](https://image.ibb.co/gjD9L7/module.png)


## Run server

```bash
    git clone https://github.com/OwlLaboratory/graphql_mongodb_golang_server $HOME/go/src/github.com/OwlLaboratory/graphql_mongodb_golang_server
    cd $HOME/go/src/github.com/OwlLaboratory/graphql_mongodb_golang_server/
    go get ./...
    go build server.go
    ./server
```