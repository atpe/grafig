import { styled, StyleProps } from '../style'
import { Vec2 } from '../vector'

export interface RectProps extends StyleProps {
  position: Vec2
  dimensions: Vec2
}

export function rect(canvas: HTMLCanvasElement): (props: RectProps) => void {
  const context = canvas.getContext('2d')
  if (context === null) throw new Error('canvas context cannot be null')

  return styled(context, (props) => {
    context.beginPath()
    context.rect(
      props.position.x,
      props.position.y,
      props.dimensions.x,
      props.dimensions.y
    )
    context.closePath()
    context.fill()
    context.stroke()
  })
}
