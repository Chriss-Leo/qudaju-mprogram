const { chromium } = require(process.env.QUDADJU_PLAYWRIGHT_PATH || "playwright");
const assert = require("node:assert/strict");
(async () => {
  const browser = await chromium.launch({ headless: true });
  for (const [name, width, height] of [["iphone", 393, 852], ["small", 320, 640], ["desktop", 1200, 900]]) {
    const page = await browser.newPage({ viewport: { width, height } });
    const errors = [];
    page.on("pageerror", error => errors.push(error.message));
    await page.goto("http://127.0.0.1:5173/", { waitUntil: "networkidle" });
    await page.waitForSelector(".entry-cover__image--ready");
    await page.evaluate(() => document.fonts.ready);
    assert(await page.evaluate(() => document.fonts.check('20px "QudajuPlay"')), "Display font not loaded");
    const result = await page.evaluate(() => {
      const box = selector => {
        const r = document.querySelector(selector).getBoundingClientRect();
        return { x: r.x, y: r.y, width: r.width, height: r.height, right: r.right, bottom: r.bottom };
      };
      return { hero: box(".entry-cover"), title: box(".entry-headline"), categories: box(".entry-categories"), nav: box(".tabbar"), overflow: document.documentElement.scrollWidth > innerWidth };
    });
    assert(!result.overflow, name + " horizontal overflow");
    assert(result.title.right <= width + 1, name + " headline overflow");
    assert(result.categories.bottom < result.nav.y - 10, name + " category/navigation overlap");
    await page.screenshot({ path: "qudaju-ui/entry-concepts-20261009/code-" + name + ".png" });
    if (name === "iphone") {
      await page.getByLabel("陪拍", { exact: true }).click();
      await page.waitForURL(/pages\/search\/index/);
      assert(page.url().includes("keyword"), "Category did not pass search keyword");
      await page.goto("http://127.0.0.1:5173/", { waitUntil: "networkidle" });
      await page.getByLabel("发现同城活动", { exact: true }).click();
      await page.waitForTimeout(500);
      assert(await page.evaluate(() => document.querySelector("#activity-discovery").getBoundingClientRect().top < 5), "Discovery did not scroll to feed");
      await page.getByLabel("首页", { exact: true }).click();
      await page.waitForTimeout(500);
      assert(await page.evaluate(() => scrollY < 5), "Home did not return to top");
    }
    assert.equal(errors.length, 0, name + ": " + errors.join(", "));
    console.log(name, JSON.stringify(result));
    await page.close();
  }
  await browser.close();
})().catch(error => { console.error(error); process.exit(1); });
