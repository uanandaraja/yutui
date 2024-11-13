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
                                                 
A dead simple way to download youtube videos
`
	fmt.Printf("\033[1;36m%s\033[0m\n", banner) 
}

func main() {
	printBanner()

	urlPrompt := promptui.Prompt{
		Label: "🔗 Input YouTube URL",
		Templates: &promptui.PromptTemplates{
			Prompt:  "{{ . }} ❯ ",
			Valid:   "{{ . | green }} ❯ ",
			Invalid: "{{ . | red }} ❯ ",
			Success: "{{ . | bold }} ❯ ",
		},
		Validate: func(input string) error {
			if !strings.Contains(input, "youtube.com") && !strings.Contains(input, "youtu.be") {
				return fmt.Errorf("invalid YouTube URL")
			}
			return nil
		},
	}

	formatPrompt := promptui.Select{
		Label: "📦 Select Format",
		Items: []string{"mp4", "webm", "mp3"},
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "\033[1;36m❯ {{ . | cyan }}\033[0m",
			Inactive: "  {{ . | white }}",
			Selected: "\033[1;36m✔ {{ . | cyan }}\033[0m",
		},
	}

	url, err := urlPrompt.Run()
	if err != nil {
		fmt.Printf("\033[1;31m✗ Error: %v\033[0m\n", err)
		return
	}

	_, format, err := formatPrompt.Run()
	if err != nil {
		fmt.Printf("\033[1;31m✗ Error: %v\033[0m\n", err)
		return
	}

	downloadPath := filepath.Join(os.Getenv("HOME"), "Downloads")
	fmt.Printf("\n\033[1;36m⧗ Downloading to: %s\033[0m\n", downloadPath)

	// Rest of your download logic remains the same
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

	fmt.Println("\n\033[1;36m⧗ Starting download...\033[0m")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("\033[1;31m✗ Download failed: %v\033[0m\n", err)
		return
	}

	fmt.Println("\033[1;32m✔ Download complete!\033[0m")
}
