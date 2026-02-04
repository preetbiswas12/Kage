#!/bin/bash
set -e

# Build the binary for Linux ARM64 (Termux)
echo "Building kage for Termux..."
GOOS=linux GOARCH=arm64 go build -o kage

# Create debian package structure
mkdir -p debian/DEBIAN
mkdir -p debian/data/data/com.termux/files/usr/bin

# Copy binary
cp kage debian/data/data/com.termux/files/usr/bin/
chmod +x debian/data/data/com.termux/files/usr/bin/kage

# Copy control file
cp debian/control debian/DEBIAN/

# Build the .deb package
dpkg-deb --build debian kage_4.0.6_arm64.deb

echo "Package created: kage_4.0.6_arm64.deb"
echo ""
echo "To install on Termux:"
echo "  1. Copy the .deb file to your device"
echo "  2. Run: dpkg -i kage_4.0.6_arm64.deb"
echo ""
echo "Or to create a Termux repository:"
echo "  1. Host this .deb on a web server"
echo "  2. Create a packages.json with the package metadata"
echo "  3. Add the repository to Termux package sources"
