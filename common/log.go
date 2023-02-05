package common

import (
	"fmt"
	"log"
)

func LogError(message string) {
	fmt.Println("")
	log.Printf("\n[!!] %s", message)
}

func LogInfo(message string) {
	fmt.Println("")
	log.Printf("\n[*] %s", message)
}

func LogSuccess(message string) {
	fmt.Println("")
	log.Printf("\n[*] %s", message)
}

func LogWarning(message string) {
	fmt.Println("")
	log.Printf("\n[!] %s", message)
}
