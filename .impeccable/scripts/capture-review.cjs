const { chromium } = require("playwright");

const routes = [
  ["home", "pages/index/index"],
  ["search", "pages/search/index"],
  ["detail", "pages/service-detail/index?id=city-guide-01"],
  ["order-submit", "pages/order-submit/index?id=city-guide-01"],
  ["order-detail", "pages/order-detail/index"],
  ["profile", "pages/profile/index"],
  ["provider", "pages/provider/index"],
];
const only = process.argv[2];
const port = process.argv[3] || "8090";
const captureFullPage = process.argv[4] === "both";

(async () => {
  const browser = await chromium.launch({
    executablePath: "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
    headless: true,
  });
  const page = await browser.newPage({ viewport: { width: 430, height: 720 }, deviceScaleFactor: 1 });

  for (const [name, route] of routes.filter(([name]) => !only || name === only)) {
    await page.goto(`http://127.0.0.1:${port}/#/${route}`, { waitUntil: "networkidle" });
    await page.waitForTimeout(500);
    await page.screenshot({ path: `.impeccable/review/${name}.png`, fullPage: false });
    if (captureFullPage) {
      await page.screenshot({ path: `.impeccable/review/${name}-full.png`, fullPage: true });
    }
  }

  await browser.close();
})();
