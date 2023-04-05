# SSH HTTP Proxy: Setup and Usage Guide

In this guide, we will learn how to set up and use an SSH HTTP proxy. An SSH HTTP proxy allows you to route your internet traffic through an SSH tunnel, encrypting your data and bypassing network restrictions.

## Prerequisites

To follow this guide, you will need:

1. A remote server with SSH access.
2. An SSH key pair (public and private keys) for authentication.

## Setting up the SSH HTTP Proxy

### Step 1: Install the SSH HTTP Proxy

First, you need to download and install the `ssh-http-proxy` tool on your local machine. You can find the installation instructions on the [official GitHub repository](https://github.com/qfrank/ssh-http-proxy).

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
cd bin
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
