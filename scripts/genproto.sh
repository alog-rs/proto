#!/usr/bin/env bash

set -euo pipefail

if ! [ -x "$(command -v protoc)" ]; then
    echo "genproto.sh requires protoc";

    exit 1;
fi

if ! [ -x "$(command -v protoc-gen-go)" ]; then
    echo "genproto.sh requires protoc-gen-go";

    exit 1;
fi

if ! [ -x "$(command -v protoc-gen-go-grpc)" ]; then
    echo "genproto.sh requires protoc-gen-go-grpc";

    exit 1;
fi

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROTO_DIR=$SCRIPT_DIR/../pb
DEST_DIR=$GOPATH/src

protoc \
    --experimental_allow_proto3_optional \
    --go_out=$DEST_DIR \
    --go-grpc_out=$DEST_DIR \
    -I $PROTO_DIR $PROTO_DIR/rs3.proto