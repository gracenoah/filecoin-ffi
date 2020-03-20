package ffi

// #cgo LDFLAGS: ${SRCDIR}/libfilcrypto.a
// #cgo pkg-config: ${SRCDIR}/filcrypto.pc
// #include "./filcrypto.h"
import "C"

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

// Aggregate aggregates signatures together into a new signature. If the
// provided signatures cannot be aggregated (due to invalid input or an
// an operational error), Aggregate will return nil.
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

// PrivateKeyGenerate generates a private key in a predictable manner
func PrivateKeyGenerateWithSeed(seed PrivateKeyGenSeed) PrivateKey {
	var ary generated.Fil32ByteArray
	copy(ary.Inner[:], seed[:])

	resp := generated.FilPrivateKeyGenerateWithSeed(ary)
	resp.Deref()
	resp.PrivateKey.Deref()
	defer generated.FilDestroyPrivateKeyGenerateResponse(resp)

	var out PrivateKey
	copy(out[:], resp.PrivateKey.Inner[:])
	return out
}

// PrivateKeySign signs a message
func PrivateKeySign(privateKey PrivateKey, message Message) *Signature {
	panic("")
}

// PrivateKeyPublicKey gets the public key for a private key
func PrivateKeyPublicKey(privateKey PrivateKey) PublicKey {
	panic("")
}
