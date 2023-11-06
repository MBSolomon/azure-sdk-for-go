// generate_swagger.go

package internal

import (
	"log"
	"os"
)

func main() {
	// Apply your directives or any code to modify the Swagger file
	// Assuming you have the Swagger content in a variable, e.g., swaggerContent
	swaggerContent := "modified Swagger content"

	// Save the modified Swagger content to a file
	outputPath := "modified_swagger.json" // Specify the desired output file path
	err := os.WriteFile(outputPath, []byte(swaggerContent), 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Modified Swagger saved to %s", outputPath)
}
