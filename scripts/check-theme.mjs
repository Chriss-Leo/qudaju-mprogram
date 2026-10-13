// Reject scattered visual constants in application SFCs; generated SVG/Sass excluded.
import fs from "node:fs";
import path from "node:path";
const walk = dir => fs.readdirSync(dir, { withFileTypes: true }).flatMap(entry => entry.isDirectory() ? walk(path.join(dir, entry.name)) : [path.join(dir, entry.name)]);
const errors = [];
for (const file of walk("src").filter(file => file.endsWith(".vue"))) {
  const source = fs.readFileSync(file, "utf8");
  for (const match of source.matchAll(/#[\da-fA-F]{3,8}\b|rgba\([^)]*\)/g)) {
    const prefix = source.slice(Math.max(0, match.index - 25), match.index);
    // Platform contract accepts only black/white frontColor, not arbitrary theme colors.
    if (/frontColor:\s*["']$/.test(prefix) && ["#ffffff", "#000000"].includes(match[0])) continue;
    errors.push(file + ": hardcoded color " + match[0]);
  }
  for (const style of source.matchAll(/<style[^>]*>([\s\S]*?)<\/style>/g)) {
    for (const match of style[1].matchAll(/(font-size|font-weight|line-height|border-radius):\s*([^;}]+)/g)) {
      const value = match[2].trim();
      if (!value.startsWith("$") && !["inherit", "normal", "auto"].includes(value)) errors.push(file + ": hardcoded " + match[1] + " " + value);
    }
    for (const match of style[1].matchAll(/font-family:\s*([^;}]+)/g)) {
      if (!match[1].trim().startsWith("$") && match[1].trim() !== "inherit") errors.push(file + ": hardcoded font family");
    }
  }
}
if (errors.length) { console.error(errors.join("\n")); process.exit(1); }
console.log("Theme check passed: no scattered colors, typography or radii in application SFCs.");
