package main

import (
    "crypto/rand"
    "fmt"
    "os"
    "os/exec"

    "github.com/urfave/cli"
)

func doAction(context *cli.Context) {
    imagePath := makeImagePath()
    defer os.Remove(imagePath)

    capture(imagePath)
    compose(imagePath)
    copyToClipboard(imagePath)
}

func capture(imagePath string) {
    exec.Command("screencapture", "-i", imagePath).Run()
}

func copyToClipboard(imagePath string) {
    // NOTE: PNGが指定できなかったためJPEGとしている
    command := fmt.Sprintf("set the clipboard to read(POSIX file \"%s\") as JPEG picture", imagePath)
    exec.Command("osascript", "-e", command).Run()
}

func makeImagePath() string {
    name := make([]byte, 10)
    rand.Read(name)
    return fmt.Sprintf("/tmp/%X.png", name)
}
