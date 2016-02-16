package solarterms

import (
    "testing"
)

func Test_1(t *testing.T) {
    t.Log(XiaoHan)
    t.Log(DaHan)
    t.Log(LiChun)
    t.Log(YuShui)
    if YuShui.Name == "雨水" &&
        YuShui.Order == 23 &&
        YuShui.Month == 2 &&
        YuShui.EstimateDate == 22 {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}
