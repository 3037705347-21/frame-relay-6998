# frame-relay__004 Docker 交付说明

## 项目概览
- Frame Relay is an in-memory Go library for coordinating ordered protocol frames
- Go module: `example.com/frame-relay`

## 标准命令

```bash
go build ./...
go test ./...
```

## Docker 构建

```bash
./build_benzhi_docker.sh frame-relay__004-benzhi linux/amd64
docker run --rm -it frame-relay__004-benzhi bash
```

## 环境

- 基础镜像: `golang:1.26`
- 依赖在镜像构建阶段预下载，容器内可直接执行 Go 构建和测试命令。
- 代码目录: `/app`
