package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	hostnameStr := "IMAGE_GET_ENDPOINT"

	// Get the value of the environment variable
	hostname, exists := os.LookupEnv(hostnameStr)

	if !exists {
		fmt.Printf("Environment variable %s not found\n", hostnameStr)
	}

	// Replace with the actual URL and ID
	url := "http://" + hostname + "/api/v1/screenshot/window_title/pilot-1000"

	// Send an HTTP GET request to the URL
	fmt.Println("URL: ", url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error sending GET request: %s\n", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("HTTP request failed with status: %s\n", response.Status)
		return
	}

	// Read the image data from the response body
	imageBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading image data: %s\n", err)
		return
	}

	if len(imageBytes) < 1024 {
		fmt.Printf("Image too small (%d bytes)\n", len(imageBytes))
		return
	}

	// Print the size of the image
	fmt.Printf("Image size: %d bytes\n", len(imageBytes))
}
