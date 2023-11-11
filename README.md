# Reality SNI Fetch Script

查找多个同一AS下的可用于 Reality SNI 的域名

中文 | [Engligh](/README_EN.md)

使用本脚本筛选出的域名：
- 与目标IP处于同一AS下（IP相似）
- 支持TLS v1.3
- 支持HTTP2

![Screenshot](/screenshots/main.png)

# 使用

在 Release 中下载可执行文件

运行脚本(以Windows为例)

```bash
./sni-fetch-windows.exe -t 192.168.1.0
```
可选参数：

```
-n 请求的SNI数量，默认为1
-c 单轮轮询的并发数量，默认为5 
```


# 其他

[目前使用的查询源](https://bgp.he.net)

## 开源许可

MIT