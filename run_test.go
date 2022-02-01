package main

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"path"
	"testing"

	"golang.org/x/xerrors"
	"sigs.k8s.io/yaml"
)

type Item struct {
	Domain   string `json:"domain"`
	Evidence string `json:"evidence"`
	Original string `json:"original"`
	Note     string `json:"note"`
}

func TestDomainList(t *testing.T) {
	filename := os.Getenv("YAML")
	if len(filename) == 0 {
		filename = "domain-list.yml"
	}
	tests, err := loadItems(filename)
	if err != nil {
		t.Fatal(err)
	}
	for i, tt := range tests {
		i := i
		tt := tt
		name := fmt.Sprintf("%d-th item %s", i+1, tt.Domain)
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// domain, evidence, and original should not be empty
			if len(tt.Domain) == 0 {
				t.Errorf("domain is empty")
			}
			if len(tt.Original) == 0 {
				t.Errorf("original is empty")
			}
			if len(tt.Evidence) == 0 {
				t.Errorf("evidence is empty")
			}

			// domain should match evidence
			u, err := url.Parse(tt.Evidence)
			if err != nil {
				t.Errorf("cannot parse evidence: %+v", err)
			}

			if matched, err := path.Match(tt.Domain, u.Hostname()); err != nil {
				t.Errorf("cannot check match: %+v", err)
			} else if !matched {
				t.Errorf("'%s' does not match '%s'", tt.Domain, u.Hostname())
			}

			// domain should not match Stack Exchange
			if tt.Domain == "*.stackexchange.com" || tt.Domain == "stackexchange.com" {
				t.Errorf("'%s' should not match Stack Exchange", tt.Domain)
			}
			if matched, err := path.Match("*.stackexchange.com", tt.Domain); err != nil {
				t.Errorf("cannot check match: %+v", err)
			} else if matched {
				t.Errorf("'%s' should not match Stack Exchange", tt.Domain)
			}

			stackExchanges := []string{
				"askubuntu.com",
				"stackoverflow.com",
				"superuser.com",
			}
			for _, stackExchange := range stackExchanges {
				if matched, err := path.Match(tt.Domain, stackExchange); err != nil {
					t.Errorf("cannot check match: %+v", err)
				} else if matched {
					t.Errorf("'%s' should not match Stack Exchange", tt.Domain)
				}
			}

			// domain should not duplicate others
			for j, other := range tests {
				if i == j {
					continue
				}
				if matched, err := path.Match(tt.Domain, other.Domain); err != nil {
					t.Errorf("cannot check match: %+v", err)
				} else if matched {
					t.Errorf("%d-th item ('%s') and %d-th item ('%s') are duplicated", i+1, tt.Domain, j+1, other.Domain)
				}
			}
		})
	}
}

func loadItems(filename string) ([]Item, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, xerrors.Errorf("failed to open: %w", err)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, xerrors.Errorf("failed to read: %w", err)
	}

	domains := []Item{}
	if err := yaml.UnmarshalStrict(b, &domains); err != nil {
		return nil, xerrors.Errorf("failed to unmarshal: %w", err)
	}

	return domains, nil
}
