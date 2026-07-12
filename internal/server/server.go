package server

import (
	"encoding/json"
	"net/http"

	"github.com/kanojwarv/microcloud-observability/internal/generate"

	"github.com/kanojwarv/microcloud-observability/internal/doctor"
)

func Start(
	addr string,
) error {

	http.HandleFunc(
		"/",
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			if r.URL.Path == "/" {

				http.Redirect(
					w,
					r,
					"/graph.html",
					http.StatusFound,
				)

				return
			}

			http.FileServer(
				http.Dir(
					"artifacts",
				),
			).ServeHTTP(
				w,
				r,
			)
		},
	)

	http.HandleFunc(
		"/api/graph",
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			data, err := generate.GraphJSON(
				".",
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

			w.Write(
				data,
			)
		},
	)

	http.HandleFunc(
		"/api/doctor",
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			health := doctor.Run(".")

			data, err := json.MarshalIndent(
				health,
				"",
				"  ",
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

			w.Write(
				data,
			)
		},
	)

	http.HandleFunc(
		"/api/impact/",
		handleImpact,
	)

	http.HandleFunc(
		"/api/search",
		handleSearch,
	)

	return http.ListenAndServe(
		addr,
		nil,
	)
}
