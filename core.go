package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"bytes"
)

const (
	GRIT_VER = "1.1.0"
	AUTHOR   = "shad1k1"
	LICENSE  = "GNU GPL v3"
	GRIT_LOGO= `
  ██████╗  ██████╗  ██╗ ████████╗
 ██╔════╝  ██╔══██╗ ██║ ╚══██╔══╝
 ██║  ███╗ ██████╔╝ ██║    ██║
 ██║   ██║ ██╔══██╗ ██║    ██║
 ╚██████╔╝ ██║  ██║ ██║    ██║
  ╚═════╝  ╚═╝  ╚═╝ ╚═╝    ╚═╝
	`
)

func sysInfo(){
	cmd := exec.Command("sh", "-c", "uname -a && uptime")
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
	if err != nil {
		fmt.Printf("\n[!] Command executed with error: %v\n", err)
	}
}

func discMap(){
	fmt.Println("\n--- Partition Map ---\n")
	cmd := exec.Command("lsblk", "-p")
	output, _ := cmd.Output()
	fmt.Println(string(output))
}

func memStat(){
	fmt.Println("\n--- RAM Status ---\n")
	cmd := exec.Command("free", "-h")
	output, _ := cmd.Output()
	fmt.Print(string(output))
}

func serviceStatus(name string) bool {
	cmd := exec.Command("systemctl", "is-active", name )
	err := cmd.Run()
	if err == nil {
		fmt.Printf("[+] Service [%s] is active.\n", name)
		return true
	} else {
		fmt.Println("[!] Service stopped or not found")
		return false
	}
}

func checkLogs(){
	fmt.Println("\n--- Last critical errors ---\n")
	cmd := exec.Command("journalctl", "-p", "3", "-n", "10", "--no-pager")
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
	if err != nil {
		fmt.Printf("\n[!] Command executed with error: %v\n", err)
	}
}

func smartMount() {
	fmt.Println("\n\x1b[1;34m[*] Searching for partitions for mount...\x1b[0m\n")
	cmd := exec.Command("sh", "-c", "lsblk -lnp -o NAME,SIZE,TYPE,FSTYPE,MOUNTPOINT | grep 'part'")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("[!] Error executing lsblk: %v\n", err)
		return
	}

	parts := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(parts) == 0 || parts[0] == "" {
		fmt.Println("[!] Partitions not found.")
		return
	}

	fmt.Println("ID\tDEVICE\t\tSIZE\tFS\tMOUNT")
	for id, line := range parts {
		cols := strings.Fields(line)
		if len(cols) < 2 {
			continue
		}

		dev := cols[0]
		size := cols[1]

		fs := "unknown"
		if len(cols) >= 4 {
			fs = cols[3]
		}

		mount := "-"
		if len(cols) >= 5 {
			mount = cols[4]
		}

		fmt.Printf("[%d]\t%s\t%s\t%s\t%s\n", id, dev, size, fs, mount)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nChoose partition ID for mounting /mnt (or 'q'): ")
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)
	if choiceStr == "q" {
		return
	}

	choice, err := strconv.Atoi(choiceStr)
	if err != nil || choice < 0 || choice >= len(parts) {
		fmt.Println("[!] Invalid ID.")
		return
	}

	selected := strings.Fields(parts[choice])[0]
	fmt.Printf("[*] Mounting %s to /mnt...\n", selected)

	var stderr bytes.Buffer
	mountCmd := exec.Command("mount", selected, "/mnt")
	mountCmd.Stderr = &stderr

	err = mountCmd.Run()
	if err == nil {
		fmt.Println("\033[1;32m[+] Ready! Disk files available at /mnt\033[0m")
	} else {
		fmt.Printf("\033[1;31m[!] Mount error: %s\033[0m\n", stderr.String())
	}
}
