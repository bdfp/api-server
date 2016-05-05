package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"encoding/json"
	"github.com/julienschmidt/httprouter"
)

// Env Functionalities required by the entire Application
type Env struct {
	Db *sql.DB
}

// RegisterHandlers adds the route handlers for various calls
func RegisterHandlers(router *httprouter.Router, env Env) {

}
