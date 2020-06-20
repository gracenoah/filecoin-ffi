//+build cgo

package ffi

import (
	"os"
	"unsafe"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/specs-actors/actors/runtime/proof"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/filecoin-ffi/generated"
)

// VerifySeal returns true if the sealing operation from which its inputs were
// derived was valid, and false if not.
func VerifySeal(info proof.SealVerifyInfo) (bool, error) {
	panic("")
}

// VerifyWinningPoSt returns true if the Winning PoSt-generation operation from which its
// inputs were derived was valid, and false if not.
func VerifyWinningPoSt(info proof.WinningPoStVerifyInfo) (bool, error) {
	panic("")
}

// VerifyWindowPoSt returns true if the Winning PoSt-generation operation from which its
// inputs were derived was valid, and false if not.
func VerifyWindowPoSt(info proof.WindowPoStVerifyInfo) (bool, error) {
	panic("")
}

// GeneratePieceCommitment produces a piece commitment for the provided data
// stored at a given path.
func GeneratePieceCID(proofType abi.RegisteredSealProof, piecePath string, pieceSize abi.UnpaddedPieceSize) (cid.Cid, error) {
	panic("")
}

// GenerateDataCommitment produces a commitment for the sector containing the
// provided pieces.
func GenerateUnsealedCID(proofType abi.RegisteredSealProof, pieces []abi.PieceInfo) (cid.Cid, error) {
	panic("")
}

// GeneratePieceCIDFromFile produces a piece CID for the provided data stored in
//a given file.
func GeneratePieceCIDFromFile(proofType abi.RegisteredSealProof, pieceFile *os.File, pieceSize abi.UnpaddedPieceSize) (cid.Cid, error) {
	panic("")
}

// WriteWithAlignment
func WriteWithAlignment(
	proofType abi.RegisteredSealProof,
	pieceFile *os.File,
	pieceBytes abi.UnpaddedPieceSize,
	stagedSectorFile *os.File,
	existingPieceSizes []abi.UnpaddedPieceSize,
) (leftAlignment, total abi.UnpaddedPieceSize, pieceCID cid.Cid, retErr error) {
	panic("")
}

// WriteWithoutAlignment
func WriteWithoutAlignment(
	proofType abi.RegisteredSealProof,
	pieceFile *os.File,
	pieceBytes abi.UnpaddedPieceSize,
	stagedSectorFile *os.File,
) (abi.UnpaddedPieceSize, cid.Cid, error) {
	panic("")
}

// SealPreCommitPhase1
func SealPreCommitPhase1(
	proofType abi.RegisteredSealProof,
	cacheDirPath string,
	stagedSectorPath string,
	sealedSectorPath string,
	sectorNum abi.SectorNumber,
	minerID abi.ActorID,
	ticket abi.SealRandomness,
	pieces []abi.PieceInfo,
) (phase1Output []byte, err error) {
	panic("")
}

// SealPreCommitPhase2
func SealPreCommitPhase2(
	phase1Output []byte,
	cacheDirPath string,
	sealedSectorPath string,
) (sealedCID cid.Cid, unsealedCID cid.Cid, err error) {
	panic("")
}

// SealCommitPhase1
func SealCommitPhase1(
	proofType abi.RegisteredSealProof,
	sealedCID cid.Cid,
	unsealedCID cid.Cid,
	cacheDirPath string,
	sealedSectorPath string,
	sectorNum abi.SectorNumber,
	minerID abi.ActorID,
	ticket abi.SealRandomness,
	seed abi.InteractiveSealRandomness,
	pieces []abi.PieceInfo,
) (phase1Output []byte, err error) {
	panic("")
}

// SealCommitPhase2
func SealCommitPhase2(
	phase1Output []byte,
	sectorNum abi.SectorNumber,
	minerID abi.ActorID,
) ([]byte, error) {
	panic("")
}

// Unseal
func Unseal(
	proofType abi.RegisteredSealProof,
	cacheDirPath string,
	sealedSector *os.File,
	unsealOutput *os.File,
	sectorNum abi.SectorNumber,
	minerID abi.ActorID,
	ticket abi.SealRandomness,
	unsealedCID cid.Cid,
) error {
	panic("")
}

// UnsealRange
func UnsealRange(
	proofType abi.RegisteredSealProof,
	cacheDirPath string,
	sealedSector *os.File,
	unsealOutput *os.File,
	sectorNum abi.SectorNumber,
	minerID abi.ActorID,
	ticket abi.SealRandomness,
	unsealedCID cid.Cid,
	unpaddedByteIndex uint64,
	unpaddedBytesAmount uint64,
) error {
	panic("")
}

// GenerateWinningPoStSectorChallenge
func GenerateWinningPoStSectorChallenge(
	proofType abi.RegisteredPoStProof,
	minerID abi.ActorID,
	randomness abi.PoStRandomness,
	eligibleSectorsLen uint64,
) ([]uint64, error) {
	panic("")
}

// GenerateWinningPoSt
func GenerateWinningPoSt(
	minerID abi.ActorID,
	privateSectorInfo SortedPrivateSectorInfo,
	randomness abi.PoStRandomness,
) ([]proof.PoStProof, error) {
	panic("")
}

// GenerateWindowPoSt
func GenerateWindowPoSt(
	minerID abi.ActorID,
	privateSectorInfo SortedPrivateSectorInfo,
	randomness abi.PoStRandomness,
) ([]proof.PoStProof, []abi.SectorNumber, error) {
	panic("")
}

// GetGPUDevices produces a slice of strings, each representing the name of a
// detected GPU device.
func GetGPUDevices() ([]string, error) {
	panic("")
}

// GetSealVersion
func GetSealVersion(proofType abi.RegisteredSealProof) (string, error) {
	panic("")
}

// GetPoStVersion
func GetPoStVersion(proofType abi.RegisteredPoStProof) (string, error) {
	panic("")
}

// ClearCache
func ClearCache(sectorSize uint64, cacheDirPath string) error {
	panic("")
}

func FauxRep(proofType abi.RegisteredSealProof, cacheDirPath string, sealedSectorPath string) (cid.Cid, error) {
	panic("")
}

func FauxRep2(proofType abi.RegisteredSealProof, cacheDirPath string, existingPAuxPath string) (cid.Cid, error) {
	panic("")
}

func toFilExistingPieceSizes(src []abi.UnpaddedPieceSize) ([]uint64, uint) {
	panic("")
}

func toFilPublicPieceInfos(src []abi.PieceInfo) ([]generated.FilPublicPieceInfo, uint, error) {
	panic("")
}

func toFilPublicReplicaInfos(src []proof.SectorInfo, typ string) ([]generated.FilPublicReplicaInfo, uint, error) {
	panic("")
}

func toFilPrivateReplicaInfos(src []PrivateSectorInfo, typ string) ([]generated.FilPrivateReplicaInfo, uint, func(), error) {
	panic("")
}

func fromFilPoStFaultySectors(ptr []uint64, l uint) ([]abi.SectorNumber, error) {
	panic("")
}

func fromFilPoStProofs(src []generated.FilPoStProof) ([]proof.PoStProof, error) {
	panic("")
}

func toFilPoStProofs(src []proof.PoStProof) ([]generated.FilPoStProof, uint, func(), error) {
	panic("")
}

func to32ByteArray(in []byte) generated.Fil32ByteArray {
	panic("")
}

func toProverID(minerID abi.ActorID) (generated.Fil32ByteArray, error) {
	panic("")
}

func fromFilRegisteredPoStProof(p generated.FilRegisteredPoStProof) (abi.RegisteredPoStProof, error) {
	panic("")
}

func toFilRegisteredPoStProof(p abi.RegisteredPoStProof) (generated.FilRegisteredPoStProof, error) {
	panic("")
}

func toFilRegisteredSealProof(p abi.RegisteredSealProof) (generated.FilRegisteredSealProof, error) {
	panic("")
}

func to32ByteCommD(unsealedCID cid.Cid) (generated.Fil32ByteArray, error) {
	panic("")
}

func to32ByteCommR(sealedCID cid.Cid) (generated.Fil32ByteArray, error) {
	panic("")
}

func to32ByteCommP(pieceCID cid.Cid) (generated.Fil32ByteArray, error) {
	panic("")
}

func toGoStringCopy(raw string, rawLen uint) string {
	panic("")
}

type stringHeader struct {
	Data unsafe.Pointer
	Len  int
}
