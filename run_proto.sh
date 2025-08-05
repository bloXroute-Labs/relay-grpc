export PATH=$PATH:$(go env GOPATH)/bin

# Compile the Go protobuf files
protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative ./blxr-mev.proto

# Install the necessary Rust protobuf tools.
cargo install protoc-gen-prost
cargo install protoc-gen-tonic
# Compile the Rust protobuf files and output them all into the 'rust/proto/src' directory.
protoc -I . blxr-mev.proto --prost_out=rust/proto/src --tonic_out=rust/proto/src

# Rename Rust proto files and replace generated file include to point to the correct file name.
# This is a temporary workaround until we can coordinate with builders and add a 'package' declaration to the proto file.
mv rust/proto/src/_ rust/proto/src/blxrmev.rs
mv rust/proto/src/.tonic.rs rust/proto/src/blxrmev.tonic.rs
sed -i '' 's/\.tonic\.rs/blxrmev.tonic.rs/g' rust/proto/src/blxrmev.rs

