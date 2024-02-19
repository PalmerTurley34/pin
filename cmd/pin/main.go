package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)

func readAppData() (map[string]string, error) {
	filePath, err := xdg.DataFile("pin/data.json")
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]string), nil
		}
		return nil, err
	}
	directories := make(map[string]string)
	err = json.Unmarshal(data, &directories)
	return directories, err
}

func writeAppData(data map[string]string) error {
	filePath, err := xdg.DataFile("pin/data.json")
	if err != nil {
		return err
	}
	binData, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, binData, 0644)
	return err
}

func main() {
	dirs, err := readAppData()
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		fmt.Println("Subcommands are: new | rm | list | get")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "new":
		// pin new [pin name] [directory]
		if len(os.Args) != 4 {
			fmt.Println("new expects two arguments: [pin name] [directory]")
			os.Exit(1)
		}
		pinName := os.Args[2]
		dir := os.Args[3]
		if _, ok := dirs[pinName]; ok {
			fmt.Printf("%v is already pinned\n", pinName)
			os.Exit(1)
		}
		path, err := filepath.Abs(dir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		dirs[pinName] = path
		err = writeAppData(dirs)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	case "rm":
		// pin rm [pin name]
		if len(os.Args) != 3 {
			fmt.Println("rm expects one argument: [pin name]")
			os.Exit(1)
		}
		pinName := os.Args[2]
		if _, ok := dirs[pinName]; !ok {
			fmt.Printf("%v is not a pinned directory", pinName)
			os.Exit(1)
		}
		delete(dirs, pinName)
		err = writeAppData(dirs)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	case "list":
		// pin list
		for pinName, dir := range dirs {
			fmt.Printf("Name: %s, directory: %s\n", pinName, dir)
		}

	case "get":
		// pin get [pin name]
		if len(os.Args) != 3 {
			fmt.Println("get expects one argument: [pin name]")
			os.Exit(1)
		}
		pinName := os.Args[2]
		dir, ok := dirs[pinName]
		if !ok {
			fmt.Printf("%v is not a pinned directory", pinName)
			os.Exit(1)
		}
		fmt.Println(dir)

	default:
		fmt.Println("Command:", os.Args[1], "not recognized")
	}

}
