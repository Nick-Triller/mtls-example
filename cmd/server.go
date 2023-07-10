package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	commonName := r.TLS.PeerCertificates[0].Subject.CommonName
	io.WriteString(w, fmt.Sprintf("Hello, %s!\n", commonName))
}

func main() {
	// Set up a /hello resource handler
	log.Print("Starting server")
	http.HandleFunc("/hello", helloHandler)

	// Read the client CA certificate that will be used to verify client certificates
	clientCaCert, err := os.ReadFile("certs/client-ca.crt")
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(clientCaCert)

	// Create the TLS Config with the CA pool and enable Client certificate validation
	tlsConfig := &tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}

	// Create a Server instance to listen on port 8443 with the TLS config
	server := &http.Server{
		Addr:      "localhost:8443",
		TLSConfig: tlsConfig,
	}

	// Start server
	log.Fatal(server.ListenAndServeTLS("certs/server.crt", "certs/server.key"))
}
