// Export only the Phosphor glyphs used by the fixed navigation.
import fs from "node:fs";
import path from "node:path";
const { icons, width, height } = JSON.parse(fs.readFileSync("node_modules/@iconify-json/ph/icons.json", "utf8"));
const { colors } = JSON.parse(fs.readFileSync("src/theme/tokens.json", "utf8"));
const targets = {
  home: ["house-fill", colors.primary],
  "home-muted": ["house-fill", colors.iconMuted],
  profile: ["user", colors.primary],
  "profile-muted": ["user", colors.iconMuted],
  plus: ["plus", colors.accentForeground],
};
for (const [file, [name, color]] of Object.entries(targets)) {
  const icon = icons[name];
  if (!icon) throw new Error("Missing Phosphor glyph: " + name);
  const svg = '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 ' + (icon.width || width) + ' ' + (icon.height || height) + '" fill="' + color + '">' + icon.body.replaceAll("currentColor", color) + "</svg>";
  fs.writeFileSync(path.join("src/static/entry/icons", file + ".svg"), svg);
}
