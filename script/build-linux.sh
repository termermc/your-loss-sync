#!/usr/bin/env bash

set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$SCRIPT_DIR/.."

fyne package --appID io.github.your-loss-sync --name 'Your Loss! Sync' --icon ../icon.png --src ./cmd/ --executable 'your-loss-sync' --release
