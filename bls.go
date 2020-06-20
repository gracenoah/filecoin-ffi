//+build cgo

package ffi

import (
)


// Hash computes the digest of a message
func Hash(message Message) Digest {
	panic("")
}

// Verify verifies that a signature is the aggregated signature of digests - pubkeys
func Verify(signature *Signature, digests []Digest, publicKeys []PublicKey) bool {
	panic("")
}

// HashVerify verifies that a signature is the aggregated signature of hashed messages.
func HashVerify(signature *Signature, messages []Message, publicKeys []PublicKey) bool {
	panic("")
}

// Aggregate aggregates signatures together into a new signature
func Aggregate(signatures []Signature) *Signature {
	panic("")
}

// PrivateKeyGenerate generates a private key
func PrivateKeyGenerate() PrivateKey {
	panic("")
}

// PrivateKeyGenerate generates a private key in a predictable manner
func PrivateKeyGenerateWithSeed(seed PrivateKeyGenSeed) PrivateKey {
	panic("")
}

// PrivateKeySign signs a message
func PrivateKeySign(privateKey PrivateKey, message Message) *Signature {
	panic("")
}

// PrivateKeyPublicKey gets the public key for a private key
func PrivateKeyPublicKey(privateKey PrivateKey) PublicKey {
	panic("")
}
