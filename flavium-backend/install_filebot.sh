#!/bin/sh -xu

PACKAGE_VERSION=4.8.5
PACKAGE_SHA256=$(curl -fsSL https://raw.githubusercontent.com/filebot/website/master/get.filebot.net/filebot/FileBot_$PACKAGE_VERSION/FileBot_$PACKAGE_VERSION-portable.tar.xz.sha256)

PACKAGE_FILE=FileBot_$PACKAGE_VERSION-portable.tar.xz
PACKAGE_URL=https://get.filebot.net/filebot/FileBot_$PACKAGE_VERSION/$PACKAGE_FILE


# Download latest portable package
curl -o "$PACKAGE_FILE" -z "$PACKAGE_FILE" "$PACKAGE_URL"

# Check SHA-256 checksum
echo "$PACKAGE_SHA256 *$PACKAGE_FILE" | sha256sum -c || exit 1

# Extract *.tar.xz archive
tar xvf "$PACKAGE_FILE"

# Link into default $PATH
ln -sf "$PWD/filebot.sh" /usr/local/bin/filebot
