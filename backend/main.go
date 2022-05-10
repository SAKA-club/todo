package main

import (
	"github.com/SAKA-club/todo/backend/gen/restapi"
	"github.com/SAKA-club/todo/backend/gen/restapi/operations"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

var host = "None"

func main() {
	// Get configs
	cfg := LoadConfig()
	initLog(cfg)

	//swagger
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load swagger spec")
	}

	//local debug
	if cfg.LocalDebug {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		log.Error().Msg("debug mode")
	}

	api := operations.NewTodoAPI(swaggerSpec)

	api.ApplicationJSONProducer = runtime.JSONProducer()
	api.ApplicationJSONConsumer = runtime.JSONConsumer()

	server := restapi.NewServer(api)
	defer func(server *restapi.Server) {
		err := server.Shutdown()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to shutdown server")
		}
	}(server)

	server.Port = cfg.Port
	server.Host = cfg.Host

	if err := server.Serve(); err != nil {
		log.Panic().Err(err).Msg("server error")
	}

}

func initLog(cfg *Config) {
	var err error

	host, err = os.Hostname()
	if err != nil {
		log.Panic().Err(err).Msg("unable to get hostname")
	}

	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Log().Str("Host", host).Msg("Service Startup")
}
