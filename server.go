package main

import (
	"log"
	"net/http"
	"github.com/graph-gophers/graphql-go/relay"
	"./graph"
	"github.com/OwlLaboratory/graphql_mongodb_golang_server/DB"
)

func main() {
	graph.RegisterNodes()
	defer DB.Session.CloseSession()

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	}))

	http.Handle("/query", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
		w.Header().Set( "Accept", "application/json")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Range, Content-Disposition, Content-Type")

		// Handle OPTIONS for preflight request / Angular app
		if r.Method == "OPTIONS" {
			return
		}

		handlerScheme := relay.Handler{Schema: graph.GQLSchema}
		handlerScheme.ServeHTTP(w, r)
	}))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

var page = []byte(`
<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.css" />
		<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.1.0/fetch.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.js"></script>
	</head>
	<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
		<div id="graphiql" style="height: 100vh;">Loading...</div>
		<script>
			function graphQLFetcher(graphQLParams) {
				return fetch("/query", {
					method: "post",
					body: JSON.stringify(graphQLParams),
					credentials: "include",
				}).then(function (response) {
					return response.text();
				}).then(function (responseBody) {
					try {
						return JSON.parse(responseBody);
					} catch (error) {
						return responseBody;
					}
				});
			}

			ReactDOM.render(
				React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
				document.getElementById("graphiql")
			);
		</script>
	</body>
</html>
`)
