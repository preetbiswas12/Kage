# Kage - Usage Summary

## 🎯 Two Main Ways to Use Kage

### 1️⃣ Interactive Mode (Default)
For humans who want a beautiful, interactive experience

```bash
kage
```

```
┌─────────────────────────────────────┐
│  Search for manga...                │
│  ▶ One Piece                        │
│  ▶ Jujutsu Kaisen                   │
│  ▶ Demon Slayer                     │
└─────────────────────────────────────┘
```

**Use when:**
- Exploring manually
- Browsing manga
- Interactive downloads
- Reading in terminal

---

### 2️⃣ Inline Mode (Scripting)
For automation, scripts, and JSON output

```bash
kage inline -q "Death Note" -j
```

**Use when:**
- Running scripts
- Automation
- Batch processing
- Getting JSON data
- CI/CD pipelines

---

## 📋 Quick Command Reference

### Search
```bash
# Interactive search
kage

# Get JSON results (no download)
kage inline -q "Death Note" -j

# Search specific source
kage inline -q "Death Note" -S Mangadex -j
```

### Download
```bash
# Download first manga's first chapter
kage inline -q "Death Note" -m first -c first -d

# Download chapters 1-10
kage inline -q "Death Note" -m first -c 1-10 -d

# Download all chapters
kage inline -q "Death Note" -m first -c all -d

# Choose format (pdf, cbz, zip, plain)
kage inline -q "Death Note" -m first -c all -d -F pdf
```

### Utilities
```bash
# List sources
kage sources list

# Show directories
kage where

# Clear cache
kage clear

# Version
kage version
```

---

## 🔄 Real-World Examples

### Example 1: Termux User
```bash
# Install
bash termux-install.sh

# Search
kage inline -q "Jujutsu Kaisen" -j

# Download to phone
kage inline -q "Jujutsu Kaisen" -m first -c all -d -F pdf
```

### Example 2: Linux Automation
```bash
#!/bin/bash
# Auto-download favorite manga

for manga in "Death Note" "Demon Slayer" "Bleach"; do
  echo "Downloading $manga..."
  kage inline -q "$manga" -m first -c all -d -F cbz
done

echo "All manga downloaded!"
```

### Example 3: JSON Processing
```bash
# Get results
kage inline -q "One Piece" -j > results.json

# Parse with jq
jq '.result[0].mangal.name' results.json
# Output: "One Piece"

# Get chapter count
jq '.result[0].mangal.chapters' results.json
# Output: 1100+
```

### Example 4: Batch Download with Script
```bash
#!/bin/bash
# manga_download.sh

while read manga; do
  echo "⬇️  Downloading: $manga"
  kage inline -q "$manga" -m first -c 1-50 -d -F pdf
done < manga_list.txt

echo "✅ Done!"
```

```bash
# Usage
echo -e "Death Note\nJujutsu Kaisen\nDemon Slayer" > manga_list.txt
bash manga_download.sh
```

---

## 🌍 Supported Sources

| Source | Type | Speed | Quality |
|--------|------|-------|---------|
| **Mangadex** | API | ⚡ Fast | ⭐ Excellent |
| **Mangapill** | HTML | ⚡ Fast | ⭐ Good |

---

## 📥 Download Formats

| Format | Best For | Command |
|--------|----------|---------|
| **PDF** | Ebooks, sharing | `-F pdf` |
| **CBZ** | Comic readers | `-F cbz` |
| **ZIP** | Archive, portable | `-F zip` |
| **Plain** | Custom processing | `-F plain` |

---

## 🎮 Interactive Mode Controls

| Key | Action |
|-----|--------|
| ↑↓ | Navigate |
| ← → | Scroll |
| Enter | Select |
| Space | Toggle |
| q | Quit |
| ? | Help |

---

## 💡 Pro Tips

1. **Use JSON for scripting**
   ```bash
   kage inline -q "Death Note" -j | jq '.result[0]'
   ```

2. **Specify source for speed**
   ```bash
   kage inline -q "Death Note" -S Mangadex -d  # Faster
   ```

3. **Batch with ranges**
   ```bash
   kage inline -q "One Piece" -m first -c 1-100 -d -F pdf
   ```

4. **Custom download location**
   ```bash
   export KAGE_DOWNLOAD_DIR="$HOME/MyManga"
   kage inline -q "Death Note" -m first -d
   ```

5. **Check available space**
   ```bash
   kage where  # Shows download directory
   ```

---

## 📦 Installation Quick Links

- **Termux (Android)**: See `TERMUX_SETUP.md`
- **Linux/macOS**: Build from source
- **Windows**: Build from source or download binary
- **Official Termux Repo**: See `TERMUX_SUBMISSION_QUICK.md`

---

## 🆘 Getting Help

```bash
kage --help          # General help
kage inline --help   # Inline mode help
kage sources --help  # Sources help
kage config --help   # Config help
```

---

## 🚀 Next Steps

1. Install Kage on your platform
2. Try: `kage inline -q "Death Note" -j`
3. Download: `kage inline -q "Death Note" -m first -c first -d`
4. Explore: `kage` (interactive mode)
5. Automate: Create scripts!

**Enjoy downloading manga! 🎉**
