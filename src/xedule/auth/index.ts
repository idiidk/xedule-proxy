import Logger, { LoggerLogLevel } from '@root/logger';
import * as chalk from 'chalk';
import { readFile, writeFile } from 'fs/promises';
import { join } from 'path';
import * as puppeteer from 'puppeteer';
import { Stream } from 'stream';
import XeduleEncryption, { XeduleEncryptionHash } from './encryption';

export default class XeduleAuthMonitor {
  private endpoint: string;
  private logger: Logger;
  private page: puppeteer.Page;
  private cookieFile: string = join(
    process.cwd(),
    'secrets',
    'xedule-cookie.json'
  );
  public cookieStream: Stream;

  constructor(endpoint: string) {
    this.endpoint = endpoint;
    this.logger = Logger.getInstance('AuthProxy', chalk.cyan.bold);
    this.cookieStream = new Stream.Writable();
  }

  public async monitor() {
    this.logger.log('Starting Xedule auth monitor...');

    // Start puppeteer
    const browser = await puppeteer.launch({ headless: true });
    this.page = await browser.newPage();

    // Load cookies from file
    this.logger.log('Loading cookies from file...');

    try {
      // Read the file and decrypt the cookies
      const encryptedCookieBuffer = await readFile(this.cookieFile);
      const encryptedCookies = JSON.parse(
        encryptedCookieBuffer.toString()
      ) as XeduleEncryptionHash;
      const decryptedCookieString = XeduleEncryption.decrypt(encryptedCookies);
      const decryptedCookies = JSON.parse(decryptedCookieString);

      // Set the cookies to be used by puppeteer
      this.page.setCookie(...decryptedCookies);

      // Send the cookies over the stream
      this.cookieStream.emit('cookies', decryptedCookies);
    } catch (err) {
      if (err.code === 'ENOENT') {
        this.logger.log('No cookie file found, skipping...', LoggerLogLevel.Error);
      } else {
        this.logger.log(err, LoggerLogLevel.Error);
      }
    }

    this.refreshCookie();
    // Start timer to refresh cookies
    setInterval(() => {
      this.logger.log('Refreshing Xedule cookie...');
      this.refreshCookie();
    }, Number.parseInt(process.env.XEDULE_INTERVAL) * 60 * 1000);
  }

  private async refreshCookie() {
    await this.page.goto(this.endpoint);

    // Not logged in, start the refresh procedure
    if (await this.page.url().startsWith('https://login.microsoft')) {
      this.logger.log('Relogging into Xedule...', LoggerLogLevel.Warn);

      await this.page.waitForNavigation();
      await this.typeUsername(process.env.XEDULE_USERNAME);
      await this.typePassword(process.env.XEDULE_PASSWORD);
      await this.staySignedIn();
      await this.signIn();
      await this.page.waitForNavigation();
    }

    while (!this.page.url().startsWith(this.endpoint)) {
      await new Promise((resolve) => setTimeout(resolve, 100));
    }

    const cookies = await this.getCookiesForAllDomains(this.page);
    this.cookieStream.emit('cookies', cookies);
    await this.saveCookies(cookies);
    
    this.logger.log('Refreshed auth cookie');
  }

  private async saveCookies(cookies) {
    this.logger.log('Saving cookies...');

    const encryptedCookies = XeduleEncryption.encrypt(JSON.stringify(cookies));
    await writeFile(this.cookieFile, JSON.stringify(encryptedCookies));

    this.logger.log('Saved cookies...');
  }

  private async getCookiesForAllDomains(page: puppeteer.Page) {
    const client = await page.target().createCDPSession();
    return (await client.send('Network.getAllCookies')).cookies;
  }

  private async signIn() {
    await this.page.waitForSelector('input[name=DontShowAgain]', {
      visible: true,
    });
    await this.page.click('input[data-report-event=Signin_Submit]');
  }

  private async staySignedIn() {
    await this.page.waitForSelector('input[name=DontShowAgain]', {
      visible: true,
    });
    await this.page.click('input[name=DontShowAgain]');
  }

  private async typeUsername(username: string) {
    await this.page.waitForSelector(
      'input[name=loginfmt]:not(.moveOffScreen)',
      {
        visible: true,
      }
    );
    await this.page.type('input[name=loginfmt]', username, { delay: 50 });
    await this.page.click('input[type=submit]');
  }

  private async typePassword(password: string) {
    await this.page.waitForSelector(
      'input[name=Password]:not(.moveOffScreen),input[name=passwd]:not(.moveOffScreen)',
      { visible: true }
    );
    await this.page.type('input[name=passwd]', password, { delay: 50 });
    await this.page.click('input[type=submit]');
  }
}
