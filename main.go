package main

import (
	"fmt"
	"github.com/frankffenn/go-utils/log"
	"github.com/frankffenn/trading-assistants/config"
	"github.com/frankffenn/trading-assistants/routers"
	"net/http"
)

func main() {
	confJson := "config.toml"
	if err := config.InitConfig(confJson); err != nil {
		log.Fatal("init config: %v", err)
	}

	if err := config.InitLog("debug"); err != nil {
		log.Fatal("init log: %v", err)
	}

	router := routers.InitRouter()
	srv := &http.Server{
		Addr:  config.Cfg.API.ListenAddress,
		Handler: router,
	}

	fmt.Printf("[API] listen on %s \n", config.Cfg.API.ListenAddress)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("listen: %s\n", err)
	}
}
