package cmd

import (
	"fmt"
	"os"
	"strings"
	"os/exec"
)

func Run() {
	args := os.Args

	out, _:= exec.Command(strings.Join(args[1:]," ")).Output()
	fmt.Println(string(out))
}
