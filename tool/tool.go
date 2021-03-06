package tool

import (
	"fmt"
	"os"
	"os/exec"
)

// contain tool's install, remove, search commands
type Tool struct {
	ToolName       string `json:"name"`
	InstallCommand string `json:"install"`
	RemoveCommand  string `json:"remove"`
	SearchCommand  string `json:"search"`
	UpgradeCommand string `json:"upgrade"`
	InfoCommand    string `json:"info"`
	SetupCommand   string `json:"setup"`
}

// run script in bash
func (this *Tool) runScript(script string) {
	fmt.Println("[+] Runned runScript method in Tool Struct")
	cmd := exec.Command("/bin/bash", "-c", script)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

// install package with Tool
func (this *Tool) Install(packagename string) {
	fmt.Println("[+] Runned Install method in Tool Struct")
	script := this.InstallCommand + " " + packagename
	this.runScript(script)
}

// remove package with Tool
func (this *Tool) Remove(packagename string) {
	fmt.Println("[+] Runned Remove method in Tool Struct")
	script := this.RemoveCommand + " " + packagename
	this.runScript(script)
}

// search package with Tool
func (this *Tool) Search(searchstring string) {
	fmt.Println("[+] Runned Search method in Tool Struct")
	script := this.SearchCommand + " " + searchstring
	this.runScript(script)
}

// upgrade packages with Tool
func (this *Tool) Upgrade() {
	fmt.Println("[+] Runned Upgrade method in Tool Struct")
	this.runScript(this.UpgradeCommand)
}

// info package with Tool
func (this *Tool) Info(packagename string) {
	fmt.Println("[+] Runned Info method in Tool Struct")
	script := this.InfoCommand + " " + packagename
	this.runScript(script)
}

func (this *Tool) Setup(path string) {
	fmt.Println("[+] Runned Setup method in Tool Struct")
	script := this.SetupCommand + " " + path
	this.runScript(script)
}
