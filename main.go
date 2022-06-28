package main

import (
	"log"

	"github.com/godbus/dbus/v5"
)

func main() {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	screencast := conn.BusObject().Call("org.freedesktop.portal.ScreenCast.CreateSession", 0, make(map[string]string))
	if screencast.Err != nil {
		log.Panicln(screencast.Err)
	}

	log.Println(screencast.Body...)

	log.Println("Hello, world!")
}
