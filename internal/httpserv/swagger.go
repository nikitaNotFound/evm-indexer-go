package httpserv

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/swaggo/swag"
	"gopkg.in/yaml.v3"
)

// RegisterSwagger registers routes for swagger documentation
func RegisterSwagger(e *echo.Group) {
	log.Debug().Msg("Registering Swagger")

	// Define the file paths
	swaggerFilePath := filepath.Join("spec", "openapi.yaml")

	// Load and register the OpenAPI specification
	if err := loadAndRegisterSpec(swaggerFilePath); err != nil {
		log.Error().Err(err).Msg("Failed to register OpenAPI spec with swag")
		// Continue anyway to prevent server from not starting
	}

	// Serve the Swagger UI with explicit configuration
	// Point to the JSON version since that's what the UI expects
	e.GET("/swagger/*", echoSwagger.EchoWrapHandler(
		echoSwagger.URL("/swagger/doc.json"),
		echoSwagger.DocExpansion("list"),
		echoSwagger.DomID("swagger-ui"),
	))

	// Add a redirect for convenience
	e.GET("/swagger", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	// Serve the OpenAPI spec directly at both root and swagger-ui paths
	e.GET("/swagger.yaml", serveSwaggerFile(swaggerFilePath, "application/x-yaml"))
	e.GET("/swagger/swagger.yaml", serveSwaggerFile(swaggerFilePath, "application/x-yaml"))
	e.GET("/swagger/doc.yaml", serveSwaggerFile(swaggerFilePath, "application/x-yaml"))

	// JSON format alternatives (some tools prefer these)
	e.GET("/swagger.json", serveSwaggerAsJSON(swaggerFilePath))
	e.GET("/swagger/swagger.json", serveSwaggerAsJSON(swaggerFilePath))
	e.GET("/swagger/doc.json", serveSwaggerAsJSON(swaggerFilePath))

	// Add diagnostic endpoint
	e.GET("/swagger-check", func(c echo.Context) error {
		fileExists := "not found"
		if _, err := os.Stat(swaggerFilePath); err == nil {
			fileExists = "exists"
		}

		// Get current path to help with debugging
		currentPath, _ := os.Getwd()

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":           "Swagger API is working",
			"file":             swaggerFilePath,
			"fileExists":       fileExists,
			"currentDirectory": currentPath,
			"fullPath":         filepath.Join(currentPath, swaggerFilePath),
		})
	})
}

// loadAndRegisterSpec loads the OpenAPI specification from file and registers it with swag
func loadAndRegisterSpec(filePath string) error {
	// Read the YAML file
	yamlData, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read OpenAPI spec: %w", err)
	}

	// Parse YAML to a map
	var yamlObj interface{}
	if err := yaml.Unmarshal(yamlData, &yamlObj); err != nil {
		return fmt.Errorf("failed to parse YAML: %w", err)
	}

	// Convert to JSON
	jsonData, err := json.Marshal(yamlObj)
	if err != nil {
		return fmt.Errorf("failed to convert to JSON: %w", err)
	}

	// Register with swag
	info := &swag.Spec{
		InfoInstanceName: "swagger",
		SwaggerTemplate:  string(jsonData),
	}
	swag.Register(info.InstanceName(), info)
	return nil
}

// Helper to serve the swagger file as YAML
func serveSwaggerFile(filePath string, contentType string) echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := os.ReadFile(filePath)
		if err != nil {
			return c.String(http.StatusInternalServerError,
				fmt.Sprintf("Failed to read Swagger spec: %s", err.Error()))
		}
		return c.Blob(http.StatusOK, contentType, data)
	}
}

// Helper to serve the swagger file as JSON (converts from YAML)
func serveSwaggerAsJSON(filePath string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Read the YAML file
		yamlData, err := os.ReadFile(filePath)
		if err != nil {
			return c.String(http.StatusInternalServerError,
				fmt.Sprintf("Failed to read Swagger spec: %s", err.Error()))
		}

		// Parse YAML to a map
		var yamlObj interface{}
		if err := yaml.Unmarshal(yamlData, &yamlObj); err != nil {
			return c.String(http.StatusInternalServerError,
				fmt.Sprintf("Failed to parse YAML: %s", err.Error()))
		}

		// Convert to JSON
		jsonData, err := json.Marshal(yamlObj)
		if err != nil {
			return c.String(http.StatusInternalServerError,
				fmt.Sprintf("Failed to convert to JSON: %s", err.Error()))
		}

		return c.Blob(http.StatusOK, "application/json", jsonData)
	}
}
