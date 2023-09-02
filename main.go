package main

import (
	"github.com/ashep/go-apprun"

	"github.com/ashep/ghwh/internal/app"
)

func main() {
	apprun.Run("ghwh", app.New, app.Config{})
}
