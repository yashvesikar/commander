#!/bin/bash

set -e

REPO="yashvesikar/commander"
INSTALL_DIR="/usr/local/bin"
CMD_NAME="commander"

# Detect OS and ARCH
OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$ARCH" in
  x86_64) ARCH="amd64" ;;
  armv6l) ARCH="armv6" ;;
  armv7l) ARCH="armv7" ;;
  aarch64) ARCH="arm64" ;;
  *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

BINARY="${CMD_NAME}-${OS}-${ARCH}"

echo "Detected platform: $OS/$ARCH"
echo "Fetching latest release for $BINARY..."

# Get latest release tag from GitHub API
LATEST_TAG=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"v([^"]+)".*/v\1/')

if [ -z "$LATEST_TAG" ]; then
  echo "Failed to fetch latest release tag."
  exit 1
fi

URL="https://github.com/$REPO/releases/download/$LATEST_TAG/$BINARY"

echo "Downloading $URL..."
curl -L "$URL" -o "$CMD_NAME"

chmod +x "$CMD_NAME"
sudo mv "$CMD_NAME" "$INSTALL_DIR"

echo "✅ Installed $CMD_NAME to $INSTALL_DIR"

