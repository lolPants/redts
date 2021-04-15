import execa from 'execa'
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
  stdout: string
}

export const runScript: (
  script: AllowedScript,
  args: readonly string[]
) => Promise<ScriptReturn> = async (script, args) => {
  try {
    const { stderr, all, exitCode } = await execa(
      'python3',
      [`./edts/${script}.py`, ...args],
      {
        all: true,
      }
    )

    const error = stderr !== '' || exitCode !== 0
    const success = !error

    return { success, stdout: all ?? '' }
  } catch (error: unknown) {
    if (isExecaError(error)) {
      console.log(error)
      return { success: false, stdout: error.all ?? '' }
    }

    throw error
  }
}
// #endregion
