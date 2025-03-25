package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	frame "github.com/Nexo/cmd/frameworks"
	"github.com/urfave/cli/v2"
)



func main() {
	cmd := exec.Command("go", "version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		exec.Command("winget" , "install --id=GoLang.Go  -e")
	}
	fmt.Printf("%s version: %s", "go", string(output))

	app := cli.NewApp()
	app.Name = "This is the name of the application "
	app.Description = "This is the description of the application"
	app.Version = "1.0.0"
	app.Commands = []*cli.Command{
		frame.NodeFunc(),
		frame.CheckStructe(),
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
