
// main.go
package main

import (
	"fmt"
	"iot_motion_detection_system/pkg/camera"
	"iot_motion_detection_system/pkg/motion"
	"iot_motion_detection_system/pkg/transfer"
	"sync"
	"time"
)

func main() {
	// Create a channel to signal when motion is detected.
	motionDetected := make(chan bool)
	// Create a channel to signal the motion detection goroutine to stop.
	done := make(chan struct{})

	// Use a WaitGroup to wait for all goroutines to finish.
	var wg sync.WaitGroup

	fmt.Println("Starting IoT Motion Detection Surveillance System...")

	// Start the motion detection goroutine.
	wg.Add(1)
	go func() {
		defer wg.Done()
		motion.DetectMotion(motionDetected, done)
	}()

	// Start the video recording and file transfer goroutine.
	wg.Add(1)
	go func() {
		defer wg.Done()
		for motionEvent := range motionDetected {
			if motionEvent {
				fmt.Println("Motion detected! Starting video recording...")
				// Record video and get the file path.
				filePath, err := camera.RecordVideo()
				if err != nil {
					fmt.Println("Error recording video:", err)
					continue
				}
				fmt.Println("Video recording stopped. File saved at:", filePath)

				// Transfer the recorded video file.
				fmt.Println("Transferring file:", filePath)
				err = transfer.TransferFile(filePath)
				if err != nil {
					fmt.Println("Error transferring file:", err)
					continue
				}
				fmt.Println("File transferred successfully.")
			}
		}
	}()

	// Simulate the application running for a period of time.
	// In a real application, this would be an infinite loop or managed by a service.
	time.Sleep(20 * time.Second)

	// Signal the motion detection goroutine to stop.
	close(done)
	// Wait for all goroutines to finish.
	wg.Wait()

	fmt.Println("Shutting down IoT Motion Detection Surveillance System.")
}
