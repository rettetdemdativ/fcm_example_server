package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

const port = ":8080"

func main() {
	// Prepare the FCM connection by setting the API key used for authorization
	// during requests
	prepareFCM()

	// Set up gin router
	r := gin.Default()
	r.POST("/fcm/register", registerNewFCMID)

	// Run the loop to wait for user input
	go waitForInput()

	// Run the HTTP server
	if err := r.Run(port); err != nil {
		panic(err)
	}
}

// Loops, waiting for user input. Messages entered will be sent to the list of
// clients via Firebase Cloud Messaging.
func waitForInput() {
	reader := bufio.NewReader(os.Stdin)

	var message string

	for message != "/q" {
		fmt.Printf("Enter a message: ")
		message, _ = reader.ReadString('\n')

		sendFCMMessageToAll(message)
	}
}
