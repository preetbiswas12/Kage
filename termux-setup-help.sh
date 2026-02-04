#!/bin/bash
# =============================================================================
# Kage Termux Quick Setup
# =============================================================================
# This script helps you set up Kage for Termux installation
# =============================================================================

echo "╔═══════════════════════════════════════════════════════════════════════╗"
echo "║                    Kage Termux Package Setup                          ║"
echo "╚═══════════════════════════════════════════════════════════════════════╝"
echo ""

# Check if we're on a system that can build the package
if ! command -v dpkg-deb &> /dev/null; then
    echo "⚠️  dpkg-deb not found. You can still build for Termux."
    echo ""
    echo "To build the Debian package (.deb), you need to:"
    echo "  • Run this on Linux or WSL"
    echo "  • Have dpkg-deb installed (apt install dpkg)"
    echo ""
else
    echo "✓ dpkg-deb found - you can build Debian packages"
fi

echo ""
echo "📦 Available setup options:"
echo ""
echo "1. Build from source (recommended for Termux users)"
echo "   → Automatically builds and installs kage"
echo "   → Run on your Termux device with:"
echo "     bash termux-install.sh"
echo ""
echo "2. Build Debian package (.deb)"
echo "   → Creates kage_4.0.6_arm64.deb"
echo "   → Transfer and install with: dpkg -i kage_4.0.6_arm64.deb"
echo "   → Run with: bash build-deb.sh"
echo ""
echo "3. Build custom repository"
echo "   → See: TERMUX_INSTALLATION.md"
echo ""
echo "═══════════════════════════════════════════════════════════════════════"
echo ""
echo "📖 Documentation:"
echo "   • TERMUX_SETUP.md - Quick start guide"
echo "   • TERMUX_INSTALLATION.md - Detailed repository setup"
echo "   • README.md - General Kage documentation"
echo ""
echo "═══════════════════════════════════════════════════════════════════════"
