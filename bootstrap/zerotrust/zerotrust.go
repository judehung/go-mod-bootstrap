package zerotrust

import (
	"github.com/edgexfoundry/go-mod-bootstrap/v4/bootstrap/interfaces"
	"github.com/edgexfoundry/go-mod-bootstrap/v4/config"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/clients/logger"
	"net/http"
)

const (
	OpenZitiControllerKey = "OpenZitiController"
	ZeroTrustMode         = "zerotrust"
	OpenZitiServicePrefix = "edgex."
)

func HttpTransportFromService(secretProvider interfaces.SecretProviderExt, serviceInfo config.ServiceInfo, lc logger.LoggingClient) (http.RoundTripper, error) {
	roundTripper := http.DefaultTransport
	return roundTripper, nil
}

func HttpTransportFromClient(secretProvider interfaces.SecretProviderExt, clientInfo *config.ClientInfo, lc logger.LoggingClient) (http.RoundTripper, error) {
	roundTripper := http.DefaultTransport
	return roundTripper, nil
}
