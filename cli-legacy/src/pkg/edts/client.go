package edts

import (
	"fmt"
	"net/http"
	"time"

	"github.com/lolPants/redts/cli/src/pkg/version"
)

var (
	client = &http.Client{
		Timeout:   time.Second * 120,
		Transport: agentTransport(getAgent()),
	}
)

func getAgent() string {
	app := "redts"
	ver := version.AppVersion()

	if version.IsDev() && version.CommitHash() != "unknown" {
		ver = version.CommitHash()[:7]
	}

	return fmt.Sprintf("%s/%s", app, ver)
}

func agentTransport(agent string) userAgentTransport {
	return userAgentTransport{
		agent: agent,
		rt:    http.DefaultTransport,
	}
}

type userAgentTransport struct {
	agent string
	rt    http.RoundTripper
}

func (uat userAgentTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("User-Agent", uat.agent)
	return uat.rt.RoundTrip(r)
}
