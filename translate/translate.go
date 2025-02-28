package translate

import (
	"aTranslate/chatgpt"
	"aTranslate/utils"
	"bytes"
	"encoding/base64"
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gen2brain/go-fitz"
)

func Pdf2img(pdfDir string, outputDir string) error {
	doc, err := fitz.New(pdfDir)
	if err != nil {
		panic(err)
	}
	defer doc.Close()
	var wg sync.WaitGroup
	numPages := doc.NumPage()
	wg.Add(numPages)
	TotTime := time.Now()
	for n := 0; n < numPages; n++ {
		go func(pageNum int) {
			defer wg.Done()
			DOC, _ := fitz.New(pdfDir)
			img, err := DOC.Image(pageNum)
			if err != nil {
				log.Errorf("failed to get image for page %d: %v\n", pageNum, err)
				return
			}
			var buf bytes.Buffer
			if err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: 90}); err != nil {
				log.Errorf("failed to encode image for page %d: %v\n", pageNum, err)
				return
			}
			encoded := "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(buf.Bytes())

			var retry int
			for retry = 0; retry < 3; retry++ {
				err = chatgpt.Img2MarkdownOut(encoded, filepath.Join(outputDir, fmt.Sprintf("Page-%d.md", pageNum)), pageNum,numPages)
				if err == nil {
					break
				}
				log.Errorf("failed to process page %d, retrying %d/3: %v\n", pageNum, retry+1, err)
				time.Sleep(5 * time.Second) // wait for 5 seconds before retrying
			}
			if err != nil {
				log.Errorf("failed to process page %d after 3 retries: %v\n", pageNum, err)
			}
		}(n)
	}
	wg.Wait()
	log.Printf("Finished processing all pages in %v\n", time.Since(TotTime))
	err = Combine2Md(outputDir, numPages)
	return err
}

/*
开始翻译流程：
1. 读取工作目录文件
2. 打开文件
3. 求文件md5
4. 早工作目录下outputs/md5(file)下将pdf拆分成图片
5. 将图片转换成base64
6. 调用openai翻译
7. 将翻译结果写入文件
8. 合并所有文件
*/
func DoTranslate(pdfDir string) error {
	// 读取工作目录文件
	workDir, _ := os.Getwd()
	outputDir := filepath.Join(workDir, "outputs")
	filemd5, err := utils.CalcFileMD5(pdfDir) // 验证文件是否存在并返回文件md5
	if err != nil {
		return fmt.Errorf("file '%s' not exist", pdfDir)
	}
	outputDir = filepath.Join(outputDir, filemd5) // 生成输出目录
	if !utils.IsDirExist(filepath.Join(outputDir, filemd5)) {
		err = os.Mkdir(outputDir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
	}
	err = Pdf2img(pdfDir, outputDir)
	return err
}
