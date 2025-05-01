import { StyleProps, TypographyProps } from './props'

export function styled<P extends StyleProps & TypographyProps>(
  context: CanvasRenderingContext2D,
  fn: (props: P) => void
): (props: P) => void {
  return (props) => {
    context.fillStyle = props.fill || 'rgba(0,0,0,0)'
    context.strokeStyle = props.stroke || 'rgba(0,0,0,0)'
    context.lineWidth = props.weight || 1
    context.font = props.font || '12px sans-serif'
    fn(props)
  }
}
