import { styled, StyleProps } from '../style'
import { Vec2 } from '../vector'

/** The rect props object */
export interface RectProps extends StyleProps {
  /** The position of the top-left vertex */
  position: Vec2
  /** The dimensions of the rectangle */
  dimensions: Vec2
}

/**
 * Creates a function bound to the given canvas that draws a rectangle.
 *
 * @param canvas the target canvas element
 * @returns the rect function
 */
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
