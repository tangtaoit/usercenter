package main

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	jwt "github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var privateKey string ="-----BEGIN RSA PRIVATE KEY-----MIICXAIBAAKBgQDDZV7T3ZhIE03PxAN6CXMtDTKQ9vCBu0XrK0UD9UCPoaBJgy5LHz+z7+Pc+nCQT0AbvbroKveSdmOhz7zs6RulNjQQPUP8175Ci8p5n0jouAKV5s4K4BtjwuMzxKyZNHReJ1yTXQ3FjHpOzr0JzqBZ3S4EB7CBkQUb4Te0WfSf7QIDAQABAoGAFXaqLv21f51XO85dT2eAVl+PwWrOyoFm0clkAGZNXDm14L1fNXNOTRa54glEmiWKdkGmKWCm51jH4vtt1lxY4+CoEj3x2qnJ3SsqxE2gJiC2ASIxgpmkV1T14MEzwJO9LCVEzefHcEdtJwTs5o9CIiCFMH5mpw0OzvMTz/4d86ECQQD3ehzpmdA0xw/s/CBIs9Upe1qkmTCQVRUmyz8iWZPjVLsAjBwn0ZtDPXze5ln4pxopvtoqdbCimW2yuWXy+TS5AkEAyiAVdt3AQsqSIn0lzA7hPl7E6ACM9k54GR3ngmiUCgwMLn9GXNZRXzLY52nj2u1t1c9aqvEnfvGHo1qHqDzS1QJBALWYO4MGxQsVTxBc6euvWil4RMknR8WBSWYQGiHAjY5w7E+4gCiP3Fh41BpT+Y1GQSKE0134wkZuQ1q0RKUITLECQCZ8q3mhyd0t81uL1umfH7aflwDSMgUodefacN29CgtLtfoYlA5TZNUqunB+Ejv6n8JppEsOdkXOudQaBeC8DC0CQDmQQFXx36Ip4jTUj/CaiZTXpfbszSYr8zQ8PJUqC4oZr2N2wBMiJps4zIaQau6zOOYRegJx/waJL6p+y1dRzwo=-----END RSA PRIVATE KEY-----"

type JWTAuthenticationBackend struct {
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

const (
	tokenDuration = 72
	expireOffset  = 3600
)

var authBackendInstance *JWTAuthenticationBackend = nil

func InitJWTAuthenticationBackend() *JWTAuthenticationBackend {
	if authBackendInstance == nil {
		authBackendInstance = &JWTAuthenticationBackend{
			privateKey: getPrivateKey(),
			PublicKey:  getPublicKey(),
		}
	}

	return authBackendInstance
}

func (backend *JWTAuthenticationBackend) GenerateToken(openId string) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)
	token.Claims["exp"] = time.Now().Add(time.Hour * time.Duration(GetSetting().JWTExpirationDelta)).Unix()
	token.Claims["iat"] = time.Now().Unix()
	token.Claims["sub"] = openId
	tokenString, err := token.SignedString(backend.privateKey)
	if err != nil {
		panic(err)
		return "", err
	}
	return tokenString, nil
}



func (backend *JWTAuthenticationBackend) getTokenRemainingValidity(timestamp interface{}) int {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return int(remainer.Seconds() + expireOffset)
		}
	}
	return expireOffset
}




func getPrivateKey() *rsa.PrivateKey {


	privateKeyFile, err := os.Open(GetSetting().PrivateKeyPath)
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	privateKeyFile.Close()

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	return privateKeyImported
}

func getPublicKey() *rsa.PublicKey {
	publicKeyFile, err := os.Open(GetSetting().PublicKeyPath)
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := publicKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	publicKeyFile.Close()

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)

	if !ok {
		panic(err)
	}

	return rsaPub
}
