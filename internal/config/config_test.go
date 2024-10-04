package config

import (
	"testing"
)

func TestRead(t *testing.T) {
	config, err := Read()
	if err != nil {
		t.Errorf("trying to read: %v", err)
	}

	if config.Db_url != "postgres://example" {
		t.Errorf("db_url doesn't match up. must be 'postgres://example' but it is %s", config.Db_url)
	}

	test_user := "test"
	err = config.SetUser(test_user)
	if err != nil {
		t.Errorf("trying to set user: %v", err)
	}

	if config.Current_user_name != test_user {
		t.Errorf("user name doesn't match up. must be 'canko' but it is %s", config.Current_user_name)
	}

	if config.Db_url != "postgres://example" {
		t.Errorf("db_url doesn't match up. must be 'postgres://example' but it is %s", config.Db_url)
	}
}
