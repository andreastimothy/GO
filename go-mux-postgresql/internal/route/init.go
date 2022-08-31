package route

import (
	"github.com/gorilla/mux"
)

func InitRoute(r *mux.Router) {
	InitProductRoute(r)
}