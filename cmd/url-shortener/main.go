package main

import (
	"fmt"

	"github.com/antoneka/url-shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Printf("%+v\n", cfg)

	// TODO: init logger: slog

	// TODO: init storage: sqlite

	// TODO: init router: chi, "chi render"

	// TODO: run server
}
