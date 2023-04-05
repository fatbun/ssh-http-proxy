## Setting up the SSH HTTP Proxy

### Step 1: Install the SSH HTTP Proxy

First, you need to download and install the `ssh-http-proxy` tool on your local machine. Since it's a Golang repository, follow these steps to install the tool:

1. Make sure you have [Golang](https://golang.org/dl/) installed on your local machine.
2. Open a terminal or command prompt and run the following command to download and install the `ssh-http-proxy` tool:

```bash
go get -u github.com/qfrank/ssh-http-proxy
```

3. This command will install the `ssh-http-proxy` binary in your `$GOPATH/bin` directory. Make sure the `$GOPATH/bin` directory is included in your system's `PATH` environment variable.

### Step 2: Configure the SSH HTTP Proxy

Before starting the service, you need to configure the proxy with the SSH server details. Create a configuration file named `config.yml` in the `ssh-http-proxy` directory with the following content:

```yaml
ssh_server:
  addr: example.com:22
  user: root
  cert: /path/to/your/pem

proxy:
  port: 8080
```

Replace `example.com`, `root`, and `/path/to/your/pem` with the correct values for your SSH server, user, and certificate.

### Step 3: Start the SSH HTTP Proxy Service

To start the SSH HTTP proxy service, navigate to the `bin` directory and run the following command:

```bash
ssh-http-proxy -config=config.yml
```

This command starts the HTTP proxy server on port 8080 and uses the specified SSH server, user, and certificate to create a tunnel for incoming HTTP requests.

## Client Usage

Once the service is started, clients can connect to the proxy server. The following example demonstrates how to configure an iPhone to use the proxy server:

1. Go to Settings -> Wi-Fi.
2. Tap the "i" icon next to the connected Wi-Fi network.
3. Scroll down and tap on "Configure Proxy".
4. Select "Manual" and enter the IP and port number for the proxy server (in this case, port 8080).
5. Tap "Save" to apply the changes.

Remember to replace the IP and port number with the correct values for your specific setup.

Now, you have successfully set up and started using the SSH HTTP proxy. All your internet traffic from the configured device will be routed through the SSH tunnel, providing you with a secure and private browsing experience.
