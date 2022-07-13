package main

import (
	"time"

	"github.com/krishpranav/flash"
)

type Clock struct{}

func (c *Clock) Render() (string, error) {
	return `
		The time now is {{ call .Now }}.
	`, nil
}

func (c *Clock) Now() string {
	return time.Now().Format(time.RFC1123)
}

func main() {
	panic(flash.NewServer().Start(func(conn *flash.Connection) flash.Component {
		go func() {
			for range time.NewTicker(time.Second).C {
				conn.Update()
			}
		}()

		return &Clock{}
	}))
}
