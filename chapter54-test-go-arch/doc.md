## 代码输出
```bash
Linux xxx 5.10.0-27.1.uelc20.x86_64 #1 SMP Fri May 12 00:03:52 CST 2023 x86_64 x86_64 x86_64 GNU/Linux
GOARCH: amd64

Linux ningan 5.4.0-150-generic #167~18.04.1-Ubuntu SMP Wed May 24 00:51:42 UTC 2023 x86_64 x86_64 x86_64 GNU/Linux
GOARCH: amd64

Linux xxx 4.19.0-91.82.42.uelc20.aarch64 #1 SMP Tue Jul 26 10:32:01 CST 2022 aarch64 aarch64 aarch64 GNU/Linux
GOARCH: arm64
```

## 系统架构
```bash
# uname -a | awk '{print $12}'
x86_64
# uname -a | awk '{print $12}'
aarch64
```