import { styled, StyleProps } from '../style'
import { Vec2 } from '../vector'

export interface SquareProps extends StyleProps {
  position: Vec2
  size: number
}

export function square(
  canvas: HTMLCanvasElement
): (props: SquareProps) => void {
  const context = canvas.getContext('2d')
  if (context === null) throw new Error('canvas context cannot be null')

  return styled(context, (props) => {
    context.beginPath()
    context.rect(props.position.x, props.position.y, props.size, props.size)
    context.closePath()
    context.fill()
    context.stroke()
  })
}
