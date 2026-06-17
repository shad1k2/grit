package main

import(
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func showBanner() {
	fmt.Println("\033[1;32m")
	fmt.Println("  ██████╗  ██████╗  ██╗ ████████╗")
	fmt.Println(" ██╔════╝  ██╔══██╗ ██║ ╚══██╔══╝")
	fmt.Println(" ██║  ███╗ ██████╝ ██║    ██║")
	fmt.Println(" ██║   ██║ ██╔══██╗ ██║    ██║")
	fmt.Println(" ╚██████╔╝ ██║  ██║ ██║    ██║")
	fmt.Println("  ╚═════╝  ╚═╝  ╚═╝ ═╝    ╚═╝")
	fmt.Println("\033[0m-------------------------------")
}

func main(){
	showBanner()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("grit> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}
		
		parts := strings.Fields(input)
		command := parts[0]

		var arg string
		if len(parts) > 1 {
			arg = parts[1]
		}

		switch command {
		case "sys", "info":
			sysInfo()
			
		case "disks":
			discMap()
			
		case "mem":
			 memStat()
			
		case "service":
			if arg != "" {
				 serviceStatus(arg)
			} else {
				fmt.Println("[!] Enter service name: service <name>")
			}
			
		case "logs":
			 checkLogs()
			 			
		case "mount":
			smartMount()
			
		case "help":
			fmt.Println("Availible commands: info, disks, mount, mem, service <name>, logs, clear, exit")
			
		case "clear":
			 exec.Command("clear").Run()
			
		case "ver", "version":
			fmt.Printf("GRIT - Go Recovery and Inspection Tool v%s, Made by %s, licensed under %s \n", GRIT_VER, AUTHOR, LICENSE)
			
		case "exit":
			fmt.Println("Shutting down GRIT. Safe travels.")
			os.Exit(0)
			
		default:
			fmt.Printf("[?] Unknown command: %s. Print 'help'.\n", command)
		}
	} 
}
