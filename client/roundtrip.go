package client

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type loggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

type autorisationRoundTripper struct {
	next http.RoundTripper
}

func (l loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "[%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL)
	return l.next.RoundTrip(r)
}

func (a autorisationRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {

	adminName, adminPassword := Autorisation()

	r.Header.Add("admin-name", adminName)
	r.Header.Add("admin-password", adminPassword)

	return a.next.RoundTrip(r)
}
