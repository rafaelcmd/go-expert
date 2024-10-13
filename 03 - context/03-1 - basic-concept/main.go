package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking is canceled. Timeout exceeded")
		return
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Hotel booking is done")
	}
}
