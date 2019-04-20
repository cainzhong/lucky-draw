package lucky_drawn

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

var insecure *bool

func init() {
	insecure = flag.Bool("insecure-ssl", false, "Accept/Ignore all server SSL certificates")
	flag.Parse()
}

func InitSystemCertPool(localCertFile string) *http.Client {
	// Get the SystemCertPool, continue with an empty pool on error
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		logger.Warn("The current SystemCertPool is empty.")
		rootCAs = x509.NewCertPool()
	}

	if len(localCertFile) > 0 {
		// Read in the cert file
		certs, err := ioutil.ReadFile(localCertFile)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to append %q to RootCAs: %v", localCertFile, err))
		}
		// Append the local cert file to the system pool
		if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
			logger.Info("No certs appended, using system certs only")
		}
		logger.Info(fmt.Sprintf("Append the local cert file %s to the system pool", localCertFile))
	}

	// Trust the augmented cert pool in our client
	config := &tls.Config{
		InsecureSkipVerify: *insecure,
		RootCAs:            rootCAs,
	}
	cookieJar, _ := cookiejar.New(nil)
	tr := &http.Transport{TLSClientConfig: config}
	client := &http.Client{
		Jar:       cookieJar,
		Transport: tr}
	return client
}
