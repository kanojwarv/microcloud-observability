package server

import (
	"net/http"
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

	return http.ListenAndServe(
		addr,
		nil,
	)
}
