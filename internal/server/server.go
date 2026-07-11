package server

import (
	"fmt"
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

			fmt.Fprintf(
				w,
				"MicroCloud Observability",
			)
		},
	)

	return http.ListenAndServe(
		addr,
		nil,
	)
}
