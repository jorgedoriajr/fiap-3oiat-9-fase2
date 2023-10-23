package sql

import (
	"testing"
)

func Test_build_url_conn(t *testing.T) {

	configs := []struct {
		title       string
		config      Config
		expectedUrl string
	}{
		{"Given config with schema should return correct url",
			Config{
				Host:         "localhost",
				Port:         5432,
				DatabaseName: "database",
				User:         "card",
				Password:     "card",
				MaxPoolSize:  1,
				Schema:       "card_schema",
			}, "postgres://card:card@localhost:5432/database?pool_max_conns=1&search_path=card_schema"},
		{"Given config without schema should return correct url",
			Config{
				Host:         "localhost",
				Port:         5432,
				DatabaseName: "database",
				User:         "card",
				Password:     "card",
				MaxPoolSize:  1,
			}, "postgres://card:card@localhost:5432/database?pool_max_conns=1"},
	}

	for _, c := range configs {
		t.Log(c.title)
		url, _ := getConnString(c.config)
		if url != c.expectedUrl {
			t.Errorf("Url incorreta %s", url)
		}
	}
}
