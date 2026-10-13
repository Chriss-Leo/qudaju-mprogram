import { defineConfig } from "vite";
import uni from "@dcloudio/vite-plugin-uni";
// @ts-ignore Node-side build helper; no client runtime dependency.
import { generateTheme } from "./scripts/generate-theme.mjs";

generateTheme();

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    {
      name: "qudaju-theme-tokens",
      configureServer(server) {
        server.watcher.add("src/theme/tokens.json");
        const update = (file: string) => { if (file.replace(/\\/g, "/").endsWith("/src/theme/tokens.json")) generateTheme(); };
        server.watcher.on("change", update);
        server.httpServer?.once("close", () => server.watcher.off("change", update));
      },
    },
    uni(),
  ],
});
