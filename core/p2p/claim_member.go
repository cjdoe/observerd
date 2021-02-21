package p2p

import (
	"github.com/vTCP-Foundation/observerd/common/crypto"
	"github.com/vTCP-Foundation/observerd/core/common"
	"github.com/vTCP-Foundation/observerd/core/marshalling"
)

const (
	ClaimMemberBinarySize = common.Uint16ByteSize + crypto.LamportPubKeySize
)

type ClaimMember struct {
	ID     uint16
	PubKey *crypto.LamportPubKey
}

func NewClaimMember(memberID uint16) *ClaimMember {
	return &ClaimMember{
		ID:     memberID,
		PubKey: &crypto.LamportPubKey{},
	}
}

func (member *ClaimMember) MarshalBinary() (data []byte, err error) {
	encoder := marshalling.NewEncoder()
	err = encoder.PutUint16(member.ID)
	if err != nil {
		return
	}

	err = encoder.PutFixedSizeDataSegment(member.PubKey[:])
	if err != nil {
		return
	}

	data = encoder.CollectDataAndReleaseBuffers()
	return
}

func (member *ClaimMember) UnmarshalBinary(data []byte) (err error) {
	decoder := marshalling.NewDecoder(data)

	member.ID, err = decoder.GetUint16()
	if err != nil {
		return
	}

	pubKeyBinary, err := decoder.GetDataSegment(crypto.LamportPubKeySize)
	if err != nil {
		return
	}

	pubKey := &crypto.LamportPubKey{}
	copy(pubKey[:], pubKeyBinary)
	return
}
