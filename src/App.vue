<script setup lang="ts">
import { onLaunch } from "@dcloudio/uni-app";
import { getPlatformAdapter } from "./platforms";

onLaunch(() => {
  getPlatformAdapter().bootstrap();

  // H5 is also our desktop review surface. uni-app resolves rpx against the
  // browser viewport, so a wide browser needs a proportional phone-canvas
  // scale instead of only constraining the outer container width.
  if (typeof window !== "undefined") {
    const syncPreviewScale = () => {
      // uni-app caps H5 rpx calculation at a 960px design viewport.
      const scale = window.innerWidth >= 600 ? 430 / 960 : 1;
      document.documentElement.style.setProperty("--qudaju-preview-scale", String(scale));
      document.documentElement.style.setProperty(
        "--qudaju-preview-layout-height",
        `${window.innerHeight / scale}px`,
      );
    };

    syncPreviewScale();
    window.addEventListener("resize", syncPreviewScale, { passive: true });
  }
});
</script>

<style lang="scss">
page {
  min-height: 100%;
  background: #f3f5f0;
  color: #153f35;
  font-family: -apple-system, BlinkMacSystemFont, "PingFang SC", "Microsoft YaHei", sans-serif;
}

view,
text,
button,
input,
scroll-view {
  box-sizing: border-box;
}

button::after {
  border: 0;
}

button,
input,
textarea { font-family: inherit; }

button[disabled] { opacity: 0.48; }

::selection { color: #fffdf6; background: #ef644a; }

@media (min-width: 600px) {
  /* uni H5 compiles rpx to rem against a 375px canvas. Match the 960px
     layout before the existing 430/960 preview transform, not twice-shrink. */
  html { font-size: 40.96px !important; }
  html,
  body {
    overflow-x: hidden;
    background: #dce3dd;
  }

  uni-app {
    position: absolute;
    top: 0;
    left: 50%;
    width: 960px;
    min-height: var(--qudaju-preview-layout-height, 100vh);
    transform: scale(var(--qudaju-preview-scale, 1)) translateX(-50%);
    transform-origin: top left;
    background: #f3f5f0;
    box-shadow: 0 0 52px rgba(21, 63, 53, 0.16);
  }
}
</style>
