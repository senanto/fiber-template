package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html"
    "github.com/senanto/fiber-template/routes"
)

func main() {
    engine := html.New("./views", ".html")
    app := fiber.New(fiber.Config{
        Views: engine,
    })
    app.Static("/", "./public")
    routes.Setup(app)
    log.Fatal(app.Listen(":80"))
}
