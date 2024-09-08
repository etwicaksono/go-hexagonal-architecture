package utils

import (
	"path/filepath"
	"strings"
)

func IsValidExtension(allowedExtension []string, fileName string) bool {
	// Extract the file extension
	ext := strings.ToLower(filepath.Ext(fileName))
	for _, validExt := range allowedExtension {
		if ext == validExt {
			return true
		}
	}
	return false
}
