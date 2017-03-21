package main

import (
  "os"
  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "lgtm-capture"
  app.Usage = "Capture screen, then compose image for LGTM, then copy it to the clipboard."
  app.Action = doAction
  app.Run(os.Args)
}
