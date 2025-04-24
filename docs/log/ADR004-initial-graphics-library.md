# ADR004 - Initial Graphics Library

![image](https://badgen.net/badge/status/proposed/blue) ![image](https://badgen.net/badge/date/08-04-25/grey)

| Summary |
| --- |
| The first iteration will implement a core rendering service and an initial vector module within the graphics library, providing the initial foundations for the framework. |

## Context

To begin manipulating canvases, there should be a means with which to access their contexts and find or insert them into the DOM.

## Considerations

The framework will be built upon the graphics library and, as such, it initialises the architecture of the entire framework. The library should be generic enough to offer a wide variety of functionality for different uses, but also reduce the complexity of canvas interactions. The consistent use of typing should be used and offered by the library and be suitably modular to allow extension in the future. The development should additionally consider the potential for future functionality such as animation and shaders.

## Decision

Firstly, in order to have consistent typing throughout the library, the use of TypeScript is imperative, and allows the use of types and interfaces that remove the dependency on classes. The library will be as modular as possible, splitting concerns into minimally dependent modules that can be used in isolation where possible.

The first iteration will implement the minimum functionality for the following modules:

### Rendering

The rendering module will include any functionality that in core to the use of the framework. Where possible, functionality should not be separated into its own module. It will intially contain functionality for creating and accessing canvases and their contexts, as well as basic styling functionality such as position and size.

### Vector

Given that the setup of a canvas will require providing dimension values, an initial vector library will be developed to promote the consistency of types from the offset. The vector types will aim to be inclusive of several syntaxes, which should support the extended syntax when developed.

## Consequences

This initial version will permit the intuative addition and manipulation of a canvas element within the DOM using JavaScript or TypeScript. It will offer a small functionality set for the creation and manipulation of vectors, providing a baseline for their extension in future iterations.
