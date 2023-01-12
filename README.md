# gin_web_webcam 结构

```
├── README.md
├── go.mod
├── go.sum
├── image # 图像存储文件
├── main.go
├── src
│   └── router # 后端路由
│       └── route.go
└── static # 静态文件
    ├── index.html
    ├── index.js
    └── webcam.min.js
```
## 先决条件

 - goline编译器

### 初始化

```shell
go mod download
go mod tidy
```

## 运行命令

```shell
go run main.go
```

## 构建命令

```shell 
go build -tags netgo -ldflags '-s -w' -o gin_web_webcam main.go
```

