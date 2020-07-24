#!/bin/bash

for i in "$@"
do :
  echo "generating ${i} service"

  echo "  - grpc bindings"
  protoc \
    --proto_path=protos/ \
    --go_out=plugins=grpc:. \
    protos/${i}.proto

  # move generated sources
  mkdir -p generated/${i}
  mv protos/${i}.*.go generated/${i}/.
done

echo "services generated"
