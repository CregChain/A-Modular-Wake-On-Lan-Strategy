# A-Modular-Wake-On-Lan-Strategy
A modular WOL strategy that use a linux or any device to start another specific device in your local network.
# WOL Server – 通过 Web 一键唤醒服务器

一个轻量级的 Web 服务，提供友好的按钮界面，通过 HTTP 请求触发 Wake‑on‑LAN（WOL）魔术包，远程唤醒局域网内的设备。  
后端使用 Go 编写，前端为纯静态 HTML，适合部署在树莓派、PiKVM 等 ARMv7l 设备上。

## 功能特性

- 🌐 **Web 操作界面** – 浏览器访问即用，点击按钮发送唤醒信号  
- 🚀 **轻量高效** – 单个二进制文件，资源占用极低  
- 🔧 **配置灵活** – 通过环境变量设置 MAC 地址、唤醒命令和端口，或在go源文件中指定 MAC 地址、唤醒命令和端口
- 🖥️ **跨平台支持** – 可编译为 x86、ARM 等多种架构  
- 🛡️ **安全可控** – 支持搭配 Nginx 增加密码保护或限制 IP 访问  

## 文件结构

```
wol/
├── main.go           # Go 后端源码
├── static/
│── LICENSE           # LICENSE
│   └── index.html    # 前端页面（按钮）
└── README.md         # 本文件
```

编译后生成 `wol-server` 可执行文件，与 `static/` 目录一起部署。

## 部署步骤（以 PiKVM / ARMv7l 为例）

### 1. 自行编译（任意平台）
在部署设备上编译：
```cmd
go build -o wol-server main.go
```

### 2. 赋予执行权限
```bash
cd /opt/wol
chmod +x wol-server
```

### 3. 测试运行
```bash
./wol-server
```
访问 `http://设备IP:11451/`，看到红色按钮即为成功。按 `Ctrl+C` 停止。

### 4. 配置 systemd 服务（开机自启）
将wol.service放入`/etc/systemd/system/`
启用并启动：
```bash
systemctl daemon-reload
systemctl enable wol
systemctl start wol
systemctl status wol
```

### 6. 防火墙放行端口（PiKVM 使用 iptables）
```bash
iptables -A INPUT -p tcp --dport 11451 -j ACCEPT
iptables-save > /etc/iptables/iptables.rules
```
若使用 UFW（Ubuntu/Debian）：
```bash
ufw allow 11451/tcp
```

## 环境变量说明

| 变量 | 说明 | 默认值 |
|------|------|--------|
| `MAC_ADDRESS` | 目标设备的 MAC 地址（格式 `xx:xx:xx:xx:xx:xx`） | `34:5A:60:AD:C0:85` |
| `WAKE_CMD` | 唤醒命令（如 `wakeonlan`、`etherwake`、`wakeonline`） | `wakeonline` |
| `PORT` | 服务监听端口 | `11451` |

**许可证**：Apache License  
**作者**：CregChain  
**反馈**：欢迎提交 Issue 或 Pull Request

# English:

# A-Modular-Wake-On-Lan-Strategy
A modular WOL strategy that use a linux or any device to start another specific device in your local network.
# WOL Server – Wake Up Servers with One Click via Web

A lightweight web service that provides a friendly button interface to trigger Wake‑on‑LAN (WOL) magic packets via HTTP requests, remotely waking up devices on your local network.  
The backend is written in Go, and the frontend is pure static HTML, suitable for deployment on ARMv7l devices such as Raspberry Pi or PiKVM.

## Features

- 🌐 **Web Operation Interface** – Access via browser, click the button to send the wake-up signal  
- 🚀 **Lightweight & Efficient** – Single binary, extremely low resource usage  
- 🔧 **Flexible Configuration** – Set MAC address, wake-up command, and port via environment variables, or specify them directly in the Go source file  
- 🖥️ **Cross-Platform Support** – Can be compiled for x86, ARM, and other architectures  
- 🛡️ **Secure & Controllable** – Can be paired with Nginx for password protection or IP restriction  

## File Structure

```
wol/
├── main.go           # Go backend source code
├── static/
│── LICENSE           # LICENSE
│   └── index.html    # Frontend page (button)
└── README.md         # This file
```

After compilation, the `wol-server` executable is generated and deployed together with the `static/` directory.

## Deployment Steps (e.g., on PiKVM / ARMv7l)

### 1. Build Yourself (on any platform)
Compile on the target device:
```cmd
go build -o wol-server main.go
```

### 2. Grant Execute Permissions
```bash
cd /opt/wol
chmod +x wol-server
```

### 3. Test Run
```bash
./wol-server
```
Visit `http://device-ip:11451/` in your browser; you should see a red button indicating success. Press `Ctrl+C` to stop.

### 4. Configure systemd Service (Auto-start on Boot)
Place `wol.service` into `/etc/systemd/system/`
Enable and start:
```bash
systemctl daemon-reload
systemctl enable wol
systemctl start wol
systemctl status wol
```

### 6. Allow Port in Firewall (PiKVM uses iptables)
```bash
iptables -A INPUT -p tcp --dport 11451 -j ACCEPT
iptables-save > /etc/iptables/iptables.rules
```
If using UFW (Ubuntu/Debian):
```bash
ufw allow 11451/tcp
```

## Environment Variables

| Variable | Description | Default Value |
|----------|-------------|---------------|
| `MAC_ADDRESS` | MAC address of the target device (format `xx:xx:xx:xx:xx:xx`) | `34:5A:60:AD:C0:85` |
| `WAKE_CMD` | Wake-up command (e.g., `wakeonlan`, `etherwake`, `wakeonline`) | `wakeonline` |
| `PORT` | Service listening port | `11451` |

**License**: Apache License  
**Author**: CregChain  
**Feedback**: Issues and Pull Requests are welcome