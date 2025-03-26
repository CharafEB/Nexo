package frameworks

import (
	"fmt"
	"strings"

	//"os"
	"os/exec"

	"github.com/Nexo/cmd/midel"
	"github.com/urfave/cli/v2"
)

var (
	Stract string
	save   string
	//Blue   string
	libs string
	valF []string
	valS []string
)
//Nexo.exe Node --Stract S --Save S project1 express mongoose //build a struct and save it
//Nexo.exe Check --Stract G //build your own struct
//Nexo.exe Node --Stract G ----Save S project2 express mongoose //using the struct that you build
//Nexo.exe Node project2 //call the saved struct
//Nexo.exe Node project1 //call the saved struct
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
			name := ctx.Args().Get(0)
			//In this case the name is the first value so the user have to enter the name of the blueprint to use it
			fmt.Print(ctx.NArg())
			if ctx.NArg() == 1 && Stract == "Free" {
				// fmt.Printf("Blueprint value is : %s The structe value is : %s \n", ctx.Args().Get(0), Stract)
				fmt.Println(fmt.Sprint(searchJSON("name", "test.json", ctx.Args().Get(0))))
				res := searchJSON("name", "test.json", ctx.Args().Get(0))
				libsJSON := fmt.Sprint(res[0]["libs"])
				switch res[0]["struct"] {
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

			//this case to check if the user went to create a saved struct
			if Stract != "S" && Stract != "A" {
				res := searchJSON("struct", "Blue.json", Stract)[0]["structbuild"]
				valStemp := (res.(map[string]interface{})["file"]).(string)
				valFtemp := (res.(map[string]interface{})["folder"]).(string)
				fmt.Printf("The value of the struct is : %s \n", valStemp)
				CreateNewStructure(valStemp, valFtemp)
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
				fmt.Print("Dont save\n")
			case "S":
				if n := ctx.NArg(); n != 0 {
					for i := range n {
						a := ctx.Args().Get(i + 1)

						libs = libs + " " + a
						
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
				fmt.Printf("Name of blueprint %s , Libs are %s the struct is : %s\n", Nwelib.Name, Nwelib.Libs , Nwelib)
				if err := WritJSON(Nwelib, "test.json") ; err != nil {
					return err
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
			&cli.StringFlag{Destination: &Stract, Name: "Stract", Required: true, Usage: "Enter A name of type and save to build a struct"},
		},
		Action: func(ctx *cli.Context) error {
			fmt.Printf("%s\n", Stract)
			if Stract != " " && ctx.NArg() == 0 {
				// mydir, err := os.Getwd()
				// if err != nil {
				// 	fmt.Println(err)
				// }
				if err := BuildStructer(&valS, &valF, "./"); err != nil {
					return err
				}
				StructBuild := midel.StructBuildF{File: strings.Join(valS, ","), Folder: strings.Join(valF, ",")}
				newS := midel.Structrs{StructCall: Stract, StructBuild: StructBuild}
				fmt.Printf("Folder:%s,File:%s", valF, valS)
				//fmt.Println(mydir)

				if err := WritJSON(newS, "Blue.json"); err != nil {
					return err
				}
			} else {
				return fmt.Errorf("Make sure to enter 2 value [Struct call]")
			}

			return nil
		},
	}
}
