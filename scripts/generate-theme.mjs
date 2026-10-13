// Generated outputs are derived from tokens.json; never edit them directly.
import fs from "node:fs";
import path from "node:path";
import { fileURLToPath } from "node:url";
const root = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "..");
export function generateTheme() {
  const tokens = JSON.parse(fs.readFileSync(path.join(root, "src/theme/tokens.json"), "utf8"));
  const groups = { colors: "color", fontFamilies: "font", fontSizes: "font-size", fontWeights: "font-weight", lineHeights: "line-height", radii: "radius", spacing: "spacing", iconSizes: "icon-size", opacity: "opacity" };
  const lines = ["// GENERATED from src/theme/tokens.json. Do not edit."];
  for (const [group, prefix] of Object.entries(groups)) {
    for (const [key, value] of Object.entries(tokens[group])) lines.push("$" + prefix + "-" + key + ": " + value + ";");
  }
  const out = path.join(root, "src/theme/generated.scss");
  const body = lines.join("\n") + "\n";
  if (!fs.existsSync(out) || fs.readFileSync(out, "utf8") !== body) fs.writeFileSync(out, body);
  // Navigation glyph colors use the same source as CSS and Vue color props.
  for (const [file, key] of Object.entries({ home: "primary", "home-muted": "iconMuted", profile: "primary", "profile-muted": "iconMuted", plus: "accentForeground" })) {
    const target = path.join(root, "src/static/entry/icons", file + ".svg");
    const old = fs.readFileSync(target, "utf8");
    const next = old.replace(/fill="#[0-9a-fA-F]{3,8}"/g, 'fill="' + tokens.colors[key] + '"');
    if (old !== next) fs.writeFileSync(target, next);
  }
  // One Phosphor family; generated glyph colors come only from theme tokens.
  const { icons, width, height } = JSON.parse(fs.readFileSync(path.join(root, "node_modules/@iconify-json/ph/icons.json"), "utf8"));
  const glyphs = { message: "chat-circle-dots", calendar: "calendar-check", launch: "paper-plane-tilt", favorite: "star", wallet: "wallet", support: "headset", settings: "gear-six", chevron: "caret-right", back: "caret-left", clock: "clock", location: "map-pin", people: "users", info: "info", edit: "pencil-simple", share: "share-network", heart: "heart", close: "x" };
  const dir = path.join(root, "src/static/ui-icons");
  fs.mkdirSync(dir, { recursive: true });
  // Exact vector seam, kept separate from text and the generated ticket illustration.
  fs.writeFileSync(path.join(dir, "identity-seam.svg"), `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 700 48" preserveAspectRatio="none"><path fill="${tokens.colors.white}" d="M0 48V34Q0 2 34 3L329 18L350 39L371 18L666 3Q700 2 700 34V48Z"/></svg>`);
  for (const [name, glyph] of Object.entries(glyphs)) {
    const icon = icons[glyph];
    if (!icon) throw new Error("Missing Phosphor icon: " + glyph);
    for (const [tone, key] of Object.entries({ blue: "primary", muted: "iconMuted", black: "text" })) {
      const color = tokens.colors[key];
      const svg = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 ${icon.width || width} ${icon.height || height}" fill="${color}">${icon.body.replaceAll("currentColor", color)}</svg>`;
      fs.writeFileSync(path.join(dir, `${name}-${tone}.svg`), svg);
    }
  }
}
if (process.argv[1] && path.resolve(process.argv[1]) === fileURLToPath(import.meta.url)) generateTheme();
