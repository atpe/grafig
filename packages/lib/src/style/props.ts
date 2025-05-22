/** The typography props object */
export interface TypographyProps {
  /** The text font */
  font?: CanvasRenderingContext2D['font']
}

/** The style props object */
export interface StyleProps {
  /** The fill color */
  fill?: CanvasRenderingContext2D['fillStyle']
  /** The stroke color */
  stroke?: CanvasRenderingContext2D['strokeStyle']
  /** The stroke weight */
  weight?: CanvasRenderingContext2D['lineWidth']
}
