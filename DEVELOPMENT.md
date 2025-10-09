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

First login to Upbound:

```shell
up login
```

### authenticate

If your session has expired you need to authenticate against the private `xpkg.upbound.io` repository first

Session token can be found like this [(see details)](https://github.com/crossplane/crossplane/issues/4785#issuecomment-1759953441):

```shell
cat ~/.up/config.json|jq -r .upbound.profiles.default.session
```

Then login:

```shell
$ docker login xpkg.upbound.io
username: _token
password: # paste the session token
```

Finally, publish:

```shell
make build.all publish.artifacts XPKG_REG_ORGS=xpkg.upbound.io/philips-software
```
