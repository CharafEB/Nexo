package frameworks

import (
	"fmt"
	"os/exec"

	"github.com/urfave/cli/v2"
)

var (
	Stract string
	save   string
)

// Commande Place
func NodeFunc() *cli.Command {
	return &cli.Command{
		Name: "Node",
		Flags: []cli.Flag{
			&cli.StringFlag{Destination: &Stract, Name: "Stract", Value: "Free", Usage: "--Stract F -free- , --Stract S -Stander- --Stract A -Adevance- "},
			&cli.StringFlag{Destination: &save, Name: "Save", Value: "N", Usage: "--Save S [Name of the blue printe]"},
		},
		Action: func(ctx *cli.Context) error {
			cmd := exec.Command("node", "-v")
			output, err := cmd.CombinedOutput()
			if err != nil {
				exec.Command("winget", "install -e --id OpenJS.NodeJS")
			}
			fmt.Printf("%s version: %s", "Node", string(output))

			switch Stract {
			case "Free":
				fmt.Print("this is Free \n")
			case "S":
				CreatStanderNode()
			case "A":
				CreatAdvancedNode()

			}
			switch save {
			case "N":
				fmt.Print("This is No value \n")
			case "S":
				if n := ctx.NArg(); n != 0 {
					for i:= range n{
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
			}

			return nil
		},
	}
}
