package frameworks

import (
	"fmt"
	//"os"
	"os/exec"

	"github.com/Nexo/cmd/midel"
	"github.com/urfave/cli/v2"
)

var (
	Stract string
	save   string
	Blue   string
	libs   string
	valF   []string
	valS   []string
)

// Commande Place
func NodeFunc() *cli.Command {
	return &cli.Command{
		Name: "Node",
		Flags: []cli.Flag{
			&cli.StringFlag{Destination: &Stract, Name: "Stract", Value: "Free", Usage: "--Stract F -free- , --Stract S -Stander- --Stract A -Adevance- "},
			&cli.StringFlag{Destination: &save, Name: "Save", Value: "N", Usage: "--Save S [Name of the blue printe]"},
			&cli.StringFlag{Destination: &Blue, Name: "Blue", Value: "N", Usage: "--Blue [Name of the blue printe]"},
		},
		Action: func(ctx *cli.Context) error {
			cmd := exec.Command("node", "-v")
			output, err := cmd.CombinedOutput()
			if err != nil {
				exec.Command("winget", "install -e --id OpenJS.NodeJS")
			}
			fmt.Printf("%s version: %s", "Node", string(output))
			name := ctx.Args().Get(0)
			if err := CheckJSON(name); err != nil {
				return err
			}
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
					for i := range n {
						a := ctx.Args().Get(i + 1)

						libs = libs + " " + a
						fmt.Printf("the value is:%s\n", a)
					}
				} else {
					return fmt.Errorf("you have to select libs to start")
				}
				cmd := exec.Command("powershell", "npm install ", libs)
				out, err := cmd.Output()
				if err != nil {
					return fmt.Errorf("err")
				}
				fmt.Print(string(out))
				fmt.Printf("Name of blueprint %s , Libs are %s \n", name, libs)
				Nwelib := midel.Blueprints{Name: name, Libs: libs, StructType: Stract}
				WritJSON(Nwelib, "test.json")
			}
			if ctx.String("Blue") != " " {
				if ctx.NArg() == 1 {
					fmt.Printf("Blueprint value is : %s The structe value is : %s \n", ctx.Args().Get(0), Stract)
					fmt.Println(searchJSON(ctx.Args().Get(0))[0]["libs"])
					libsJSON := fmt.Sprint(searchJSON(ctx.Args().Get(0))[0]["libs"])
					switch searchJSON(ctx.Args().Get(0))[0]["struct"] {
					case "S":
						CreatStanderNode()
					case "A":
						CreatAdvancedNode()
					}
					cmd := exec.Command("powershell", "npm install ", libsJSON)
					out, err := cmd.Output()
					if err != nil {
						return fmt.Errorf("err")
					}
					fmt.Print(string(out))
				}
			}

			return nil
		},
	}
}

func CheckStructe() *cli.Command {
	return &cli.Command{
		Name: "Check",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "Stract", Value: "Free", Usage: "Enter A name of type and save to build a struct"},
		},
		Action: func(ctx *cli.Context) error {
			if ctx.String("Stract") != " " && ctx.NArg() ==2{
				// mydir, err := os.Getwd()
				// if err != nil {
				// 	fmt.Println(err)
				// }
				if err :=BuildStructer(&valS , &valF , "./"); err != nil {
					return err
				}
				StructBuild := midel.StructBuild{File: fmt.Sprint(valS) , Folder: fmt.Sprint(valF)}
				newS := midel.Blueprints{Name: ctx.Args().Get(0),StructType: ctx.Args().Get(1),StructBuild:StructBuild}
				fmt.Printf("Folder:%s,File:%s",valF,valS)
				//fmt.Println(mydir)
				
				if err := WritJSON(newS , "Blue.json"); err != nil {
					return err
				}
			}else{
				return fmt.Errorf("Make sure to enter 2 value [Name] [Struct refrance]")
			}

			return nil
		},
	}
}
