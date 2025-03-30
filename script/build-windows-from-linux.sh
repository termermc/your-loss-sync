#!/usr/bin/env bash

set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$SCRIPT_DIR/.."

REAL_USER=$(whoami)
ROOT_DIR="$(pwd)"
BUILD_DIR="/tmp/yls-build"

trap "sudo rm -rf $BUILD_DIR" EXIT

cp -r "$ROOT_DIR" "$BUILD_DIR"
cd "$BUILD_DIR"

sudo fyne-cross windows --arch="*"
sudo mv fyne-cross/dist/ /tmp/yls-dist
sudo chown -R "$REAL_USER" /tmp/yls-dist
mv /tmp/yls-dist "$ROOT_DIR/dist"

cd "$ROOT_DIR"
zip -r dist/windows-386/your-loss-sync.exe.zip ffmpeg-bin/
zip -r dist/windows-amd64/your-loss-sync.exe.zip ffmpeg-bin/
zip -r dist/windows-arm64/your-loss-sync.exe.zip ffmpeg-bin/

mv dist/windows-386/your-loss-sync.exe.zip dist/your-loss-sync-windows-386.zip
mv dist/windows-amd64/your-loss-sync.exe.zip dist/your-loss-sync-windows-amd64.zip
mv dist/windows-arm64/your-loss-sync.exe.zip dist/your-loss-sync-windows-arm64.zip
rm -rf dist/yls-dist || true
rm -r dist/windows-*

echo "Built files available in $ROOT_DIR/dist"
