import * as dotenv from 'dotenv';
dotenv.config();

import * as chalk from 'chalk';
import fastify from 'fastify';
import { registerRoutes } from '@root/routes';
import Xedule from './xedule';
import Logger, { LoggerLogLevel } from './logger';

const PORT = Number.parseInt(process.env.PORT) || 8080;
const LOGGER = !!process.env.LOGGER || false;

const app = fastify({ logger: LOGGER });

const main = async () => {
  const logger = Logger.getInstance('Main', chalk.magenta.bold);

  try {
    registerRoutes(app);
    Xedule.getInstance().startAuthMonitor();

    logger.log(`Listening to ${PORT}`);
    await app.listen({ port: PORT });
  } catch (err) {
    logger.log(err, LoggerLogLevel.Warn);
    process.exit(1);
  }
};

main();
