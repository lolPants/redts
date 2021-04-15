import 'source-map-support/register.js'

import { PORT } from './env/index.js'
import { app } from './server.js'

app.listen(PORT)
