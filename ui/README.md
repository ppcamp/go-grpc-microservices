To generate protos

```shell
protoc -I=protos protos/* \
    --js_out=import_style=commonjs,binary:web \
    --grpc-web_out=import_style=typescript,mode=grpcwebtext:web
```

