package main

// https://dev.to/koddr/build-a-restful-api-on-go-fiber-postgresql-jwt-and-swagger-docs-in-isolated-docker-containers-475j

import (
	"fmt"
	_ "pkg/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title                      Swagger Example API
// @version                    1.0
// @description                This is a sample server celler server.
// @termsOfService             http://swagger.io/terms/
//
// @contact.name               API Support
// @contact.url                http://www.swagger.io/support
// @contact.email              support@swagger.io
//
// @license.name               Apache 2.0
// @license.url                http://www.apache.org/licenses/LICENSE-2.0.html
//
// @host                       localhost:8080
// @BasePath                   /api/v1
//
// @accept                     json
// @produce                    json
// @schemes                    http https
//
// @securityDefinitions.basic  BasicAuth
//
// @securityDefinitions.apikey ApiKeyAuth
// @in                         header
// @name                       Authorization
// @description                Description for what is this security definition being used
func main() {
	app := fiber.New()

	app.Get("/*", swagger.HandlerDefault)


	api := app.Group("/api/v1")

	api.Get("/", handler)

	app.Listen(":3000")
}

// ListAccounts godoc
// @Summary     List accounts
// @Description get accounts
// @Tags        accounts
// @Accept      json
// @Produce     json
// @Param       group_id      path     int    true "Group ID"
// @Param       user_id       path     int    true "User ID"
// @Param       Authorization header   string true "Authentication header"
// @Success     200           {array}  Account
// @Failure     400           {string} string "ok"
// @Failure     404           {string} string "ok"
// @Failure     500           {string} string "ok"
// @Security    ApiKeyAuth
// @Router      /{group_id}/users/{user_id}/accounts [get]
func handler(c *fiber.Ctx) error {
	q := c.Query("q")
	fmt.Printf("q: %v\n", q)
	return c.SendString("Hello, World!")
}

type Account struct {
	ID   int    `json:"id" example:"1"`
	Name string `json:"name" example:"account name"`
}
