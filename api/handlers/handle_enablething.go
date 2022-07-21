package handlers

import (
	"encoding/json"
	"net/http"
)

func (env *Env) EnableThing(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// define this close to its usage
	var input struct {
		ThingID int `json:"thing_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = env.EnableThingService.EnableThing(input.ThingID)
	if err != nil {
		if err.Error() == "thing not found" {
			http.Error(w, "Not found", http.StatusNotFound)
		}
		http.Error(w, "Internal error", http.StatusInternalServerError)
	}

	http.Error(w, "OK", http.StatusOK)
}
