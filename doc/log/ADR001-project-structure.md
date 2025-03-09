# ADR001 - Project Structure

![image](https://badgen.net/badge/status/accepted/green) ![image](https://badgen.net/badge/date/02-03-25/grey)

| Summary |
| --- |
| The project will use a monorepo structure, to adhere to the project submission rules, with a top-level package that leverages workspaces to facilitate nested package management. |

## Context

The project contains multiple distinct projects and, due to it being an assessed peice of work, all parts are required to be submitted within a single repository - i.e. a monorepo.

## Considerations

Package managers such as `npm` and `yarn` allow packages to use so-called [workspaces](https://docs.npmjs.com/cli/v11/using-npm/workspaces) which "_provides support for managing multiple packages from your local file system from within a singular top-level, root package_". Tools such as [Turbo](https://turbo.build) and [Nx](https://nx.dev) provide utilities suchs as caching and performant bundling, which can facilitate monorepo development, albeit at the expense of simplicity.

## Decision

Given that the use of a monorepo is imperative, native workspaces will be used to facilitate development. The use of additional tooling will not be included, but rather remain a possibility for the future should the need arise.

## Consequences

### Expected Benefits

- improves package dependency management and installation
- simplifies common boiler-plate interactions
- provides hoisting for global CI/CD dependencies

### Potential Impacts

- requires strict standards due to increased project complexity
- creates more potential for conflict when working collaboratively
- provides a platform for potentially restrictive rules on subpackages
