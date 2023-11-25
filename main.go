package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	var target string
	fmt.Print("Enter the IP address to ping: ")
	fmt.Scanln(&target)

	file, err := os.Create("ping_results.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer file.Close()

	for {
		result, err := ping(target)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			now := time.Now()
			log := fmt.Sprintf("Time: %s\n%s\n", now.Format("2006-01-02 15:04:05"), result)
			fmt.Print(log)
			_, err := file.WriteString(log)
			if err != nil {
				fmt.Println("Error writing to file:", err)
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func ping(target string) (string, error) {
	cmd := exec.Command("cmd", "/c", "chcp 65001 && ping "+target)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
