# SSH HTTP Proxy
This is a Go program that creates an HTTP proxy server that uses SSH to connect to another server.

## Usage
To use the program, run it with the following command-line arguments:
```shell
ssh-http-proxy -ssh_addr <SSH server address> -ssh_user <SSH server user> -ssh_cert <SSH server certificate path> -proxy_port <HTTP proxy server port>
```
For example:
```shell
ssh-http-proxy -ssh_addr example.com:22 -ssh_user root -ssh_cert /path/to/pem -proxy_port 8080
```
This will start the HTTP proxy server on port 8080, and it will use the specified SSH server, user, and certificate to create a tunnel for incoming HTTP requests.

# 操作步骤

## 生成pem秘钥文件

> 以阿里云为例

- ECS——》密钥对——》创建密钥对
- 下载pem秘钥文件

- 绑定密钥对——》选择实例
- 重启ecs

## 安装 golang 环境

- `brew install go`

- 修改镜像源

  ```bash
  echo "export GO111MODULE=on" >> ~/.zshrc
  echo "export GOPROXY=https://goproxy.cn,direct" >> ~/.zshrc
  source ./zshrc
  ```

## 克隆仓库

```bash
gcl https://github.com/qfrank/ssh-http-proxy.git
```

## 编译代码

```bash
cd build
./build
```

## 启动服务

```bash
cd bin
ssh-http-proxy -ssh_addr example.com:22 -ssh_user root -ssh_cert /path/to/pem -proxy_port 8080
```

# 使用

客户端如iPhone，在设置——》无线局域网——》配置代理，输入ip、端口即可。
