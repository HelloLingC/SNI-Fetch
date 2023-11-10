# Reality SNI Fetch Script

查找多个同一AS下的可用于 Reality SNI 的域名

中文 | [Engligh](/README_EN.md)

使用本脚本筛选出的域名：
与目标IP处于同一AS下（IP相似）
支持TLS v1.3
支持HTTP2

# 使用

安装 Golang

```bash
git clone
```

安装依赖

```bash
go get ./,,,
```

运行脚本

```bash
go run . -t 目标IP
```
可选参数：
-n 请求的SNI数量

# 其他

[目前使用的查询源](https://bgp.he.net)