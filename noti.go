package main

import (
	"fmt"
	"time"

	"github.com/go-toast/toast"
)

func main() {
	// Function to display notification
	notify := func() {
		notification := toast.Notification{
			AppID:    "Reminder",
			Title:    "Reminder",
			Message:  "It's time to drink water!",
			Duration: toast.Long,
		}
		err := notification.Push()
		if err != nil {
			fmt.Println("Error displaying notification:", err)
		}
	}

	// Schedule reminder every 20 minutes
	for {
		// Trigger notification
		notify()

		// Wait for 20 minutes
		time.Sleep(20 * time.Minute)
	}
}
