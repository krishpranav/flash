# flash
A golang frontend web framework

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

## Features:
- Fast ⚡
- Easy To Use 🙂
- Cross Platform 💻 

## Installation:
```
$ go get -u https://github.com/krishpranav/flash
```

## Usage:
- import flash into your project:
```golang
import "github.com/krishpranav/flash"
```

- form view:
```golang
package main

import (
	"strconv"

	"github.com/krishpranav/flash"
)

type People struct {
	Names []string
	Name  string
}

func (c *People) Render() (string, error) {
	return `
		<table>
			{{ range $i, $name := .Names }}
				<tr><td>
					{{ $name }}
					<button key="{{ $i }}" @click="Delete">Delete</button>
				</td></tr>
			{{ end }}
		</table>
		Add name: <input type="text" @value="Name">
		<button @click="Add">Add</button>
	`, nil
}

func (c *People) Delete(key string) {
	index, _ := strconv.Atoi(key)
	c.Names = append(c.Names[:index], c.Names[index+1:]...)
}

func (c *People) Add() {
	c.Names = append(c.Names, c.Name)
	c.Name = ""
}

func main() {
	panic(flash.NewServer().Start(func(_ *flash.Connection) flash.Component {
		return &People{
			Names: []string{"NameOne", "NameTwo"},
		}
	}))
}
```

- clock view:
```golang
package main

import (
	"time"

	"github.com/krishpranav/flash"
)

type Clock struct{}

func (c *Clock) Render() (string, error) {
	return `
		The time now is {{ call.Now }}.
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

```

- counter view:
```golang
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

```

## License:
- flash is licensed under [GPL-2.0 License](https://github.com/krishpranav/flash/blob/main/LICENSE)
