import { styled, StyleProps } from '../style'
import { Vec2 } from '../vector'

/** The square props object */
export interface SquareProps extends StyleProps {
  /** The position of the top left vertex */
  position: Vec2
  /** The size of the square */
  size: number
}

/**
 * Creates a function bound to the given canvas that draws a square.
 *
 * @param canvas the target canvas element
 * @returns the square function
 */
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
