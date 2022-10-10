require("dotenv").config();
const fastify = require("fastify")({ logger: true });

const { parse: parseUrl } = require("url");
const { existsSync, writeFileSync, readFileSync } = require("fs");
const { azureAdSingleSignOn } = require("./azure");

const COOKIES_FILE = ".cookies.json";

(async () => {
  const { PORT, SECRET, XEDULE_URL, USERNAME, PASSWORD, REFRESH_INTERVAL_M } =
    process.env;
  if (
    !PORT ||
    !SECRET ||
    !XEDULE_URL ||
    !USERNAME ||
    !PASSWORD ||
    !REFRESH_INTERVAL_M
  ) {
    console.error("not all env variables initialised");
    process.exit(1);
  }

  let cookies = [];
  if (existsSync(COOKIES_FILE)) {
    cookies = JSON.parse(readFileSync(COOKIES_FILE, "utf-8"));
  }

  async function refreshCookies() {
    console.log(`${new Date()} - refreshing cookies`);

    const ssoResponse = await azureAdSingleSignOn({
      username: USERNAME,
      password: PASSWORD,
      loginUrl: XEDULE_URL,
      postLoginSelector: "#page-title",
      getAllBrowserCookies: true,
      headless: true,
      cookies: cookies,
    });
    cookies = ssoResponse.cookies;
    writeFileSync(COOKIES_FILE, JSON.stringify(cookies), "utf-8");
  }

  await refreshCookies();
  setInterval(refreshCookies, 1000 * 60 * REFRESH_INTERVAL_M);

  const parsedUrl = parseUrl(XEDULE_URL);
  fastify.get("/", (req) => {
    if (req.query.secret !== SECRET) {
      return { error: "invalid secret" };
    }

    const filteredCookies = cookies
      .filter((e) => e.domain.includes(parsedUrl.hostname))
      .map((e) => ({
        name: e.name,
        value: e.value,
        domain: e.domain,
      }));
    const config = { xeduleUrl: XEDULE_URL };
    return { cookies: filteredCookies, config };
  });

  await fastify.listen({ port: PORT });
})();
