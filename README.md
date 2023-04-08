# Openapi for 5GCoreNetSDK
    
[![Go Reference](https://pkg.go.dev/badge/github.com/5GCoreNet/openapi.svg)](https://pkg.go.dev/github.com/5GCoreNet/openapi)
[![Go Report Card](https://goreportcard.com/badge/github.com/5GCoreNet/openapi)](https://goreportcard.com/report/github.com/5GCoreNet/openapi)
[![License](https://img.shields.io/github/license/5GCoreNet/openapi)](LICENSE)
[![Release](https://img.shields.io/github/v/release/5GCoreNet/openapi)](https://github.com/5GCoreNet/openapi/releases/latest)

> The code in this repository is generated automatically using a GitHub Action workflow. DO NOT EDIT THE CODE IN THIS REPOSITORY.
> Note that for the moment the automatic code generation is not working, we will be able to provide it when [this PR](https://github.com/OpenAPITools/openapi-generator/pull/13472) will be merged. 
This repository contains the Go code generated from the 5GC OpenAPI specifications provided [here](https://github.com/jdegre/5GC_APIs) used by the [5GCoreNetSDK](https://github.com/5GCoreNet/5GCoreNetSDK).

We use the [OpenAPI Generator](https://openapi-generator.tech/) to generate the code from the OpenAPI specifications automatically using a GitHub Action workflow 
that runs every month and creates a Pull Request with the generated code if there are any changes. 

Feel free to create an issue if there are any updates to the OpenAPI specifications we will update the code in this repository manually.

Also, it is possible to run the code generation locally, see [Run locally](#run-locally) for more information.

> Feel free to use the code in this repository for your own projects. If you do, please consider giving us a star :star:.

## Run locally

Sometimes you may want to generate the code locally. This is possible by following the steps below.

### Prerequisites

The following tools are required to run the code generation locally:

- [docker](https://docs.docker.com/get-docker/)
- [act](https://github.com/nektos/act#installation)

> Act is a tool to run GitHub Actions locally. It is used to run the GitHub Action workflow that generates the code.

### Clone

As the code generation is done in a GitHub Action workflow, the repository needs to be cloned first.

```bash
git clone https://github.com/5GCoreNet/openapi
git submodule update --init --recursive
```

### Run

To run the code generation locally, run the following command:

```bash
cd openapi
act
```

This will generate the code in the current directory. Be aware that the code generation may overwrite the existing code.

## License

This project uses the [Apache 2.0](LICENSE) license.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).
