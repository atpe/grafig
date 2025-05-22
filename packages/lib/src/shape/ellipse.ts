import { styled, StyleProps } from '../style'
import { Vec2 } from '../vector'

/** The ellipse props object */
export interface EllipseProps extends StyleProps {
  /** The center position of the ellipse */
  position: Vec2
  /** The ellipse radius */
  radius: Vec2
  /** The rotation of the ellipse */
  rotation?: number
  /** The angle at which to start */
  start?: number
  /** The angle at which to stop */
  stop?: number
}

/**
 * Creates a function bound to the given canvas that draws a ellipse.
 *
 * @param canvas the target canvas element
 * @returns the ellipse function
 */
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
