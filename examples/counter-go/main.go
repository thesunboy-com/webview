package main

import (
	"fmt"
	"github.com/zserge/webview"
	"time"
)

// Counter is a simple example of automatic Go-to-JS data binding
type Counter struct {
	Value int `json:"value"`
}

// Add increases the value of a counter by n
func (c *Counter) Add(n int) {
	fmt.Println(n)
	c.Value = c.Value + int(n)
}

// Reset sets the value of a counter back to zero
func (c *Counter) Reset() {
	c.Value = 0
}

func main() {
	w := webview.New(webview.Settings{
		Title: "Click counter: " + uiFrameworkName,
	})
	defer w.Exit()

	w.Dispatch(func() {
		// Inject controller
		counter := new(Counter)
		go func() {
			for {
				time.Sleep(time.Second)
				fmt.Println("累加:1")
				counter.Add(1)
			}
		}()
		w.Bind("counter", counter)

		// Inject CSS
		w.InjectCSS(string(MustAsset("js/styles.css")))

		// Inject web UI framework and app UI code
		loadUIFramework(w)
	})
	w.Run()
}
