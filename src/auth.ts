import { ReasonPhrases, StatusCodes } from 'http-status-codes'
import type { Middleware } from 'koa'
import { createHmac } from 'node:crypto'
import { HMAC_SECRET } from './env/index.js'

export const auth: Middleware = async (ctx, next) => {
  if (HMAC_SECRET === undefined) return next()

  const { authorization } = ctx.request.headers
  if (authorization === undefined || authorization === '') {
    ctx.status = StatusCodes.UNAUTHORIZED
    ctx.body = ReasonPhrases.UNAUTHORIZED

    return
  }

  const [username, password] = authorization.replace(/^Bearer /, '').split(':')
  if (username === '' || password === '') {
    ctx.status = StatusCodes.UNAUTHORIZED
    ctx.body = ReasonPhrases.UNAUTHORIZED

    return
  }

  const digest = createHmac('sha256', HMAC_SECRET)
    .update(username)
    .digest('base64')

  if (password !== digest) {
    ctx.status = StatusCodes.UNAUTHORIZED
    ctx.body = ReasonPhrases.UNAUTHORIZED

    return
  }

  return next()
}
