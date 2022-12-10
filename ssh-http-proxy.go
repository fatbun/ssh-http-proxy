package main

import (
	"flag"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
	"time"
)

func main() {
	// Parse the command line arguments
	sshAddr := flag.String("ssh", "example.com:22", "SSH server address")
	sshUser := flag.String("user", "root", "SSH server user")
	sshCertPath := flag.String("cert", "/path/to/pem", "SSH server certificate path")
	httpPort := flag.Int("http", 8080, "HTTP proxy server port")
	sshTimeout := flag.Int("timeout", 5, "SSH client connection timeout in seconds")
	flag.Parse()

	// Read the SSH certificate
	cert, err := os.ReadFile(*sshCertPath)
	if err != nil {
		log.Fatal("error reading SSH certificate: ", err)
	}

	// 解析pem证书
	key, err := ssh.ParsePrivateKey(cert)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Dial the SSH server
	sshConf := &ssh.ClientConfig{
		User: *sshUser,
		Auth: []ssh.AuthMethod{
			// 设置pem证书作为认证方式
			ssh.PublicKeys(key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout: time.Duration(*sshTimeout) * time.Second,
	}
	sshConn, err := ssh.Dial("tcp", *sshAddr, sshConf)
	if err != nil {
		log.Fatal("error tunnel to server: ", err)
	}
	defer sshConn.Close()

	// Create a Dial function that uses the ssh dialer
	dial := func(network, addr string) (net.Conn, error) {
		return sshConn.Dial(network, addr)
	}

	// Create a new HTTP proxy
	httpProxy := &httputil.ReverseProxy{
		Director: func(r *http.Request) {
		},
		Transport: &http.Transport{
			Dial: dial,
		},
	}
	// Create a new HTTP handler that supports the CONNECT method
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodConnect {
			// Get the target host and port
			host, port, err := net.SplitHostPort(r.URL.Host)
			println("request host", host, ":", port)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Dial the target host and port
			conn, err := dial("tcp", net.JoinHostPort(host, port))
			if err != nil {
				http.Error(w, err.Error(), http.StatusServiceUnavailable)
				return
			}
			defer conn.Close()

			// Respond to the CONNECT request
			w.WriteHeader(http.StatusOK)
			if flusher, ok := w.(http.Flusher); ok {
				flusher.Flush()
			}

			// Hijack the connection
			hj, ok := w.(http.Hijacker)
			if !ok {
				http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
				return
			}
			clientConn, _, err := hj.Hijack()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer clientConn.Close()

			// Copy data between the client and the target host
			go func() {
				defer conn.Close()
				io.Copy(conn, clientConn)
			}()
			io.Copy(clientConn, conn)
		} else {
			// Use the HTTP proxy to handle the request
			httpProxy.ServeHTTP(w, r)
		}
	})

	// Start the HTTP proxy server
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*httpPort), handler))

}
