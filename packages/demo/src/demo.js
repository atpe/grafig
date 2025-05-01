// Code generated from src/demo.fs by @grafig/syntax. DO NOT EDIT.

import FigScript, { Canvas, Renderer } from '@grafig/lib'

function scale(dimensions) {
  return FigScript.div(FigScript.from(dimensions, 'x', 'x'), { x: 2, y: 8 })
}

const position = { x: 100, y: 250 }
const transparent = 'rgba(255, 255, 255, 0.5)'

function Text(color, shadow) {
  return Renderer((it) => {
    it.text({
      text: `this canvas is ${color}`,
      position: shadow ? FigScript.sub(position, 10) : position,
      width: scale(it.dimensions).x,
      fill: shadow ? transparent : 'white',
      font: '32px sans-serif'
    })
  })
}

function Message(fill) {
  return Renderer((it) => {
    it.background({ fill })
    it.apply(Text(fill, true))
    it.apply(Text(fill))
  })
} 

Canvas({ parent: document.getElementById('canvas-a') }, Message('black'))()
Canvas({ parent: document.getElementById('canvas-b') }, (it) => {
  it.apply(Message('blue'))
  it.rect({ 
    position: FigScript.add(position, { x: -10, y: 40 }),
    dimensions: { x: 220, y: 30 },
    stroke: 'white',
    weight: 2,
  })
  it.text({
    text: 'but this one has extra stuff!', 
    position: FigScript.add(position, { x: 0, y: 60 }),
    fill: 'white'
  })
  it.apply((ctx) => {
    for (let i = 1; i <= 10; i++) {
      const center = FigScript.add(position, { x: FigScript.add(230, FigScript.mult(15, i)), y: 55 })
      ctx.circle({
        position: center,
        radius: 15,
        stroke: 'white',
        fill: `rgba(255, 255, 255, ${FigScript.sub(1,FigScript.div(i,10))})`
      })
    }
  })
})()