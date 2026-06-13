#!/usr/bin/env bash
# Crush OCR Script
# 
# This script is called by Crush when pasting an image and the current model
# doesn't support images. It should read an image file and output extracted text.
#
# Usage: ocr.sh <image-path>
#
# Requirements:
# - macOS: Install tesseract via: brew install tesseract tesseract-lang
# - Linux: Install tesseract via: apt-get install tesseract-ocr
#
# The script receives the image path as the first argument and should output
# the extracted text to stdout.

set -euo pipefail

IMAGE_PATH="${1:-}"

if [[ -z "$IMAGE_PATH" ]]; then
    echo "Error: No image path provided" >&2
    exit 1
fi

if [[ ! -f "$IMAGE_PATH" ]]; then
    echo "Error: Image file not found: $IMAGE_PATH" >&2
    exit 1
fi

# Check if tesseract is installed
if ! command -v tesseract &> /dev/null; then
    echo "Error: tesseract not found. Install it first:" >&2
    echo "  macOS: brew install tesseract" >&2
    echo "  Linux: apt-get install tesseract-ocr" >&2
    exit 1
fi

# Run OCR with tesseract
# Output to stdout (-), use English language, output plain text
tesseract "$IMAGE_PATH" stdout -l eng 2>/dev/null || {
    echo "Error: OCR failed" >&2
    exit 1
}
