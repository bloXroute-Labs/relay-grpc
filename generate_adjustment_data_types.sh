#go get github.com/ferranbt/fastssz/sszgen

sszgen --path ./bidadjustment/types.go --objs AdjustmentData,AdjustmentDataV2,AdjustmentDataV3 --output ./bidadjustment/adjustment_data_ssz.go
