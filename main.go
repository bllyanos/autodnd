package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"slices"
	"strings"

	_ "embed"
)

const MODE_NAME = "auto_dnd"

//go:embed template.txt
var template string

func main() {
	template = strings.Trim(template, "\n")

	EnsureModeInstalled()

	cmd := exec.Command("makoctl", "mode", "-t", "auto_dnd")

	var out strings.Builder
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	modes := parseResult(out.String())

	index := slices.Index(modes, MODE_NAME)

	if index >= 0 {
		log.Println("DND is active")
	} else {
		log.Println("DND inactive")
		exec.Command("notify-send", "welcome back").Run()
	}
}

func parseResult(output string) []string {
	modeStrings := strings.Split(strings.Trim(output, "\n"), "\n")
	return modeStrings
}

func EnsureModeInstalled() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("could not open your homedir")
	}

	configPath := path.Join(homedir, ".config", "mako", "config")
	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("could not open the %s", configPath)
	}

	log.Println("checking if auto_dnd is installed in the mako config")

	configFile, err := os.OpenFile(configPath, os.O_RDWR, os.ModeAppend)

	if err != nil {
		log.Fatalf("could not open the config file")
	}

	defer configFile.Close()

	TryAppendToConfig(configFile, template)
}

func TryAppendToConfig(file *os.File, template string) {
	configData, err := io.ReadAll(file)

	if err != nil {
		log.Fatalf("could not read the config file")
	}

	if strings.Contains(string(configData), template) {
		log.Println("auto_dnd already installed in config")
	} else {
		log.Println("auto_dnd is missing from config, appending the config")

		strToWrite := "\n" + template + "\n"
		file.WriteString(strToWrite)

		log.Println("reloading makoctl")
		exec.Command("makoctl", "reload")
		log.Println("makoctl reloaded")
	}
}
