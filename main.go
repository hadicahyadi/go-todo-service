package main

import(
  "fmt"
  "strings"
	"net/http"

	"github.com/gorilla/mux"
  "github.com/hadicahyadi/todo-service/api"
)

func main() {
  // Declare a new router
	r := mux.NewRouter()
	// Define available router from api package
	r.HandleFunc("/", defaultHandler).Methods("GET")
  mount(r, "/api", api.Router())
  // Up and running
	http.ListenAndServe(":8080", r)
}

func mount(r *mux.Router, path string, handler http.Handler) {
    r.PathPrefix(path).Handler(
        http.StripPrefix(
            strings.TrimSuffix(path, "/"),
            handler,
        ),
    )
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Please go to your API request name!")
}
