
import { Context } from './Context'


class CrisisCoreFusionError extends Error {

  isCrisisCoreFusionError = true

  sdk = 'CrisisCoreFusion'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  CrisisCoreFusionError
}

