# Architecture

This project uses a monorepo architecture whereby the distinct compmonents of the application suite are held within a single, top-level package using [workspaces](https://classic.yarnpkg.com/lang/en/docs/workspaces/)[^1].

## Packages

### `lib`

The core framework functionality - e.g. vector mathematics, graphics rendering, etc.

### `syntax`

Extended syntax grammar and parser generator program.

### `test`

Program for measuring the complexity of code written using various graphics libraries.

### `example`

An example application that uses the developed framework.

[^1]: [ADR001 - Project Structure](./log/ADR001-project-structure.md)
