import { onMounted, onUnmounted, ref } from "vue";
import { onLoad, onShow } from "@dcloudio/uni-app";

// H5 can reuse a cached page for query-only navigation. Sync only its own route.
export function useFlowId(route: string) {
  const id = ref("");
  onLoad(query => { id.value = String(query?.id || ""); });
  function sync() {
    if (typeof window === "undefined") return;
    const [path, query = ""] = window.location.hash.replace(/^#\/?/, "").split("?");
    if (path === `pages/${route}/index`) id.value = new URLSearchParams(query).get("id") || "";
  }
  onShow(sync);
  onMounted(() => { if (typeof window !== "undefined") { sync(); window.addEventListener("hashchange", sync); } });
  onUnmounted(() => { if (typeof window !== "undefined") window.removeEventListener("hashchange", sync); });
  return id;
}
