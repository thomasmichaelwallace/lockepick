package main

import "github.com/gopherjs/gopherjs/js"

func main() {
	js.Module.Get("exports").Set("fromJsonBytes", FromJsonBytes)
	js.Module.Get("exports").Set("toJsonBytes", ToJsonBytes)
}
