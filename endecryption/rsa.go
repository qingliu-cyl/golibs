// Package endecryption
// @Description: rsa加解密
package endecryption

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"math/big"
)

const defaultExponent = 65537

type RSA struct {
	modules string
	exponent int
	pubKey []byte
}

// SetModules
//  @Description: 设置模数， 用于解密与转换为公钥
//  @receiver r
//  @param m 模数
func (r *RSA) SetModules(m string) {
	r.modules = m
}

// SetExponent
//  @Description: 设置指数， 用于解密与转换为公钥
//  @receiver r
//  @param e
func (r *RSA) SetExponent(e int) {
	r.exponent = e
}

// PubKey2String
//  @Description: 将公钥转换为string
//  @receiver r
//  @return string
func (r *RSA) PubKey2String() string {
	return string(r.pubKey)
}

// Modules2Key
//  @Description: 将模数转换为公钥
//  @receiver r
//  @return []byte
//  @return error
func (r *RSA) Modules2Key() ([]byte, error) {
	if r.exponent == 0 {
		r.exponent = defaultExponent
	}

	bigN := new(big.Int)
	_, ok := bigN.SetString(r.modules, 16)
	if !ok {
		return nil, errors.New("failed convert modeles to string")
	}
	pub := rsa.PublicKey{
		N: bigN,
		E: r.exponent,
	}

	derStream,err := x509.MarshalPKIXPublicKey(&pub)
	if err != nil{
		return nil, errors.New("failed to marshal pub key")
	}
	block := &pem.Block{
		Type:"PUBLIC KEY",
		Bytes:derStream,
	}

	r.pubKey = pem.EncodeToMemory(block)
	return r.pubKey, nil
}
