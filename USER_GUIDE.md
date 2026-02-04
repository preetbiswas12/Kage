# How Users Will Use Kage

## Installation

### On Termux (Android)
```bash
pkg install go git
git clone https://github.com/preetbiswas12/Kage.git
cd Kage/mangal
bash termux-install.sh
```

### On Linux/macOS
```bash
# From source
git clone https://github.com/preetbiswas12/Kage.git
cd Kage/mangal
go build -o kage
sudo mv kage /usr/local/bin/
```

### On Windows
```powershell
# Download from releases or build from source
git clone https://github.com/preetbiswas12/Kage.git
cd Kage/mangal
go build -o kage.exe
# Add to PATH or run from directory
.\kage.exe
```

---

## Basic Usage

### 1. Interactive TUI Mode (Default)

Just run `kage` to open the terminal UI:

```bash
kage
```

This opens an interactive menu where you can:
- 🔍 **Search** for manga by title
- 📖 **Browse** chapters
- ⬇️ **Download** with various formats
- 📖 **Read** in terminal
- 💾 **Save** reading progress

**Controls:**
- Arrow keys - Navigate
- Enter - Select
- Space - Toggle
- q - Quit

---

### 2. Inline Mode (Scripting/Automation)

Perfect for scripts, automation, and getting JSON output:

#### Search and get JSON results
```bash
kage inline -q "Death Note" -j
```

Output:
```json
{
  "query": "Death Note",
  "result": [
    {
      "source": "Mangadex",
      "mangal": {
        "name": "Death Note",
        "url": "https://mangadex.org/title/...",
        "chapters": 108,
        ...
      }
    }
  ]
}
```

#### Download first manga's first chapter
```bash
kage inline -q "Death Note" -m first -c first -d
```

#### Download specific chapters (1-10)
```bash
kage inline -q "Death Note" -m first -c 1-10 -d
```

#### Download all chapters
```bash
kage inline -q "Death Note" -m first -c all -d
```

#### Download from specific source
```bash
kage inline -q "Death Note" -S "Mangadex" -m first -c first -d
```

#### Custom download format
```bash
kage inline -q "Death Note" -m first -c first -d -F pdf
# Formats: pdf, cbz, zip, plain
```

---

### 3. Available Commands

#### List Available Sources
```bash
kage sources list
```

Output:
```
Builtin:
Mangadex
Mangapill

Custom:
```

#### Show Configuration
```bash
kage config
```

#### Check Version
```bash
kage version
```

#### Show Directories
```bash
kage where
```

Shows paths to:
- Downloads directory
- Cache directory
- Config file

#### Clear Cache
```bash
kage clear
```

---

## Common Use Cases

### 📱 Termux User (Android)

```bash
# Search for manga
kage inline -q "Jujutsu Kaisen" -j

# Download first manga, all chapters to phone
kage inline -q "Jujutsu Kaisen" -m first -c all -d

# Read downloaded manga
kage inline -q "Jujutsu Kaisen" -m first
```

### 💻 Linux/macOS User

```bash
# Interactive mode
kage

# Batch download multiple manga
for manga in "Death Note" "Demon Slayer" "Bleach"; do
  kage inline -q "$manga" -m first -c all -d -F pdf
done

# Export as different formats
kage inline -q "Death Note" -m first -c 1-5 -d -F cbz  # Comic book format
```

### 🖥️ Windows User

```powershell
# Interactive mode
.\kage.exe

# Download specific chapters as PDF
.\kage.exe inline -q "One Piece" -m first -c 1-10 -d -F pdf

# Get results as JSON for other programs
.\kage.exe inline -q "Naruto" -j | jq .
```

### 🤖 Automation/Scripts

```bash
#!/bin/bash
# Auto-download new manga releases

MANGA_LIST="Death Note,Jujutsu Kaisen,Demon Slayer,Bleach"

for manga in ${MANGA_LIST//,/ }; do
  echo "Downloading $manga..."
  kage inline -q "$manga" -m first -c all -d -F pdf -S Mangadex
done

echo "Downloads complete!"
```

---

## Inline Mode Flags Explained

| Flag | Description | Example |
|------|-------------|---------|
| `-q, --query` | Search query | `-q "Death Note"` |
| `-m, --manga` | Manga selector | `-m first`, `-m 0`, `-m last` |
| `-c, --chapter` | Chapter selector | `-c first`, `-c 1-10`, `-c all` |
| `-d, --download` | Download manga | `-d` |
| `-F, --format` | Output format | `-F pdf`, `-F cbz`, `-F zip`, `-F plain` |
| `-S, --source` | Select source | `-S Mangadex`, `-S Mangapill` |
| `-j, --json` | JSON output | `-j` |
| `-H, --write-history` | Save progress | `-H true` (default) |

---

## Output Formats

### PDF
Best for: Ebooks, sharing, archiving
```bash
kage inline -q "Death Note" -m first -c all -d -F pdf
```

### CBZ (Comic Book Archive)
Best for: Comic readers, detailed viewing
```bash
kage inline -q "Death Note" -m first -c all -d -F cbz
```

### ZIP
Best for: Portable, cross-platform
```bash
kage inline -q "Death Note" -m first -c all -d -F zip
```

### Plain Images
Best for: Custom processing, web uploads
```bash
kage inline -q "Death Note" -m first -c all -d -F plain
```

---

## Configuration

Edit config file:
```bash
kage config
```

Or environment variables:
```bash
export KAGE_DOWNLOAD_DIR="/custom/path"
kage inline -q "Death Note" -m first -d
```

---

## Features Summary

✅ **Multi-Source Support** - Mangadex (API), Mangapill (HTML)
✅ **Multiple Formats** - PDF, CBZ, ZIP, Plain Images
✅ **Interactive TUI** - Beautiful terminal interface
✅ **Scripting Ready** - JSON output for automation
✅ **Fast Downloading** - Parallel downloads
✅ **Progress Tracking** - Resume reading history
✅ **Cross-Platform** - Linux, macOS, Windows, Termux
✅ **No Dependencies** - Single binary, everything built-in

---

## Example Workflows

### Workflow 1: Find and Download
```bash
# Search
kage inline -q "Jujutsu Kaisen" -j

# Download specific manga
kage inline -q "Jujutsu Kaisen" -m 0 -c 1-50 -d -F pdf
```

### Workflow 2: Batch Processing
```bash
# Create a list of manga
echo "Death Note
Demon Slayer
Jujutsu Kaisen" > manga_list.txt

# Download all
while read manga; do
  kage inline -q "$manga" -m first -c all -d -F pdf
done < manga_list.txt
```

### Workflow 3: JSON to Script
```bash
# Get JSON
kage inline -q "One Piece" -j > results.json

# Parse with jq
cat results.json | jq '.result[0].mangal.name'

# Download based on JSON
MANGA_URL=$(cat results.json | jq -r '.result[0].mangal.url')
echo "Downloading from: $MANGA_URL"
```

---

## Getting Help

```bash
# General help
kage --help

# Command-specific help
kage inline --help
kage sources --help
kage config --help

# Check version
kage version

# Show directories
kage where
```

---

## Tips & Tricks

1. **Fast Search**: Use short queries
   ```bash
   kage inline -q "jjk" -j  # Shorter than "Jujutsu Kaisen"
   ```

2. **Always Specify Source** for faster results:
   ```bash
   kage inline -q "Death Note" -S Mangadex  # Faster
   ```

3. **Use JSON for automation**:
   ```bash
   kage inline -q "Bleach" -j | jq '.result[0].mangal.chapters'
   ```

4. **Batch with ranges**:
   ```bash
   kage inline -q "One Piece" -m first -c 1-100 -d -F pdf
   ```

5. **Keep downloads organized**:
   ```bash
   export KAGE_DOWNLOAD_DIR="$HOME/Manga"
   kage inline -q "Death Note" -m first -d
   ```

---

## System Requirements

- **Termux**: ARM64 device with Android 5.0+
- **Linux**: Any distribution with glibc
- **macOS**: Intel or Apple Silicon
- **Windows**: Windows 7 or newer
- **Space**: ~10MB for binary, 100MB+ for manga
- **Network**: Internet connection for downloads

---

## Next Steps

1. **Install**: Choose your platform and follow installation guide
2. **Try It**: Run `kage` to see interactive mode
3. **Explore**: Use `kage inline -q "Your favorite manga" -j` to search
4. **Download**: Add `-d` flag to download
5. **Automate**: Use inline mode in your scripts

Enjoy! 🎉
