package config

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"
)

const (
	defaultAddr   = ":80"
	defaultUseSSL = false
)

type ServerConfigBuilder struct {
	Addr          string
	UseSSL        bool
	CertChainPath string
	CertKeyPath   string
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
}

func (b *ServerConfigBuilder) WithAddr(addr string) *ServerConfigBuilder {
	log.Printf("set addr: %s", addr)
	b.Addr = addr
	return b
}

func (b *ServerConfigBuilder) WithSSL(certPath, keyPath string) *ServerConfigBuilder {
	log.Printf("set ssl config: cert %s, key %s", certPath, keyPath)
	b.UseSSL = true
	b.CertChainPath = certPath
	b.CertKeyPath = keyPath
	return b
}

func (b *ServerConfigBuilder) WithTimeout(readTimeout, writeTimeout time.Duration) *ServerConfigBuilder {
	log.Printf("set timeout: read %v, write %v", readTimeout, writeTimeout)
	b.ReadTimeout = readTimeout
	b.WriteTimeout = writeTimeout
	return b
}

func (b *ServerConfigBuilder) Build() *http.Server {
	srv := &http.Server{
		Addr:         defaultAddr,
		ReadTimeout:  b.ReadTimeout,
		WriteTimeout: b.WriteTimeout,
	}

	if b.Addr != "" {
		srv.Addr = b.Addr
	}

	if b.UseSSL {
		cert, err := tls.LoadX509KeyPair(
			b.CertChainPath,
			b.CertKeyPath,
		)
		if err != nil {
			log.Fatalf("load certificate failed: %v", err)
		}

		srv.TLSConfig = &tls.Config{
			Certificates: []tls.Certificate{cert},
		}
	}

	return srv
}
