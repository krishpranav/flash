package main

import "github.com/krishpranav/flash"

type Counter struct {
	Number int
}

func (c *Counter) Render() (string, error) {
	return `
		Counter: {{ .Number }}
		<button @click="AddOne">+</button>
	`, nil
}

func (c *Counter) AddOne() {
	c.Number++
}

func main() {
	panic(flash.NewServer().Start(func(_ *flash.Connection) flash.Component {
		return &Counter{}
	}))
}
