# First protobuf example
## Writing a .proto file

## Compiling

### Install Compiler

- Download Protocol Buffers

> release list in [here](https://github.com/protocolbuffers/protobuf/releases/)

```
wget https://github.com/protocolbuffers/protobuf/releases/tag/v3.6.1
```

```
unzip protoc-3.6.1-osx-x86_64.zip
cp ./bin/protoc /usr/local/bin
```

- Install Go protocol buffers plugin

```
go get -u github.com/golang/protobuf/protoc-gen-go
```

- Add include files

```
cp -r include/ /usr/local/include/
```

### Compile

```
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto
```