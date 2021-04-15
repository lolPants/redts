import 'source-map-support/register.js'

import { field } from '@lolpants/jogger'
import { PORT } from './env/index.js'
import { logger } from './logger.js'
import { app } from './server.js'
import { syncDB } from './update.js'

const boot = async () => {
  await syncDB(true)

  app.listen(PORT).on('listening', () => {
    logger.info(field('status', 'ready'), field('port', PORT))
  })
}

void boot()
