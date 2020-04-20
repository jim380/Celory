package util

import "fmt"

// TitlePrint prints out the titles
func TitlePrint(perform string, target string) {
	switch perform {
	case "stop":
		title := "Stopping node on"
		fmt.Println("\n---------------------------------------------------")
		fmt.Printf("	%s %s machine", title, target)
		fmt.Println("\n---------------------------------------------------")
	case "keyCheck":
		title := "Checking keys on"
		fmt.Println("\n---------------------------------------------------")
		fmt.Printf("	%s %s machine", title, target)
		fmt.Println("\n---------------------------------------------------")
	case "delete":
		title := "Deleting chain data on"
		fmt.Println("\n---------------------------------------------------")
		fmt.Printf("	%s %s machine", title, target)
		fmt.Println("\n---------------------------------------------------")
	case "start":
		title := "Starting node on"
		fmt.Println("\n---------------------------------------------------")
		fmt.Printf("	%s %s machine", title, target)
		fmt.Println("\n---------------------------------------------------")
	}
}
