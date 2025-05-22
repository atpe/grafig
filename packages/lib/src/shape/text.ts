import { styled, StyleProps, TypographyProps } from '../style'
import { Vec2 } from '../vector'

/** The text props object */
export interface TextProps extends StyleProps, TypographyProps {
  /** The text to render */
  text: string
  /** The position of the text */
  position: Vec2
  /** The maximum width of the text */
  width?: number
}

/**
 * Creates a function bound to the given canvas that draws text.
 *
 * @param canvas the target canvas element
 * @returns the text function
 */
export function text(canvas: HTMLCanvasElement): (props: TextProps) => void {
  const context = canvas.getContext('2d')
  if (context === null) throw new Error('canvas context cannot be null')

  return styled(context, (props) => {
    context.fillText(
      props.text,
      props.position.x,
      props.position.y,
      props.width
    )
    context.strokeText(
      props.text,
      props.position.x,
      props.position.y,
      props.width
    )
  })
}
