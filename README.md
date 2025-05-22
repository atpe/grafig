<p align="center">
  <img src="packages/demo/img/Chess%20board%20graphic.png" alt="Chess board graphic" width="200" align="center"/>
  <h1 align="center">grafig</h1>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/github-blue?logo=github" alt="GitHub" align="center"/>
  <img src="https://img.shields.io/github/package-json/v/atpe/grafig" alt="Version" align="center"/>
</p>

<p align="center" font="italic">
  A graphics rendering framework for client-side web applications.
</p>

_This project was developed alongside a [dissertation](docs/Exploring%20a%20graphics%20rendering%20framework%20for%20client-side%20web%20applications.pdf) submitted in partial fulfilment for the degree of Batchelor of Science in Computer Science._

## Quick Start

Clone the repo and run the demo application:

```sh
git clone https://github.com/atpe/grafig /path/to/repo
cd /path/to/repo && yarn demo
```

## Background

Computer graphics are an essential component within modern technology and their use in client-side applications has greatly improved web standards, allowing the flexibil-ity to dynamically generate complex animations, data visualisations, even online multi-player games. The HTML5 Canvas API, provides the foundation for many of these appli-cations however it can impact development and maintainability of code due to its very low-level implementation. This has prompted the development of libraries that all aim to simplify these interactions and similarly, this project aims to reduce complexity using a rendering framework that includes a cohesive graphics library alongside a domain-specific syntax that extends vanilla JavaScript.

This project explores the impact that programming languages can have, quantitatively comparing the code complexity of the developed system against that of the Canvas API and other libraries. In doing so the developed framework yielded an average reduction of 36% compared to its equivalent implementation using the Canvas API, enhancing the reduction made by the p5.js library by a further 25%. This effectively demonstrates that, by introducing syntax and patterns employed by other rendering solutions, the devel-opment experience and evaluated maintainability of client-side applications can be greatly influenced by the programming languages used to write them.

## Installation

Clone the repository:

```sh
git clone https://github.com/atpe/grafig /path/to/repo
```

If using Visual Studio Code, it may be useful at this point to open the `project.code-workspace` file (see [here](.vscode/project.code-workspace)) and open the project as a workspace with better formatting.

Build the internal packages:

```sh
cd /path/to/repo && yarn build
```

Note that this may fail if the resources required for building the transpiler binary are not available. The latest binary version is included in source control such that the project will work out of the box.

If only running the demo application, this can be done from anywhere within the project with:

```sh
yarn workspace @grafig/demo start
```

Similarly, to use the transpiler binary from anywhere within the project:

```sh
yarn workspace @grafig/syntax transpile [flags]
```

Both the `@grafig/lib` and `grafig/syntax` have comprehensive usage information, which can be seen by running:

```sh
yarn workspace @grafig/lib run help
# or
yarn workspace @grafig/syntax run help
```

### Packages

| Package | Description |
| --- | --- |
| [@grafig/demo](packages/demo/README.md) | An example application using the framework. |
| [@grafig/lib](packages/lib/README.md) | The core functionality for the framework. |
| [@grafig/syntax](packages/syntax/README.md) | The transpiler for the extended syntax. |
| [@grafig/test](packages/test/README.md) | The testing package for the framework. |

See [architecture.md](docs/architecture.md) for more detailed information.

### FigScript

See [figscript.md](docs/figscript.md) for the extended syntax reference.
