#!/bin/bash

# Snap Install Service Installation Script

set -e

echo "Building snap-install service..."
go build -o snap-install main.go


echo "Building snap-install service..."
go build -o snap-install main.go

echo "Creating installation directory..."
sudo mkdir -p /home/admin/snap-install

echo "Copying binary..."
sudo cp install-snap /home/admin/snap-install/

echo "Setting permissions..."
sudo chmod +x /home/admin/snap-install/install-snap

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
