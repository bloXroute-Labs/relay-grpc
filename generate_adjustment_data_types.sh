#go get github.com/ferranbt/fastssz/sszgen

go run sszgen/*.go --path ./bidadjustment [--objs AdjustmentData,AdjustmentDataV2,AdjustmentDataV3] --output ./bidadjustment/adjustment_data_ssz.go


