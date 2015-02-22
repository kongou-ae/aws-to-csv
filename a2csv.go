package main

import (
  "os"
  "github.com/codegangsta/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "a2csv"
  app.Version = Version
  app.Usage = ""
  app.Author = "kongou-ae"
  app.Email = "y.matsumoto.ae@gmail.com"
  app.Commands = Commands

  app.Run(os.Args)
}
