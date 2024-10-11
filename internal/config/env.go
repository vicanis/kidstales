package config

import "os"

type envKey string
type envValue string

const (
	Addr        envKey = "ADDR"
	SSLEnabled  envKey = "SSL_ENABLED"
	SSLCertPath envKey = "SSL_CERT"
	SSLKeyPath  envKey = "SSL_KEY"
)

func Env(key envKey) envValue {
	return envValue(os.Getenv(string(key)))
}

func (v envValue) String() string {
	return string(v)
}

func (v envValue) Bool() bool {
	return string(v) == "true"
}
