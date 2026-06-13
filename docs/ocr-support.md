# Crush 截图粘贴与 OCR 支持

## 功能说明

Crush 现在完全支持粘贴截图：

1. **支持图片的模型**（如 GPT-4o、Claude Sonnet）
   - 直接粘贴，图片作为附件发送给模型

2. **不支持图片的模型**（如纯文本模型）
   - 自动调用 OCR 脚本提取文字
   - 将 OCR 结果作为文本附件发送

## 快捷键

- `ctrl+v`: 粘贴剪贴板中的截图
- `ctrl+shift+v`: 添加图片文件

## OCR 脚本配置

Crush 会按以下顺序查找 OCR 脚本：

1. `~/.config/crush/ocr.sh`
2. `./.crush/ocr.sh`
3. `./ocr.sh`

脚本必须：
- 可执行（`chmod +x ocr.sh`）
- 接受图片路径作为第一个参数
- 输出提取的文本到 stdout

## 安装 OCR 依赖

### macOS

```bash
brew install tesseract tesseract-lang
```

### Linux (Debian/Ubuntu)

```bash
sudo apt-get install tesseract-ocr tesseract-ocr-chi-sim
```

### 配置示例脚本

```bash
# 复制到配置目录
mkdir -p ~/.config/crush
cp ocr.sh ~/.config/crush/
chmod +x ~/.config/crush/ocr.sh
```

## 使用示例

### 1. 截图并粘贴

```text
1. 截图（macOS: cmd+shift+4, Windows: Win+Shift+S）
2. 在 Crush 输入框按 ctrl+v
3. 如果当前模型支持图片 → 直接发送图片
   如果当前模型不支持 → 自动 OCR 并发送文本
```

### 2. 自定义 OCR 脚本

你可以替换默认的 tesseract，使用其他 OCR 服务：

```bash
#!/usr/bin/env bash
# ~/.config/crush/ocr.sh

IMAGE_PATH="$1"

# 使用云端 OCR API（示例）
# curl -X POST https://api.example.com/ocr \
#   -F "image=@$IMAGE_PATH" \
#   -H "Authorization: Bearer YOUR_TOKEN"

# 或使用本地 OCR
tesseract "$IMAGE_PATH" stdout -l eng+chi_sim 2>/dev/null
```

## 故障排查

### ctrl+v 没反应

1. 确认剪贴板有图片内容
2. 检查 OCR 脚本是否可执行：`ls -l ~/.config/crush/ocr.sh`
3. 手动测试 OCR 脚本：`~/.config/crush/ocr.sh /path/to/test.png`

### OCR 输出为空

1. 检查 tesseract 是否安装：`which tesseract`
2. 测试 tesseract：`tesseract test.png stdout`
3. 查看 Crush 日志中的 OCR 错误信息

### 中文识别不准

安装中文语言包：

```bash
# macOS
brew install tesseract-lang

# Linux
sudo apt-get install tesseract-ocr-chi-sim tesseract-ocr-chi-tra
```

修改 `ocr.sh` 添加中文支持：

```bash
tesseract "$IMAGE_PATH" stdout -l eng+chi_sim 2>/dev/null
```

## 进阶配置

### 使用 PaddleOCR（更高精度）

```bash
#!/usr/bin/env bash
# ~/.config/crush/ocr.sh

IMAGE_PATH="$1"

# 需要安装: pip install paddleocr
python3 -c "
from paddleocr import PaddleOCR
ocr = PaddleOCR(use_angle_cls=True, lang='ch')
result = ocr.ocr('$IMAGE_PATH', cls=True)
for line in result[0]:
    print(line[1][0])
" 2>/dev/null
```

### 使用 Apple Vision API（macOS）

```bash
#!/usr/bin/env bash
# ~/.config/crush/ocr.sh

osascript -l JavaScript << EOF
const vision = $.NSVision.alloc.init
const image = $.NSImage.alloc.initWithContentsOfFile('$1')
const request = $.VNRecognizeTextRequest.alloc.init
// ... (需要更多 Objective-C 桥接代码)
EOF
```

## 工作原理

1. 用户按 `ctrl+v` 粘贴截图
2. Crush 读取剪贴板图片数据
3. 检查当前模型是否支持图片
4. **支持** → 直接作为图片附件添加
5. **不支持** → 调用 OCR 脚本：
   - 保存图片到临时文件
   - 执行 `ocr.sh <temp-file>`
   - 读取 stdout 输出
   - 删除临时文件
   - 将 OCR 文本作为文本附件添加
6. 用户按回车发送消息（含附件）

## 相关文件

- `internal/ui/model/ocr.go` - OCR 核心逻辑
- `internal/ui/model/ui.go` - 粘贴处理和模型检查
- `internal/ui/model/clipboard.go` - 剪贴板读取
- `ocr.sh` - 默认 OCR 脚本示例
