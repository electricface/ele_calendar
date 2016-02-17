package calendarutil

import (
    "time"
)
/**
 * Gregorian历法实施年份
 */
const GREGORIAN_FIRST_YEAR = 1582
/**
 * Gregorian历法实施月份
 */
const GREGORIAN_FIRST_MONTH = 10
/**
 * Gregorian历法实施月份
 */
const GREGORIAN_FIRST_DATE = 15

/**
 * Julian历法废弃年份
 */
const JULIAN_LAST_YEAR = 1582
/**
 * Julian历法废弃月份
 */
const JULIAN_LAST_MONTH = 10
/**
 * Julian历法废弃日期
 */
const JULIAN_LAST_DATE = 4

/**
 * Gregorian历TT1582年10月15日中午12点的儒略日
 */
const JULIAN_GREGORIAN_BOUNDARY = 2299161;

/**
 * Gregorian历TT2000年1月1日中午12点的儒略日
 */
const J2000                     = 2451545.0;

/**
 * 1000年的日数
 */
const DAYS_OF_1000_YEARS        = 365250.0;

/**
 * 100年的日数
 */
const DAYS_OF_CENTURY           = 36525.0;

/**
 * Gregorian闰年判断
 *
 * @param y
 *            年份
 * @return 闰年返回true，平年返回false
 */
func IsGregorianLeapYear(y int) bool {
    return (y & 3) == 0 && y % 100 != 0 || y % 400 == 0;
}

/**
 * Julian闰年判断
 *
 * @param y
 *            年份
 * @return 闰年返回true，平年返回false
 */
func IsJulianLeapYear(y int) bool {
    return (y & 3) == 0
}

/**
 * {@value #JULIAN_LAST_YEAR}年(含)以前按照Julian历法，{@value #JULIAN_LAST_YEAR}年以后按照Gregorian历法。
 *
 * @param y
 *            年份
 * @return 闰年返回true，平年返回false
 */
func IsLeapYear(y int) bool {
    if y <= JULIAN_LAST_YEAR {
        return IsJulianLeapYear(y)
    }
    return IsGregorianLeapYear(y)
}

/**
 * 计算Gregorian日期的儒略日数，以TT当天中午12点为准(结果是整数)。 算法摘自<a href=
 * "http://en.wikipedia.org/wiki/Julian_day" >英文维基百科<i>Julian Day</i>词条</a>。
 *
 * @param year
 *            年份
 * @param month
 *            月份
 * @param day
 *            日期
 * @return 返回以Gregorian历法计算的儒略日数
 */

func ToJulianDateInGregorian(year, month, day int) int {
    var a int = (14 - month ) / 12
    var y int = year + 4800 - a
    var m int = month + 12 * a - 3
    return day + (153 * m + 2) / 5 + 365 * y + y / 4 - y / 100 + y / 400 - 32045;
}

/**
 * 计算Julian日期的儒略日数，以TT当天中午12点为准(结果是整数)。 算法摘自<a href= "http://en.wikipedia.org/wiki/Julian_day"
 * >英文维基百科<i>Julian Day</i>词条</a>。
 *
 * @param year
 *            年份
 * @param month
 *            月份
 * @param day
 *            日期
 * @return 返回以Julian历法计算的儒略日数
 */
func ToJulianDateInJulian(year, month, day int) int {
    var a int = (14 - month) / 12
    var y int = year + 4800 - a
    var m int = month + 12 * a - 3
    return day + (153 * m + 2) / 5 + 365 * y + y / 4 - 32083
}

/**
 * 计算儒略日数，以TT当天中午12点为准，{@value #JULIAN_LAST_YEAR}年{@value #JULIAN_LAST_MONTH}月
 * {@value #JULIAN_LAST_DATE}日及以前按照Julian历法，{@value #GREGORIAN_FIRST_YEAR}年
 * {@value #GREGORIAN_FIRST_MONTH}月{@value #GREGORIAN_FIRST_DATE}日及以后按照Gregorian历法，中间的日期按照
 * {@value #GREGORIAN_FIRST_YEAR}年{@value #GREGORIAN_FIRST_MONTH}月
 * {@value #GREGORIAN_FIRST_DATE}日计算。
 *
 * @param y
 *            年份
 * @param m
 *            月份
 * @param d
 *            日期
 * @return 返回相应历法的儒略日数
 */
func ToJulianDate(y, m, d int ) int {
    if y < JULIAN_LAST_YEAR {
        return ToJulianDateInJulian(y, m, d);
    }
    if y == JULIAN_LAST_YEAR && m < JULIAN_LAST_MONTH {
        return ToJulianDateInJulian(y, m, d);
    }
    if y == JULIAN_LAST_YEAR && m == JULIAN_LAST_MONTH && d <= JULIAN_LAST_DATE {
        return ToJulianDateInJulian(y, m, d);
    }

    if y > GREGORIAN_FIRST_YEAR {
        return ToJulianDateInGregorian(y, m, d);
    }
    if y == GREGORIAN_FIRST_YEAR && m > GREGORIAN_FIRST_MONTH {
        return ToJulianDateInGregorian(y, m, d);
    }
    if y == GREGORIAN_FIRST_YEAR && m == GREGORIAN_FIRST_MONTH && d >= GREGORIAN_FIRST_DATE {
        return ToJulianDateInGregorian(y, m, d);
    }

    // 剩下的都是中间的，以Gregorian历法实施第一天计算
    return ToJulianDateInGregorian(GREGORIAN_FIRST_YEAR, GREGORIAN_FIRST_MONTH,
            GREGORIAN_FIRST_DATE)
}

/**
 * 计算Gregorian时间的儒略日数。 算法摘自<a href= "http://en.wikipedia.org/wiki/Julian_day" >英文维基百科<i>Julian
 * Day</i>词条</a>。
 *
 * @param year
 *            年份
 * @param month
 *            月份
 * @param day
 *            日期
 * @param hour
 *            小时
 * @param minute
 *            分钟
 * @param second
 *            秒数
 * @return 返回以Gregorian历法计算的儒略日数
 */
func ToJulianDateInGregorianHMS(year, month, day, hour, minute int, second float64) float64 {
    jdn := ToJulianDateInGregorian(year, month, day)
    return float64(jdn) + ( float64(hour) - 12) / 24.0 + float64(minute) / 1440.0 + second / 86400.0
}

/**
 * 计算Julian时间的儒略日数。 算法摘自<a href= "http://en.wikipedia.org/wiki/Julian_day" >英文维基百科<i>Julian
 * Day</i>词条</a>。
 *
 * @param year
 *            年份
 * @param month
 *            月份
 * @param day
 *            日期
 * @param hour
 *            小时
 * @param minute
 *            分钟
 * @param second
 *            秒数
 * @return 返回以Julian历法计算的儒略日数
 */
func ToJulianDateInJulianHMS(year, month, day, hour, minute int, second float64) float64 {
    jdn := ToJulianDateInJulian(year, month, day)
    return float64(jdn) + ( float64(hour) - 12) / 24.0 + float64(minute) / 1440.0 + second / 86400.0
}

/**
 * 计算儒略日数，{@value #JULIAN_LAST_YEAR}年{@value #JULIAN_LAST_MONTH}月 {@value #JULIAN_LAST_DATE}
 * 日及以前按照Julian历法，{@value #GREGORIAN_FIRST_YEAR}年 {@value #GREGORIAN_FIRST_MONTH}月
 * {@value #GREGORIAN_FIRST_DATE}日及以后按照Gregorian历法，中间的日期按照 {@value #GREGORIAN_FIRST_YEAR}年
 * {@value #GREGORIAN_FIRST_MONTH}月 {@value #GREGORIAN_FIRST_DATE}日计算。
 *
 * @param year
 *            年份
 * @param month
 *            月份
 * @param day
 *            日期
 * @param hour
 *            小时
 * @param minute
 *            分钟
 * @param second
 *            秒数
 * @return 返回相应历法的儒略日数
 */
func ToJulianDateHMS(year, month, day, hour, minute int, second float64) float64 {
    jdn := ToJulianDate(year, month, day)
    return float64(jdn) + ( float64(hour) - 12) / 24.0 + float64(minute) / 1440.0 + second / 86400.0
}

func fromJulianDateHelper(a int, f float64, tz *time.Location) time.Time {
    b := a + 1524
    c := int( (float64(b) - 122.1) / 365.25 )
    d := int( 365.25 * float64(c) )
    e := int( float64(b - d) / 30.6001 )
    dd := float64( b - d - int(30.6001 * float64(e)) ) + f
    var month int
    if e <= 13 {
        month = e - 1
    } else {
        month = e - 13
    }
    var year int
    if month <= 2 {
        year = c - 4715
    } else {
        year = c - 4716
    }
    day := int(dd)
    dd = ( dd - float64(day) ) * 24
    hour := int(dd)
    dd = (dd - float64(hour)) * 60
    minute := int(dd)
    dd = (dd - float64(minute)) * 60
    second := int(dd)
    dd = (dd - float64(second)) * 1000000000
    nanosecond := int(dd)
    cal := time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, time.UTC)
    // set TimeZone
    cal = cal.In( tz )
    return cal
}

/**
 * 由儒略日计算对应的Gregorian历日期时间，算法参考<i>Jean Meeus</i>的<i>Astronomical Formulae for Calculators</i>
 *
 * @param jd
 *            儒略日
 * @param tz
 *            要使用的时区
 * @return 对应的Gregorian历TT日期时间的一个Calendar对象，时区为 tz。
 */

func fromJulianDateInGregorian(jd float64, tz *time.Location) time.Time {
    z := int(jd + 0.5)
    f := jd + 0.5 - float64(z)
    alpha := int( (float64(z) - 1867216.25) / 36524.25 )
    a := z + 1 + alpha - alpha / 4
    return fromJulianDateHelper(a, f, tz)
}

/**
 * 由儒略日计算对应的Julian历日期时间，算法参考<i>Jean Meeus</i>的<i>Astronomical Formulae for Calculators</i>
 *
 * @param jd
 *            儒略日
 * @param tz
 *            要使用的时区
 * @return 对应的Julian历TT日期时间的一个Calendar对象，时区为 tz。
 */
func fromJulianDateInJulian(jd float64, tz *time.Location) time.Time {
    z := int(jd + 0.5);
    f := jd + 0.5 - float64(z)
    a := z
    return fromJulianDateHelper(a, f, tz)
}


/**
 * 由儒略日计算对应的日期时间，算法参考<i>Jean Meeus</i>的<i>Astronomical Formulae for Calculators</i> 当儒略日小于<i>
 * {@value #JULIAN_GREGORIAN_BOUNDARY} - 0.5</i>时按Julian历法计算，以后按Gregorian历计算。
 *
 * @param jd
 *            儒略日
 * @param tt2ut
 *            是否要做TT到UTC的转换
 * @return 对应的时间的一个Calendar对象，时区为系统默认时区。
 */
func FromJulianDateLocal(jd float64, tt2utc bool) time.Time {
    return FromJulianDate(jd, time.Local, tt2utc )
}

/**
 * 由儒略日计算对应的日期时间，算法参考<i>Jean Meeus</i>的<i>Astronomical Formulae for Calculators</i> 当儒略日小于<i>
 * {@value #JULIAN_GREGORIAN_BOUNDARY} - 0.5</i>时按Julian历法计算，以后按Gregorian历计算。
 *
 * @param jd
 *            儒略日
 * @param tz
 *            使用的时区
 * @param tt2utc
 *            是否要做TT到UTC的转换
 * @return 对应的日期时间的一个Calendar对象，时区为tz。
 */

func FromJulianDate(jd float64, tz *time.Location, tt2utc bool) time.Time {
    if tt2utc {
        jd -= GetDeltaTJD(jd) / 86400
    }
    z := int(jd + 0.5)
    if z < JULIAN_GREGORIAN_BOUNDARY {
        fromJulianDateInJulian(jd, tz)
    }
    return fromJulianDateInGregorian(jd, tz)
}

/**
 * 计算儒略千年数
 *
 * @param jd
 *            要计算的儒略日
 * @return 儒略千年数
 */

func GetJulianThousandYears(jd float64) float64 {
    return (jd - J2000) / DAYS_OF_1000_YEARS
}


/**
 * 计算儒略世纪数
 *
 * @param jd
 *            要计算的儒略日
 * @return 儒略世纪数
 */
func GetJulianCentury(jd float64) float64 {
    return (jd - J2000) / DAYS_OF_CENTURY;
}

/**
 * 计算Gregorian日历的星期几。算法摘自<a href= "http://en.wikipedia.org/wiki/Zeller%27s_congruence"
 * >英文维基百科<i>Zeller's congruence</i>词条</a>。
 *
 * @param y
 *            年份
 * @param m
 *            月份
 * @param d
 *            日期
 * @return 星期几的数字表示，1-6表示星期一到星期六，0表示星期日
 */
func getWeekdayForGregorian(y, m, d int)  int {
    if m <= 2 {
        y -= 1
        m += 12
    }
    c := int(y / 100)
    y = y % 100
    w := (d + 13 * (m + 1) / 5 + y + (y / 4) + (c / 4) - 2 * c - 1) % 7;
    if w < 0 {
        w += 7
    }
    return w
}

/**
 * 计算Julian日历的星期几。算法摘自<a href= "http://en.wikipedia.org/wiki/Zeller%27s_congruence"
 * >英文维基百科<i>Zeller's congruence</i>词条</a>。
 *
 * @param y
 *            年份
 * @param m
 *            月份
 * @param d
 *            日期
 * @return 星期几的数字表示，1-6表示星期一到星期六，0表示星期日
 */
func getWeekdayForJulian(y, m, d int) int {
    if m <= 2 {
        y -= 1
        m += 12
    }
    c := int(y / 100)
    y = y % 100
    w := (d + 13 * (m + 1) / 5 + y + y / 4 + 4 - c ) % 7
    if w < 0 {
        w += 7
    }
    return w
}

/**
 * 计算星期几，{@value #JULIAN_LAST_YEAR}年{@value #JULIAN_LAST_MONTH}月 {@value #JULIAN_LAST_DATE}
 * 日及以前按照Julian历法，{@value #GREGORIAN_FIRST_YEAR}年 {@value #GREGORIAN_FIRST_MONTH}月
 * {@value #GREGORIAN_FIRST_DATE}日及以后按照Gregorian历法，中间的日期按照 {@value #GREGORIAN_FIRST_YEAR}年
 * {@value #GREGORIAN_FIRST_MONTH}月 {@value #GREGORIAN_FIRST_DATE}日计算。
 *
 * @param y
 *            年份
 * @param m
 *            月份
 * @param d
 *            日期
 * @return 星期几的数字表示，1-6表示星期一到星期六，0表示星期日
 */
func GetWeekday(y,m,d int) int {
    if y < JULIAN_LAST_YEAR {
        return getWeekdayForJulian(y, m, d)
    }
    if y == JULIAN_LAST_YEAR && m < JULIAN_LAST_MONTH {
        return getWeekdayForJulian(y, m, d)
    }
    if y == JULIAN_LAST_YEAR && m == JULIAN_LAST_MONTH && d <= JULIAN_LAST_DATE {
        return getWeekdayForJulian(y, m, d)
    }

    if y > GREGORIAN_FIRST_YEAR {
        return getWeekdayForGregorian(y, m, d)
    }
    if y == GREGORIAN_FIRST_YEAR && m > GREGORIAN_FIRST_MONTH {
        return getWeekdayForGregorian(y, m, d)
    }
    if y == GREGORIAN_FIRST_YEAR && m == GREGORIAN_FIRST_MONTH && d >= GREGORIAN_FIRST_DATE {
        return getWeekdayForGregorian(y, m, d)
    }

    // 剩下的都是中间的，以Gregorian历法实施第一天计算
    return getWeekdayForGregorian(GREGORIAN_FIRST_YEAR, GREGORIAN_FIRST_MONTH,
            GREGORIAN_FIRST_DATE)
}

/**
 * 计算地球时和UTC的时差，算法摘自<i><a
 * href="http://eclipse.gsfc.nasa.gov/SEhelp/deltatpoly2004.html">NASA网站</a></i><br />
 * ∆T = TT - UT 此算法在-1999年到3000年有效
 *
 * @param jd
 *            儒略日
 * @return ∆T的值，单位为秒
 */

func GetDeltaTJD(jd float64) float64 {
    cal := FromJulianDate(jd, time.UTC, false)
    year, month, _ := cal.Date()
    return GetDeltaT(year, int(month) )
}

/**
 * 计算地球时和UTC的时差，算法摘自<i><a
 * href="http://eclipse.gsfc.nasa.gov/SEhelp/deltatpoly2004.html">NASA网站</a></i><br />
 * ∆T = TT - UT 此算法在-1999年到3000年有效
 *
 * @param year
 *            要计算的年份
 * @param month
 *            要计算的月份
 * @return ∆T的值，单位为秒
 */
func GetDeltaT(year, month int) float64 {
    y := float64(year) + (float64(month) - 0.5) / 12
    if year < -500 {
        u :=  (float64(year) - 1820) / 100.0
        return -20 + 32 * u * u
    } else if year < 500 {
        u := y / 100
        u2 := u * u
        u3 := u2 * u
        u4 := u3 * u
        u5 := u4 * u
        u6 := u5 * u
        return 10583.6 - 1014.41 * u + 33.78311 * u2 - 5.952053 * u3 - 0.1798452 * u4 + 0.022174192 * u5 + 0.0090316521 * u6
    } else if year < 1600 {
        u := (y - 1000) / 100
        u2 := u * u
        u3 := u2 * u
        u4 := u3 * u
        u5 := u4 * u
        u6 := u5 * u
        return 1574.2 - 556.01 * u + 71.23472 * u2 + 0.319781 * u3 - 0.8503463 * u4 - 0.005050998 * u5 + 0.0083572073 * u6
    } else if year < 1700 {
        t := y - 1600
        t2 := t * t
        t3 := t2 * t
        return 120 - 0.9808 * t - 0.01532 * t2 + t3 / 7129
    } else if year < 1800 {
        t := y - 1700
        t2 := t * t
        t3 := t2 * t
        t4 := t3 * t
        return 8.83 + 0.1603 * t - 0.0059285 * t2 + 0.00013336 * t3 - t4 / 1174000
    } else if year < 1860 {
        t := y - 1800
        t2 := t * t
        t3 := t2 * t
        t4 := t3 * t
        t5 := t4 * t
        t6 := t5 * t
        t7 := t6 * t
        return 13.72 - 0.332447 * t + 0.0068612 * t2 + 0.0041116 * t3 - 0.00037436 * t4 + 0.0000121272 * t5 - 0.0000001699 * t6 + 0.000000000875 * t7
    } else if year < 1900 {
        t := y - 1860
        t2 := t * t
        t3 := t2 * t
        t4 := t3 * t
        t5 := t4 * t
        return 7.62 + 0.5737 * t - 0.251754 * t2 + 0.01680668 * t3 - 0.0004473624 * t4 + t5 / 233174
    } else if year < 1920 {
        t := y - 1900
        t2 := t * t
        t3 := t2 * t
        t4 := t3 * t
        return -2.79 + 1.494119 * t - 0.0598939 * t2 + 0.0061966 * t3 - 0.000197 * t4;
    } else if year < 1941 {
        t := y - 1920
        t2 := t * t
        t3 := t2 * t
        return 21.20 + 0.84493 * t - 0.076100 * t2 + 0.0020936 * t3;
    } else if year < 1961 {
        t := y - 1950
        t2 := t * t
        t3 := t2 * t
        return 29.07 + 0.407 * t - t2 / 233 + t3 / 2547;
    } else if year < 1986 {
        t := y - 1975
        t2 := t * t
        t3 := t2 * t
        return 45.45 + 1.067 * t - t2 / 260 - t3 / 718;
    } else if year < 2005 {
        t := y - 2000
        t2 := t * t
        t3 := t2 * t
        t4 := t3 * t
        t5 := t4 * t
        return 63.86 + 0.3345 * t - 0.060374 * t2 + 0.0017275 * t3 + 0.000651814 * t4 + 0.00002373599 * t5
    } else if year < 2050 {
        t := y - 2000
        t2 := t * t
        return 62.92 + 0.32217 * t + 0.005589 * t2;
    } else if year < 2150 {
        u := (y - 1820) / 100
        u2 := u * u
        return -20 + 32 * u2 - 0.5628 * (2150 - y)
    } else {
        u := (y - 1820) / 100
        u2 := u * u
        return -20 + 32 * u2;
    }
}
