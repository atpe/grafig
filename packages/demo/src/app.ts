import {
  createRenderer,
  createCanvas,
  useElementById,
  background,
  text,
} from '@grafig/lib'

const canvas = createCanvas(useElementById('canvas-container'))
const renderer = createRenderer(canvas)

renderer.apply(
  background('black'),
  text('hello world!', 'white', { x: 100, y: 100 }, { x: 200, y: 40 })
)
