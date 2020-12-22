//+build cgo

package ffi

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/specs-actors/actors/runtime/proof"
)

type FallbackChallenges struct {
	Sectors    []abi.SectorNumber
	Challenges map[abi.SectorNumber][]uint64
}

type VanillaProof []byte

// GenerateWinningPoStSectorChallenge
func GeneratePoStFallbackSectorChallenges(
	proofType abi.RegisteredPoStProof,
	minerID abi.ActorID,
	randomness abi.PoStRandomness,
	sectorIds []abi.SectorNumber,
) (*FallbackChallenges, error) {
	panic("")
}

func GenerateSingleVanillaProof(
	replica PrivateSectorInfo,
	challange []uint64,
) ([]byte, error) {
	panic("")
}

func GenerateWinningPoStWithVanilla(
	proofType abi.RegisteredPoStProof,
	minerID abi.ActorID,
	randomness abi.PoStRandomness,
	proofs [][]byte,
) ([]proof.PoStProof, error) {
	panic("")
}

func GenerateWindowPoStWithVanilla(
	proofType abi.RegisteredPoStProof,
	minerID abi.ActorID,
	randomness abi.PoStRandomness,
	proofs [][]byte,
) ([]proof.PoStProof, error) {
	panic("")
}
