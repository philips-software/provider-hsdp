# Crossplane HSDP Provider

`provider-hsdp` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for the
HSDP API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/philips-software/provider-hsdp):
```
kubectl crossplane install provider loafoe/provider-hsdp:v0.3.0 provider-hsdp
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-hsdp
spec:
  package: xpkg.upbound.io/philips-software/provider-hsdp:v0.4.6
EOF
```

## Credentials

Provider secrets are passed via the `ProviderConfig` resource which in turn 
refers to a Kubernetes secret holding HSDP credentials

```yaml
apiVersion: hsdp.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: example-creds
      namespace: crossplane-system
      key: credentials
```

### Secret example

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: example-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "service_id": "iam.service.id.here",
      "service_private_key": "----PRIVATE KEY HERE----",
      "region": "us-east",
      "environment": "client-test",
      "debug_log": "/tmp/crossplane.log"
    }
```

### Supported credential keys

| credential            | description                                                         | Example                 |
|-----------------------|---------------------------------------------------------------------|-------------------------|
| `service_id`          | Service ID of the IAM Service account to use                        |                         |
| `service_private_key` | The RSA private key associated with the IAM Service account         |                         |
| `region`              | The HSDP Region to use                                              | `us-east` or `eu-west`  |
| `environment`         | The HSDP Environment to use                                         | `client-test` or `prod` |
| `debug_log`           | Optional path where debug API traffic is logged in the provider Pod |                         |


## API Reference

You can see the API reference [here](https://doc.crds.dev/github.com/philips-software/provider-hsdp)

## Known limitations

* Fields which are marked with `ForceNew` in the Terraform provider do not trigger recreation of 
resources via Crossplane currently. This is a [known issue](https://github.com/upbound/upjet/issues/78) and will be addressed once [CRD Validation rules suport](https://kubernetes.io/blog/2022/09/23/crd-validation-rules-beta/) becomes
widely available starting in Kubernetes 1.25+

## Developing the provider

Refer to the [DEVELOPMENT](DEVELOPMENT.md) page for details.

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/philips-software/provider-hsdp/issues).
