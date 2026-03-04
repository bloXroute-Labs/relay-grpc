package bidadjustment

import (
	d "github.com/attestantio/go-builder-client/api/deneb"
	v1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/deneb"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the AdjustableSubmitBlockRequest object
func (a *DenebAdjustableSubmitBlockRequest) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(a)
}

// MarshalSSZTo ssz marshals the AdjustableSubmitBlockRequest object to a target array
func (a *DenebAdjustableSubmitBlockRequest) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(344)

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

	// Field (3) 'Signature'
	dst = append(dst, a.Signature[:]...)

	// Offset (4) 'AdjustmentData'
	dst = ssz.WriteOffset(dst, offset)
	if a.AdjustmentData == nil {
		a.AdjustmentData = new(VersionedAdjustmentData)
	}
	offset += a.AdjustmentData.SizeSSZ() //nolint:ineffassign

	// Field (1) 'ExecutionPayload'
	if dst, err = a.ExecutionPayload.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (2) 'BlobsBundle'
	if dst, err = a.BlobsBundle.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (4) 'AdjustmentData'
	if dst, err = a.AdjustmentData.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the AdjustableSubmitBlockRequest object
func (a *DenebAdjustableSubmitBlockRequest) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 344 {
		return ssz.ErrSize
	}

	tail := buf
	var o1, o2, o4 uint64

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

	if o1 < 344 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (2) 'BlobsBundle'
	if o2 = ssz.ReadOffset(buf[240:244]); o2 > size || o1 > o2 {
		return ssz.ErrOffset
	}

	// Field (3) 'Signature'
	copy(a.Signature[:], buf[244:340])

	// Offset (4) 'AdjustmentData'
	if o4 = ssz.ReadOffset(buf[340:344]); o4 > size || o4 < o2 {
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
		buf = tail[o2:o4]
		if a.BlobsBundle == nil {
			a.BlobsBundle = new(d.BlobsBundle)
		}
		if err = UnmarshalBlobsBundleReuse(a.BlobsBundle, buf); err != nil {
			return err
		}
	}

	// Field (4) 'AdjustmentData'
	{
		buf = tail[o4:]
		if a.AdjustmentData == nil {
			a.AdjustmentData = new(VersionedAdjustmentData)
		}
		if err = a.AdjustmentData.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the AdjustableSubmitBlockRequest object
func (a *DenebAdjustableSubmitBlockRequest) SizeSSZ() (size int) {
	size = 344

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

	// Field (4) 'AdjustmentData'
	if a.AdjustmentData == nil {
		a.AdjustmentData = new(VersionedAdjustmentData)
	}
	size += a.AdjustmentData.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the AdjustableSubmitBlockRequest object
func (a *DenebAdjustableSubmitBlockRequest) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(a)
}

// HashTreeRootWith ssz hashes the AdjustableSubmitBlockRequest object with a hasher
func (a *DenebAdjustableSubmitBlockRequest) HashTreeRootWith(hh ssz.HashWalker) (err error) {
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

	// Field (3) 'Signature'
	hh.PutBytes(a.Signature[:])

	// Field (4) 'AdjustmentData'
	if a.AdjustmentData == nil {
		a.AdjustmentData = new(VersionedAdjustmentData)
	}
	if err = a.AdjustmentData.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the AdjustableSubmitBlockRequest object
func (a *DenebAdjustableSubmitBlockRequest) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(a)
}
