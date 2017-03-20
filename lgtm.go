package main

import (
  "os"
  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "lgtm"
  app.Usage = "Make the LGTM image"
  app.Action = doAction
  app.Run(os.Args)
}