package main

import (
	"os"
	"os/signal"
	"syscall"

	"bitcrunchy.com/zenfighter-api/adapters/http"
	"bitcrunchy.com/zenfighter-api/engine"
	"bitcrunchy.com/zenfighter-api/providers/database"
)

func main() {
	provider := database.NewProvider()
	e := engine.NewEngine(provider)

	adapter := http.NewHTTPAdapter(e)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	defer close(stop)

	adapter.Start()

	<-stop

	adapter.Stop()
	provider.Close()
}
