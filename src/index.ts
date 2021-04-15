import 'source-map-support/register.js'

import { field } from '@lolpants/jogger'
import { PORT } from './env/index.js'
import { logger } from './logger.js'
import { app } from './server.js'

app.listen(PORT).on('listening', () => {
  logger.info(field('status', 'ready'), field('port', PORT))
})
