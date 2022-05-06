import {
  createConsoleSink,
  createField,
  createLogger,
  field,
} from '@lolpants/jogger'
import type { Field } from '@lolpants/jogger'
import type { Middleware } from 'koa'
import { ENABLE_LOGGING, HMAC_SECRET, IS_DEV } from './env/index.js'

export const wrapLogger = (name: string) =>
  createLogger({
    name,
    sink: createConsoleSink(IS_DEV),
  })

const httpLogger = wrapLogger('http')
export const logger = wrapLogger('app')

export const errorField: <T extends Error>(
  error: T
) => Readonly<Field> = error => {
  const array: Array<Readonly<Field>> = [
    field('type', error.name),
    field('message', error.message),
  ]

  if (error.stack) array.push(field('stack', error.stack))
  return field('error', array[0], ...array.slice(1))
}

const httpVersionField = createField('httpVersion')
const methodField = createField('method')
const urlField = createField('url')
const statusField = createField('status')
const sizeField = createField('size')
const durationField = createField('duration')
const uaField = createField('userAgent')
const referrerField = createField('referrer')

export const middleware: Middleware = async (ctx, next) => {
  const before = Date.now()
  await next()
  const duration = (Date.now() - before) / 1000

  if (IS_DEV === true || ENABLE_LOGGING === true) {
    const fields = [
      httpVersionField(
        `${ctx.req.httpVersionMajor}.${ctx.req.httpVersionMinor}`
      ),
      methodField(ctx.method),
      urlField(ctx.url),
      statusField(ctx.status),
      sizeField(ctx.status === 204 ? 0 : ctx.response.length ?? -1),
      durationField(duration),
      uaField(ctx.headers['user-agent'] ?? '-'),
      referrerField(ctx.headers.referer ?? ctx.headers.referrer ?? '-'),
    ] as const

    const extraFields: Field[] = []

    if (HMAC_SECRET !== undefined) {
      const { authorization } = ctx.request.headers
      const split = authorization?.replace(/^Bearer /, '').split(':')

      const username = split?.shift()
      if (username) extraFields.push(field('username', username))
    }

    if (typeof ctx.request.body === 'string') {
      extraFields.push(field('query', ctx.request.body))
    }

    httpLogger.info(...fields, ...extraFields)
  }
}
