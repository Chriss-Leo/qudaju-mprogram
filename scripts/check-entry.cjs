const { chromium } = require(process.env.QUDADJU_PLAYWRIGHT_PATH || "playwright");
const assert = require("node:assert/strict");
const base = process.env.QUDADJU_BASE_URL || "http://127.0.0.1:5173/";
(async () => {
  const browser = await chromium.launch({ headless: true });
  try {
    for (const [name, width, height] of [["iphone", 393, 852], ["small", 320, 640], ["desktop", 1200, 900]]) {
      const page = await browser.newPage({ viewport: { width, height } });
      const errors = [];
      page.on("pageerror", error => errors.push(error.message));
      await page.goto(base, { waitUntil: "networkidle" });
      await page.waitForSelector(".entry-cover__image--ready");
      await page.evaluate(() => document.fonts.ready);
      assert(await page.evaluate(() => document.fonts.check('20px "QudajuRoundedBold"')), "Headline font missing");
      await page.mouse.wheel(0, 800);
      assert(await page.evaluate(() => scrollY < 2), "Home must not scroll");
      assert.equal(await page.locator(".discovery-location:visible").count(), 0);
      assert(!await page.evaluate(() => document.documentElement.scrollWidth > innerWidth), name + " overflow");
      await page.locator(".entry-discover__link").click();
      await page.waitForURL(/pages\/groups\/index/);
      const header = await page.locator(".groups-header:visible").boundingBox();
      await page.locator(".groups-scroll:visible").hover();
      await page.mouse.wheel(0, 650);
      await page.waitForTimeout(200);
      const after = await page.locator(".groups-header:visible").boundingBox();
      assert(Math.abs(header.y - after.y) < 1, "Header moved");
      assert(await page.locator(".groups-scroll .uni-scroll-view").evaluateAll(els => els.some(el => el.scrollTop > 0)), "Feed not scrolling");
      await page.locator(".rail-search:visible").click();
      const dialog = await page.locator(".search-dialog:visible").boundingBox();
      assert(Math.abs(dialog.y + dialog.height / 2 - height / 2) < 2, "Search not centered");
      await page.locator(".search-choice").nth(1).getByText("免费", { exact: true }).click();
      await page.locator(".search-dialog__actions").getByText("搜索好局", { exact: true }).click();
      assert.equal(await page.locator(".group-card:visible").count(), 3, "Free filter");
      await page.locator(".rail-search:visible").click();
      await page.locator(".search-choice").nth(0).getByText("交友", { exact: true }).click();
      await page.locator(".search-dialog__header uni-button").click();
      assert.equal(await page.locator(".group-card:visible").count(), 3, "Cancel changed filters");
      await page.locator(".rail-search:visible").click();
      await page.locator(".search-dialog__actions").getByText("重置", { exact: true }).click();
      await page.locator(".search-dialog__actions").getByText("搜索好局", { exact: true }).click();
      assert.equal(await page.locator(".group-card:visible").count(), 5, "Reset filter");
      assert(!await page.evaluate(() => document.documentElement.scrollWidth > innerWidth), name + " feed overflow");
      assert.deepEqual(errors, [], name + " browser errors");
      console.log(name + ": home, navigation, fixed header, scrolling, search, cancel/reset passed");
      await page.close();
    }
  } finally { await browser.close(); }
})().catch(error => { console.error(error); process.exit(1); });
