# Cloud native example

This is an example of how to 'become a cloud native' (Have a familiarity with the development pipeline used and orchestration)

_This repo is intended for educational purposes_

The following technologies are used within this repo:
- Skaffold
- Docker
- Kubernetes
- Helm

### Overview

The project contains source code to a simple docker image
`tibbar/cloud-native-example` which runs a simple blacklist API.
You can interact with this API when it's running via the spec `swagger/swagger.yaml`

This project contains the dockerfile to build from the source code
and the Skaffold configuration to develop against a local k8s cluster.

### Orchestration

The k8s folder has three examples
- local (for skaffold to use and run a pod)
- vanilla (a barebones simple kubernetes deployment)
- helm (a helmified k8s release)
```
k8s
├── helm
│   └── cloud-native-example
│       ├── Chart.yaml
│       ├── charts
│       ├── templates
│       │   ├── NOTES.txt
│       │   ├── _helpers.tpl
│       │   ├── deployment.yaml
│       │   ├── ingress.yaml
│       │   ├── service.yaml
│       │   └── tests
│       │       └── test-connection.yaml
│       └── values.yaml
├── local
│   └── k8s-local-dev.yaml
└── vanilla
    ├── deployment.yaml
    └── service.yaml

```

### Usage

- I recommend you run `skaffold dev` make some code changes and see the image rebuild.
- You could then tag and push the new image.
- Following this you could try a kubernetes apply on a local cluster.
- Subsequently using helm to create then patch a release.