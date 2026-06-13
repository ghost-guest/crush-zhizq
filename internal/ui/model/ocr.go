package model

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/charmbracelet/crush/internal/message"
)

// ocrImage runs OCR on an image attachment and returns a text attachment.
// If OCR fails, returns the original image attachment unchanged.
func ocrImage(ctx context.Context, att message.Attachment) message.Attachment {
	// Look for OCR script in config dir or project .crush/
	ocrScript := findOCRScript()
	if ocrScript == "" {
		return att // No OCR script configured
	}

	// Write image to temp file
	tmpFile, err := os.CreateTemp("", "crush-ocr-*.png")
	if err != nil {
		return att
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(att.Content); err != nil {
		tmpFile.Close()
		return att
	}
	tmpFile.Close()

	// Run OCR script with 10s timeout
	ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctxTimeout, ocrScript, tmpFile.Name())
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return att // OCR failed, keep original
	}

	text := stdout.String()
	if text == "" {
		return att // Empty result
	}

	// Return text attachment
	return message.Attachment{
		FilePath: att.FileName + ".txt",
		FileName: att.FileName + " (OCR)",
		MimeType: "text/plain",
		Content:  []byte(fmt.Sprintf("[OCR from %s]\n\n%s", att.FileName, text)),
	}
}

// findOCRScript searches for an OCR script in:
// 1. ~/.config/crush/ocr.sh
// 2. ./.crush/ocr.sh
// 3. ./ocr.sh
func findOCRScript() string {
	candidates := []string{
		filepath.Join(os.Getenv("HOME"), ".config", "crush", "ocr.sh"),
		filepath.Join(".crush", "ocr.sh"),
		"ocr.sh",
	}
	for _, path := range candidates {
		if info, err := os.Stat(path); err == nil && !info.IsDir() && info.Mode()&0111 != 0 {
			abs, _ := filepath.Abs(path)
			return abs
		}
	}
	return ""
}

// isImageAttachment returns true if the attachment is an image.
func isImageAttachment(att message.Attachment) bool {
	return att.MimeType != "" &&
		(att.MimeType == "image/png" ||
			att.MimeType == "image/jpeg" ||
			att.MimeType == "image/jpg" ||
			att.MimeType == "image/gif" ||
			att.MimeType == "image/webp" ||
			att.MimeType[:6] == "image/")
}
