package main

import (
  "os"
  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()
  app.Version = "1.0.0"
  app.Name = "lgtm-capture"
  app.Usage = "Capture screen, then compose image for LGTM, then copy it to the clipboard."
  app.Action = doAction
  app.Run(os.Args)
}
