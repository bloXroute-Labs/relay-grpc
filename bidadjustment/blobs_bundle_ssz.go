package bidadjustment

import (
	api "github.com/attestantio/go-builder-client/api/deneb"
	"github.com/attestantio/go-eth2-client/spec/deneb"
	ssz "github.com/ferranbt/fastssz"
)

// UnmarshalBlobsBundleReuse decodes SSZ into b reusing slice capacity to avoid
// per-call reallocations. Call this instead of b.UnmarshalSSZ.
func UnmarshalBlobsBundleReuse(b *api.BlobsBundle, buf []byte) error {
	if len(buf) < 12 {
		return ssz.ErrSize
	}
	size := uint64(len(buf))
	tail := buf
	var o0, o1, o2 uint64

	// Offsets
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size || o0 < 12 {
		return ssz.ErrOffset
	}
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}
	if o2 = ssz.ReadOffset(buf[8:12]); o2 > size || o1 > o2 {
		return ssz.ErrOffset
	}

	// Commitments
	seg := tail[o0:o1]
	num, err := ssz.DivideInt2(len(seg), 48, 4096)
	if err != nil {
		return err
	}
	if cap(b.Commitments) >= num {
		b.Commitments = b.Commitments[:num]
	} else {
		b.Commitments = make([]deneb.KZGCommitment, num)
	}
	for i := 0; i < num; i++ {
		copy(b.Commitments[i][:], seg[i*48:(i+1)*48])
	}

	// Proofs
	seg = tail[o1:o2]
	num, err = ssz.DivideInt2(len(seg), 48, 4096)
	if err != nil {
		return err
	}
	if cap(b.Proofs) >= num {
		b.Proofs = b.Proofs[:num]
	} else {
		b.Proofs = make([]deneb.KZGProof, num)
	}
	for i := 0; i < num; i++ {
		copy(b.Proofs[i][:], seg[i*48:(i+1)*48])
	}

	// Blobs (128 KiB each)
	seg = tail[o2:]
	num, err = ssz.DivideInt2(len(seg), 131072, 4096)
	if err != nil {
		return err
	}
	if cap(b.Blobs) >= num {
		b.Blobs = b.Blobs[:num]
	} else {
		b.Blobs = make([]deneb.Blob, num)
	}
	for i := 0; i < num; i++ {
		copy(b.Blobs[i][:], seg[i*131072:(i+1)*131072])
	}
	return nil
}
