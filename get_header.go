package relay_grpc

func NewGetHeaderRequest(slot uint64, parentHash string, proposerPubkey string) *GetHeaderRequest {
	return &GetHeaderRequest{
		Slot:       slot,
		ParentHash: parentHash,
		Pubkey:     proposerPubkey,
	}
}
