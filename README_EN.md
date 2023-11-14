# Reality SNI Fetch Script

Quickly find multiple SNI within the same AS for Reality

Domains selected by this script:
- Is under the same AS with the target IP (Ensure IP simularity)
- Support TLS v1.3
- Support HTTP2

# How to Use

Go to Release page and download the execuable file.

Execute the script (As an example for Windwos)

```bash
./sni-fetch.exe -t 142.251.46.206 -n 0
```

Optional parameters:

```
-n The number of SNIs requests needed. Default: 1
If set to 0, will check every domains that can be checked
-c The number of allowed concurrent checking tasks. Please adjust with your system resources. Default: 50

```


# More

[Current query source](https://bgp.he.net)

# License

MIT

