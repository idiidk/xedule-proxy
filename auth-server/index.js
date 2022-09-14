require("dotenv").config();
const puppeteer = require("puppeteer");
const fastify = require("fastify")({ logger: true });

(async () => {
  const { PORT, SECRET, XEDULE_URL, INIT_COOKIE, REFRESH_INTERVAL_M } =
    process.env;
  if (!PORT || !SECRET || !XEDULE_URL || !INIT_COOKIE || !REFRESH_INTERVAL_M) {
    console.error("not all env variables initialised");
    process.exit(1);
  }

  const cookies = INIT_COOKIE.split("; ")
    .map((e) => e.split("="))
    .map((e) => ({ name: e[0], value: e[1], domain: "sa-han.xedule.nl" }));

  const browser = await puppeteer.launch();
  const page = await browser.newPage();
  await page.setCookie(...cookies);
  const response = await page.goto(XEDULE_URL);

  const responseUrl = response.url();
  if (!responseUrl.startsWith(XEDULE_URL)) {
    console.error(`init cookie invalid?, got redirected to ${responseUrl}`);
    process.exit(1);
  }

  const statusCode = response.status();
  if (statusCode !== 200) {
    console.error(`init cookie invalid?, got status code ${statusCode}`);
    process.exit(1);
  }

  let currentCookies = await page.cookies();
  console.log("fetched first iteration cookies, we're in business");

  setInterval(async () => {
    console.log(`${new Date()} - refreshing cookies`);
    await page.reload({ waitUntil: ["networkidle0", "domcontentloaded"] });
    currentCookies = await page.cookies();
  }, 1000 * 60 * REFRESH_INTERVAL_M);

  fastify.get("/", (req) => {
    if (req.query.secret !== SECRET) {
      return { error: "invalid secret" };
    }

    const cookies = currentCookies.map((e) => ({
      name: e.name,
      value: e.value,
      domain: e.domain,
    }));
    const config = { xeduleUrl: XEDULE_URL };
    return { cookies, config };
  });

  await fastify.listen({ port: PORT });
})();
