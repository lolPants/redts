import execa from 'execa'
import { SCRIPT_TIMEOUT } from './env/index.js'
import { isExecaError } from './execa.js'

// #region Allowed Scripts
const ALLOWED_SCRIPTS = [
  'close_to',
  'coords',
  'direction',
  'distance',
  'edts',
  'find',
  'fuel_usage',
  'galmath',
] as const

type AllowedScript = typeof ALLOWED_SCRIPTS[number]

export const resolveScript = (string: string) => string.replace(/-/g, '_')
export function isScript(string: string): string is AllowedScript {
  // @ts-expect-error
  return ALLOWED_SCRIPTS.includes(string)
}
// #endregion

// #region Script Runner
interface ScriptReturn {
  success: boolean
  timeout: boolean
  stdout: string
}

export const runScript: (
  script: AllowedScript,
  args: readonly string[]
) => Promise<ScriptReturn> = async (script, args) => {
  try {
    const job = execa('python3', [`./edts/${script}.py`, ...args], {
      all: true,
    })

    const timeout = setTimeout(() => {
      job.kill('SIGKILL')
    }, SCRIPT_TIMEOUT * 1000)

    const { stderr, all, exitCode } = await job
    clearTimeout(timeout)

    const error = stderr !== '' || exitCode !== 0
    const success = !error

    return { success, stdout: all ?? '', timeout: false }
  } catch (error: unknown) {
    if (isExecaError(error)) {
      return {
        success: false,
        stdout: error.all ?? '',
        timeout: error.killed || error.isCanceled,
      }
    }

    throw error
  }
}
// #endregion
