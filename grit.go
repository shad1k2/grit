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
	fmt.Println(" ██║  ███╗ ██████╝  ██║    ██║")
	fmt.Println(" ██║   ██║ ██╔══██╗ ██║    ██║")
	fmt.Println(" ╚██████╔╝ ██║  ██║ ██║    ██║")
	fmt.Println("  ╚═════╝  ╚═╝  ╚═╝  ═╝    ╚═╝")
	fmt.Println("\033[0m-------------------------------")
}

func main(){

	if len(os.Args) > 1 {
		switch os.Args[1] {
			case "-v", "--version", "version":
				fmt.Printf("GRIT v%s by %s (%s)\n", GRIT_VER, AUTHOR, LICENSE)
				os.Exit(0)
			case "-h", "--help", "help":
				fmt.Println("GRIT - Go Recovery and Inspection Tool")
				fmt.Println("\nUsage:")
				fmt.Println("  grit            Start interactive recovery shell")
				fmt.Println("  grit -h, --help Show this help message")
				fmt.Println("  grit -v, --ver  Show version info")
				os.Exit(0)
		}
	}

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
			 fmt.Print("\033[H\033[2J")
			
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
