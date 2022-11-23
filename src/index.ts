import * as dotenv from 'dotenv';
dotenv.config();

import * as chalk from 'chalk';
import fastify from 'fastify';
import { registerRoutes } from '@root/routes';
import Xedule from './xedule';

const PORT = Number.parseInt(process.env.PORT) || 8080;
const LOGGER = !!process.env.LOGGER || false;

const app = fastify({ logger: LOGGER });

const main = async () => {
  try {
    registerRoutes(app);
    Xedule.getInstance().startAuthMonitor();

    console.log(chalk.blue(`Listening to ${PORT}`));
    await app.listen({ port: PORT });
  } catch (err) {
    console.error(chalk.bgRed.white(err));

    process.exit(1);
  }
};

main();
