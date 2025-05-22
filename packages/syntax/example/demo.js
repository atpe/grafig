// Code generated from example/demo.fs by @grafig/syntax. DO NOT EDIT.

const position = FigScript.add({ x: 200, y: FigScript.mult(100, 2) }, 100)

/*
  This is a multiline comment.

  With many lines of text.
*/
const canvas = createCanvas(useElementById('canvas-container'))
const bg = createRenderer(canvas, (it) => {
  it.background('black')
})

const scale = (dimensions) =>
  FigScript.div(FigScript.from(dimensions, 'x', 'x'), { x: 4, y: 10 })

const shadow = () => {
  canvas.text(
    'hello world!', // this is an inline comment
    'grey',
    FigScript.sub(position, 10),
    scale(canvas.dimensions)
  )
}

useRenderer(canvas, bg.render, shadow, (it) => {
  it.text('hello world!', 'white', position, scale(canvas.dimensions))
})

// this is a trailing comment
