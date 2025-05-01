import { styled, StyleProps, TypographyProps } from '../style'
import { Vec2 } from '../vector'

export interface TextProps extends StyleProps, TypographyProps {
  text: string
  position: Vec2
  width?: number
}

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
