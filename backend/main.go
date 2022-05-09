package main

import (
	"github.com/SAKA-club/todo/backend/gen/restapi"
	"github.com/SAKA-club/todo/backend/gen/restapi/operations"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/rs/zerolog/log"
)

func main() {
	//swagger
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load swagger spec")
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

	server.Port = 8080
	server.Host = "0.0.0.0"

	if err := server.Serve(); err != nil {
		log.Panic().Err(err).Msg("server error")
	}

}
