TERMUX_PKG_HOMEPAGE=https://github.com/preetbiswas12/Kage
TERMUX_PKG_DESCRIPTION="Manga scraper and downloader CLI with support for Mangadex and Mangapill"
TERMUX_PKG_LICENSE="MIT"
TERMUX_PKG_MAINTAINER="Preet Biswas @preetbiswas12"
TERMUX_PKG_VERSION=4.0.6
TERMUX_PKG_REVISION=0
TERMUX_PKG_SRCURL=https://github.com/preetbiswas12/Kage/archive/refs/tags/v${TERMUX_PKG_VERSION}.tar.gz
TERMUX_PKG_SHA256=REPLACE_WITH_ACTUAL_SHA256
TERMUX_PKG_DEPENDS="ca-certificates"
TERMUX_PKG_BUILD_DEPENDS="golang"
TERMUX_PKG_BLACKLISTED_ARCHES="i686"

termux_step_pre_configure() {
	# Navigate to the mangal subdirectory where the actual code is
	cd "$TERMUX_PKG_SRCDIR"/mangal
}

termux_step_make() {
	# Build the binary for the target architecture
	go build \
		-ldflags "-X github.com/preetbiswas12/Kage/constant.Version=$TERMUX_PKG_VERSION" \
		-o kage
}

termux_step_make_install() {
	# Install the binary to the Termux bin directory
	install -Dm700 kage "$TERMUX_PREFIX"/bin/kage
}
