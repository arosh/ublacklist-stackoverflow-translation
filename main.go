package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"

	"golang.org/x/xerrors"
	"sigs.k8s.io/yaml"
)

type Domain struct {
	Domain   string `json:"domain"`
	Evidence string `json:"evidence"`
	Original string `json:"original"`
	Note     string `json:"note"`
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}

func run() error {
	filename := "domain-list.yml"
	domains, err := loadDomains(filename)
	if err != nil {
		return xerrors.Errorf("failed to load %s: %w", filename, err)
	}
	ctx := context.Background()
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.MaxConnsPerHost = 1
	h := &http.Client{
		Timeout:   5 * time.Second,
		Transport: tr,
	}
	for _, domain := range domains {
		if err := verify(ctx, h, domain, domains); err != nil {
			fmt.Fprintf(os.Stderr, "%+v\n", err)
		}
	}
	return nil
}

func isHttpTimeout(err error) bool {
	for err != nil {
		if err.Error() == "context deadline exceeded (Client.Timeout exceeded while awaiting headers)" {
			return true
		}
		err = xerrors.Unwrap(err)
	}
	return false
}

func matchAny(hostname string, domains []Domain) (bool, error) {
	for _, domain := range domains {
		matched, err := path.Match(domain.Domain, hostname)
		if err != nil {
			return false, xerrors.Errorf("pattern = %s: %w", domain.Domain, err)
		}
		if matched {
			return true, nil
		}
	}
	return false, nil
}

func verify(ctx context.Context, h *http.Client, domain Domain, domains []Domain) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, domain.Evidence, nil)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36")
	resp, err := h.Do(req)
	if err != nil {
		if isHttpTimeout(err) {
			fmt.Printf("timeout %s\n", domain.Domain)
			return nil
		} else {
			return xerrors.Errorf("%w", err)
		}
	}
	defer func() {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}()
	matched, err := path.Match(domain.Domain, resp.Request.URL.Hostname())
	if err != nil {
		return xerrors.Errorf("%w", err)
	}
	if !matched {
		matchedAny, err := matchAny(resp.Request.URL.Hostname(), domains)
		if err != nil {
			return xerrors.Errorf("%w", err)
		}
		if !matchedAny {
			fmt.Printf("%d %s %s\n", resp.StatusCode, domain.Domain, resp.Request.URL.Hostname())
		}
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("%d %s %s\n", resp.StatusCode, domain.Domain, resp.Request.URL.Hostname())
	}
	return nil
}

func loadDomains(filename string) ([]Domain, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, xerrors.Errorf("failed to open: %w", err)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, xerrors.Errorf("failed to read: %w", err)
	}

	domains := []Domain{}
	if err := yaml.UnmarshalStrict(b, &domains); err != nil {
		return nil, xerrors.Errorf("failed to unmarshal: %w", err)
	}

	return domains, nil
}
