package frameworks

import (
	"fmt"
	"os"
	"os/exec"
)
//Structer Stander Place

func CreatStanderNode() {
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

//Structer Advanced Place
func CreatAdvancedNode() {
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

	foldersP := []string{"controllers", "models", " routes", "services", "utils" , "middlewares" , "tests"}

	for _, folderp := range foldersP {

		path := fmt.Sprintf("src/%s", folderp)
		if err := os.Mkdir(path, 0755); err != nil {
			fmt.Print(err)
		}
	}

	exec.Command("powershell", "ni", "README.md")
}