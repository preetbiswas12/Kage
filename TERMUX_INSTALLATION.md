# Termux Package Repository Setup

To make `kage` installable via `pkg install kage` in Termux, you need to set up a custom Termux package repository.

## Option 1: Manual Installation (Easiest)

```bash
# On Linux/WSL, build the package:
cd /path/to/kage
bash build-deb.sh

# This creates: kage_4.0.6_arm64.deb

# Transfer to your Termux device and install:
dpkg -i kage_4.0.6_arm64.deb
```

Then use it:
```bash
kage search "Death Note"
kage download "Death Note" -S Mangadex
```

## Option 2: Create a Custom Termux Repository

To make it installable via `pkg install kage`, you need to create a Termux package repository:

### Step 1: Build the package
```bash
bash build-deb.sh
```

### Step 2: Create a web-hosted repository

Host the `.deb` file on a web server and create a `packages.json`:

```json
{
  "name": "Kage Repository",
  "suites": {
    "stable": {
      "location": "https://your-domain.com/kage/stable",
      "components": "main",
      "architectures": "arm64"
    }
  }
}
```

### Step 3: Add repository to Termux

In Termux, edit `/data/data/com.termux/files/etc/apt/sources.list` and add:

```
deb https://your-domain.com/kage/stable stable main
```

Then:
```bash
apt update
apt install kage
```

## Option 3: Submit to Official Termux Repository

See: https://github.com/termux/termux-packages

This is the most user-friendly option but requires following Termux packaging guidelines.

## Building for Multiple Architectures

The current script builds for ARM64 (most common). To support more:

```bash
# ARM (32-bit)
GOOS=linux GOARCH=arm go build -o kage-arm

# x86_64 (rarely used in Termux)
GOOS=linux GOARCH=amd64 go build -o kage-amd64
```

Then create separate .deb packages for each architecture.
