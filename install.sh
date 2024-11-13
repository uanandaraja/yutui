#!/bin/bash

# Check if Go is installed
if ! command -v go &>/dev/null; then
  echo "Go is not installed. Please install Go first."
  exit 1
fi

# Check if yt-dlp is installed
if ! command -v yt-dlp &>/dev/null; then
  echo "yt-dlp is not installed. Installing..."
  if [[ "$OSTYPE" == "darwin"* ]]; then
    brew install yt-dlp
  elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
    sudo apt update && sudo apt install yt-dlp -y
  else
    echo "Unsupported OS. Please install yt-dlp manually."
    exit 1
  fi
fi

# Build the program
echo "Building yutui..."
go build -o ytdl-wrapper

# Install to system or user bin
if [ -w "/usr/local/bin" ]; then
  echo "Installing to /usr/local/bin..."
  mv ytdl-wrapper /usr/local/bin/
elif [ -d "$HOME/bin" ]; then
  echo "Installing to $HOME/bin..."
  mv ytdl-wrapper "$HOME/bin/"
else
  echo "Creating $HOME/bin..."
  mkdir -p "$HOME/bin"
  mv ytdl-wrapper "$HOME/bin/"
  # Add to PATH if not already there
  if [[ ":$PATH:" != *":$HOME/bin:"* ]]; then
    echo 'export PATH="$HOME/bin:$PATH"' >>"$HOME/.bashrc"
    echo 'export PATH="$HOME/bin:$PATH"' >>"$HOME/.zshrc"
    echo "Added $HOME/bin to PATH. Please restart your terminal."
  fi
fi

echo "Installation complete!"
