# Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Update provider version

Edit `Makefile` and bump the `TERRAFORM_PROVIDER_VERSION` value accordingly, then:

Update the provider package:

```shell
go get -u  github.com/philips-software/terraform-provider-hsdp@vX.Y.Z
```

```shell
rm -fr .work
make generate build
```

## Publish to marketplace

```shell
up login

make build.all publish.artifacts XPKG_REG_ORGS=xpkg.upbound.io/philips-software
```
