package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"

	"go-web-skeleton/app/routes"
	"go-web-skeleton/app/shared/database"
	"go-web-skeleton/app/shared/email"
	"go-web-skeleton/app/shared/jsonconfig"
	"go-web-skeleton/app/shared/webengine"
)

// *****************************************************************************
// Application Logic
// *****************************************************************************

func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)
}

func main() {
	// Load the configuration file
	jsonconfig.Load("config/config.json", config)

	// Connect to database
	database.Connect(config.Database)

	// Initialize gin application
	webengine.Router = gin.Default()
	// Call routes
	routes.Test()
	// Run application on localhost:8080
	webengine.Router.Run(":8080")
}

// config the settings variable
var config = &configuration{}

// configuration contains the application settings
type configuration struct {
	Database database.PgInfo `json:"Database"`
	Email    email.SMTPInfo  `json:"Email"`
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
