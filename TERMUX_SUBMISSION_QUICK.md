# Submitting Kage to Termux - Quick Reference

## 🚀 Quick Submission Steps

### 1. Fork the Termux Packages Repository
```
Go to: https://github.com/termux/termux-packages
Click "Fork" button
```

### 2. Clone Your Fork
```bash
git clone https://github.com/YOUR_USERNAME/termux-packages.git
cd termux-packages
```

### 3. Create Package Directory
```bash
mkdir -p packages/kage/patches
touch packages/kage/patches/.gitkeep
```

### 4. Add build.sh
Copy the content from `termux-build.sh` in this repository and save it as `packages/kage/build.sh`

### 5. Get SHA256 Hash
```bash
# Download the source
wget https://github.com/preetbiswas12/Kage/archive/refs/tags/v4.0.6.tar.gz

# Get SHA256
sha256sum v4.0.6.tar.gz

# Copy the hash and update TERMUX_PKG_SHA256 in packages/kage/build.sh
```

### 6. Test Build Locally
```bash
cd termux-packages
./build-package.sh kage
```

### 7. Test on Termux Device
```bash
# Transfer the .deb file to your Termux device
# Then run:
dpkg -i kage_4.0.6_0_aarch64.deb

# Verify installation
kage sources list
kage --help
```

### 8. Commit and Push
```bash
git add packages/kage/
git commit -m "Add kage package v4.0.6 - Manga scraper and downloader CLI"
git push origin main
```

### 9. Create Pull Request
1. Go to https://github.com/termux/termux-packages
2. Click "New Pull Request"
3. Select your fork as source
4. Write description:
```
## Package: kage v4.0.6

Manga scraper and downloader CLI supporting Mangadex and Mangapill sources.

### Features
- Multi-source manga search and download
- Mangadex API integration
- Mangapill HTML scraping
- Fast and lightweight

### Testing
- ✓ Builds on ARM64
- ✓ Binary executes correctly
- ✓ Sources list available
- ✓ Search and download tested

### References
- Repository: https://github.com/preetbiswas12/Kage
- Maintainer: @preetbiswas12
```

## 📋 Checklist

- [ ] Forked termux-packages
- [ ] Created packages/kage/ directory
- [ ] Added build.sh with correct SHA256
- [ ] Tested local build
- [ ] Tested on Termux device
- [ ] Committed changes
- [ ] Pushed to fork
- [ ] Created PR with good description

## 📚 Files to Use

In this repository:
- **termux-build.sh** - Ready-to-use build.sh template
- **SUBMIT_TO_TERMUX.md** - Detailed step-by-step guide
- **TERMUX_SETUP.md** - Installation instructions for users

## 🔗 Important Links

- **Termux Packages Repo**: https://github.com/termux/termux-packages
- **Contribution Guide**: https://github.com/termux/termux-packages/blob/master/README.md
- **Your Repository**: https://github.com/preetbiswas12/Kage

## ❓ Frequently Asked Questions

**Q: How long does approval take?**
A: Usually 1-2 weeks, depends on maintainer availability

**Q: Do I need to maintain the package?**
A: Yes, you should update it when new versions are released

**Q: Can I use my own repository instead?**
A: Yes, see TERMUX_INSTALLATION.md for custom repository setup

**Q: What if the build fails?**
A: Check the error message and refer to SUBMIT_TO_TERMUX.md troubleshooting section

**Q: Can I test build.sh before submitting?**
A: Yes, clone termux-packages and run `./build-package.sh kage` locally

## 🎯 After Approval

Once merged, users can install with:
```bash
pkg install kage
```

And use it immediately on Termux!
