# SSH HTTP Proxy

This is a Go program that creates an HTTP proxy server using SSH to connect to another server.

## Usage

To use the program, run it with the following command-line arguments:

```shell
ssh-http-proxy -ssh_addr <SSH server address> -ssh_user <SSH server user> -ssh_cert <SSH server certificate path> -proxy_port <HTTP proxy server port>
```

For example:

```shell
ssh-http-proxy -ssh_addr example.com:22 -ssh_user root -ssh_cert /path/to/pem -proxy_port 8080
```

This command starts the HTTP proxy server on port 8080 and uses the specified SSH server, user, and certificate to create a tunnel for incoming HTTP requests.

# Preparation

## Generating a PEM Key File

If you are using Alibaba Cloud, follow these steps to generate a PEM key file:

1. Go to ECS -> Key Pair -> Create Key Pair.
2. Download the PEM key file.
3. Bind the key pair -> Select instance.
4. Restart ECS.

## Installing the Golang Environment

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

# Client Usage

For clients such as iPhone, configure the proxy settings by following these steps:

1. Go to Settings -> Wi-Fi.
2. Tap the "i" icon next to the connected Wi-Fi network.
3. Scroll down and tap on "Configure Proxy".
4. Select "Manual" and enter the IP and port number for the proxy server.
5. Tap "Save" to apply the changes.
