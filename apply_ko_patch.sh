#!/usr/bin/env bash
set -euo pipefail

ROOT="${1:-.}"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

cp -R "$SCRIPT_DIR/src" "$ROOT/"
mkdir -p "$ROOT/App/MiyooPod/assets"
cp "$SCRIPT_DIR/App/MiyooPod/assets/ui_font.ttf" "$ROOT/App/MiyooPod/assets/ui_font.ttf"

echo "한글화 소스와 폰트를 적용했습니다."
echo "빌드하려면 저장소 루트에서: make go"
