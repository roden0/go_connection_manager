package httpService

import (
  "fmt"
  "time"
)

func SendNotification() {
  time.Sleep(5 * time.Second)
  fmt.Println("Notification sent")
}
