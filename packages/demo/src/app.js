// Code generated from src/app.fs by @grafig/syntax. DO NOT EDIT.

import FigScript, { Canvas } from '@grafig/lib'
import { Spiral } from './ulam'

Canvas({ parent: document.body }, (it) => {
  it.background('black')
  it.apply(Spiral(it.dimensions, 4))
})()