package optimisticv3

import (
	ssz "github.com/ferranbt/fastssz"
)

// appendHashTreeRoot appends v's hash tree root to hh, which is what a nested
// v.HashTreeRootWith(hh) leaves there. optimistic_v3_ssz.go uses it for
// go-eth2-client types: from v0.30 those hash with dynamic-ssz's HashWalker,
// which fastssz's does not satisfy. Pinned by TestHeaderSubmissionHashTreeRoot.
func appendHashTreeRoot(hh ssz.HashWalker, v interface{ HashTreeRoot() ([32]byte, error) }) error {
	root, err := v.HashTreeRoot()
	if err != nil {
		return err
	}
	hh.AppendBytes32(root[:])
	return nil
}
