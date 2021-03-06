// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

// Package aes implements AES encryption (formerly Rijndael), as defined in
// U.S. Federal Information Processing Standards Publication 197.

// aes包实现了AES加密算法，参见U.S. Federal Information Processing Standards
// Publication 197。
package aes

import (
	"crypto/cipher"
	"strconv"
)

// The AES block size in bytes.

// AES字节块大小。
const BlockSize = 16

type KeySizeError int

// NewCipher creates and returns a new cipher.Block.
// The key argument should be the AES key,
// either 16, 24, or 32 bytes to select
// AES-128, AES-192, or AES-256.

// 创建一个cipher.Block接口。参数key为密钥，长度只能是16、24、32字节，用以选择
// AES-128、AES-192、AES-256。
func NewCipher(key []byte) (cipher.Block, error)

func (k KeySizeError) Error() string

