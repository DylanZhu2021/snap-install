#!/bin/bash

# Snap Install Service Installation Script

set -e

BINARY_URL="http://10.0.5.217/snap-install/install-snap"
INSTALL_DIR="/home/admin/snap-install"
SERVICE_FILE_URL="http://10.0.5.217/snap-install/snap-install.service"

echo "Creating installation directory..."
sudo mkdir -p "$INSTALL_DIR"

echo "Downloading install-snap binary..."
sudo curl -L -o "$INSTALL_DIR/install-snap" "$BINARY_URL"

echo "Setting permissions..."
sudo chmod +x "$INSTALL_DIR/install-snap"

echo "Downloading systemd service file..."
sudo curl -L -o "$INSTALL_DIR/snap-install.service" "$SERVICE_FILE_URL"

echo "Installing systemd service..."
sudo cp "$INSTALL_DIR/snap-install.service" /etc/systemd/system/

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
