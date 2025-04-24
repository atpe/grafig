# ADR002 - Complexity Analysis

![image](https://badgen.net/badge/status/accepted/green) ![image](https://badgen.net/badge/date/07-03-25/grey)

| Summary |
| --- |
| An initial complexity measurement tool is needed evaluate baseline metrics for comparable JavaScript libraries so that complexity can be continuously evaulated throughout development. |

## Context

In order to discern the success of the developed library and syntax, there must be a effective method for quantifiably measuring the complexity of projects that use the framework alongside libraries that provide equivalent functionality.

## Considerations

Code complexity is an extensively researched topic with several methodologies and metrics that have been suggested for its measurement. As such, there exist many tools that use these means to provide reporting on source code complexity and enforce complexity thresholds during development.

### Complexity Metrics

#### Source Lines Of Code (SLOC)

- correlates line count with complexity
- does not account for control structures

```text
D = 4.86 + 0.0018 * L

where:
  D = estimated defects in code
  L = lines of code
```

#### Cyclomatic Complexity

- examines the program control graph structure
- measures the number of paths through a program

```text
C = e - n + p

where:
  C = cyclomatic complexity
  e = number of edges in the program
  n = number of vertices in the program
  p = connected components
```

#### Cyclomatic Complexity Density

- examines both cyclomatic complexity and line count
- measures the ratio of complexity to lines of code

```text
Mᵨ = M / L

where:
  Mᵨ = cyclomatic complexity density
  C  = cyclomatic complexity
  L  = lines of code
```

#### Halstead Complexity Metrics

- examines the total and distinct quantities of operators and operands
- estimates vocabulary, volume, level, trouble, difficulty, and time

```text
η = η₁ + η₂
N = N₁ + N₂
V = N * log₂(η)
D = (η / 2) * (N₂ / η₂)
E = D * V

where:
  η  = program vocabulary
  N  = program length
  V  = volume
  D  = difficulty
  E  = effort
  η₁ = number of distinct operands
  η₂ = number of distinct operators
  N₁ = total number of operands
  N₂ = total number of operators
```

#### Maintainablility Index

- examines three dimensions:
  - control structure
  - information structure
  - typography, naming, and commenting
- estimates the maintainability of a system or component

```text
M = 171 - 5.2ln(V̄) - 0.23M̄ - 16.2 ln(L̄) + 50 sin(√2.46C)

where:
  M = maintainability
  V̄ = average Halstead volume
  M̄ = average cyclomatic complexity
  L̄ = average lines of code
  C = percentage of comments
```

### Fan In/Out and Efferent Coupling

- examines the number of references between components
- measures functionality abstraction and dependency coupling

```text
I = Fₒ / (Fᵢ + Fₒ)

where:
  I  = component instability
  Fₒ = number of outgoing references
  Fᵢ = number of incoming references
```

## Decision

To calculate complexity metrics for existing libraries that can be compared with that of a novel syntax, a generic method of measurement is required. The measurements should be quantitative and repeatable, implementing the considered methods to mimic the capability of similar analysis tools. The system would initially incorporate SLOC, cyclomatic complexity, and Halstead metrics, but could be expanded to include other metrics in the future.

To be able to gain a baseline evaluation of other libraries prior to creating the extended syntax, an initial parser for native JavaScript will need developing. Additionally, a visitor able to evaluate the complexity measurements is required to analyse programs written with each library.

## Consequences

### Expected Benefits

- facilitates objective comparison between libraries
- provides baseline values for complexity
- creates initial ANTLR parser program

### Potential Impacts

- requires considerable effort and time
- omits validation of calculated complexity metrics
