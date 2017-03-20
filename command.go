package main

import (
    "fmt"
    "github.com/urfave/cli"

    "os/exec"
)

func doAction(context *cli.Context) {
    imagePath := "/tmp/lgtm.png"

    capture(imagePath)
    compose(imagePath)
    copyToClipboard(imagePath)
}

func capture(imagePath string) {
    exec.Command("screencapture", "-i", imagePath).Run()
}

func compose(imagePath string) {
    fmt.Println("compose")
}

func copyToClipboard(imagePath string) {
    command := fmt.Sprintf("set the clipboard to POSIX file \"%s\"", imagePath)
    exec.Command("osascript", "-e", command).Run()
}
