package main

import (
    "fmt"
    "github.com/urfave/cli"
)

func doAction(context *cli.Context) {
    imagePath := "/tmp/lgtm.png"

    capture(imagePath)
    compose(imagePath)
    copyToClipboard(imagePath)
}

func capture(imagePath string) {
    fmt.Println("capture")
}

func compose(imagePath string) {
    fmt.Println("compose")
}

func copyToClipboard(imagePath string) {
    fmt.Println("copyToClipboard")
}
