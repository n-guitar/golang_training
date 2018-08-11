// A Basic Model
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todo []Todo

unc TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}
  
	json.NewEncoder(w).Encode(todos)
  }