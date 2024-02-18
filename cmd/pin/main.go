package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func readAppData() (map[string]string, error) {
	// filePath, err := xdg.DataFile("pin/data.json")
	// if err != nil {
	// 	return nil, err
	// }
	filePath := "data.json"
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
	// filePath, err := xdg.DataFile("pin/data.json")
	// if err != nil {
	// 	return err
	// }
	filePath := "data.json"
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

	newCmd := flag.NewFlagSet("new", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("Subcommands are: new | go | rm | ls | show")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "new":
		// pin new [pin name] [directory]
		newCmd.Parse(os.Args[2:])
		if len(newCmd.Args()) != 2 {
			fmt.Println("new expects two arguments: [pin name] [directory]")
			os.Exit(1)
		}
		name := newCmd.Arg(0)
		dir := newCmd.Arg(1)
		if _, ok := dirs[name]; ok {
			fmt.Printf("%v is already pinned\n", name)
			os.Exit(1)
		}
		path, err := filepath.Abs(dir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		dirs[name] = path
		err = writeAppData(dirs)
		if err != nil {
			log.Fatal(err)
		}
	case "go":
		fmt.Println("Running go...")
	case "rm":
		fmt.Println("Running rm...")
	case "ls":
		fmt.Println("Running ls...")
	case "show":
		fmt.Println("Running show...")
	default:
		fmt.Println("Command:", os.Args[1], "not recognized")
	}

}
