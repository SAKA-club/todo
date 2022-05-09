package main

import (
	"github.com/go-openapi/runtime"
	"github.com/rs/zerolog/log"

	"github.com/go-openapi/loads"

	"github.com/saka-club/todo/gen/restapi"
	"github.com/saka-club/todo/gen/restapi/operations"
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

}
