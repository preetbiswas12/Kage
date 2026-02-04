#!/bin/bash
# Quick installer for Termux

set -e

ARCH=$(getprop ro.product.cpu.abi2 2>/dev/null || getprop ro.product.cpu.abi)
VERSION="4.0.6"

echo "Installing Kage v${VERSION} for Termux..."

# Detect architecture
if [[ "$ARCH" == "arm64-v8a" ]] || [[ "$ARCH" == "aarch64" ]]; then
    echo "Detected: ARM64"
    GOARCH="arm64"
elif [[ "$ARCH" == "armeabi-v7a" ]]; then
    echo "Detected: ARM"
    GOARCH="arm"
elif [[ "$ARCH" == "x86_64" ]]; then
    echo "Detected: x86_64"
    GOARCH="amd64"
else
    echo "Unknown architecture: $ARCH"
    echo "Defaulting to ARM64..."
    GOARCH="arm64"
fi

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Go is required but not installed."
    echo "Install it with: pkg install golang"
    exit 1
fi

# Build for Termux
GOOS=linux GOARCH=$GOARCH go build -o kage

# Install to Termux bin
mkdir -p $PREFIX/bin
cp kage $PREFIX/bin/
chmod +x $PREFIX/bin/kage

echo ""
echo "✓ Kage installed successfully!"
echo ""
echo "Usage:"
echo "  kage sources list                    - List available sources"
echo "  kage search 'Death Note'             - Search for manga"
echo "  kage download 'Death Note'           - Download manga"
echo ""
