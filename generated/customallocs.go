package generated

import (
)

// AllocateProxy allocates a FilPrivateReplicaInfo proxy object in the C heap,
// returning a function which, when called, frees the allocated memory. This
// method exists because the default c-for-go allocation strategy allocates a
// C struct with a field whose values is a pointer into the Go heap, which is
// not permitted by the most strict CGO check (cgocheck=2).
func (x *FilPrivateReplicaInfo) AllocateProxy() func() {
	panic("")
}

// AllocateProxy allocates a FilPoStProof proxy object in the C heap,
// returning a function which, when called, frees the allocated memory.
func (x *FilPoStProof) AllocateProxy() func() {
	panic("")
}
