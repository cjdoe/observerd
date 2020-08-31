package crypto

const (
	PublicKeySize  = 1024 * 16
	PrivateKeySize = 1024 * 16
	SigSize        = 1024 * 8
)

type PrivateKey [PublicKeySize]byte
type PublicKey [PrivateKeySize]byte
type Signature [SigSize]byte
