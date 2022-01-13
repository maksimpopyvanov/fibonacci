package handler

import (
	"encoding/json"
	"fibonacci"
	"io/ioutil"
	"net/http"
)

func (h *Handler) getSequence(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		input := new(fibonacci.Input)
		err = json.Unmarshal([]byte(body), &input)

		if err != nil || input.Start < 0 || input.End > 93 || input.End < input.Start {
			http.Error(w, "invalid parametrs", http.StatusBadRequest)
			return
		}

		output := h.repos.GetSequence(input)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(output)
	} else {
		http.Error(w, "Need post-request", http.StatusBadRequest)
		return
	}
}
