import type { ExecaError } from 'execa'

export function isExecaError(object: unknown): object is ExecaError {
  if (typeof object !== 'object') return false
  if (object === null) return false

  // @ts-expect-error Type Assertion
  const record: Record<string, unknown> = object
  if ('shortMessage' in record === false) return false
  if ('command' in record === false) return false
  if ('exitCode' in record === false) return false
  if ('signal' in record === false) return false
  if ('stdout' in record === false) return false
  if ('stderr' in record === false) return false

  return true
}
