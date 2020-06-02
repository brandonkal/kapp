## build//\_default

Quick build

```sh
set -e -x -u

# makes builds reproducible
export CGO_ENABLED=0
repro_flags="-ldflags=-buildid= -trimpath"
go build $repro_flags -o kapp ./cmd/kapp/...
./kapp version
```
