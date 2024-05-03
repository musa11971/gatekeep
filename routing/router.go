package routing

import (
	"fmt"
	"musa11971/gatekeep/controllers"
	"musa11971/gatekeep/policy"
	"net/http"

	"github.com/gorilla/mux"
)

func Initialize() {
	r := mux.NewRouter()

	go registerPolicyRoutes(r)

	r.NotFoundHandler = http.HandlerFunc(controllers.NotFoundHandler)

	http.Handle("/", r)
}

func registerPolicyRoutes(router *mux.Router) {
	// To do: it should register each policy as a go-routine.
	for _, p := range policy.Policies {
		fmt.Println("Registering RoutingPath", p.RoutingPath)

		router.PathPrefix(p.RoutingPath).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			controllers.ForwardToPolicyHandler(p, w, r)
		})
	}
}
