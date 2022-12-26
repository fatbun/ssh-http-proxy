package main

import (
	"flag"
	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

type Config struct {
	SshAddr    string `yaml:"ssh_addr"`
	SshUser    string `yaml:"ssh_user"`
	SshCert    string `yaml:"ssh_cert"`
	SshTimeout int    `yaml:"ssh_timeout"`
	ProxyPort  int    `yaml:"proxy_port"`
}

func parseYaml(yamlFile string) *Config {
	bytes, err := os.ReadFile(yamlFile)
	if err != nil {
		log.Fatal(err)
	}
	var config Config
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("read config: %+v\n", config)
	return &config
}

func parseConfig() *Config {
	// Parse the command line arguments
	configFile := flag.String("config", "", "e.g. /path/to/config.yaml")
	sshAddr := flag.String("ssh_addr", "", "SSH server address, e.g. example.com:22")
	sshUser := flag.String("ssh_user", "root", "SSH server user")
	sshCertPath := flag.String("ssh_cert", "", "SSH server certificate path")
	proxyPort := flag.Int("proxy_port", 8080, "http proxy server port")
	sshTimeout := flag.Int("timeout", 2, "SSH client connection timeout in seconds")

	flag.Parse()

	config := new(Config)
	if *configFile != "" {
		config = parseYaml(*configFile)
	}
	if config.SshAddr == "" {
		config.SshAddr = *sshAddr
	}
	if config.SshUser == "" {
		config.SshUser = *sshUser
	}
	if config.SshCert == "" {
		config.SshCert = *sshCertPath
	}
	if config.SshTimeout == 0 {
		config.SshTimeout = *sshTimeout
	}
	if config.ProxyPort == 0 {
		config.ProxyPort = *proxyPort
	}
	return config
}

func createSshConfig(config *Config) *ssh.ClientConfig {
	// Read the SSH certificate
	cert, err := os.ReadFile(config.SshCert)
	if err != nil {
		log.Fatal("error reading SSH certificate: ", err)
	}

	// 解析pem证书
	key, err := ssh.ParsePrivateKey(cert)
	if err != nil {
		log.Fatal(err)
	}

	// Dial the SSH server
	sshConf := &ssh.ClientConfig{
		User: config.SshUser,
		Auth: []ssh.AuthMethod{
			// 设置pem证书作为认证方式
			ssh.PublicKeys(key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Duration(config.SshTimeout) * time.Second,
	}
	return sshConf
}
