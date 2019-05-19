// Package getter will fetch URLs over HTTP and return the Content-Length of the response.
package getter

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Getter struct {
	client *http.Client
}

// New returns a new Getter with the HTTP client timeout set to the given number of seconds.
func New(tSec uint) *Getter {
	g := &Getter{}

	g.client = &http.Client{
		Timeout: time.Duration(tSec) * time.Second,
	}

	return g
}

func (g *Getter) Get(u string) uint {
	url, errParse := url.Parse(u)
	if errParse != nil {
		log.Fatalf("error parsing URL '%s': %v", u, errParse)
	}

	req, errReq := http.NewRequest(http.MethodHead, url.String(), nil)
	if errReq != nil {
		log.Fatalf("error creating request: %v", errReq)
	}

	resp, errResp := g.client.Do(req)
	if errResp != nil {
		// In the explicit case where the error was a timeout, return cleanly.
		if strings.Contains(errResp.Error(), "(Client.Timeout exceeded while awaiting headers)") {
			return 0
		}
		log.Fatalf("error performing request: %v", errResp)
	}

	rawCL := resp.Header.Get("Content-Length")
	if rawCL == "" {
		return 0
	}
	nCL, errConv := strconv.Atoi(rawCL)
	if errConv != nil {
		log.Fatalf("error converting value '%v' of 'Content-Length' header: %v", rawCL, errConv)
	}

	// Some implementations go against the HTTP spec and return a negative value:
	// https://tools.ietf.org/html/rfc7230#section-3.3.2
	if nCL < 0 {
		return 0
	}

	return uint(nCL)
}
