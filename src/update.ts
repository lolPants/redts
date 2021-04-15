import { field } from '@lolpants/jogger'
import execa from 'execa'
import fs from 'node:fs'
import type { PathLike } from 'node:fs'
import { access } from 'node:fs/promises'
import { errorField, logger } from './logger.js'

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

const ctxField = field('ctx', 'update-db')

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
      const string = buf.toString('utf-8').trim()
      logger.info(
        ctxField,
        field('status', 'updating'),
        field('content', string)
      )
    })

    await job
    logger.info(ctxField, field('status', 'complete'))
  } catch (error: unknown) {
    if (error instanceof Error) {
      logger.error(ctxField, errorField(error))
    } else {
      logger.error(ctxField, field('status', 'unknown error'))
    }
  }
}
