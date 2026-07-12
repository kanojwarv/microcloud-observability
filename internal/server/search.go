package server

import (
	"encoding/json"
	"net/http"

	"github.com/kanojwarv/microcloud-observability/internal/search"
)

func handleSearch(
	w http.ResponseWriter,
	r *http.Request,
) {

	query := r.URL.Query().Get(
		"q",
	)

	if query == "" {

		http.Error(
			w,
			"missing query parameter: q",
			http.StatusBadRequest,
		)

		return
	}

	response, err := search.Find(
		".",
		query,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	err = json.NewEncoder(
		w,
	).Encode(
		response,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
	}
}
