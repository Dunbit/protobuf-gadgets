# Protobuf Gadgets

[![Go Report Card](https://goreportcard.com/badge/github.com/dunbit/protobuf-gadgets)](https://goreportcard.com/report/github.com/dunbit/protobuf-gadgets)

## Tools

### protoc-gen-docmaster

Documentation generator plugin for Google Protocol Buffers (`protoc`). The plugin supports multiple documentation formats: Markdown, HTML, Custom Templates. It supports Uber's prototool, and Protocol Buffers version proto2 and proto3.

#### Prerequisites

- [Protocol Buffers](https://github.com/protocolbuffers/protobuf/releases)
- [Prototool](https://github.com/uber/prototool) (optional)

#### Installation

```bash
go get -u github.com/dunbit/protobuf-gadgets/cmd/protoc-gen-docmaster
```

#### Invoke plugin

The plugin is invoked by passing `--docmaster_out` with options to the `protoc` compiler. 

```bash
--docmaster_out=OPTIONS
```

The options are KEY=VALUE delimited by comma.

```bash
KEY=VALUE,KEY=VALUE,KEY=VALUE,...
```

##### Simple Invokation
```bash
protoc --docmaster_out=format=markdown:./docs proto/foo/v1/bar.proto
```
output:
```
./docs/bar.md
```

Full Documention: [protoc-gen-docmaster.md](https://github.com/dunbit/protobuf-gadgets/blob/master/docs/protoc-gen-docmaster.md)

---