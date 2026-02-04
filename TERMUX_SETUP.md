# Kage Termux Package Setup

## Quick Start for Termux Users

### Easiest Method: Build and Install from Source

On your Termux terminal:

```bash
# Install Go (if not already installed)
pkg install go git

# Clone the repository
git clone https://github.com/preetbiswas12/Kage.git
cd Kage/mangal

# Run the installer
bash termux-install.sh
```

Then use it:
```bash
kage search "Death Note"
kage download "Death Note" -S Mangadex
kage sources list
```

## For Package Maintainers: Creating a Termux Repository

If you want to make `kage` installable via `pkg install kage`, you need to:

### 1. Build the Debian Package

```bash
bash build-deb.sh
```

This creates: `kage_4.0.6_arm64.deb`

### 2. Submit to Official Termux Repository

Follow the guidelines at https://github.com/termux/termux-packages:
- Fork the repository
- Create a package definition following Termux standards
- Submit a pull request

### 3. Alternative: Create a Custom Repository

Host the `.deb` file and metadata on a web server, then users can add your repository to Termux.

## What's Included

- `termux-install.sh` - Quick installer script for Termux
- `build-deb.sh` - Builds a Debian package for ARM64
- `debian/` - Debian package metadata
- `TERMUX_INSTALLATION.md` - Detailed setup instructions

## Supported Architectures

Currently builds for:
- ARM64 (most common) ✓
- Can be extended to ARM, x86_64, etc.

## Available Sources in Kage

- **Mangadex** (API-based) - Fully functional
- **Mangapill** (HTML scraping) - Fully functional

## Usage Examples

```bash
# List available sources
kage sources list

# Search for manga
kage search "Jujutsu Kaisen"

# Download manga
kage download "Jujutsu Kaisen" -S Mangadex

# Show help
kage --help
```

## Troubleshooting

### "Go is required but not installed"
```bash
pkg install golang
```

### Permission denied when running termux-install.sh
```bash
chmod +x termux-install.sh
bash termux-install.sh
```

### Build fails
Ensure you have:
- Go 1.24+ installed
- Git installed
- At least 500MB free space

## More Information

- See [TERMUX_INSTALLATION.md](TERMUX_INSTALLATION.md) for detailed repository setup
- See [README.md](README.md) for general Kage documentation
