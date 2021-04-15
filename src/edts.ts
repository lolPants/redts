import execa from 'execa'
import { isExecaError } from './execa.js'

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
      'python',
      [`${script}.py`, ...args],
      {
        cwd: '../edts',
        all: true,
      }
    )

    const error = stderr !== '' || exitCode !== 0
    const success = !error

    return { success, stdout: all ?? '' }
  } catch (error: unknown) {
    if (isExecaError(error)) {
      return { success: false, stdout: error.all ?? '' }
    }

    throw error
  }
}
