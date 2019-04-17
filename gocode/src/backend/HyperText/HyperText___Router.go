package HyperText

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(routes []Route) *mux.Router {
/*________________________________________TESTING FUNCTION________________________________________*/log.Println("\n\n...\n")

	router := mux.NewRouter()
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
