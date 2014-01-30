package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	open(combine(parse()))
}

// parse translates command-line arguments to a destination and parameters
func parse() (string, string) {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "usage: doc (go|py) (pkg|module)")
		os.Exit(1)
	}

	return strings.ToLower(os.Args[1]), strings.Join(os.Args[2:], " ")
}

// combine builds a URL from a destination language and package/module/etc.
func combine(lang, args string) string {
	switch lang {
	case "go", "golang":
		return "http://golang.org/pkg/" + args + "/"
	case "py", "python":
		return "http://docs.python.org/3/library/" + args
	default:
		// Google it instead
		return "https://www.google.com/search?q=" + args
	}
}

// open launches a URL in the user's default Web browser
func open(url string) {
	switch runtime.GOOS {
	case "linux":
		exec.Command("xdg-open", url).Run()
	case "darwin":
		exec.Command("open", url).Run()
	case "windows":
		exec.Command("start", url).Run()
	}
}
