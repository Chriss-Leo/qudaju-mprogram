// Illustrative group-feed data only; legacy service detail routes remain unchanged.
export interface GroupCardData {
  title: string;
  host: string;
  category: string;
  admission: "局主审核" | "直接加入";
  description: string;
  time: string;
  area: string;
  fund: number;
  joined: number;
  capacity: number;
  gallery: string;
}

export const groupCards: Record<string, GroupCardData> = {
  "game-01": {
    title: "周六桌游局，新手也能来", host: "阿橙", category: "桌游", admission: "局主审核",
    description: "带你从轻策略开始，一起认识新朋友。", time: "周六 14:00", area: "静安寺",
    fund: 0, joined: 4, capacity: 6, gallery: "/static/groups/game-strip.jpg",
  },
  "photo-01": {
    title: "梧桐街区，边走边互拍", host: "夏夏", category: "陪拍", admission: "局主审核",
    description: "手机或相机都可以，彼此做模特，记录城市里的好时光。", time: "周日 15:00", area: "衡山路",
    fund: 35, joined: 3, capacity: 5, gallery: "/static/groups/photo-strip.jpg",
  },
  "outdoor-01": {
    title: "苏州河轻徒步，走走聊聊", host: "木木", category: "户外", admission: "直接加入",
    description: "走一段路，认识几个新朋友。慢一点也没关系。", time: "周日 09:30", area: "苏州河",
    fund: 0, joined: 5, capacity: 8, gallery: "/static/groups/outdoor-strip.jpg",
  },
  "city-guide-01": {
    title: "梧桐街区漫步，发现街角故事", host: "林野", category: "向导", admission: "局主审核",
    description: "一起走进熟悉又陌生的街区，聊聊建筑与城市日常。", time: "周六 15:00", area: "武康路",
    fund: 25, joined: 2, capacity: 6, gallery: "/static/groups/photo-strip.jpg",
  },
  "guide-02": {
    title: "夜游苏河湾，聊聊建筑故事", host: "周屿", category: "向导", admission: "局主审核",
    description: "沿着河岸慢慢走，在夜色里发现城市的另一面。", time: "周六 19:00", area: "苏河湾",
    fund: 0, joined: 3, capacity: 6, gallery: "/static/groups/outdoor-strip.jpg",
  },
};
