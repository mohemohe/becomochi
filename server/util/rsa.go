package util

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
)

func GenerateKeyPair() (privatePem string, publicPem string, lastErr error) {
	key, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return "", "", err
	}

	privateKey := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	privatePemBytes := new(bytes.Buffer)
	if err := pem.Encode(privatePemBytes, privateKey); err != nil {
		return "", "", err
	}

	asn1Bytes, err := asn1.Marshal(key.PublicKey)
	if err != nil {
		return "", "", err
	}
	publicKey := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}
	publicPemBytes := new(bytes.Buffer)
	if err := pem.Encode(publicPemBytes, publicKey); err != nil {
		return "", "", err
	}

	privatePem = privatePemBytes.String()
	publicPem = publicPemBytes.String()
	return
}