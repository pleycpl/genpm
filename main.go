package main

import (
	"fmt"
	"os"

	"github.com/pleycpl/genpm/genpm"
)

// https://stackoverflow.com/questions/35343707/linux-apt-get-command-not-found-how-to-install-a-package-in-arch-linux , Thnaks for https://stackoverflow.com/users/523100/czechnology
func main() {
	fmt.Println("[+] Runned Main function")
	home := os.Getenv("HOME")
	path := home + "/.genpmrc"
	GenpmInstance := genpm.NewGenpm(path)

	args := os.Args[1:]
	if len(args) > 0 {
		tool := GenpmInstance.Tool
		switch args[0] {
		case "install":
			tool.Install(args[1])
		case "remove":
			tool.Remove(args[1])
		case "search":
			tool.Search(args[1])
		case "upgrade":
			tool.Upgrade()
		case "info":
			tool.Info(args[1])
		case "setup":
			tool.Setup(args[1])
		case "version", "--version":
			fmt.Println("Genpm Version 1.0.0")
		default:
			fmt.Println("Not founded!")
		}
	}
}
