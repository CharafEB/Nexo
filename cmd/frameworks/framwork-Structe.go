package frameworks

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

// Structer Advanced Place
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

	foldersP := []string{"controllers", "models", " routes", "services", "utils", "middlewares", "tests"}

	for _, folderp := range foldersP {

		path := fmt.Sprintf("src/%s", folderp)
		if err := os.Mkdir(path, 0755); err != nil {
			fmt.Print(err)
		}
	}

	exec.Command("powershell", "ni", "README.md")
}

// Structer newstruct Place
func BuildStructer(valS, valF *[]string, path string) error {
	err := filepath.WalkDir(path, func(fullPath string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() {
			*valF = append(*valF, fullPath)
		} else {
			*valS = append(*valS, fullPath)
		}

		return err
	})
	return err
}

// Structer newstruct Place
func CreateNewStructure(valS, valF string) error {

	if err := os.MkdirAll(valF, 0755); err != nil {
		return fmt.Errorf("failed to create folder: %v", err)
	}


	files := strings.Split(valS, ",")


	for _, file := range files {
		filePath := fmt.Sprintf("%s/%s", valF, strings.TrimSpace(file))
		if _, err := os.Create(filePath); err != nil {
			return fmt.Errorf("failed to create file %s: %v", filePath, err)
		}
	}


	cmd := exec.Command("powershell", "ni", fmt.Sprintf("%s/README.md", valF))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create README.md: %v", err)
	}

	return nil
}
