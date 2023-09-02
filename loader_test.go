package senv

import (
	"os"
	"strconv"
	"testing"
)

const testDBURL string = "url"
const testName string = "name"
const testCount int = 134

type TestCfg struct {
	ConnectionURL string `senv:"TEST_DB_URL" json:"connectionUrl" xdf:"lol"`
	Name          string `senv:"TEST_NAME" json:"name"`
	Count         int    `senv:"TEST_COUNT" json:"name"`
}

func TestLoader(t *testing.T) {
	os.Setenv("TEST_DB_URL", testDBURL)
	os.Setenv("TEST_NAME", testName)
	os.Setenv("TEST_COUNT", strconv.Itoa(testCount))
	cfg := TestCfg{}
	err := Load(cfg)
	if err != nil {
		t.Error(err)
		return
	}
	if cfg.ConnectionURL != testDBURL || cfg.Name != testName || cfg.Count != testCount {
		t.Error("cfg vales are incorrect")
	}
}
