package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
)

func printBanner() {
	banner := `
                __        _ 
   __  ____  __/ /___  __(_)
  / / / / / / / __/ / / / / 
 / /_/ / /_/ / /_/ /_/ / /  
 \__, /\__,_/\__/\__,_/_/   
/____/                      
`
	fmt.Println(banner)
}

func main() {
  printBanner()

	// Get URL
	urlPrompt := promptui.Prompt{
		Label: "Enter YouTube URL",
		Validate: func(input string) error {
			if !strings.Contains(input, "youtube.com") && !strings.Contains(input, "youtu.be") {
				return fmt.Errorf("invalid YouTube URL")
			}
			return nil
		},
	}

	url, err := urlPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// Format selection
	formats := []string{"mp4", "webm", "mp3"}
	formatPrompt := promptui.Select{
		Label: "Select format",
		Items: formats,
	}

	_, format, err := formatPrompt.Run()
	if err != nil {
		fmt.Printf("Format selection failed %v\n", err)
		return
	}

	downloadPath := filepath.Join(os.Getenv("HOME"), "Downloads")

	// Build yt-dlp command with headers
	var cmd *exec.Cmd
	switch format {
	case "mp3":
		cmd = exec.Command("yt-dlp",
			"--user-agent", "Mozilla/5.0",
			"--add-header", "Accept:text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
			"-x", "--audio-format", "mp3",
			"-o", filepath.Join(downloadPath, "%(title)s.%(ext)s"),
			url)
	default:
		cmd = exec.Command("yt-dlp",
			"--user-agent", "Mozilla/5.0",
			"--add-header", "Accept:text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
			"-f", fmt.Sprintf("bestvideo[ext=%s]+bestaudio[ext=m4a]/best[ext=%s]", format, format),
			"-o", filepath.Join(downloadPath, "%(title)s.%(ext)s"),
			url)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Downloading...")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Download failed: %v\n", err)
		return
	}

	fmt.Println("Download complete!")
}
