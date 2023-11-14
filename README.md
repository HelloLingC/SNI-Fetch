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
./sni-fetch.exe -t 142.251.46.206 -n 0
```

可选参数：

```
-n 请求的SNI数量，默认为1
如设置为0则检测全部可检测的域名
-c 检测任务的并发数量，请根据系统资源自行调整，默认为50
```


# 其他

[目前使用的查询源](https://bgp.he.net)

## 开源许可

MIT