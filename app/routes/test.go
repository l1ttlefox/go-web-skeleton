package routes

import (
	"go-web-skeleton/app/controllers"
	"go-web-skeleton/app/shared/webengine"
)

func Test() {
	// Group by user route
	users := webengine.Router.Group("/user")
	users.GET("/:id", controllers.Test)
}
