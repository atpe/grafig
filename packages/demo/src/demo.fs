import FigScript, { Canvas, Renderer } from '@grafig/lib'

function scale(dimensions) {
  return dimensions.xx / (2, 8)
}

const position = (100, 250)
const transparent = 'rgba(255, 255, 255, 0.5)'

function Text(color, shadow) {
  return Renderer {
    text({
      text: `this canvas is ${color}`,
      position: shadow ? position - 10 : position,
      width: scale(it.dimensions).x,
      fill: shadow ? transparent : 'white',
      font: '32px sans-serif'
    })
  }
}

function Message(fill) {
  return Renderer {
    background({ fill })
    apply(Text(fill, true))
    apply(Text(fill))
  }
} 

Canvas({ parent: document.getElementById('canvas-a') }, Message('black'))()
Canvas({ parent: document.getElementById('canvas-b') }) {
  apply(Message('blue'))
  rect({ 
    position: position + (-10, 40),
    dimensions: (220, 30),
    stroke: 'white',
    weight: 2,
  })
  text({
    text: 'but this one has extra stuff!', 
    position: position + (0, 60),
    fill: 'white'
  })
  apply((ctx) => {
    for (let i = 1; i <= 10; i++) {
      const center = position + (230 + 15 * i, 55)
      ctx.circle({
        position: center,
        radius: 15,
        stroke: 'white',
        fill: `rgba(255, 255, 255, ${1-i/10})`
      })
    }
  })
}()