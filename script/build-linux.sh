#!/usr/bin/env bash

set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$SCRIPT_DIR/.."

CGO_ENABLED=1 fyne package --appID io.github.your-loss-sync --name 'Your Loss! Sync' --icon ./Icon.png --executable 'your-loss-sync' --release
