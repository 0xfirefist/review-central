// package main

// import (
// 	"log"
// 	"net/http"

// 	mux "github.com/gorilla/mux"
// 	server "github.com/kalradev/review-central/graphql-server"
// 	auth "github.com/kalradev/review-central/internal/auth"
// 	"github.com/kalradev/review-central/internal/db"
// 	"github.com/kalradev/review-central/internal/users"
// )

// const (
// 	uiDirectory     = "./ui/build"
// 	graphqlEndpoint = "/graphql"
// )

// func main() {
// 	// initialize database
// 	db.Initialize(&users.User{}, &users.IPFSHash{})

// 	router := mux.NewRouter()
// 	router.Use(auth.Middleware())
// 	// graphql server
// 	router.Handle(graphqlEndpoint, server.GetGraphQLHandler())
// 	// graphql playground to test API
// 	router.HandleFunc("/graphql/playground", server.GetPlaygroundHandler(graphqlEndpoint))

// 	// This will serve files under http://localhost:8000/<filename>
// 	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(uiDirectory))))

// 	// running server
// 	// srv := &http.Server{
// 	// 	Handler: router,
// 	// 	Addr:    "127.0.0.1:8000",
// 	// 	// Good practice: enforce timeouts for servers you create!
// 	// 	WriteTimeout: 15 * time.Second,
// 	// 	ReadTimeout:  15 * time.Second,
// 	// }
// 	log.Println("Started Listening on Port : 8000")
// 	log.Fatal(http.ListenAndServe(":8000", router))
// }

package main

import (
	"log"

	contracts "github.com/kalradev/review-central/contracts"
)

func main() {
	// contracts.Deploy()
	// log.Println(contracts.Put("token", []string{"1", "2"}))
	log.Println(contracts.GetTokens())
	log.Println(contracts.Get("token"))
}
