package solarterms

type SolarTerm struct {
    Order int
    Name string
    Month int
    EstimateDate int
}

var (
    XiaoHan = &SolarTerm{20, "小寒", 1, 5}
    DaHan = &SolarTerm{21, "大寒", 1, 22}
    LiChun = &SolarTerm{22, "立春", 2, 5}
    YuShui = &SolarTerm{23, "雨水", 2, 22}
    JingZhe = &SolarTerm{24, "惊蛰", 3, 5}
    ChunFen = &SolarTerm{1, "春风", 3, 22}
    QingMing = &SolarTerm{2, "清明", 4, 5}
    GuYu = &SolarTerm{3, "谷雨", 4, 22}
    LiXia = &SolarTerm{4, "立夏", 5, 5}
    XiaoMan = &SolarTerm{5, "小满", 5, 22}
    MangZhong = &SolarTerm{6, "芒种", 6, 5}
    XiaZhi = &SolarTerm{7, "夏至", 6, 22}
    XiaoShu = &SolarTerm{8, "小暑", 7, 5}
    DaShu = &SolarTerm{9, "大暑", 7, 22}
    LiQiu =  &SolarTerm{10, "立秋", 8, 5}
    ChuShu = &SolarTerm{11, "处暑", 8, 22}
    BaiLu = &SolarTerm{12, "白露", 9, 5}
    QiuFen = &SolarTerm{13, "秋分", 9, 22}
    HanLu = &SolarTerm{14, "寒露", 10, 5}
    ShuangJiang = &SolarTerm{15, "霜降", 10, 22}
    LiDong = &SolarTerm{16, "立冬", 11, 5}
    XiaoXue = &SolarTerm{17, "小雪", 11, 22 }
    DaXue = &SolarTerm{18, "大雪", 12, 5}
    DongZhi = &SolarTerm{19, "冬至", 12, 22}
)
