package utilities

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var ConnectConfiguration Configuration = ImportConfig()
var ConnectClient *http.Client = createClient(ConnectConfiguration)

func ImportConfig() Configuration {
	fmt.Println("I am importing the configuration file") // control statement print - TOREMOVE
	file, err := os.Open(os.Getenv("CONNECTCFG"))        // previously used hardcoded ./connect-config.json
	if err != nil {
		fmt.Println("Please add the configuration file as an environment variable named CONNECTCFG")
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return configuration
}

func createTlsClient(capath string, certpath string, keypath string) *http.Client {

	// the CertPool wants to add a root as a []byte so we read the file ourselves
	caCert, err := ioutil.ReadFile(capath)
	if err != nil {
		log.Fatal(err)
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caCert)

	// LoadX509KeyPair reads files, so we give it the paths
	clientCert, err := tls.LoadX509KeyPair(certpath, keypath)
	if err != nil {
		log.Fatal(err)
	}

	tlsConfig := tls.Config{
		RootCAs:      pool,
		Certificates: []tls.Certificate{clientCert},
	}
	transport := http.Transport{
		TLSClientConfig: &tlsConfig,
	}
	client := http.Client{
		Transport: &transport,
	}

	return &client
}

func createClient(conf Configuration) *http.Client {
	fmt.Println("I am creating the connect client") // control statement print - TOREMOVE
	if conf.TlsEnable {
		return createTlsClient(conf.CaPath, conf.CertPath, conf.KeyPath)
	} else {
		return &http.Client{}
	}
}

func PrettyPrint(data []byte) {
	var prettyData bytes.Buffer
	json.Indent(&prettyData, data, "", "  ")
	fmt.Println(prettyData.String())
}