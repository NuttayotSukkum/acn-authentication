package main

import (
	"context"
	"fmt"
	"github.com/NuttayotSukkum/acn/acn-authentication/configs"
	"github.com/NuttayotSukkum/acn/acn-authentication/configs/dbconfigs"
	"github.com/NuttayotSukkum/acn/acn-authentication/internal/handlers/rest"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func CMD(e *echo.Echo) {
	cfg := configs.InitConfig()
	ctx := context.Background()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	client := dbconfigs.InitDB(cfg)
	log.Debugf("Port for execution is: %v", cfg.App.Port)
	log.Infof("Executed: 1")
	rest.UserInitRouter(e, client)
	execute(cfg, ctx, e)

}

func execute(cfg *configs.Configs, ctx context.Context, e *echo.Echo) {
	svPort := fmt.Sprintf(":%v", cfg.App.Port)
	if err := e.Start(svPort); err != nil {
		log.Fatal(ctx, "shutting down the server")
	}

}
