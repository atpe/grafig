import FigScript, {
  createRenderer,
  createCanvas,
  useElementById,
  useRenderer,
} from '@grafig/lib'

const position = (200, 100 * 2) + 100


/*
  This is a multiline comment.

  With many lines of text.
*/
const canvas = createCanvas(useElementById('canvas-container'))
const bg = createRenderer(canvas) {
  background('black')
}

const scale = (dimensions) => dimensions.xx / (4, 10)

const shadow = () => {
  canvas.text(
    'hello world!', // this is an inline comment
    'grey',
    position - 10,
    scale(canvas.dimensions)
  )
}

useRenderer(canvas, bg.render, shadow) {
  text('hello world!', 'white', position, scale(canvas.dimensions))
}

// this is a trailing comment