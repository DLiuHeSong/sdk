# sdk

# University SDK

这是一个简单的中国大学查询SDK，包含 Python 和 Go 两种语言版本。

## 目录结构
- `data/`：大学数据（universities.json）
- `python/`：Python SDK 代码和示例
- `go/`：Go SDK 代码和示例

## 使用说明

### Python

```bash
cd python
python example.py
```

### Go

```bash
cd go
go mod tidy
go run example.go
```
### json文件示例
[Json示例文件](https://raw.githubusercontent.com/DLiuHeSong/sdk/refs/heads/master/data/universities_example.json)

## 项目结构

```
university-sdk/
├── data/
│   └── universities.json         # 示例大学数据
├── python/
│   ├── university_sdk.py         # Python SDK 核心代码
│   ├── example.py                # Python 使用示例
│   └── requirements.txt          # Python 依赖（只有标准库）
├── go/
│   ├── university/               # Go SDK 包
│   │   ├── university.go
│   │   └── university_test.go
│   ├── example.go                # Go 使用示例
│   └── go.mod                   # Go module 初始化文件
└── README.md                    # 项目说明
```

