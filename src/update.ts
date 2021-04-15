import { field } from '@lolpants/jogger'
import execa from 'execa'
import fs from 'node:fs'
import type { PathLike } from 'node:fs'
import { access } from 'node:fs/promises'
import { errorField, wrapLogger } from './logger.js'

export const exists = async (path: PathLike) => {
  try {
    await access(path, fs.constants.F_OK)
    return true
  } catch (error: unknown) {
    // @ts-expect-error
    if (error instanceof Error && error.code === 'ENOENT') return false
    throw error
  }
}

const logger = wrapLogger('updater')

const logRX = /^\[\d{2}:\d{2}:\d{2}\.\d{3}] \[update *] \[ *INFO] (.+)$/
const parseLogEntry = (line: string) => {
  const trimmed = line.trim()
  const entry = logRX.exec(trimmed)
  if (entry !== null) {
    return entry[1]
  }

  return trimmed
}

export const syncDB = async (boot: boolean) => {
  if (boot) {
    const dbExists = await exists('./edts/edtslib/data/edts.db')
    if (dbExists) return
  }

  try {
    const job = execa('python3', ['./edts/update.py'], {
      all: true,
      buffer: false,
    })

    job.all?.on('data', (buf: Buffer) => {
      const string = buf.toString('utf-8')
      const message = parseLogEntry(string)

      logger.info(field('status', 'updating'), field('message', message))
    })

    await job
    logger.info(field('status', 'complete'))
  } catch (error: unknown) {
    if (error instanceof Error) {
      logger.error(errorField(error))
    } else {
      logger.error(field('status', 'unknown error'))
    }
  }
}
