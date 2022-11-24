import { ChalkFunction } from 'chalk';
import chalk = require('chalk');

export enum LoggerLogLevel {
  Log = 'log',
  Warn = 'warn',
  Error = 'err',
}

export default class Logger {
  private static instance: Logger;
  private color: Function;
  private prefix: string;

  constructor(prefix: string, color: ChalkFunction) {
    this.prefix = prefix;
    this.color = color;
  }

  log(message: string, level: LoggerLogLevel = LoggerLogLevel.Log) {
    const levelColor = this.getColorByLevel(level);

    console[level](`${this.color(this.prefix)} ${levelColor(message)}`);
  }

  private getColorByLevel(level: LoggerLogLevel): ChalkFunction {
    if (level === LoggerLogLevel.Log) {
      return chalk.white;
    }

    if (level === LoggerLogLevel.Warn) {
      return chalk.yellow;
    }

    if (level === LoggerLogLevel.Error) {
      return chalk.red;
    }
  }

  // Singleton
  static getInstance(prefix: string, color: ChalkFunction): Logger {
    return new Logger(prefix, color);
  }
}
