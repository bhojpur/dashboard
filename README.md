# Bhojpur Dashboard - Data Visualization

The `Bhojpur Dashboard` is a high-performance dashboard engine applied within the
[Bhojpur.NET Platform](https://github.com/bhojpur/platform) for delivery of web-scale
distributed applications or services. It is a web-based user interface for
[Bhojpur Application](https://github.com/bhojpur/application), allowing users to see
information, view logs, and more for the [Bhojpur Application](https://github.com/bhojpur/application),
components, and configurations running either locally or in a Kubernetes cluster.

## Key Features

The [Bhojpur Dashboard](https://github.com/bhojpur/dashboard) provides information about
[Bhojpur Application](https://github.com/bhojpur/application), components, configurations,
and control plane services. Users can view metadata, manifests, and deployment files, actors,
logs, and more on both Kubernetes and self-hosted platforms.

## Getting started

### Prerequisites

You need the following to be able to run the dashboard
- [Bhojpur Application](https://github.com/bhojpur/application) Runtime
- [Bhojpur Application](https://github.com/bhojpur/application) CLI

### User Interface Build

For compiling the web-based user interface developed using `Angular Frameework`, you must
have `node.js`, `angular.io`, etc installed. Then, type the following commands

```bash
$ cd pkg/webui
$ npm install
$ ng build
$ ng serve
```

Alternativerly, you can run `./build_standalone.sh` or `./build_kubernetes.sh` script.
### Installation

The [Bhojpur Dashboard](https://github.com/bhojpur/dashboard) comes pre-packaged with the
[Bhojpur Application](https://github.com/bhojpur/application) CLI. To learn more about the
dashboard command, use the CLI command `appctl dashboard -h`.

#### Kubernetes

Run `appctl dashboard -k`, or if you installed [Bhojpur Application](https://github.com/bhojpur/application) in a non-default namespace, `appctl dashboard -k -n <your-namespace>`.

#### Standalone

Run `appctl dashboard`, and navigate to http://localhost:8080.
