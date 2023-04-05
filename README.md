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

# Preparation

## Generating PEM Key File

If you are using Alibaba Cloud, follow these steps:
- Go to ECS -> Key Pair -> Create Key Pair.
- Download the PEM key file.
- Bind the key pair -> Select instance.
- Restart ECS.


## Installing Golang Environment

To install Golang, run `brew install go` on the command line. After installation, modify the mirror source with the following commands:
```bash
  echo "export GO111MODULE=on" >> ~/.zshrc
  echo "export GOPROXY=https://goproxy.cn,direct" >> ~/.zshrc
  source ./zshrc
```


## Cloning the Repository

Clone the repository by running the following command:
```bash
git clone https://github.com/qfrank/ssh-http-proxy.git
```


## Compiling the Code

Navigate to the build directory and compile the code:
```bash
cd build
./build
```


## Starting the Service

Navigate to the bin directory and start the service:
```bash
cd bin
ssh-http-proxy -ssh_addr example.com:22 -ssh_user root -ssh_cert /path/to/pem -proxy_port 8080
```


# Usage

For clients such as iPhone, go to Settings -> Wi-Fi -> Configure Proxy and enter the IP and port number.
