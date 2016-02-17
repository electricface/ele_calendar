package calendarutil
import (
    "testing"
    "time"
)
// 西元年分逢4的倍数闰、100的倍数不闰、400的倍数闰。
// 例如：公元1992、1996年等为4的倍数，故为闰年
// 公元1800、1900、2100年为100的倍数，当年不闰；
// 公元1600、2000、2400年为400的倍数，故有闰。

func Test_IsGregorianLeapYear(t *testing.T) {
    // IsLeapYear := IsGregorianLeapYear(1992)
    if IsGregorianLeapYear(1992) && IsGregorianLeapYear(1996) &&
        IsGregorianLeapYear(1600) && IsGregorianLeapYear(2000) &&
        IsGregorianLeapYear(2400) &&
        !IsGregorianLeapYear(1800) && !IsGregorianLeapYear(1900) && !IsGregorianLeapYear(2100) {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}

func Test_IsJulianLeapYear(t *testing.T) {
    if IsJulianLeapYear(1400) && IsJulianLeapYear(1444) && !IsJulianLeapYear(1445) {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}

func Test_IsLeapYear(t *testing.T) {
    if IsLeapYear(1400) && IsLeapYear(1996) &&
        IsLeapYear(2000) && !IsLeapYear(2001) &&
        !IsLeapYear(1800) {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}

// 1977年3月27日的儒略日是2443230，2005年5月31日的儒略日是2453522
func Test_ToJulianDateInGregorian(t *testing.T) {
    if  2443230 == ToJulianDateInGregorian(1977, 3, 27) &&
        2453522 == ToJulianDateInGregorian(2005, 5, 31) {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}

func Test_ToJulianDateInJulian(t *testing.T) {
    t.Log(ToJulianDateInJulian(1400, 3, 27))
    if 2232494 == ToJulianDateInJulian(1400, 3, 27) {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}

func Test_ToJulianDate(t *testing.T) {
    if  2443230 == ToJulianDate(1977, 3, 27) &&
        2453522 == ToJulianDate(2005, 5, 31) &&
        2232494 == ToJulianDateInJulian(1400, 3, 27) {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}

func Test_ToJulianDateInJulianHMS(t *testing.T) {
    t.Log( ToJulianDateInJulianHMS(1400,3,27, 6,6,6) )
    if  2.2324937542361114e+06 == ToJulianDateInJulianHMS(1400,3,27, 6,6,6) {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}

func Test_ToJulianDateInGregorianHMS(t *testing.T) {
    t.Log( ToJulianDateInGregorianHMS(1977,3,27, 6,6,6) )
    if  2.4432297542361114e+06 == ToJulianDateInGregorianHMS(1977,3,27, 6,6,6) {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}

func Test_ToJulianDateHMS(t *testing.T) {
    t.Log(ToJulianDateHMS(1400,3,27, 6,6,6))
    t.Log(ToJulianDateHMS(1977,3,27, 6,6,6))
    if 2.2324937542361114e+06 == ToJulianDateHMS(1400,3,27, 6,6,6) &&
        2.4432297542361114e+06 == ToJulianDateHMS(1977,3,27, 6,6,6) {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}


func Test_fromJulianDateInJulian(t *testing.T) {
    cal := fromJulianDateInJulian(2.2324937542361114e+06, time.UTC)
    timestr := cal.Format(time.UnixDate)
    t.Log(timestr)
    if "Thu Mar 27 06:06:06 UTC 1400" == timestr {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}

func Test_fromJulianDateInGregorian(t *testing.T) {
    cal := fromJulianDateInGregorian(2.4432297542361114e+06, time.UTC)
    timestr := cal.Format(time.UnixDate)
    t.Log(timestr)
    if "Sun Mar 27 06:06:06 UTC 1977" == timestr {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}


func Test_FromJulianDateLocal(t *testing.T) {
    cal := FromJulianDateLocal( 2.4432297542361114e+06, true)
    timestr := cal.Format(time.UnixDate)
    t.Log(timestr)
    if "Sun Mar 27 14:05:18 CST 1977" == timestr {
        t.Log("ok")
    } else {
        t.Error("fail")
    }

    cal = FromJulianDateLocal(2.2324937542361114e+06, true)
    timestr = cal.Format(time.UnixDate)
    t.Log(timestr)
    if "Sat Apr  5 14:06:27 LMT 1400" == timestr {
        t.Log("ok")
    } else {
        t.Error("fail")
    }

}

func Test_FromJulianDate(t *testing.T) {
    cal := FromJulianDate( 2.4432297542361114e+06, time.UTC, true)
    timestr := cal.Format(time.UnixDate)
    t.Log(timestr)
    if "Sun Mar 27 06:05:18 UTC 1977" == timestr {
        t.Log("ok")
    } else {
        t.Error("fail")
    }

    cal = FromJulianDate(2.2324937542361114e+06, time.UTC, true)
    timestr = cal.Format(time.UnixDate)
    t.Log(timestr)
    if "Sat Apr  5 06:00:44 UTC 1400" == timestr {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}

func Test_GetJulianThousandYears(t *testing.T ) {
    t.Log(GetJulianThousandYears(2.2324937542361114e+06))
    if -0.5997296256369298 == GetJulianThousandYears(2.2324937542361114e+06) {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}

func Test_GetJulianCentury(t *testing.T) {
    t.Log(GetJulianCentury(2.2324937542361114e+06))
    if -5.997296256369297 == GetJulianCentury(2.2324937542361114e+06) {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}


func Test_getWeekdayForGregorian(t *testing.T) {
    t.Log(getWeekdayForGregorian(2016,2,16))
    t.Log(getWeekdayForGregorian(2015,1,1))
    if 2 == getWeekdayForGregorian(2016,2,16) &&
        4 == getWeekdayForGregorian(2015,1,1) {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}

func Test_getWeekdayForJulian(t *testing.T) {
    t.Log(getWeekdayForJulian(1416,2,16))
    t.Log(getWeekdayForJulian(1415,1,1))
    if 0 == getWeekdayForJulian(1416,2,16) &&
        2 == getWeekdayForJulian(1415,1,1) {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}


func Test_GetWeekday(t *testing.T) {
    if 0 == GetWeekday(1416,2,16) &&
        4 == GetWeekday(2015,1,1) {
        t.Log("ok")
    } else {
        t.Error("fail")
    }
}
