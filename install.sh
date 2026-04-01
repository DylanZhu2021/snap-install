#!/bin/bash

# Snap Install Service Installation Script

set -e

REPO_URL="https://github.com/DylanZhu2021/snap-install.git"
INSTALL_DIR="/home/admin/snap-install"

echo "Cloning repository from GitHub..."
if [ -d "$INSTALL_DIR" ]; then
    echo "Directory already exists, pulling latest changes..."
    cd "$INSTALL_DIR"
    sudo git pull
else
    sudo git clone "$REPO_URL" "$INSTALL_DIR"
    cd "$INSTALL_DIR"
fi

echo "Building snap-install service..."
sudo go build -o install-snap main.go

echo "Setting permissions..."
sudo chmod +x "$INSTALL_DIR/install-snap"

echo "Installing systemd service..."
sudo cp snap-install.service /etc/systemd/system/

echo "Reloading systemd..."
sudo systemctl daemon-reload

echo "Enabling service..."
sudo systemctl enable snap-install.service

echo "Starting service..."
sudo systemctl start snap-install.service

echo "Checking service status..."
sudo systemctl status snap-install.service

echo ""
echo "Installation complete!"
echo "Service is running on http://127.0.0.1:8347"
echo ""
echo "Useful commands:"
echo "  sudo systemctl status snap-install.service   - Check service status"
echo "  sudo systemctl stop snap-install.service     - Stop service"
echo "  sudo systemctl start snap-install.service    - Start service"
echo "  sudo systemctl restart snap-install.service  - Restart service"
echo "  sudo journalctl -u snap-install.service -f   - View logs"
