package main

import (
	"wallet/internal/app"
)

func main() {
	result, err := app.Bootstrap()
	if err != nil {
		panic(err)
	}
	defer result.Log.Sync()

	startHTTPServer(result.Engine, result.Cfg, result.Log)
}
