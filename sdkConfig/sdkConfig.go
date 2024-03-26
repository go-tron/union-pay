package sdkConfig

import (
	"crypto/rsa"
	"github.com/go-tron/crypto/rsaUtil"
	"strings"
)

type SDKConfig struct {
	SignCertificate     map[string]string
	SignPublicKey       map[string]*rsa.PublicKey
	EncryptCertId       string
	EncryptPublicKeyPem string
	EncryptPublicKey    *rsa.PublicKey
	FrontTransPath      string
	BackTransPath       string
	BatchTransPath      string
	SingleQueryPath     string
	CardRequestPath     string
	AppRequestPath      string
	FileQueryPath       string
	IssBackTransPath    string
}

func New(config *SDKConfig) *SDKConfig {

	config.SignPublicKey = make(map[string]*rsa.PublicKey)
	for key, val := range config.SignCertificate {
		publicKey, err := rsaUtil.GetPublicKeyFromCertificate([]byte(val))
		if err != nil {
			panic(err)
		}
		config.SignPublicKey[key] = publicKey
		config.SignCertificate[key] = strings.ReplaceAll(val, "\r", "")
		config.SignCertificate[key] = strings.ReplaceAll(val, "\n", "")
	}

	encryptPublicKey, err := rsaUtil.GetPublicKeyPem([]byte(config.EncryptPublicKeyPem))
	if err != nil {
		panic(err)
	}
	config.EncryptPublicKey = encryptPublicKey

	return config
}
