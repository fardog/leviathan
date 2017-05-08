# leviathan

Easily encode/decode Kubernetes "Secrets" yaml.

## Installation

```bash
go get -u github.com/fardog/leviathan
```

## Usage

Decode secrets in a yaml file output from Kubernetes

```bash
kubectl get secret/my-secret -o yaml > output.yaml
leviathan output.yaml
# some editing to output.yaml occurrs
leviathan --encode output.yaml
```

You can also pipe into `leviathan` directly:

```bash
kubectl get secret/my-secret -o yaml | leviathan
```

## License

[Apache-2.0](./LICENSE)
