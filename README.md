# aTranslate

**aTranslate** 是一款基于 OpenAI API 平台的 PDF 翻译工具。通过将 PDF 文件转化为图片，然后调用 OpenAI API 实现高精度翻译。该工具旨在解决 PDF 文件翻译中公式复杂、内容提取困难的问题，提供高效、精准的翻译体验。


## 使用说明

```bash
./aTranslate translate -f demo.pdf 
```

### 全局参数
- `--config string`  
  指定配置文件路径（默认为 `aTranslate.yaml`）。

## 环境要求

由于 [go-fitz](https://github.com/gen2brain/go-fitz) 库在 Windows 平台存在编译问题，本工具仅支持在 `Linux` 平台下运行。

在使用前，需配置好 `aTranslate.yaml` 文件，内容包括以下部分：

```yaml
general:
    openai_key: sk-xxxx
    openai_url: https://api.openai.com/v1
    openai_model: chatgpt-4o-latest
```
- **openai_key**  
 OpenAI API 密钥，用于调用 ChatGPT 服务。
  
- **openai_url**  
  OpenAI API 的请求地址，默认为 `https://api.openai.com/v1`。
  
- **openai_model**  
  指定使用的 OpenAI 模型，默认为 `chatgpt-4o-latest`。

## 编译

```
git clone https://github.com/ljlvink/aTranslate
cd aTranslate
go build "-s -w"
./aTranslate
```