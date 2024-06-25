package main

import "net/http"

func main() {
	http.ListenAndServe(":8686", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := hello("foo", "bar").Render(r.Context(), w); err != nil {
			panic(err)
		}
	}))
}
