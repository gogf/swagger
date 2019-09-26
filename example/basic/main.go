package main

import (
	"github.com/gogf/gf-swagger"
	"github.com/gogf/gf/frame/g"
	"github.com/swaggo/files"

	_ "github.com/gogf/gf-swagger/example/basic/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	s := g.Server()
	s.SetPort(8199)
	//url := gfSwagger.URL("http://localhost:8199/swagger/doc.json")
	s.BindHandler("/swagger/*any", gfSwagger.WrapHandler(swaggerFiles.Handler))
	//s.BindHandler("/swagger/*any", gfSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))
	s.Run()
}
