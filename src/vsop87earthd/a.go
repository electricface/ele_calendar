package vsop87earthd

import (
    "calendarutil"
    "mathutil"
    "math"
)
/**
 * 按儒略日计算地球的日心黄经
 *
 * @param jd
 *            儒略日
 * @return 地球的日心黄经，单位是弧度(rad)
 */

func GetSunEclipticLongitudeForEarth(jd float64) float64 {
    t := calendarutil.GetJulianThousandYears(jd)
    L0 := GetEarthL0(t)
    L1 := GetEarthL1(t)
    L2 := GetEarthL2(t)
    L3 := GetEarthL3(t)
    L4 := GetEarthL4(t)
    L5 := GetEarthL5(t)
    L := ((((L5 * t + L4) * t + L3) * t + L2) * t + L1) * t + L0
    return mathutil.Mod2Pi( L )
}


/**
 * 按儒略日计算地球的日心黄纬
 *
 * @param jd
 *            儒略日
 * @return 地球的日心黄纬，单位是弧度(rad)
 */

func GetSunEclipticLatitudeForEarth(jd float64) float64 {
    t := calendarutil.GetJulianThousandYears(jd)
    B0 := GetEarthB0(t)
    B1 := GetEarthB1(t)
    B2 := GetEarthB2(t)
    B3 := GetEarthB3(t)
    B4 := GetEarthB4(t)
    B := ((((B4 * t) + B3) * t + B2) * t + B1) * t + B0
    return B
}

/**
 * 按照儒略日计算地球和太阳的距离
 *
 * @param jd
 *            儒略日
 * @return 地球和太阳的距离，单位是天文单位(au)
 */
func GetSunRadiusForEarth(jd float64) float64 {
    t := calendarutil.GetJulianThousandYears(jd)
    R0 := GetEarthR0(t)
    R1 := GetEarthR1(t)
    R2 := GetEarthR2(t)
    R3 := GetEarthR3(t)
    R4 := GetEarthR4(t)
    R5 := GetEarthR5(t)
    R := ((((R5 * t + R4) * t + R3) * t + R2) * t + R1) * t + R0
    return R
}


/**
     * 用于把vsop87理论算出来的经度转换成fk5目视系统的经度的修正值，参考<i>Jean Meeus</i>的<i>Astronomical
     * Algorithms</i>第二版(1998)第32章219页(32.3)式
     *
     * @param l
     *            vsop87经度(rad)
     * @param b
     *            vsop87纬度(rad)
     * @param t
     *            儒略日
     * @return 修正量(rad)
     */
func Vsop2Fk5LongitudeCorrection(l float64, b float64, jd float64) float64 {
    t := calendarutil.GetJulianCentury(jd)
    lp := l - mathutil.ToRadians(1.397) * t - mathutil.ToRadians(0.00031) * t * t
    return mathutil.SecondsToRadians(-0.09033 + 0.03916 * (math.Cos(lp) + math.Sin(lp)) * math.Tan(b))
}

/**
 * 用于把vsop87理论算出来的纬度转换成fk5目视系统的纬度的修正值，参考<i>Jean Meeus</i>的<i>Astronomical
 * Algorithms</i>第二版(1998)第32章219页(32.3)式
 *
 * @param l
 *            vsop87经度(rad)
 * @param b
 *            vsop87纬度(rad)
 * @param t
 *            儒略日
 * @return 修正量(rad)
 */
 // 也许没什么用
func Vsop2Fk5LatitudeCorrection(l float64, b float64, jd float64) float64 {
    t := calendarutil.GetJulianCentury(jd)
    lp :=  l - mathutil.ToRadians(1.397) * t - mathutil.ToRadians(0.00031) * t * t;
    return mathutil.SecondsToRadians(0.03916) * (math.Cos(lp) - math.Sin(lp));
}

/**
 * 计算修正后的太阳的地心视黄经
 *
 * @param jd
 *            儒略日
 * @return 修正后的地心黄经(rad)
 */
func GetEarthEclipticLongitudeForSun(jd float64 ) float64 {
    // 计算地球的日心黄经
    l := GetSunEclipticLongitudeForEarth(jd)

    // 计算地球的日心黄纬
    b := GetSunEclipticLatitudeForEarth(jd)


    // 修正章动
    // l += GetLongitudeNutation(jd)

    // 转换到fk5
    l += Vsop2Fk5LongitudeCorrection(l, b, jd);

    // 转换成太阳的地心黄经
    l = mathutil.Mod2Pi(l + math.Pi);

    // 计算光行差
    // 计算日地距离
    r := GetSunRadiusForEarth(jd)
    // 太阳到地球的光行差参数
    LIGHT_ABERRATION := mathutil.SecondsToRadians(20.4898)
    l -= LIGHT_ABERRATION / r;

    return l;
}
