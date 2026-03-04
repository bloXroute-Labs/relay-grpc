package bidadjustment

import (
	"fmt"

	ssz "github.com/ferranbt/fastssz"
	"github.com/pkg/errors"
)

func (v *VersionedAdjustmentData) MarshalSSZ() ([]byte, error) {
	switch v.Version {
	case AdjustmentDataVersion3:
		return v.V3.MarshalSSZ()
	case AdjustmentDataVersion2:
		return v.V2.MarshalSSZ()
	case AdjustmentDataVersion1:
		return v.V1.MarshalSSZ()
	default:
		return nil, errors.Wrap(ErrInvalidVersion, fmt.Sprintf("%d is not supported", v.Version))
	}
}

func (v *VersionedAdjustmentData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	switch v.Version {
	case AdjustmentDataVersion3:
		return v.V3.MarshalSSZTo(buf)
	case AdjustmentDataVersion2:
		return v.V2.MarshalSSZTo(buf)
	case AdjustmentDataVersion1:
		return v.V1.MarshalSSZTo(buf)
	default:
		return nil, errors.Wrap(ErrInvalidVersion, fmt.Sprintf("%d is not supported", v.Version))
	}
}

func (v *VersionedAdjustmentData) UnmarshalSSZ(input []byte) error {
	var err error

	dataV3 := new(AdjustmentDataV3)
	if err = dataV3.UnmarshalSSZ(input); err == nil {
		v.Version = AdjustmentDataVersion3
		v.V3 = dataV3
		return nil
	}

	dataV2 := new(AdjustmentDataV2)
	if err = dataV2.UnmarshalSSZ(input); err == nil {
		v.Version = AdjustmentDataVersion2
		v.V2 = dataV2
		return nil
	}

	dataV1 := new(AdjustmentData)
	if err = dataV1.UnmarshalSSZ(input); err == nil {
		v.Version = AdjustmentDataVersion1
		v.V1 = dataV1
		return nil
	}

	return errors.Wrap(err, "failed to unmarshal VersionedAdjustmentData SSZ")
}

func (v *VersionedAdjustmentData) SizeSSZ() (size int) {
	switch v.Version {
	case AdjustmentDataVersion3:
		return v.V3.SizeSSZ()
	case AdjustmentDataVersion2:
		return v.V2.SizeSSZ()
	case AdjustmentDataVersion1:
		return v.V1.SizeSSZ()
	default:
		return 0
	}
}

func (v *VersionedAdjustmentData) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	switch v.Version {
	case AdjustmentDataVersion3:
		return v.V3.HashTreeRootWith(hh)
	case AdjustmentDataVersion2:
		return v.V2.HashTreeRootWith(hh)
	case AdjustmentDataVersion1:
		return v.V1.HashTreeRootWith(hh)
	default:
		return errors.Wrap(ErrInvalidVersion, fmt.Sprintf("%d is not supported", v.Version))
	}
}
