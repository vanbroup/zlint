// lint_ext_san_uri_host_not_fqdn_or_ip.go
/*********************************************************************
When the subjectAltName extension contains a URI, the name MUST be
stored in the uniformResourceIdentifier (an IA5String).  The name
MUST NOT be a relative URI, and it MUST follow the URI syntax and
encoding rules specified in [RFC3986].  The name MUST include both a
scheme (e.g., "http" or "ftp") and a scheme-specific-part.  URIs that
include an authority ([RFC3986], Section 3.2) MUST include a fully
qualified domain name or IP address as the host.  Rules for encoding
Internationalized Resource Identifiers (IRIs) are specified in
Section 7.4.
*********************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type sanUriHost struct {
	// Internal data here
}

func (l *sanUriHost) Initialize() error {
	return nil
}

func (l *sanUriHost) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SanOID)
}

func (l *sanUriHost) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, uri := range c.URIs {
		auth := util.GetAuthority(uri)
		if auth != "" {
			host := util.GetHost(auth)
			if !util.AuthIsFqdnOrIp(host) {
				return ResultStruct{Result: Error}, nil
			}
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "ext_san_uri_host_not_fqdn_or_ip",
		Description:   "URIs that include an authority ([RFC3986], Section 3.2) MUST include a fully qualified domain name or IP address as the host.",
		Providence:    "RFC 5280: 4.2.1.7",
		EffectiveDate: util.RFC5280Date,
		Test:          &sanUriHost{}})
}
