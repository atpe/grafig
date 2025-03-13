# ADR003 - Parser Target

![image](https://badgen.net/badge/status/draft/grey) ![image](https://badgen.net/badge/date/11-03-25/grey)

| Summary |
| --- |
| The project will use a monorepo structure, to adhere to the project submission rules, with a top-level package that leverages workspaces to facilitate nested package management. |

## Context

The ANTLR parser generator provides several target languages such as Java, C#, JaveScript, and more. This allows a certain degree of freedom when deciding how to best implement the transpiler.

## Considerations

Given that this package is intended for use by developers when creating dynamic graphics, it should seamlessly integrate into the existing workflow of web application development. Throughout the project, this package will be used to analyse the complexity of applied graphics libraries and should consider ease of use, offering a simplistic interface able to be easily included in package scripts.

Performance of the transpiler is a key consideration as slow performance could hinder development more than it helps, should it be incorporated into a developer's workflow or CI/CD pipeline. While JavaScript and Typescript are both used within this project and suitable target languages, other targets such as C++ and Golang may offer greater performance due to their low-level capabilities and compilation to a binary code.

On top of the performance, cross-platform portability should be assessed as developers and containerisation tools use various platforms, of which transpilation should be independent. Target languages such as Java and python can support this, as they run inside virtual machines, while those such Swift and Golang also offer cross-platform compilation and low-level interoperability.

## Decision

Given that the two main concerns are performance and portability, Golang can provide the most performant option while still supporting cross-platfrom transpilation. Additionally, of the languages that are similarly suitable, developer experience is far greater with Golang and would better suit a time-bound development cycle.

## Consequences

### Expected Benefits

- creates a cross-platform transpiler
- creates a performant transpiler
- reduces learning needed for development

### Potential Impacts

- introduces complexity into the monorepo with a new language
- requires local installation of the Golang distribution
