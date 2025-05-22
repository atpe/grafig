import FigScript, { Renderer } from '@grafig/lib'

// renders a piece to the canvas
export function Piece({ position, radius, type, white }) {
  return Renderer {
    circle({
      position,
      radius,
      fill: white ? 'white' : '#0F4761',
      stroke: '#0F4761'
    })
    text({
      text: type,
      position: position + (-15, 15),
      fill: white ? '#0F4761' : 'white',
      font: '42px sans-serif',
    })
  }
}

// renders a tile to the canvas
export function Tile({ position, white, piece }) {
  return (context) => {
    const min = Math.min(context.dimensions.x, context.dimensions.y)
    const size = min / 8
    const offset = (context.dimensions - min) / 2

    context.square({
      position: position * size + offset,
      size,
      fill: white ? 'white' : '#999999'
    })

    if (piece) {
      Piece({
        ...piece,
        position: (position + 0.5) * size + offset,
        radius: size / 3,
      })(context)
    }
  }
}