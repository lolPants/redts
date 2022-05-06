package edts

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/lolPants/redts/cli/src/pkg/config"
)

func callAPI(cfg *config.Config, script string, args string) error {
	url := fmt.Sprintf("%s/api/%s", cfg.URL, script)
	req, err := http.NewRequest("GET", url, strings.NewReader(args))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "text/plain")

	if cfg.Username != "" && cfg.Token == "" {
		fmt.Fprintln(os.Stderr, "error: missing auth token")
		os.Exit(2)
	}

	if cfg.Username != "" && cfg.Token != "" {
		bearer := fmt.Sprintf("Bearer %s:%s", cfg.Username, cfg.Token)
		req.Header.Add("Authorization", bearer)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		if cfg.EnsureAuthFields() {
			fmt.Fprintln(os.Stderr, "error: invalid credentials")
		}

		os.Exit(2)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", body)

	if resp.StatusCode != http.StatusOK {
		os.Exit(1)
	} else {
		os.Exit(0)
	}

	return nil
}
