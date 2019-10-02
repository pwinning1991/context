package log

import (
	"context"
	"log"
	"net/http"
	"rand"
)

const requestIDKey = 42

func Println(ctx context.Context, msg string) {
	id, ok := ctx.Value(requestIDKey).(int64)
	if !ok {
		log.Println("could not find request ID in conetxt")
		return
	}

	log.Printf("[%d] %s", id, msg)
}

func Decorate(f http.HandlerFunc) http.HandlerFun {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int64()
		ctx = context.WithValue(ctx, requestIDKey, id)
		f(w, r.WithContext(ctx))

	}
}
