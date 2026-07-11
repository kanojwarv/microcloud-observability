package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/kanojwarv/microcloud-observability/internal/impact"
)

func handleImpact(
	w http.ResponseWriter,
	r *http.Request,
) {

	artifact := strings.TrimPrefix(
		r.URL.Path,
		"/api/impact/",
	)

	dependents, err := impact.Analyze(
		".",
		artifact,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusNotFound,
		)

		return
	}

	response := map[string]any{
		"artifact":     artifact,
		"referencedBy": dependents,
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
