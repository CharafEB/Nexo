package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

var (
	Stract string
	save   string
)

func main() {
	app := cli.NewApp()
	app.Name = "This is the name of the application "
	app.Description = "This is the description of the application"
	app.Version = "1.0.0"
	app.Commands = []*cli.Command{
		NodeFunc(),
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func NodeFunc() *cli.Command {
	return &cli.Command{
		Name: "Node",
		Flags: []cli.Flag{
			&cli.StringFlag{Destination: &Stract, Name: "Stract", Value: "Free", Usage: "--Stract F -free- , --Stract S -Stander- --Stract A -Adevance- "},
			&cli.StringFlag{Destination: &save, Name: "Save", Value: "No", Usage: "--Save [Name of the blue printe]"},
		},
		Action: func(ctx *cli.Context) error {
			switch Stract {
			case "Free":
				fmt.Print("this is Free \n")
			case "S":
				CreatStanderNode()
			case "A":
				fmt.Print("this is Advance \n")

			}
			switch save {
			case "No":
				fmt.Print("This is No value \n")
			}
			if n := ctx.NArg(); n != 0 {
				for i := 0; i < n; i++ {
					a := ctx.Args().Get(i)
					cmd := exec.Command("powershell", "npm install ", a)
					out, err := cmd.Output()
					if err != nil {
						return fmt.Errorf("err")
					}
					fmt.Printf("Output is: %s\n", out)
					fmt.Printf("the value is:%s\n", a)
				}
			} else {
				return fmt.Errorf("you have to select libs to start")
			}
			return nil
		},
	}
}

func CreatStanderNode() {
	//controllers,models,routes,services,utils")

	folders := []string{"src", "config", " public", "views"}

	for _, folder := range folders {
		if err := os.Mkdir(folder, 0755); err != nil {
			fmt.Print(err)
		}
	}

	files := []string{".env", ".gitignore"}

	for _, file := range files {
		if _, err := os.Create(file); err != nil {
			fmt.Print(err)
		}
	}

	foldersP := []string{"controllers", "models", " routes", "services", "utils"}

	for _, folderp := range foldersP {

		path := fmt.Sprintf("src/%s", folderp)
		if err := os.Mkdir(path, 0755); err != nil {
			fmt.Print(err)
		}
	}

	exec.Command("powershell", "ni", "README.md")
}
