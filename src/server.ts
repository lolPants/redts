import Router from '@koa/router'
import Koa from 'koa'

const app = new Koa()
const router = new Router()

app.use(router.routes())
app.use(router.allowedMethods())

export { app }
