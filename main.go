package main

import (
	"library/connection"
	"library/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := connection.Connection()

	eng := &routers.Routes{
		Db: db,
		R:  r,
	}

	eng.Routers()

	r.Run()
}
