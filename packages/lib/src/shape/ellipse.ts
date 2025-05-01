import { styled, StyleProps } from '../style'
import { Vec2 } from '../vector'

export interface EllipseProps extends StyleProps {
  position: Vec2
  radius: Vec2
  rotation?: number
  start?: number
  stop?: number
}

export function ellipse(
  canvas: HTMLCanvasElement
): (props: EllipseProps) => void {
  const context = canvas.getContext('2d')
  if (context === null) throw new Error('canvas context cannot be null')

  return styled(context, (props) => {
    context.beginPath()
    context.ellipse(
      props.position.x,
      props.position.y,
      props.radius.x,
      props.radius.y,
      props.rotation || 0,
      props.start || 0,
      props.stop || Math.PI * 2
    )
    context.closePath()
    context.fill()
    context.stroke()
  })
}
