package translate

import (
	"aTranslate/utils"
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func Combine2Md(outputDir string, totPages int) error {
	var combinedMarkdown bytes.Buffer
	vaildPages := 0
	for i := 0; i < totPages; i++ {
		pagePath := filepath.Join(outputDir, fmt.Sprintf("Page-%d.md", i))
		if utils.IsFileExist(pagePath) {
			vaildPages++
			content, err := os.ReadFile(pagePath)
			if err != nil {
				return fmt.Errorf("failed to read markdown file for page %d: %v", i, err)
			}
			combinedMarkdown.Write(content)
			combinedMarkdown.WriteString("\n\n")
		}
	}
	combinedPath := filepath.Join(outputDir, "combined.md")
	if err := os.WriteFile(combinedPath, combinedMarkdown.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write combined markdown file: %v", err)
	}
	log.Printf("Combined %d/%d pages into %s\n", vaildPages, totPages, combinedPath)
	return nil
}
