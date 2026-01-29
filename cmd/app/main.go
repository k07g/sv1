package main

import (
	"fmt"
	"os"

	"github.com/k07g/sv1/internal/app"
)

func main() {
	cmd := app.NewRootCmd()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
