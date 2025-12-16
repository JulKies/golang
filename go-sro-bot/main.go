package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {

	//logger aufsetzten
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	log.Println("Starting GO Bot")

	programPath := "C:\\Users\\kiese\\Desktop\\Sammelkiste\\golang\\sro\\silkroad.exe"

	cmd := exec.Command(programPath)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Println("Starte Programm:", programPath)

	// Programm ausführen
	if err := cmd.Run(); err != nil {
		log.Println("Fehler beim Ausführen:", err)
		os.Exit(1)
	}

	log.Println("Programm erfolgreich beendet")
}
