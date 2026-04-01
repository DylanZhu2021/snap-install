# Snap Install Web Service

一个用于通过Web界面下载和安装Snap包的Go语言服务。

## 功能

- Web界面输入HTTP URL下载.snap文件
- 自动执行`snap install XXX.snap --dangerous`安装命令
- 支持开机自启动
- 长时间稳定运行

## 技术规格

- 运行端口: 8347
- 运行地址: 127.0.0.1
- 下载目录: /tmp/snap-downloads

## 安装步骤

1. 确保系统已安装Go语言环境和snap

2. 运行安装脚本:
```bash
chmod +x install.sh
./install.sh
```

3. 访问服务:
```
http://127.0.0.1:8347
```

## 服务管理

```bash
# 查看服务状态
sudo systemctl status snap-install.service

# 停止服务
sudo systemctl stop snap-install.service

# 启动服务
sudo systemctl start snap-install.service

# 重启服务
sudo systemctl restart snap-install.service

# 查看日志
sudo journalctl -u snap-install.service -f

# 禁用开机自启
sudo systemctl disable snap-install.service
```

## 使用方法

1. 打开浏览器访问 http://127.0.0.1:8347
2. 在输入框中输入.snap文件的URL
3. 点击"Download and Install"按钮
4. 等待下载和安装完成

## 注意事项

- 需要root权限运行（因为snap install需要root权限）
- 确保输入的URL指向有效的.snap文件
- 下载的文件会保存在/home/admin/snap-downloads目录
