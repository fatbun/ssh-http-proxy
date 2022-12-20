package main

func main() {
	config := parseConfig()
	proxy := NewSshHttpProxy(config)
	proxy.Start()
}
