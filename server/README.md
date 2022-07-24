# Server

This package contains files related to the [Fiber](https://gofiber.io/) web framework.

The handlers serve the incoming HTTP requests and pass them to the appropriate business logic.

The business logic is in a separate package called `pkg`.

Some comments are on top of each handler function so that [swag](https://github.com/swaggo/swag) can use to generate RESTful API documentation with Swagger 2.0 for Go automatically.
