package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/fgunawan1995/xendit/config"
	marveldal "github.com/fgunawan1995/xendit/dal/api/marvel"
	cachedal "github.com/fgunawan1995/xendit/dal/cache"
	"github.com/fgunawan1995/xendit/handler"
	"github.com/fgunawan1995/xendit/resources"
	"github.com/fgunawan1995/xendit/usecase"
	cron "github.com/robfig/cron/v3"
)

func main() {
	// Init resources
	cfg := config.GetConfig("./config/")
	marvel := resources.InitMarvelClient(cfg)
	cache := resources.InitCache(cfg)

	// Init layers
	marvelDAL := marveldal.New(cfg, marvel)
	cacheDAL := cachedal.New(cache)
	usecaseLayer := usecase.New(cfg, marvelDAL, cacheDAL)
	handlerLayer := handler.New(usecaseLayer)

	initNecessaryData(usecaseLayer)
	initCron(usecaseLayer)
	initHTTP(cfg, handlerLayer)
}

// Init necessary data
func initNecessaryData(usecaseLayer usecase.Usecase) {
	go func() {
		defer func() { // recover go routine in case of panic
			if r := recover(); r != nil {
				log.Printf("error = %v", fmt.Errorf("%v", r))
			}
		}()
		usecaseLayer.SaveCharacters()
	}()
}

// Init cron
func initCron(usecaseLayer usecase.Usecase) {
	cronHandler := cron.New()
	cronHandler.AddFunc("@hourly", func() { usecaseLayer.SaveCharacters() }) // update marvel characters cache hourly (in case of new characters)
}

// Init routes
func initHTTP(cfg *config.Config, handlerLayer *handler.Handler) {
	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods(http.MethodGet)
	r.HandleFunc("/characters/{id}", handlerLayer.GetCharacterByID).Methods(http.MethodGet)
	r.HandleFunc("/characters", handlerLayer.GetAllCharacterIDs).Methods(http.MethodGet)

	// Start server
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type"},
	})
	handler := c.Handler(r)
	port := fmt.Sprintf(":%s", cfg.Server.Port)
	fmt.Printf("Server started at %s\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}
