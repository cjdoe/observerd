package crypto

const (
	LamportPKeySize   = 1024 * 16
	LamportPubKeySize = 1024 * 16
	LamportSigSize    = 1024 * 8
)

type LamportSig = [LamportSigSize]byte

type LamportPKey = [LamportPKeySize]byte

type LamportPubKey = [LamportPubKeySize]byte
