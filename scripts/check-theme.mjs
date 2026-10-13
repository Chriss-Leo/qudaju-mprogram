// Reject scattered visual constants in application SFCs; generated SVG/Sass excluded.
import fs from "node:fs";
import path from "node:path";
const walk = dir => fs.readdirSync(dir, { withFileTypes: true }).flatMap(entry => entry.isDirectory() ? walk(path.join(dir, entry.name)) : [path.join(dir, entry.name)]);
const errors = [];
const { colors } = JSON.parse(fs.readFileSync("src/theme/tokens.json", "utf8"));
if (colors.accentForeground !== "#FFFFFF") errors.push("accentForeground must stay white for orange buttons.");
for (const file of walk("src").filter(file => file.endsWith(".vue"))) {
  const source = fs.readFileSync(file, "utf8");
  for (const match of source.matchAll(/#[\da-fA-F]{3,8}\b|rgba\([^)]*\)/g)) {
    const prefix = source.slice(Math.max(0, match.index - 25), match.index);
    // Platform contract accepts only black/white frontColor, not arbitrary theme colors.
    if (/frontColor:\s*["']$/.test(prefix) && ["#ffffff", "#000000"].includes(match[0])) continue;
    errors.push(file + ": hardcoded color " + match[0]);
  }
  for (const style of source.matchAll(/<style[^>]*>([\s\S]*?)<\/style>/g)) {
    for (const rule of style[1].matchAll(/([^{}]+)\{([^{}]*)\}/g)) {
      const [selector, declarations] = [rule[1], rule[2]];
      if (!/(?:button|action-accent)/.test(selector) || !/background(?:-color)?:\s*\$color-accent\s*[;}]/.test(declarations + ";")) continue;
      const foreground = declarations.match(/(?:^|;)\s*color:\s*([^;]+)/)?.[1]?.trim();
      if (foreground && !["$color-white", "$color-accentForeground"].includes(foreground)) errors.push(file + ": orange button text must use accentForeground (white)");
    }
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
