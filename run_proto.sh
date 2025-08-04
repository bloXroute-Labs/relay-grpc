export PATH=$PATH:$(go env GOPATH)/bin

# Compile the Go protobuf files
protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative ./blxr-mev.proto

# Install the necessary Rust protobuf tools.
cargo install protoc-gen-prost
cargo install protoc-gen-tonic
# Compile the Rust protobuf files and output them all into the 'rust/proto/src' directory.
protoc -I . blxr-mev.proto --prost_out=rust/proto/src --tonic_out=rust/proto/src
