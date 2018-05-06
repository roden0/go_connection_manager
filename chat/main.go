package main

import (
  "chat/lib"
  "sync"
)

func main() {
  var wg sync.WaitGroup

  connections := lib.GetAllConnections()
  connectionCount := len(connections)

  wg.Add(connectionCount)

  for _, c := range connections {
    go c.Notify(&wg)
  }

  wg.Wait()
}