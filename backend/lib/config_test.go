package backend

import (
	"os"
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	got := NewConfig()

	want := &Config{
		EnvPrefix:       "APP_",
		ApiEntrypoint:   "http://localhost:8000",
		ServerBindAddr:  "127.0.0.1",
		ServerBindPort:  8080,
		MetricsBindAddr: "127.0.0.1",
		MetricsBindPort: 2112,
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}

func TestConfigPopulateFromEnv(t *testing.T) {

	prefix := "FOO_"

	envmap := map[string]string{
		prefix + "API_ENTRYPOINT":    "http://192.168.1.10:8080",
		prefix + "SERVER_BIND_ADDR":  "10.0.0.0",
		prefix + "SERVER_BIND_PORT":  "10000",
		prefix + "METRICS_BIND_ADDR": "192.168.0.0",
		prefix + "METRICS_BIND_PORT": "20000",
	}

	defer func() {
		for key := range envmap {
			if err := os.Unsetenv(key); err != nil {
				t.Errorf("error os.Unsetenv key='%v'", key)
			}
		}
	}()

	got := NewConfig()
	got.EnvPrefix = prefix

	want := &Config{
		EnvPrefix:       prefix,
		ApiEntrypoint:   "http://192.168.1.10:8080",
		ServerBindAddr:  "10.0.0.0",
		ServerBindPort:  10000,
		MetricsBindAddr: "192.168.0.0",
		MetricsBindPort: 20000,
	}

	for key, val := range envmap {
		if err := os.Setenv(key, val); err != nil {
			t.Errorf("error os.Setenv key='%v' value='%v'", key, val)
		}
	}

	got.PopulateFromEnv()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}
