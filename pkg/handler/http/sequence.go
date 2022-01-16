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

		if err != nil || input.End < input.Start || input.End > 10000 || input.Start < 0 {
			http.Error(w, "invalid parametrs", http.StatusBadRequest)
			return
		}

		if output := h.repos.GetSequence(input); output == nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(output)
		}

	} else {
		http.Error(w, "Need post-request", http.StatusBadRequest)
		return
	}
}
