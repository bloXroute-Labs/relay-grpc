package bidadjustment

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVersionedAdjustmentDataJsonMarshalUnmarshal(t *testing.T) {
	// Test that V1 version and data are unchanged from a marshal/unmarshal cycle
	versionedAdjustmentData := &VersionedAdjustmentData{
		Version: AdjustmentDataVersion1,
		V1:      testAdjustmentDataV1,
		V2:      nil,
	}

	bytes, err := json.Marshal(&versionedAdjustmentData)
	require.NoError(t, err)

	newAdjustmentData := new(VersionedAdjustmentData)
	err = json.Unmarshal(bytes, newAdjustmentData)
	require.NoError(t, err)

	require.Equal(t, versionedAdjustmentData.Version, newAdjustmentData.Version)
	require.Equal(t, *versionedAdjustmentData.V1, *newAdjustmentData.V1)
	require.Nil(t, newAdjustmentData.V2)

	// Test that V2 version and data are unchanged from a marshal/unmarshal cycle
	versionedAdjustmentData = &VersionedAdjustmentData{
		Version: AdjustmentDataVersion2,
		V1:      nil,
		V2:      testAdjustmentDataV2,
	}

	bytes, err = json.Marshal(&versionedAdjustmentData)
	require.NoError(t, err)

	newAdjustmentData = new(VersionedAdjustmentData)
	err = json.Unmarshal(bytes, newAdjustmentData)
	require.NoError(t, err)

	require.Equal(t, versionedAdjustmentData.Version, newAdjustmentData.Version)
	require.Equal(t, *versionedAdjustmentData.V2, *newAdjustmentData.V2)
	require.Nil(t, newAdjustmentData.V1)
}

func TestVersionedAdjustmentDataSszMarshalUnmarshal(t *testing.T) {
	// Test that V1 version and data are unchanged from a marshal/unmarshal cycle
	versionedAdjustmentData := &VersionedAdjustmentData{
		Version: AdjustmentDataVersion1,
		V1:      testAdjustmentDataV1,
		V2:      nil,
		V3:      nil,
	}

	bytes, err := versionedAdjustmentData.MarshalSSZ()
	require.NoError(t, err)

	newAdjustmentData := new(VersionedAdjustmentData)
	err = newAdjustmentData.UnmarshalSSZ(bytes)
	require.NoError(t, err)

	require.Equal(t, versionedAdjustmentData.Version, newAdjustmentData.Version)
	require.Equal(t, *versionedAdjustmentData.V1, *newAdjustmentData.V1)
	require.Nil(t, newAdjustmentData.V2)
	require.Nil(t, newAdjustmentData.V3)

	// Test that V2 version and data are unchanged from a marshal/unmarshal cycle
	versionedAdjustmentData = &VersionedAdjustmentData{
		Version: AdjustmentDataVersion2,
		V1:      nil,
		V2:      testAdjustmentDataV2,
		V3:      nil,
	}

	bytes, err = versionedAdjustmentData.MarshalSSZ()
	require.NoError(t, err)

	newAdjustmentData = new(VersionedAdjustmentData)
	err = newAdjustmentData.UnmarshalSSZ(bytes)
	require.NoError(t, err)

	require.Equal(t, versionedAdjustmentData.Version, newAdjustmentData.Version)
	require.Equal(t, *versionedAdjustmentData.V2, *newAdjustmentData.V2)
	require.Nil(t, newAdjustmentData.V1)
	require.Nil(t, newAdjustmentData.V3)

	// Test that V3 version and data are unchanged from a marshal/unmarshal cycle
	versionedAdjustmentData = &VersionedAdjustmentData{
		Version: AdjustmentDataVersion3,
		V1:      nil,
		V2:      nil,
		V3:      testAdjustmentDataV3,
	}

	bytes, err = versionedAdjustmentData.MarshalSSZ()
	require.NoError(t, err)

	newAdjustmentData = new(VersionedAdjustmentData)
	err = newAdjustmentData.UnmarshalSSZ(bytes)
	require.NoError(t, err)

	require.Equal(t, versionedAdjustmentData.Version, newAdjustmentData.Version)
	require.Equal(t, *versionedAdjustmentData.V3, *newAdjustmentData.V3)
	require.Nil(t, newAdjustmentData.V1)
	require.Nil(t, newAdjustmentData.V2)
}
