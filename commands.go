package main

import (
  "log"
  "os"
  "fmt"
  "github.com/codegangsta/cli"
)

var Commands = []cli.Command{
  commandSecuritygroups,

}

var commandSecuritygroupsflag = []cli.Flag{
  cli.StringFlag{
    Name: "prefix, p",
    Value: "default",
    Usage: "prefix",
  },
  cli.StringFlag{
    Name: "region, r",
    Value: "ap-northeast-1",
    Usage: "region",
  },
}

var commandSecuritygroups = cli.Command{
  Name:  "securityGroups",
  Usage: "",
  Description: `
`,
  Action: doSecuritygroups,
  Flags: commandSecuritygroupsflag,
}

func debug(v ...interface{}) {
  if os.Getenv("DEBUG") != "" {
    log.Println(v...)
  }
}

func assert(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func doSecuritygroups(c *cli.Context) {
  fmt.Print(c.String("r"))
  SecurityGroupsDetail(c.String("p"),c.String("r"))
}


