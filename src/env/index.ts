import { registerBool, registerInt, registerString } from './register.js'

// #region Globals
const NODE_ENV = registerString('NODE_ENV')
const IS_PROD = NODE_ENV?.toLowerCase() === 'production'
export const IS_DEV = !IS_PROD
// #endregion

// #region Application
export const PORT = registerInt('PORT') ?? 3000
export const ENABLE_LOGGING = registerBool('ENABLE_LOGGING') ?? false
export const HMAC_SECRET = registerString('HMAC_SECRET')
export const SCRIPT_TIMEOUT = registerInt('SCRIPT_TIMEOUT') ?? 60
// #endregion
