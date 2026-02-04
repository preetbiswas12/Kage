# Kage Help Text - Updated Complete Reference

All commands now have comprehensive help text with flags, examples, and descriptions.

## Main Help (`kage --help`)

Now includes:
- ✅ All available commands organized by category
- ✅ Interactive Mode commands
- ✅ Scripting & Automation commands
- ✅ Sources & Configuration commands
- ✅ Utilities
- ✅ Popular examples

---

## Command Reference

### Interactive Mode

#### `kage` or `kage --help`
- Opens the full interactive TUI
- Search, browse, download, and read manga
- Default mode when no arguments given

#### `kage -c, --continue`
- Resume reading from your history
- Jump to where you left off

---

### Scripting & Automation

#### `kage inline -q "manga name" [options]`
**Perfect for scripts, automation, batch processing, and JSON output**

FLAGS:
- `-q, --query STRING` - Search query (required)
- `-m, --manga STRING` - Manga selector (first, last, number, @substring@)
- `-c, --chapter STRING` - Chapter selector (first, last, all, number, range)
- `-d, --download` - Download selected chapters
- `-F, --format STRING` - Format (pdf, cbz, zip, plain)
- `-S, --source STRING` - Source (Mangadex, Mangapill)
- `-j, --json` - Output as JSON
- `-H, --write-history` - Save progress (default: true)
- `-o, --output STRING` - Save JSON to file

EXAMPLES:
```bash
# Search and get JSON
kage inline -q "Death Note" -j

# Download first manga's chapters 1-10
kage inline -q "Death Note" -m first -c 1-10 -d

# Download all chapters as PDF
kage inline -q "Jujutsu Kaisen" -m first -c all -d -F pdf

# Download as CBZ (comic format)
kage inline -q "One Piece" -m first -c all -d -F cbz

# Use specific source (faster)
kage inline -q "Naruto" -S Mangadex -m first -c all -d

# JSON output to file
kage inline -q "Bleach" -j -o results.json

# Parse with jq
kage inline -q "One Piece" -j | jq '.result[0].mangal.chapters'
```

#### `kage mini [-d] [-c]`
**Simplified, lightweight interactive mode**

FLAGS:
- `-d, --download` - Start in download mode
- `-c, --continue` - Continue from history

EXAMPLES:
```bash
kage mini              # Launch simplified interface
kage mini -d           # Start in download mode
kage mini -c           # Continue reading
```

#### `kage run [file]`
**Execute Lua scripts**

FLAGS:
- `-l, --lenient` - Skip missing function warnings

EXAMPLES:
```bash
kage run ./my-scraper.lua
kage run -l ./test.lua
```

---

### Sources & Configuration

#### `kage sources [command]`
**Manage manga sources**

SUBCOMMANDS:
- `sources list` - Show all sources (Mangadex, Mangapill, custom)
- `sources install` - Install custom scrapers
- `sources remove` - Remove custom sources
- `sources gen` - Generate new Lua scraper template

EXAMPLES:
```bash
kage sources list              # Show built-in and custom sources
kage sources list --builtin    # Show only built-in
kage sources list --custom     # Show only custom
```

#### `kage config [command]`
**Manage configuration settings**

SUBCOMMANDS:
- `config info` - Show config fields
- `config get` - Get a value
- `config set` - Set a value
- `config reset` - Reset to defaults
- `config open` - Edit config file

EXAMPLES:
```bash
kage config info               # Show all options
kage config get downloads-dir # Get a setting
kage config set default-source Mangadex
```

#### `kage env [-s] [-u]`
**List environment variables**

FLAGS:
- `-s, --set-only` - Show only set variables
- `-u, --unset-only` - Show only unset variables

EXAMPLES:
```bash
kage env                       # Show all env variables
kage env --set-only           # Show only set ones
export KAGE_DOWNLOAD_DIR="$HOME/Manga"
```

---

### Utilities

#### `kage clear [-c] [-s] [-a] [-q]`
**Clear cached data and history**

FLAGS:
- `-c, --cache` - Clear downloaded manga cache
- `-s, --history` - Clear reading history
- `-a, --anilist` - Clear Anilist bindings
- `-q, --queries` - Clear search history

EXAMPLES:
```bash
kage clear --cache            # Clear only cache
kage clear --history          # Clear history
kage clear --cache --history  # Clear both
```

#### `kage where [-d] [-c] [-h] [-l] [-s] [-b] [-q]`
**Show directory and file paths**

FLAGS:
- `-d, --downloads` - Downloads directory
- `-c, --cache` - Cache directory
- `-h, --history` - History file
- `-l, --logs` - Logs directory
- `-s, --sources` - Sources directory
- `-b, --binds` - Anilist binds file
- `-q, --queries` - Queries history file

EXAMPLES:
```bash
kage where                    # Show all paths
kage where --downloads        # Show downloads dir
kage where --cache            # Show cache dir
```

#### `kage version [-s]`
**Display version information**

FLAGS:
- `-s, --short` - Show only version number

EXAMPLES:
```bash
kage version              # Full info with build details
kage version --short      # Just the version number
```

#### `kage integration anilist [-d]`
**Configure Anilist integration**

FLAGS:
- `-d, --disable` - Disable integration

EXAMPLES:
```bash
kage integration anilist        # Set up integration
kage integration anilist -d     # Disable
```

#### `kage completion [bash|zsh|fish|powershell]`
**Generate shell completion**

EXAMPLES:
```bash
kage completion bash             # Generate bash completions
kage completion zsh | sudo tee /usr/share/zsh/site-functions
```

---

## Global Flags

Available on ALL commands:
- `-F, --format STRING` - Output format (pdf, cbz, zip, plain)
- `-I, --icons STRING` - Icons variant
- `-S, --source STRING` - Default source (Mangadex, Mangapill)
- `-H, --write-history BOOL` - Save reading progress (default: true)

---

## Common Workflows

### Quick Search
```bash
kage inline -q "Death Note" -j
```

### Download Manga
```bash
kage inline -q "Death Note" -m first -c all -d
```

### Batch Download
```bash
for manga in "Death Note" "Jujutsu Kaisen" "Bleach"; do
  kage inline -q "$manga" -m first -c all -d
done
```

### Interactive Mode
```bash
kage              # Opens TUI
```

### Continue Reading
```bash
kage -c           # Resume from history
```

---

## Tips

1. **Most commands have detailed help:**
   ```bash
   kage inline --help
   kage sources --help
   kage config --help
   ```

2. **JSON output is perfect for scripting:**
   ```bash
   kage inline -q "One Piece" -j | jq '.result[0].mangal.chapters'
   ```

3. **Use specific sources for speed:**
   ```bash
   kage inline -q "Death Note" -S Mangadex -d
   ```

4. **Specify format when downloading:**
   ```bash
   kage inline -q "Naruto" -m first -d -F pdf   # PDF
   kage inline -q "Naruto" -m first -d -F cbz   # Comic format
   ```

5. **Check your paths:**
   ```bash
   kage where              # See all directories
   kage where --downloads  # Just downloads location
   ```

---

## Getting More Help

```bash
kage --help                    # Main help
kage [command] --help         # Command-specific help
kage version                  # Version and build info
kage where                    # File locations
```

All help text is now comprehensive, with examples and clear explanations! 🎉
