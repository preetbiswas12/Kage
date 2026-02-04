# How to Submit Kage to Official Termux Repository

## Overview

To make `kage` installable via `pkg install kage`, you need to submit it to the official [Termux Packages](https://github.com/termux/termux-packages) repository.

## Step-by-Step Guide

### Step 1: Fork the Termux Packages Repository

1. Go to https://github.com/termux/termux-packages
2. Click **Fork** in the top right
3. Clone your fork locally:
```bash
git clone https://github.com/YOUR_USERNAME/termux-packages.git
cd termux-packages
```

### Step 2: Create Package Definition Directory

Create a new directory for kage:

```bash
mkdir -p packages/kage
cd packages/kage
```

### Step 3: Create `build.sh`

Create `packages/kage/build.sh`:

```bash
TERMUX_PKG_HOMEPAGE=https://github.com/preetbiswas12/Kage
TERMUX_PKG_DESCRIPTION="Manga scraper and downloader CLI"
TERMUX_PKG_LICENSE="MIT"
TERMUX_PKG_MAINTAINER="Preet Biswas @preetbiswas12"
TERMUX_PKG_VERSION=4.0.6
TERMUX_PKG_REVISION=0
TERMUX_PKG_SRCURL=https://github.com/preetbiswas12/Kage/archive/refs/tags/v${TERMUX_PKG_VERSION}.tar.gz
TERMUX_PKG_SHA256=PLACEHOLDER_SHA256  # Will be filled after first build
TERMUX_PKG_DEPENDS="ca-certificates"
TERMUX_PKG_BUILD_DEPENDS="golang"
TERMUX_PKG_BLACKLISTED_ARCHES="i686"  # Optional: if not building for 32-bit

termux_step_pre_configure() {
	# Navigate to the mangal subdirectory
	cd "$TERMUX_PKG_SRCDIR"/mangal
}

termux_step_make() {
	# Build the binary
	go build \
		-o kage \
		-ldflags "-X github.com/preetbiswas12/Kage/constant.Version=$TERMUX_PKG_VERSION"
}

termux_step_make_install() {
	# Install the binary
	install -Dm700 kage "$TERMUX_PREFIX"/bin/kage
}
```

### Step 4: Create `patches/` Directory (Optional)

If no patches are needed, create an empty directory:

```bash
mkdir -p packages/kage/patches
touch packages/kage/patches/.gitkeep
```

### Step 5: Get SHA256 Hash

Build locally to get the SHA256 hash:

```bash
cd termux-packages
./build-package.sh kage
```

If it fails on SHA256, it will show you the correct hash. Copy it and update `build.sh`:

```bash
TERMUX_PKG_SHA256=<hash_from_error_message>
```

### Step 6: Test the Build

Test that it builds correctly:

```bash
./build-package.sh kage
```

The binary should be created in `./debs/kage_*.deb`

### Step 7: Test Installation on Termux

Transfer the `.deb` to your Termux device and test:

```bash
# On Termux:
dpkg -i kage_4.0.6_0_aarch64.deb
kage --help
kage sources list
```

### Step 8: Submit Pull Request

1. Commit your changes:
```bash
git add packages/kage/
git commit -m "Add kage package - manga scraper and downloader"
```

2. Push to your fork:
```bash
git push origin main
```

3. Go to https://github.com/termux/termux-packages
4. Click **New Pull Request**
5. Select your fork and branch
6. Fill in the PR description:

```markdown
## Package: kage

- **Description**: Manga scraper and downloader CLI with support for Mangadex and Mangapill
- **Version**: 4.0.6
- **Architecture**: arm64, arm, aarch64
- **Dependencies**: Go (build-time), ca-certificates (runtime)

## Features
- Search and download manga from multiple sources
- Supports Mangadex API and Mangapill HTML scraping
- Fast and lightweight

## Testing
- [x] Builds successfully on ARM64
- [x] Binary works correctly (kage --help)
- [x] All sources functional (kage sources list)
- [x] Search and download tested

## Related
- Repository: https://github.com/preetbiswas12/Kage
- Maintainer: @preetbiswas12
```

## Termux Package Structure Reference

Your final directory structure should look like:

```
termux-packages/
└── packages/
    └── kage/
        ├── build.sh
        └── patches/
            └── .gitkeep
```

## Important Notes

1. **SHA256**: Don't use a placeholder - you need the actual SHA256 of the source tarball
2. **Maintainer**: Add yourself as the maintainer
3. **License**: Ensure you have the correct license specified
4. **Dependencies**: List all runtime dependencies
5. **Testing**: Test thoroughly before submitting PR
6. **Update Frequency**: Commit to updating the package when new versions are released

## After Approval

Once the PR is merged:
- Users can install with: `pkg install kage`
- New releases should be submitted as new PRs with updated version numbers
- The maintainer (you) should keep the package updated

## Helpful Resources

- Termux Packages Contribution Guide: https://github.com/termux/termux-packages/blob/master/README.md
- Go Package Example: https://github.com/termux/termux-packages/tree/master/packages/go
- Termux Build System: https://github.com/termux/termux-packages/wiki

## Alternative: Create Your Own Termux Repository

If you prefer to maintain your own repository without going through the official approval process:

See the detailed guide in `TERMUX_INSTALLATION.md`

---

## Quick Checklist

- [ ] Fork termux-packages repo
- [ ] Create `packages/kage/build.sh` with correct metadata
- [ ] Get SHA256 hash of source tarball
- [ ] Test build locally: `./build-package.sh kage`
- [ ] Test installation on Termux device
- [ ] Create pull request with detailed description
- [ ] Respond to review comments
- [ ] Merge and celebrate! 🎉
