package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Parse command line flags
	var clientCertFile string
	var clientKeyFile string
	flag.StringVar(&clientCertFile, "clientCertFile", "", "Path to client certificate")
	flag.StringVar(&clientKeyFile, "clientKeyFile", "", "Path to client certificate key")
	flag.Parse()

	if clientCertFile == "" || clientKeyFile == "" {
		log.Fatal("Missing required argument --clientCertFile or --clientKeyFile")
	}

	// Read CA that will be used to verify server cert
	serverCaCert, err := os.ReadFile("certs/server-ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(serverCaCert)

	// Read the client certificate that will be used for authentication
	clientCert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		log.Fatal(err)
	}

	// Create a HTTPS client and supply the created CA pool
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{clientCert},
			},
		},
	}

	r, err := client.Get("https://mtls.fbi.com:8443/hello")

	if err != nil {
		log.Fatal(err)
	}

	// Read the response body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response body to stdout
	fmt.Printf("%s", body)
}
