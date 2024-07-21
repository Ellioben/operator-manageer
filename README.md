# Operator Manager

Operator Manager 是一个用 Go 语言编写的工具，旨在简化 Kubernetes 中操作k8s资源的管理服务。它提供了一种领域特定语言 (DSL) 来定义。结合go-zero生成相关的web框架
## 特性

- **DSL 定义**: 使用简单的 DSL 来定义api和管理 Kubernetes 操作符。
- 支持多集群配置和操作

## 技术栈

- **编程语言**: Go
- **平台**: Kubernetes

## 安装

```bash
# 克隆仓库
git clone https://github.com/Ellioben/operator-manageer.git

# 进入项目目录
cd operator-manageer

# 构建项目
go build
```
# 使用
启动 Web 界面:

bash
```
./operator-manageer
```
打开浏览器，访问 http://localhost:8080。