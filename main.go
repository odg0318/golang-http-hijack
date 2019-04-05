package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hijack", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/hijack")
		hj, ok := w.(http.Hijacker)
		if !ok {
			http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
			return
		}
		conn, bufrw, err := hj.Hijack()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Don't forget to close the connection:
		defer conn.Close()

		bufrw.WriteString("HTTP/1.1 200 OK\nContent-Type: text/plain; charset=utf-8\nContent-Length: 5\n\n")
		bufrw.Flush()

		fmt.Fprintf(bufrw, "h")
		time.Sleep(time.Second * 1)
		bufrw.Flush()

		fmt.Fprintf(bufrw, "e")
		time.Sleep(time.Second * 1)
		bufrw.Flush()

		fmt.Fprintf(bufrw, "l")
		time.Sleep(time.Second * 1)
		bufrw.Flush()

		fmt.Fprintf(bufrw, "l")
		time.Sleep(time.Second * 1)
		bufrw.Flush()

		fmt.Fprintf(bufrw, "o")
		bufrw.Flush()

		fmt.Fprintf(bufrw, "\n")
		bufrw.Flush()
	})

	http.ListenAndServe(":8000", nil)
}
