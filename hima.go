package main

import (
	"os"
	"runtime"

	"github.com/nsf/termbox-go"

	"hima/cmd"
	"hima/draw"
)


func main() {
	cmd.Run()
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	draw.PollEvent()
}

func _main() int {
	if envvar := os.Getenv("GOMAXPROCS"); envvar == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	return 0
}

