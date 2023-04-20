package main

import (
	"context"
	"fmt"
	"github.com/linqcod/grade-tracker-backend/cmd/api"
	"github.com/linqcod/grade-tracker-backend/pkg/config"
	"github.com/linqcod/grade-tracker-backend/pkg/database"
	"github.com/spf13/viper"
	"log"
)

func init() {
	config.InitConfig()
}

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	port := fmt.Sprintf(":%d", viper.GetInt("app.port"))

	app := api.InitRouter(context.Background(), db)
	app.Run(port)
}
