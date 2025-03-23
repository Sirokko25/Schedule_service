package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"sheduler/internal/errorlist"
	"sheduler/internal/handlers"
	"sheduler/internal/storage"
)

func StartServer() error {
	db, err := storage.ConnectionDB()
	if err != nil {
		log.Print(errorlist.ErrorCreateTables)
		return err
	}

	handlers := handlers.Handlers{db}
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		ctx := context.Background()
		defer ticker.Stop()
		for range ticker.C {
			err := db.PingDB(ctx)
			if err != nil {
				log.Printf("База данных не доступна: %v", err)
			}
			db.Actualize()
		}
	}()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/shedule", handlers.AddShedule) //добавить проверку на повторение расписаний
	router.GET("/shedules/", handlers.GetShedules)
	router.GET("/shedule", handlers.GetShedule)
	router.GET("/next_takings", handlers.GetNextTakings)
	log.Printf("The server started at:%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	err = router.Run(fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("fail create server: %v", err)
		return err
	}
	return nil
}
