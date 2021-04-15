import Router from '@koa/router'
import { ReasonPhrases, StatusCodes } from 'http-status-codes'
import Koa from 'koa'
import koaBody from 'koa-body'
import shlex from 'shlex'
import { auth } from './auth.js'
import { isScript, resolveScript, runScript } from './edts.js'
import { middleware as logger } from './logger.js'

const app = new Koa()
const router = new Router()
router.use(logger)
router.use(auth)

const body = koaBody({
  json: false,
  urlencoded: false,
  multipart: false,
  text: true,
  parsedMethods: ['GET'],
})

router.get('/api/:script', body, async ctx => {
  const script = resolveScript(ctx.params.script)
  if (!isScript(script)) {
    ctx.status = StatusCodes.NOT_FOUND
    ctx.body = ReasonPhrases.NOT_FOUND

    return
  }

  if (typeof ctx.request.body !== 'string') {
    ctx.status = StatusCodes.BAD_REQUEST
    ctx.body = ReasonPhrases.BAD_REQUEST

    return
  }

  const args = shlex.split(ctx.request.body)
  const { success, stdout, timeout } = await runScript(script, args)
  if (success === false) {
    if (timeout) {
      ctx.status = StatusCodes.INTERNAL_SERVER_ERROR
      ctx.body = 'Execution Timeout'

      return
    }

    ctx.status = StatusCodes.BAD_REQUEST
  }

  ctx.body = stdout
})

app.use(router.routes())
app.use(router.allowedMethods())

export { app }
