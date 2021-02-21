package p2p

import (
	"github.com/vTCP-Foundation/observerd/core/common/errors"
	"github.com/vTCP-Foundation/observerd/core/common/types/transactions"
	"github.com/vTCP-Foundation/observerd/core/marshalling"
)

type Claim struct {
	TxUUID  transactions.TxID
	Members ClaimMembers
}

func (claim Claim) MarshalBinary() (data []byte, err error) {
	encoder := marshalling.NewEncoder()

	err = encoder.MarshallVariadicDataWithByteHeader(claim.TxUUID)
	if err != nil {
		return
	}

	err = encoder.MarshallVariadicDataWith2BytesHeader(claim.Members)
	if err != nil {
		return
	}

	data = encoder.CollectDataAndReleaseBuffers()
	return
}

func (claim Claim) UnmarshalBinary(data []byte) (err error) {
	decoder := marshalling.NewDecoder(data)

	err = decoder.UnmarshalDataSegmentWithByteHeader(claim.TxUUID)
	if err != nil {
		return
	}

	err = decoder.UnmarshalDataSegmentWith2BytesHeader(claim.Members)
	return
}

type Claims struct {
	At []*Claim
}

func (claims *Claims) Add(c *Claim) error {
	if c == nil {
		return errors.InvalidParameter
	}

	claims.At = append(claims.At, c)
	return nil
}

func (claims *Claims) MarshalBinary() (data []byte, err error) {
	encoder := marshalling.NewEncoder()

	err = encoder.PutUint16(uint16(len(claims.At)))
	if err != nil {
		return
	}

	for _, claim := range claims.At {
		err = encoder.MarshallVariadicDataWith2BytesHeader(claim)
		if err != nil {
			return
		}
	}

	data = encoder.CollectDataAndReleaseBuffers()
	return
}

func (claims *Claims) UnmarshalBinary(data []byte) (err error) {
	decoder := marshalling.NewDecoder(data)

	claimsLen, err := decoder.GetUint16()
	if err != nil {
		return
	}

	claims.At = make([]*Claim, int(claimsLen))
	for i, _ := range claims.At {
		claim := &Claim{}
		err = decoder.UnmarshalDataSegmentWith2BytesHeader(claim)
		if err != nil {
			return
		}

		claims.At[i] = claim
	}

	return
}
