import { Transformer } from '@parcel/plugin'
import proc from 'node:child_process'
import util from 'node:util'

const exec = util.promisify(proc.exec)
const cmd = '../syntax/bin/figsyntax transpile -s'

export default new Transformer({
  async transform({ asset }) {
    const source = await asset.getCode()
    const { stdout, stderr } = await exec(`${cmd} "${source}"`)

    if (stderr.length) {
      asset.diagnostics = parseError(asset.filePath, stderr)
      return []
    }

    asset.setCode(stdout)
    asset.type = 'js'

    return [asset]
  },
})

function parseError(filePath, message) {
  const match = message.match(/(.*)\((\d+):(\d+)\)$/s)
  if (match) {
    const [, msg, line, col] = match
    const pos = {
      line: Number(line),
      column: Number(col) + 1,
    }
    return {
      message: msg,
      codeFrames: [
        {
          filePath: filePath,
          codeHighlights: [{ start: pos, end: pos }],
        },
      ],
    }
  }
}
