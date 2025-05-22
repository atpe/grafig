import FigScript, { Canvas } from '@grafig/lib'
import { Spiral } from './ulam'

Canvas({ parent: document.body }) {
  background('white')
  apply(Spiral(it.dimensions, 10))
}()