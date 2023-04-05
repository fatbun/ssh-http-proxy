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

# Operation Steps



## Generate PEM Key File

> Taking Alibaba Cloud as an example

- ECS -> Key Pair -> Create Key Pair
- Download the PEM key file
- Bind the key pair -> Select instance
- Restart ECS



## Install Golang Environment

- `brew install go`
- Modify mirror source

```bash
  echo "export GO111MODULE=on" >> ~/.zshrc
  echo "export GOPROXY=https://goproxy.cn,direct" >> ~/.zshrc
  source ./zshrc
```



## Clone Repository

```bash
gcl https://github.com/qfrank/ssh-http-proxy.git
```



## Compile Code

```bash
cd build
./build
```



## Start Service

```bash
cd bin
ssh-http-proxy -ssh_addr example.com:22 -ssh_user root -ssh_cert /path/to/pem -proxy_port 8080
```



# Usage

For clients like iPhone, go to Settings -> Wi-Fi -> Configure Proxy and enter IP and port number.
