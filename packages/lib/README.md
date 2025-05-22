# @grafig/lib

![Version](https://img.shields.io/github/package-json/v/atpe/grafig?label=@grafig/lib&filename=packages/lib/package.json) ![TypeScript](https://img.shields.io/badge/TypeScript-grey?logo=typescript) ![Parcel](https://img.shields.io/badge/Parcel-grey) ![ESLint](https://img.shields.io/badge/eslint-grey?logo=eslint) ![Jest](https://img.shields.io/badge/jest-grey?logo=jest)

A graphics rendering library able to be used in isolation, or in tandem with the [extended syntax](../../docs/figscript.md).

## Quick Start

Clone the repository and build the library:

```sh
git clone https://github.com/atpe/grafig /path/to/repo
cd /path/to/repo/packages/lib && yarn build
```

For further help on the available package scripts, run:

```sh
yarn run help
```

## Development

To start a development server that watches for file changes, and build the library on save, run:

```sh
yarn watch
```

### Documentation

To generate and serve documentation, run:

```sh
yarn docs --serve --open
```

### Testing

![branches](./coverage/badge-branches.svg) ![lines](./coverage/badge-lines.svg) ![functions](./coverage/badge-functions.svg) ![statements](./coverage/badge-statements.svg)

To test the package, simply run the following command:

```sh
yarn test
```

>This command can accept any of the flags given in the [jest CLI documentation](https://jestjs.io/docs/cli#options).

#### Dependencies

| Dependency | Description | Notes |
|---|---|---|
| [jest](https://github.com/jestjs/jest) | Provides the core testing framework | |
| [jsdom](https://github.com/jsdom/jsdom) | Provides DOM functionality  | |
| [jest-environment-jsdom](https://github.com/jestjs/jest) | Provides jsdom environment | see [DOM Manipulation](#dom-manipulation)  |
| [canvas](https://github.com/Automattic/node-canvas) | Provides HTMLCanvasElement functionality | see [DOM Manipulation](#dom-manipulation) |

#### DOM Manipulation

To test interactions with the DOM, `jest` allows the use of a `jsdom` [test environment](https://jestjs.io/docs/configuration#testenvironment-string) to provide client-side functionality to test suites running in `node`.

The `jsdom` package does not include an implementation for the HTMLCanvasElement so to test its use, this project uses the `canvas` package to provide the required functionality.

> The `canvas` package relies on [compiled binaries](https://github.com/Automattic/node-canvas?tab=readme-ov-file#compiling) which may not work for some operating systems without additionally installing [Cairo](https://cairographics.org/).
