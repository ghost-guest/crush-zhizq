# Crush (zhizq Fork)

<p align="center">
    <a href="https://stuff.charm.sh/crush/charm-crush.png"><img width="450" alt="Charm Crush Logo" src="https://github.com/user-attachments/assets/cf8ca3ce-8b02-43f0-9d0f-5a331488da4b" /></a><br />
    <a href="https://github.com/ghost-guest/crush-zhizq/releases"><img src="https://img.shields.io/github/release/ghost-guest/crush-zhizq" alt="Latest Release"></a>
</p>

<p align="center">终端里的 AI 编程助手<br />接入你的工具、代码和工作流，兼容主流 LLM 模型</p>

<p align="center"><img width="800" alt="Crush Demo" src="https://github.com/user-attachments/assets/58280caf-851b-470a-b6f7-d5c4ea8a1968" /></p>

[English](README.md) | 简体中文

## Fork 改动

- ⚡ **禁用自动标题生成** - 节省 50% API 调用（每次对话不再重复请求）
- 🔧 基于上游 [charmbracelet/crush](https://github.com/charmbracelet/crush)

## 特性

- 🚀 **多模型支持** - Anthropic Claude、OpenAI GPT、Google Gemini、Groq、DeepSeek 等
- 🛠️ **工具集成** - 读写文件、执行命令、搜索代码
- 🔌 **LSP 支持** - 代码补全、跳转定义、重构
- 📦 **MCP 协议** - 扩展外部工具和数据源
- 💾 **会话管理** - 保存上下文，随时恢复
- 🎨 **终端 UI** - 原生 TUI 界面，无需浏览器

## 安装

### 从源码构建（推荐）

**Linux / macOS:**

```bash
# 克隆仓库
git clone https://github.com/ghost-guest/crush-zhizq.git
cd crush-zhizq

# 构建
go build -o zhizq .

# 安装到 PATH
sudo mv zhizq /usr/local/bin/
# 或用户目录安装：
mkdir -p ~/bin && mv zhizq ~/bin/ && export PATH="$HOME/bin:$PATH"

# 运行
zhizq
```

**Windows:**

```powershell
# 克隆仓库
git clone https://github.com/ghost-guest/crush-zhizq.git
cd crush-zhizq

# 构建
go build -o zhizq.exe .

# 安装到 PATH（二选一）：

# 方法 1：复制到系统目录
move zhizq.exe C:\Windows\System32\

# 方法 2：添加当前目录到 PATH（管理员权限）
$env:Path += ";" + (Get-Location).Path
[Environment]::SetEnvironmentVariable("Path", $env:Path, [EnvironmentVariableTarget]::User)

# 运行
zhizq
```

### 原项目安装方式

参考 [原项目文档](https://github.com/charmbracelet/crush) 使用包管理器安装。

## 快速开始

构建完成后直接运行：

```bash
zhizq
```

首次启动会提示：
1. 选择 AI 供应商（Anthropic、OpenAI、Gemini 等）
2. 输入 API Key
3. 选择模型

然后就可以开始对话了！

## 配置

### 配置文件位置

**Linux / macOS:**
1. `.crush.json`（项目本地）
2. `crush.json`（项目本地）
3. `$HOME/.config/crush/crush.json`（全局）

**Windows:**
1. `.crush.json`（项目本地）
2. `crush.json`（项目本地）
3. `%USERPROFILE%\.config\crush\crush.json`（全局）

### 基础配置示例

参考 `crush.json.example`：

```json
{
  "providers": {
    "custom-openai": {
      "id": "custom-openai",
      "name": "自定义 OpenAI",
      "base_url": "https://api.example.com/v1",
      "api_key": "sk-your-key-here",
      "type": "openai",
      "models": [
        {
          "id": "gpt-4",
          "name": "GPT-4",
          "context_length": 128000,
          "max_output": 4096
        }
      ]
    }
  }
}
```

### 添加自定义供应商

支持的供应商类型：
- `anthropic` - Claude API 兼容
- `openai` - OpenAI API 兼容
- `gemini` - Google Gemini API

### 环境变量

也可以通过环境变量配置：

```bash
export ANTHROPIC_API_KEY="sk-ant-xxx"
export OPENAI_API_KEY="sk-xxx"
```

## 使用技巧

### 项目上下文

在项目目录创建 `AGENTS.md` 或 `crush.md` 文件，Crush 会自动读取作为项目背景：

```markdown
# 项目说明

这是一个 Go Web 服务，使用 Gin 框架。

## 代码规范
- 使用 gofmt 格式化
- 遵循 Effective Go
```

### 常用命令

```bash
zhizq                    # 启动交互式会话
zhizq run "问题"         # 单次问答
zhizq models             # 列出可用模型
zhizq sessions           # 查看历史会话
zhizq stats              # 查看统计数据
```

## 与原项目的区别

| 功能 | 原项目 | 本 Fork |
|------|--------|---------|
| 命令名 | `crush` | `zhizq` |
| 自动标题 | ✅ 启用（每次2个请求） | ❌ 禁用（每次1个请求） |
| API 调用 | 多一次标题生成 | 仅实际问答 |

## 常见问题

### 如何切换模型？

运行时按 `Ctrl+M` 打开模型选择器。

### 如何查看历史会话？

```bash
zhizq sessions
```

### 配置文件在哪？

运行 `zhizq` 后会显示配置文件路径。

### 支持哪些编程语言？

支持所有主流语言的 LSP：Go、Python、JavaScript/TypeScript、Rust、Java 等。

## 贡献

欢迎提 Issue 和 PR！

## 许可证

MIT License - 详见 [LICENSE.md](LICENSE.md)

## 相关链接

- 上游项目：[charmbracelet/crush](https://github.com/charmbracelet/crush)
- 模型目录：[charmbracelet/catwalk](https://github.com/charmbracelet/catwalk)
