package ffi

import (
	"os"
)

// NewSortedPublicSectorInfo returns a SortedPublicSectorInfo

// VerifySeal returns true if the sealing operation from which its inputs were
// derived was valid, and false if not.
func VerifySeal(
	sectorSize uint64,
	commR [CommitmentBytesLen]byte,
	commD [CommitmentBytesLen]byte,
	proverID [32]byte,
	ticket [32]byte,
	seed [32]byte,
	sectorID uint64,
	proof []byte,
) (bool, error) {

	panic("")
}

// VerifyPoSt returns true if the PoSt-generation operation from which its
// inputs were derived was valid, and false if not.
func VerifyPoSt(
	sectorSize uint64,
	sectorInfo SortedPublicSectorInfo,
	randomness [32]byte,
	challengeCount uint64,
	proof []byte,
	winners []Candidate,
	proverID [32]byte,
) (bool, error) {
	panic("")
}

// GetMaxUserBytesPerStagedSector returns the number of user bytes that will fit
// into a staged sector. Due to bit-padding, the number of user bytes that will
// fit into the staged sector will be less than number of bytes in sectorSize.
func GetMaxUserBytesPerStagedSector(sectorSize uint64) uint64 {
	panic("")
}

// GeneratePieceCommitment produces a piece commitment for the provided data
// stored at a given path.
func GeneratePieceCommitment(piecePath string, pieceSize uint64) ([CommitmentBytesLen]byte, error) {
	panic("")
}

// GenerateDataCommitment produces a commitment for the sector containing the
// provided pieces.
func GenerateDataCommitment(sectorSize uint64, pieces []PublicPieceInfo) ([CommitmentBytesLen]byte, error) {
	panic("")
}

// GeneratePieceCommitmentFromFile produces a piece commitment for the provided data
// stored in a given file.
func GeneratePieceCommitmentFromFile(pieceFile *os.File, pieceSize uint64) (commP [CommitmentBytesLen]byte, err error) {
	panic("")
}

// WriteWithAlignment
func WriteWithAlignment(
	pieceFile *os.File,
	pieceBytes uint64,
	stagedSectorFile *os.File,
	existingPieceSizes []uint64,
) (leftAlignment, total uint64, commP [CommitmentBytesLen]byte, retErr error) {
	panic("")
}

// WriteWithoutAlignment
func WriteWithoutAlignment(
	pieceFile *os.File,
	pieceBytes uint64,
	stagedSectorFile *os.File,
) (uint64, [CommitmentBytesLen]byte, error) {
	panic("")
}

// SealPreCommit
func SealPreCommit(
	sectorSize uint64,
	poRepProofPartitions uint8,
	cacheDirPath string,
	stagedSectorPath string,
	sealedSectorPath string,
	sectorID uint64,
	proverID [32]byte,
	ticket [32]byte,
	pieces []PublicPieceInfo,
) (RawSealPreCommitOutput, error) {
	panic("")
}

// SealCommit
func SealCommit(
	sectorSize uint64,
	poRepProofPartitions uint8,
	cacheDirPath string,
	sectorID uint64,
	proverID [32]byte,
	ticket [32]byte,
	seed [32]byte,
	pieces []PublicPieceInfo,
	rspco RawSealPreCommitOutput,
) ([]byte, error) {
	panic("")
}

// Unseal
func Unseal(
	sectorSize uint64,
	poRepProofPartitions uint8,
	cacheDirPath string,
	sealedSectorPath string,
	unsealOutputPath string,
	sectorID uint64,
	proverID [32]byte,
	ticket [32]byte,
	commD [CommitmentBytesLen]byte,
) error {
	panic("")
}

// UnsealRange
func UnsealRange(
	sectorSize uint64,
	poRepProofPartitions uint8,
	cacheDirPath string,
	sealedSectorPath string,
	unsealOutputPath string,
	sectorID uint64,
	proverID [32]byte,
	ticket [32]byte,
	commD [CommitmentBytesLen]byte,
	offset uint64,
	len uint64,
) error {
	panic("")
}

// FinalizeTicket creates an actual ticket from a partial ticket.
func FinalizeTicket(partialTicket [32]byte) ([32]byte, error) {
	panic("")
}

// GenerateCandidates
func GenerateCandidates(
	sectorSize uint64,
	proverID [32]byte,
	randomness [32]byte,
	challengeCount uint64,
	privateSectorInfo SortedPrivateSectorInfo,
) ([]Candidate, error) {
	panic("")
}

// GeneratePoSt
func GeneratePoSt(
	sectorSize uint64,
	proverID [32]byte,
	privateSectorInfo SortedPrivateSectorInfo,
	randomness [32]byte,
	winners []Candidate,
) ([]byte, error) {
	panic("")
}

// SingleProofPartitionProofLen denotes the number of bytes in a proof generated
// with a single partition. The number of bytes in a proof increases linearly
// with the number of partitions used when creating that proof.
const SingleProofPartitionProofLen = 192
