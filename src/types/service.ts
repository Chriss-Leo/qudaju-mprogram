export interface ServiceItem {
  id: string;
  title: string;
  provider: string;
  distance: string;
  rating: string;
  orders: number;
  price: number;
  tags: string[];
  visual: "guide" | "photo" | "game" | "outdoor";
  mode: "线下见面" | "线上远程";
  available: Array<"今天" | "明天" | "本周末">;
}
