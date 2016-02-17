package mathutil

import (
    "testing"
    "math"
)

func Test_toRadians(t *testing.T) {
    rad := ToRadians(90)
    t.Log("rad: ", rad)
    if rad == math.Pi / 2 {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}

func Test_SecondsToRadians(t *testing.T) {
    rad := SecondsToRadians( 648000 )
    t.Log("rad: ", rad)
    if rad == math.Pi {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}


func Test_Mod2Pi_1(t *testing.T) {
    rad := Mod2Pi(3 * math.Pi)
    t.Log("rad:", rad)
    if rad == math.Pi {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}

func Test_Mod2Pi_2(t *testing.T) {
    rad := Mod2Pi( - math.Pi)
    t.Log("rad:", rad)
    if rad == math.Pi {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}


func Test_ModPi_1(t *testing.T) {
    rad := ModPi( 2 * math.Pi )
    if rad == 0 {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}

func Test_ModPi_2(t *testing.T) {
    rad := ModPi( -2 * math.Pi )
    if rad == 0 {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}

func Test_SecondsToDegrees(t *testing.T) {
    deg := SecondsToDegrees( 648000 )
    if deg == 180 {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}

func Test_DmsToDegrees(t *testing.T) {
    deg := DmsToDegrees(40, 11, 15)
    t.Log("deg: ", deg)
    if deg == 40.1875 {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}

func Test_DmsToSeconds(t *testing.T) {
    secs := DmsToSeconds(40, 11, 15)
    t.Log("secs: ", secs)
    if secs == 144675 {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}

func Test_DmsToRadians(t *testing.T) {
    rad := DmsToRadians(40, 11, 15)
    t.Log("rad:", rad)
    if rad == 0.7014041931452212 {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}

func Test_NewtonIteration(t *testing.T) {
    var n float64 = 2
    // 求根号2
    x := NewtonIteration(
        func(x float64) float64 {
            return x * x - n
        },
        1.4 )
    t.Log("x: " , x)
    t.Log("sqrt(2) = ", math.Sqrt(2) )
    if math.Abs( math.Sqrt(2) - x ) < 1e-7 {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}
