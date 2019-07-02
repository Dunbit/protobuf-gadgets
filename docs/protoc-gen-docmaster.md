# protoc-gen-docmaster

Documentation generator plugin for Google Protocol Buffers (`protoc`). The plugin supports multiple documentation formats: Markdown, HTML, Custom Templates. It supports Uber's prototool, and Protocol Buffers version proto2 and proto3.

## Prerequisites

- [Protocol Buffers](https://github.com/protocolbuffers/protobuf/releases)
- [Prototool](https://github.com/uber/prototool) (optional)

## Installation

```bash
go get -u github.com/dunbit/protobuf-gadgets/cmd/protoc-gen-docmaster
```

## Invoke plugin

The plugin is invoked by passing `--docmaster_out` with options to the `protoc` compiler. 

```bash
--docmaster_out=OPTIONS
```

The options are KEY=VALUE delimited by comma.

```bash
KEY=VALUE,KEY=VALUE,KEY=VALUE,...
```

### Simple Invokation
```bash
protoc --docmaster_out=format=markdown:./docs proto/foo/v1/bar.proto
```
output:
```
./docs/bar.md
```

### Custom filename
```bash
protoc --docmaster_out=format=markdown,file=api.md:./docs proto/foo/v1/bar.proto
```
output:
```
./docs/api.md
```

### Multiple proto files
```bash
protoc --docmaster_out=format=markdown,file=api.md:./docs proto/foo/v1/*.proto
```
output (contains documentation of all supplied protofiles):
```
./docs/api.md
```

### Option variables - `DIR`
```bash
protoc --docmaster_out=format=markdown,file=api_{{DIR}}.md proto/foo/v1/*.proto proto/foo/v2/*.proto
```
output:
```
./docs/api_v1.md
./docs/api_v2.md
```

### Option variables - `DIR_PARENT`
```bash
protoc --docmaster_out=format=markdown,file=api_{{DIR_PARENT_}.md proto/foo/v1/*.proto proto/bar/v1/*.proto
```
output:
```
./docs/api_foo.md
./docs/api_bar.md
```

### Prototool

```yaml
plugins:
    - name: docmaster
      flags: format=markdown,file={{DIR_PARENT}}
      output: ./docs
```
output:
```
./docs/foo.md
./docs/bar.md
```