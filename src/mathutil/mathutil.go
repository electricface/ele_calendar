// MathUtil.java

package mathutil

import (
    "math"
)

// math.Pi
// math.Abs
// ToRadians ???

/*
    角度转换为弧度 rad
    @param degrees
            角度
    @return 弧度
*/
func ToRadians(degrees float64) float64 {
    return degrees * math.Pi / 180
}


/**
 * 把角秒换算成弧度
 *
 * @param seconds
 *            角秒
 * @return 对应的弧度值
 */

 func SecondsToRadians(seconds float64 ) float64 {
     return ToRadians( SecondsToDegrees( seconds )  )
 }


 /**
  * 把角度限制在[0, 2π]之间
  *
  * @param r
  *            原角度(rad)
  * @return 转换后的角度(rad)
  */

func Mod2Pi(r float64 ) float64 {
    for r < 0 {
        r += math.Pi * 2
    }
    for r > 2 * math.Pi {
        r -= math.Pi * 2
    }
    return r
}

/**
     * 把角度限制在[-π, π]之间
     *
     * @param r
     *            原角度(rad)
     * @return 转换后的角度(rad)
     */

func ModPi(r float64 ) float64 {
    for r <  -math.Pi {
        r += math.Pi * 2
    }
    for r > math.Pi {
        r -= math.Pi * 2
    }
    return r
}

/**
     * 把角秒换算成角度
     *
     * @param seconds
     *            角秒
     * @return 对应的弧度值
     */

func SecondsToDegrees(seconds float64) float64 {
    return seconds / 3600
}

/**
     * 把度分秒表示的角度换算成度(deg)
     *
     * @param d
     *            度
     * @param m
     *            分
     * @param s
     *            秒
     * @return 换算成度的值
     */

func DmsToDegrees(d int, m int , s float64) float64 {
    return float64(d) + float64(m) / 60 + s / 3600
}

/**
 * 把度分秒表示的角度换算成角秒(arcsecond)
 *
 * @param d
 *            度
 * @param m
 *            分
 * @param s
 *            秒
 * @return 换算成秒的值
 */

func DmsToSeconds(d int, m int, s float64) float64 {
    return float64(d) * 3600 + float64(m) * 60 + s
}

/**
     * 把度分秒表示的角度换算成弧度(rad)
     *
     * @param d
     *            度
     * @param m
     *            分
     * @param s
     *            秒
     * @return 换算成弧度的值
     */
func DmsToRadians(d int, m int, s float64 ) float64 {
    return ToRadians( DmsToDegrees(d, m, s) )
}

/**
     * 牛顿迭代求解方程的根
     *
     * @param f
     *            方程表达式
     * @param x0
     *            对根的估值
     * @return 在x0附近的一个根
     */

func NewtonIteration(f func (float64) float64 , x0 float64) float64 {
    const Epsilon = 1e-7
    const Delta = 5e-6
    var x float64

    for {
        x = x0
        fx := f(x)
        // 导数
        fpx := ( f(x + Delta) - f( x - Delta ) ) / Delta / 2
        x0 = x - fx / fpx
        if  math.Abs(x0 - x) <= Epsilon {
            break
        }
    }
    return x

}
