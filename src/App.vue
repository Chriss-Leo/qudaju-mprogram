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
  background: $color-compatF3F5F0;
  color: $color-text;
  font-family: $font-body;
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

/* Required pairing for every orange primary button. */
.action-accent { background: $color-accent; color: $color-accentForeground; }

button[disabled] { opacity: 0.48; }

::selection { color: $color-inverseWarm; background: $color-accent; }

@media (min-width: 600px) {
  /* uni H5 compiles rpx to rem against a 375px canvas. Match the 960px
     layout before the existing 430/960 preview transform, not twice-shrink. */
  html { font-size: $font-size-previewRoot !important; }
  html,
  body {
    overflow-x: hidden;
    background: $color-compatDCE3DD;
  }

  uni-app {
    position: absolute;
    top: 0;
    left: 50%;
    width: 960px;
    min-height: var(--qudaju-preview-layout-height, 100vh);
    transform: scale(var(--qudaju-preview-scale, 1)) translateX(-50%);
    transform-origin: top left;
    background: $color-compatF3F5F0;
    box-shadow: 0 0 52px $color-compatOverlay1;
  }
}
</style>
