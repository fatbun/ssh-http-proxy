package main

import (
	"context"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type SshHttpProxy struct {
	sshConfig *ssh.ClientConfig
	sshClient *ssh.Client

	config  *Config
	sshDial func(network, addr string) (net.Conn, error)

	mutex                     sync.Mutex
	lastReCreateSshClientTime *time.Time
}

func (p *SshHttpProxy) reCreateSshClient() error {
	p.sshClient.Close()
	sshClient, err := p.dial()
	if err != nil {
		return err
	}
	*p.sshClient = *sshClient
	return nil
}

func (p *SshHttpProxy) Start() {
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(p.config.ProxyPort), p.createHandler()))
}

func (p *SshHttpProxy) createHandler() http.HandlerFunc {
	// Create a new HTTP handler that supports the CONNECT method
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("request", r.URL.String(), r.Proto, r.Method, r.URL.Path, r.UserAgent())
		if r.Method == http.MethodConnect {
			// Get the target Host and port
			host, port, err := net.SplitHostPort(r.URL.Host)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			var conn net.Conn
			//SshTimeout not working for sshDial in some cases, so use context to timeout
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(p.config.SshTimeout)*time.Second)
			go func() {
				defer cancel()
				conn, err = p.sshDial("tcp", net.JoinHostPort(host, port))
				if err != nil {
					p.mutex.Lock()
					defer p.mutex.Unlock()
					if p.lastReCreateSshClientTime == nil || p.lastReCreateSshClientTime.Before(time.Now().Add(-5*time.Second)) {
						now := time.Now()
						p.lastReCreateSshClientTime = &now
						err = p.reCreateSshClient()
						if err != nil {
							log.Println("failed to recreate ssh client", err.Error())
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
						}
						log.Println("recreate ssh client successfully")
					} else {
						log.Println("not the right time to recreate ssh client", err.Error())
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					conn, err = p.sshDial("tcp", net.JoinHostPort(host, port))
					if err != nil {
						log.Println("maybe we get some problems to connect to ssh server?", err.Error())
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
				}
			}()

			select {
			case <-ctx.Done():
				if ctx.Err() == context.DeadlineExceeded {
					http.Error(w, "timeout", http.StatusRequestTimeout)
					return
				}
			}

			if conn == nil {
				http.Error(w, "conn is nil", http.StatusInternalServerError)
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

			// Copy data between the client and the target Host
			go func() {
				defer conn.Close()
				io.Copy(conn, clientConn)
			}()
			io.Copy(clientConn, conn)
		}
	}
}

func (p *SshHttpProxy) dial() (*ssh.Client, error) {
	return ssh.Dial("tcp", p.config.SshAddr, p.sshConfig)
}

func NewSshHttpProxy(config *Config) *SshHttpProxy {
	sshHttpProxy := &SshHttpProxy{
		config: config,
	}
	sshHttpProxy.sshConfig = createSshConfig(config)
	sshClient, err := sshHttpProxy.dial()
	if err != nil {
		log.Fatal("error tunnel to server: ", err)
	}
	sshHttpProxy.sshClient = sshClient

	// Create a Dial function that uses the ssh dialer
	sshHttpProxy.sshDial = func(network, addr string) (net.Conn, error) {
		return sshHttpProxy.sshClient.Dial(network, addr)
	}

	return sshHttpProxy
}
