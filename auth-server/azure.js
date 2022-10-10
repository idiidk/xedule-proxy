"use strict";

const puppeteer = require("puppeteer");

/**
 *
 * @param {options.username} string username
 * @param {options.password} string password
 * @param {options.loginUrl} string password
 * @param {options.postLoginSelector} string a selector on the app's post-login return page to assert that login is successful
 * @param {options.headless} boolean launch puppeteer in headless more or not
 * @param {options.cookies} array array of initial cookies
 * @param {options.logs} boolean whether to log cookies and other metadata to console
 * @param {options.getAllBrowserCookies} boolean whether to get all browser cookies instead of just for the loginUrl
 */
module.exports.azureAdSingleSignOn = async function azureAdSingleSignOn(
  options = {}
) {
  validateOptions(options);

  const browser = await puppeteer.launch({ headless: !!options.headless });
  const page = await browser.newPage();

  await page.setCookie(...options.cookies);
  await page.goto(options.loginUrl);

  try {
    await page.waitForNavigation({ timeout: 1000 });
  } catch (e) {}

  if ((await page.$(options.postLoginSelector)) === null) {
    await typeUsername({ page, options });
    await typePassword({ page, options });

    try {
      await staySignedIn({ page, options });
    } catch (e) {
      if (options.logs) {
        console.log("Stay signed in skipped");
      }
    }
  }

  const cookies = await getCookies({ page, options });
  await finalizeSession({ page, browser, options });

  return {
    cookies,
  };
};

function validateOptions(options) {
  if (!options.username || !options.password) {
    throw new Error("Username or Password missing for login");
  }
  if (!options.loginUrl) {
    throw new Error("Login Url missing");
  }
  if (!options.postLoginSelector) {
    throw new Error("Post login selector missing");
  }
}

async function staySignedIn({ page, options } = {}) {
  await page.waitForSelector("input[name=DontShowAgain]", {
    visible: true,
    delay: 10000,
  });
  await page.click("input[name=DontShowAgain]");
  await page.click("input[data-report-event=Signin_Submit]");
}

async function typeUsername({ page, options } = {}) {
  await page.waitForSelector("input[name=loginfmt]:not(.moveOffScreen)", {
    visible: true,
    delay: 10000,
  });
  await page.type("input[name=loginfmt]", options.username, { delay: 50 });
  await page.click("input[type=submit]");
}

async function typePassword({ page, options } = {}) {
  await page.waitForSelector(
    "input[name=Password]:not(.moveOffScreen),input[name=passwd]:not(.moveOffScreen)",
    { visible: true, delay: 10000 }
  );
  await page.type("input[name=passwd]", options.password, { delay: 50 });
  await page.click("input[type=submit]");
}

async function getCookies({ page, options } = {}) {
  await page.waitForSelector(options.postLoginSelector, {
    visible: true,
    delay: 10000,
  });

  const cookies = options.getAllBrowserCookies
    ? await getCookiesForAllDomains(page)
    : await page.cookies(options.loginUrl);

  if (options.logs) {
    console.log(cookies);
  }

  return cookies;
}

async function getCookiesForAllDomains(page) {
  const client = await page.target().createCDPSession();
  return (await client.send("Network.getAllCookies")).cookies;
}

async function finalizeSession({ page, browser, options } = {}) {
  await browser.close();
}
