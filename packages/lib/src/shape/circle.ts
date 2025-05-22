import { styled, StyleProps } from '../style'
import { Vec2 } from '../vector'

/** The circle props object */
export interface CircleProps extends StyleProps {
  /** The center position of the circle */
  position: Vec2
  /** The circle radius */
  radius: number
}

/**
 * Creates a function bound to the given canvas that draws a circle.
 *
 * @param canvas the target canvas element
 * @returns the circle function
 */
export function circle(
  canvas: HTMLCanvasElement
): (props: CircleProps) => void {
  const context = canvas.getContext('2d')
  if (context === null) throw new Error('canvas context cannot be null')

  return styled(context, (props) => {
    context.beginPath()
    context.ellipse(
      props.position.x,
      props.position.y,
      props.radius,
      props.radius,
      0,
      0,
      Math.PI * 2
    )
    context.closePath()
    context.fill()
    context.stroke()
  })
}
