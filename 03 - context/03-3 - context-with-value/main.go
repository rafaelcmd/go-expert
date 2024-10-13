package main

import "context"

func main() {
	ctx := context.WithValue(context.Background(), "traceId", "12345")
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	traceId := ctx.Value("traceId")
	println(traceId)
}
