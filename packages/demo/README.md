# @grafig/demo

![Version](https://img.shields.io/github/package-json/v/atpe/grafig?label=@grafig/demo&filename=packages/demo/package.json) ![@grafig/lib](https://img.shields.io/badge/@grafig/lib-grey) ![@grafig/syntax](https://img.shields.io/badge/@grafig/syntax-grey) ![Parcel](https://img.shields.io/badge/Parcel-grey)

A simple example application using the grafig framework.

## Quick Start

Clone the repository and run the demo application:

```sh
git clone https://github.com/atpe/grafig /path/to/repo
cd /path/to/repo/packages/demo && yarn start
```

## Packages

The application contains two example dynamic graphics and bundles them using a parent `index.html` file.

### Ulam Spiral

Renders the Ulam Sprial to the canvas.

![Ulam spiral](img/Ulam%20spiral%20graphic.png)

### Fenstring

Renders a chess board to the canvas, showing the fenstring given as a query parameter - i.e. `/fenstring?fenstr=8/Q6p/6p1/5p2/5P2/2p3P1/3r3P/2K1k3`

![Chess board](img/Chess%20board%20graphic.png)
