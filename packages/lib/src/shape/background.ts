export function background(
  canvas: HTMLCanvasElement
): (fill?: CanvasRenderingContext2D['fillStyle']) => void {
  const context = canvas.getContext('2d')
  if (context === null) throw new Error('canvas context cannot be null')

  return (fill) => {
    context.fillStyle = fill || 'black'
    context.fillRect(0, 0, canvas.width, canvas.height)
  }
}
