package bidadjustment

import (
	d "github.com/attestantio/go-builder-client/api/deneb"
	v1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/electra"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the AdjustableSubmitBlockRequestV4 object
func (a *ElectraAdjustableSubmitBlockRequest) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(a)
}

// MarshalSSZTo ssz marshals the AdjustableSubmitBlockRequestV4 object to a target array
func (a *ElectraAdjustableSubmitBlockRequest) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(348)

	// Field (0) 'Message'
	if a.Message == nil {
		a.Message = new(v1.BidTrace)
	}
	if dst, err = a.Message.MarshalSSZTo(dst); err != nil {
		return
	}

	// Offset (1) 'ExecutionPayload'
	dst = ssz.WriteOffset(dst, offset)
	if a.ExecutionPayload == nil {
		a.ExecutionPayload = new(deneb.ExecutionPayload)
	}
	offset += a.ExecutionPayload.SizeSSZ()

	// Offset (2) 'BlobsBundle'
	dst = ssz.WriteOffset(dst, offset)
	if a.BlobsBundle == nil {
		a.BlobsBundle = new(d.BlobsBundle)
	}
	offset += a.BlobsBundle.SizeSSZ()

	// Offset (3) 'ExecutionRequests'
	dst = ssz.WriteOffset(dst, offset)
	if a.ExecutionRequests == nil {
		a.ExecutionRequests = new(electra.ExecutionRequests)
	}
	offset += a.ExecutionRequests.SizeSSZ()

	// Field (4) 'Signature'
	dst = append(dst, a.Signature[:]...)

	// Offset (5) 'AdjustmentData'
	dst = ssz.WriteOffset(dst, offset)
	if a.AdjustmentData == nil {
		a.AdjustmentData = new(VersionedAdjustmentData)
	}

	// Field (1) 'ExecutionPayload'
	if dst, err = a.ExecutionPayload.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (2) 'BlobsBundle'
	if dst, err = a.BlobsBundle.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (3) 'ExecutionRequests'
	if dst, err = a.ExecutionRequests.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (5) 'AdjustmentData'
	if dst, err = a.AdjustmentData.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the AdjustableSubmitBlockRequestV4 object
func (a *ElectraAdjustableSubmitBlockRequest) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 348 {
		return ssz.ErrSize
	}

	tail := buf
	var o1, o2, o3, o5 uint64

	// Field (0) 'Message'
	if a.Message == nil {
		a.Message = new(v1.BidTrace)
	}
	if err = a.Message.UnmarshalSSZ(buf[0:236]); err != nil {
		return err
	}

	// Offset (1) 'ExecutionPayload'
	if o1 = ssz.ReadOffset(buf[236:240]); o1 > size {
		return ssz.ErrOffset
	}

	if o1 < 348 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (2) 'BlobsBundle'
	if o2 = ssz.ReadOffset(buf[240:244]); o2 > size || o1 > o2 {
		return ssz.ErrOffset
	}

	// Offset (3) 'ExecutionRequests'
	if o3 = ssz.ReadOffset(buf[244:248]); o3 > size || o2 > o3 {
		return ssz.ErrOffset
	}

	// Field (4) 'Signature'
	copy(a.Signature[:], buf[248:344])

	// Offset (5) 'AdjustmentData'
	if o5 = ssz.ReadOffset(buf[344:348]); o5 > size || o3 > o5 {
		return ssz.ErrOffset
	}

	// Field (1) 'ExecutionPayload'
	{
		buf = tail[o1:o2]
		if a.ExecutionPayload == nil {
			a.ExecutionPayload = new(deneb.ExecutionPayload)
		}
		if err = a.ExecutionPayload.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (2) 'BlobsBundle'
	{
		buf = tail[o2:o3]
		if a.BlobsBundle == nil {
			a.BlobsBundle = new(d.BlobsBundle)
		}
		if err = UnmarshalBlobsBundleReuse(a.BlobsBundle, buf); err != nil {
			return err
		}
	}

	// Field (3) 'ExecutionRequests'
	{
		buf = tail[o3:o5]
		if a.ExecutionRequests == nil {
			a.ExecutionRequests = new(electra.ExecutionRequests)
		}
		if err = a.ExecutionRequests.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (5) 'AdjustmentData'
	{
		buf = tail[o5:]
		if a.AdjustmentData == nil {
			a.AdjustmentData = new(VersionedAdjustmentData)
		}
		if err = a.AdjustmentData.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the AdjustableSubmitBlockRequestV4 object
func (a *ElectraAdjustableSubmitBlockRequest) SizeSSZ() (size int) {
	size = 348

	// Field (1) 'ExecutionPayload'
	if a.ExecutionPayload == nil {
		a.ExecutionPayload = new(deneb.ExecutionPayload)
	}
	size += a.ExecutionPayload.SizeSSZ()

	// Field (2) 'BlobsBundle'
	if a.BlobsBundle == nil {
		a.BlobsBundle = new(d.BlobsBundle)
	}
	size += a.BlobsBundle.SizeSSZ()

	// Field (3) 'ExecutionRequests'
	if a.ExecutionRequests == nil {
		a.ExecutionRequests = new(electra.ExecutionRequests)
	}
	size += a.ExecutionRequests.SizeSSZ()

	// Field (5) 'AdjustmentData'
	if a.AdjustmentData == nil {
		a.AdjustmentData = new(VersionedAdjustmentData)
	}
	size += a.AdjustmentData.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the AdjustableSubmitBlockRequestV4 object
func (a *ElectraAdjustableSubmitBlockRequest) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(a)
}

// HashTreeRootWith ssz hashes the AdjustableSubmitBlockRequestV4 object with a hasher
func (a *ElectraAdjustableSubmitBlockRequest) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Message'
	if a.Message == nil {
		a.Message = new(v1.BidTrace)
	}
	if err = a.Message.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'ExecutionPayload'
	if a.ExecutionPayload == nil {
		a.ExecutionPayload = new(deneb.ExecutionPayload)
	}
	if err = a.ExecutionPayload.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'BlobsBundle'
	if a.BlobsBundle == nil {
		a.BlobsBundle = new(d.BlobsBundle)
	}
	if err = a.BlobsBundle.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (3) 'ExecutionRequests'
	if a.ExecutionRequests == nil {
		a.ExecutionRequests = new(electra.ExecutionRequests)
	}
	if err = a.ExecutionRequests.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (4) 'Signature'
	hh.PutBytes(a.Signature[:])

	// Field (5) 'AdjustmentData'
	if a.AdjustmentData == nil {
		a.AdjustmentData = new(VersionedAdjustmentData)
	}
	if err = a.AdjustmentData.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the AdjustableSubmitBlockRequestV4 object
func (a *ElectraAdjustableSubmitBlockRequest) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(a)
}
