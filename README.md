# Crossplane HSDP Provider

`provider-hsdp` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for the
HSDP API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/philips-software/provider-hsdp):
```
kubectl crossplane provider install philipssoftware/provider-hsdp:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-hsdp
spec:
  package: philipssoftware/provider-hsdp:v0.1.0
EOF
```

## Developing

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

## Known limitations

* Fields which are marked with `ForceNew` in the Terraform provider do not trigger recreation of 
resources via Crossplane currently. This is a [known issue](https://github.com/upbound/upjet/issues/78) and will be addressed once [CRD Validation rules suport](https://kubernetes.io/blog/2022/09/23/crd-validation-rules-beta/) becomes
widely available starting in Kubernetes 1.25+


## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/philipssoftware/provider-hsdp/issues).
