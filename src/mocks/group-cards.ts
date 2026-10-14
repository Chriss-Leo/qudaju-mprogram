// Illustrative groups and formally enrolled participants, not real registration records.
export interface GroupParticipant { id: string; nickname: string; avatar: string }
export interface GroupCardData {
  title: string;
  host: string;
  hostAvatar: string;
  category: string;
  admission: "局主审核" | "直接加入";
  description: string;
  time: string;
  area: string;
  fund: number;
  joined: number;
  capacity: number;
  photos: string[];
  participants: GroupParticipant[];
  isJoinedByMe: boolean;
}

const participants: GroupParticipant[] = [
  { id: "sample-1", nickname: "小禾", avatar: "/static/group-detail/participant-1.jpg" },
  { id: "sample-2", nickname: "阿远", avatar: "/static/group-detail/participant-2.jpg" },
  { id: "sample-3", nickname: "小林", avatar: "/static/group-detail/participant-3.jpg" },
  { id: "sample-4", nickname: "安安", avatar: "/static/group-detail/participant-4.jpg" },
  { id: "sample-5", nickname: "小夏", avatar: "/static/profile-a/profile-avatar.jpg" },
];

// The sample joined count includes the host; moving them first must not add a seat.
function enrolledWithHost(host: string, avatar: string, count: number): GroupParticipant[] {
  return [{ id: `host-${host}`, nickname: host, avatar }, ...participants.filter(person => person.avatar !== avatar)].slice(0, count);
}

export const groupCards: Record<string, GroupCardData> = {
  "game-01": {
    title: "周六桌游局，新手也能来", host: "阿橙", category: "桌游", admission: "局主审核",
    description: "带你从轻策略开始，一起认识新朋友。", time: "周六 14:00", area: "静安寺",
    hostAvatar: "/static/profile-a/host-avatar.jpg",
    fund: 0, joined: 4, capacity: 6,
    photos: ["/static/profile-a/boardgame-hero.jpg", "/static/group-detail/game-2.jpg", "/static/group-detail/game-3.jpg"],
    participants: enrolledWithHost("阿橙", "/static/profile-a/host-avatar.jpg", 4),
    isJoinedByMe: false,
  },
  "photo-01": {
    title: "梧桐街区，边走边互拍", host: "夏夏", category: "陪拍", admission: "局主审核",
    description: "手机或相机都可以，彼此做模特，记录城市里的好时光。", time: "周日 15:00", area: "衡山路",
    hostAvatar: "/static/group-detail/participant-4.jpg",
    fund: 35, joined: 3, capacity: 5,
    photos: ["/static/group-detail/photo-1.jpg", "/static/group-detail/photo-2.jpg"],
    participants: enrolledWithHost("夏夏", "/static/group-detail/participant-4.jpg", 3),
    isJoinedByMe: false,
  },
  "outdoor-01": {
    title: "苏州河轻徒步，走走聊聊", host: "木木", category: "户外", admission: "直接加入",
    description: "走一段路，认识几个新朋友。慢一点也没关系。", time: "周日 09:30", area: "苏州河",
    hostAvatar: "/static/group-detail/participant-3.jpg",
    fund: 0, joined: 5, capacity: 8, photos: ["/static/group-detail/outdoor-1.jpg"],
    participants: enrolledWithHost("木木", "/static/group-detail/participant-3.jpg", 5),
    isJoinedByMe: true,
  },
  "city-guide-01": {
    title: "梧桐街区漫步，发现街角故事", host: "林野", category: "向导", admission: "局主审核",
    description: "一起走进熟悉又陌生的街区，聊聊建筑与城市日常。", time: "周六 15:00", area: "武康路",
    hostAvatar: "/static/group-detail/participant-2.jpg",
    fund: 25, joined: 2, capacity: 6,
    photos: ["/static/group-detail/photo-1.jpg", "/static/group-detail/photo-2.jpg", "/static/group-detail/photo-3.jpg"],
    participants: enrolledWithHost("林野", "/static/group-detail/participant-2.jpg", 2),
    isJoinedByMe: false,
  },
  "guide-02": {
    title: "夜游苏河湾，聊聊建筑故事", host: "周屿", category: "向导", admission: "局主审核",
    description: "沿着河岸慢慢走，在夜色里发现城市的另一面。", time: "周六 19:00", area: "苏河湾",
    hostAvatar: "/static/group-detail/participant-2.jpg",
    fund: 0, joined: 3, capacity: 6, photos: ["/static/group-detail/outdoor-2.jpg"],
    participants: enrolledWithHost("周屿", "/static/group-detail/participant-2.jpg", 3),
    isJoinedByMe: false,
  },
};
