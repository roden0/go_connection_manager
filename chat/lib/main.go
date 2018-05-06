package lib

import (
  "chat/httpService"
  "sync"
  "time"
)

type guestConnection struct {
  ip       string
  userName string
  isAdmin  bool
}

type visitorConnection struct {
  ip       string
  connHour int
}

type notifier interface {
  Notify(*sync.WaitGroup)
}

func (g guestConnection) Notify(wg *sync.WaitGroup) {
  httpService.SendNotification()
  wg.Done()
}

func (v visitorConnection) Notify(wg *sync.WaitGroup) {
  httpService.SendNotification()
  wg.Done()
}

func GetAllConnections() []notifier {
  gConn := &guestConnection{ip: "192.168.0.10", userName: "Darth Vader"}
  vConn := &visitorConnection{ip: "192.168.0.11", connHour: time.Now().Hour()}

  return []notifier{gConn, vConn}
}
