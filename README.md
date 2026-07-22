<div align="center">

```
  ██████╗  ██████╗  ██╗ ████████╗
 ██╔════╝  ██╔══██╗ ██║ ╚══██╔══╝
 ██║  ███╗ ██████╔╝ ██║    ██║
 ██║   ██║ ██╔══██╗ ██║    ██║
 ╚██████╔╝ ██║  ██║ ██║    ██║
  ╚═════╝  ╚═╝  ╚═╝ ╚═╝    ╚═╝
```
[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![License](https://img.shields.io/badge/License-GPLv3-blue.svg?style=for-the-badge)](LICENSE)
[![GitHub stars](https://img.shields.io/github/stars/shad1k2/grit?style=for-the-badge&logo=github)](https://github.com/shad1k2/grit/stargazers)
[![GitHub release](https://img.shields.io/github/v/release/shad1k2/grit?style=for-the-badge&logo=github)](https://github.com/shad1k2/grit/releases)
[![GitHub issues](https://img.shields.io/github/issues/shad1k2/grit?style=for-the-badge&logo=github)](https://github.com/shad1k2/grit/issues)
[![Code size](https://img.shields.io/github/languages/code-size/shad1k2/grit?style=for-the-badge&logo=github)](https://github.com/shad1k2/grit)

</div>
**GRIT** is a system utility for Linux designed for system diagnostics and recovery. Written in Go for maximum speed and ease of use.

##  Features

-  **Service Check** - Quick systemd service status
-  **Log Analysis** - View critical errors from journalctl
-  **Mounting** - Interactive partition mounting
-  **System Info** - RAM, disks, uptime
-  **Instant Launch** - Single binary with no dependencies

##  Installation

### Source-based
```bash
git clone https://github.com/shad1k2/grit.git
cd grit
go build -o grit
sudo cp grit /usr/local/bin/
```

### From release
# Download the binary from the Releases section
```bash
wget https://github.com/shad1k2/grit/releases/latest/download/grit-linux-amd64
chmod +x grit-linux-amd64
sudo mv grit-linux-amd64 /usr/local/bin/grit
```

<div align="center">

GRIT - Made with 💚 in Go
</div>
