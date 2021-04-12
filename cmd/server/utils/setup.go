package utils

import (
	"gin_psql_microservice_template/cmd/server/docs"
	"gin_psql_microservice_template/configs"
)

// SetupDocuments
func SetupDocuments() {
	host := configs.Envs["SERVER_HOST"]
	docs.SwaggerInfo.Schemes = []string{"https", "http"}
	if configs.Envs["SERVER_PORT"] != "80" {
		host += ":" + configs.Envs["SERVER_PORT"]
		docs.SwaggerInfo.Schemes = []string{"http"}
	}
	docs.SwaggerInfo.Title = "API Document"
	docs.SwaggerInfo.Description = "The REST API Document"
	docs.SwaggerInfo.Host = host
	docs.SwaggerInfo.Version = configs.Envs["SERVER_VERSION"]
	docs.SwaggerInfo.BasePath = ""
}
