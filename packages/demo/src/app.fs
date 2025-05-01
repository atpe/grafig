import FigScript, { Canvas } from '@grafig/lib'
import { Spiral } from './ulam'

Canvas({ parent: document.body }) {
  background('black')
  apply(Spiral(it.dimensions, 4))
}()