// pkg/motion/motion.go
package motion

import (
        "fmt"
        "time"
)

// DetectMotion simulates detecting motion using a PIR sensor.
// In a real application, this function would interface with the hardware.
// It sends a signal to the motionDetected channel when motion is detected,
// and closes the channel when it's done.
func DetectMotion(motionDetected chan<- bool, done <-chan struct{}) {
        defer close(motionDetected)
        fmt.Println("Motion detection started.")
        // Simulate checking for motion every 2 seconds.
        for {
                select {
                case <-done:
                        fmt.Println("Motion detection stopped.")
                        return
                default:
                        // In a real implementation, you would read from the PIR sensor here.
                        // For simulation purposes, we'll just send a motion detected signal
                        // every 10 seconds.
                        time.Sleep(10 * time.Second)

                        fmt.Println("Simulating motion detection...")
                        // Use a select to avoid blocking on send if the receiver is not ready
                        // or if the done channel is closed.
                        select {
                        case motionDetected <- true:
                        case <-done:
                                fmt.Println("Motion detection stopping.")
                                return
                        }
                }
        }
}
