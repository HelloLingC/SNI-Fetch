# Reality SNI Fetch Script

Easily find available SNIs for Reality

Domains selected by this script:
- Is under the same AS with the target IP (Ensure IP simularity)
- Support TLS v1.3
- Support HTTP2

# How to Use

Go to Release page and download the execuable file.

Execute the script (As an example for Windwos)

```bash
./sni-fetch-windows.exe -t 192.168.1.0
```

Optional parameters:
```
-n The number of SNIs requests needed. Default: 1
If set to 0, will check every domains that can be checked
-c The number of concurrent checks in a single round. Default: 10

Recommended command:
./sni-fetch-windows.exe -t 192.168.1.0 -n 0 -c 50
```


# More

[Current query source](https://bgp.he.net)

# License

MIT

