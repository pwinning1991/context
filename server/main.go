package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/pwinning1991/context/log"
)

func main() {
	flag.Parse()
	http.HandleFunc("/", log.Decorate(handler))
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println(ctx, "handler started")
	defer log.Println(ctx, "handler ended")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "hello")
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	time.Sleep(5 * time.Second)
	fmt.Fprintln(w, "hello")

}
