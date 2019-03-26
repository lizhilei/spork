package sun

import (
	"fmt"
	"math"
	"time"
)

func SunRise(lat, lng float64) (t time.Time) {
	lat = toRadians(lat)

	tNow := time.Now()
	tStart, _ := time.Parse("01/02/2006", "01/01/2000")

	//2000年1月1日到今天天数
	dayCount := (tNow.Unix() - tStart.Unix()) / (3600 * 24)

	var ut0 float64 = 180
	var utStart float64 = 0
	var h float64 = math.Sin(toRadians(-0.833))

	for ; math.Abs(utStart-ut0) >= 0.1; ut0 = utStart {
		t := float64((float64(dayCount) + ut0/float64(360)) / 36525)    // 世纪数
		L := 280.460 + 36000.777*t                                      // 太阳平均黄径
		G := toRadians(357.528 + 35999.050*t)                           // 太阳平近点角
		lamda := toRadians(L + 1.915*math.Sin(G) + 0.020*math.Sin(2*G)) // 太阳黄道经度
		epc := toRadians(23.4393 - 0.0130*t)                            // 地球倾角
		sigam := math.Asin(math.Sin(epc) * math.Sin(lamda))             // 太阳的偏差

		// 格林威治时间太阳时间角
		gha := ut0 - 180 - 1.915*math.Sin(G) - 0.020*math.Sin(2*G) + 2.466*math.Sin(2*lamda) - 0.053*math.Sin(4*lamda);
		// 修正值e
		e := toDegrees(math.Acos(h - math.Tan(lat)*math.Tan(sigam)))
		utStart = ut0 - gha - lng - e
	}

	zone := 0
	if lng >= 0 {
		zone = (int)(lng/15 + 1) // 当前时区
	} else {
		zone = (int)(lng/15 - 1) // 当前时区
	}
	t, _ = time.Parse("01/02/2006 15:04", time.Now().Format("01/02/2006 ") + fmt.Sprintf( "%02d:%02d",
		int32(utStart/15+float64(zone)),
		int32(60*(utStart/15+float64(zone)-float64(int32(utStart/15+float64(zone))))),
	))
	return t
}

func SunSet(lat, lng float64) (t time.Time) {
	lat = toRadians(lat)

	tNow := time.Now()
	tStart, _ := time.Parse("01/02/2006", "01/01/2000")

	//2000年1月1日到今天天数
	dayCount := (tNow.Unix() - tStart.Unix()) / (3600 * 24)

	var ut0 float64 = 180
	var utStart float64 = 0
	var h float64 = math.Sin(toRadians(-0.833))

	for ; math.Abs(utStart-ut0) >= 0.1; ut0 = utStart {
		t := float64((float64(dayCount) + ut0/float64(360)) / 36525)    // 世纪数
		L := 280.460 + 36000.777*t                                      // 太阳平均黄径
		G := toRadians(357.528 + 35999.050*t)                           // 太阳平近点角
		lamda := toRadians(L + 1.915*math.Sin(G) + 0.020*math.Sin(2*G)) // 太阳黄道经度
		epc := toRadians(23.4393 - 0.0130*t)                            // 地球倾角
		sigam := math.Asin(math.Sin(epc) * math.Sin(lamda))             // 太阳的偏差

		// 格林威治时间太阳时间角
		gha := ut0 - 180 - 1.915*math.Sin(G) - 0.020*math.Sin(2*G) + 2.466*math.Sin(2*lamda) - 0.053*math.Sin(4*lamda);
		// 修正值e
		e := toDegrees(math.Acos(h - math.Tan(lat)*math.Tan(sigam)))
		utStart = ut0 - gha - lng + e
	}

	zone := 0
	if lng >= 0 {
		zone = (int)(lng/15 + 1) // 当前时区
	} else {
		zone = (int)(lng/15 - 1) // 当前时区
	}
	t, _ = time.Parse("01/02/2006 15:04", time.Now().Format("01/02/2006 ") + fmt.Sprintf( "%02d:%02d",
		int32(utStart/15+float64(zone)),
		int32(60*(utStart/15+float64(zone)-float64(int32(utStart/15+float64(zone))))),
	))
	return t
}

func toRadians(a float64) (r float64) {
	return math.Pi / 180 * a
}

func toDegrees(r float64) (a float64) {
	return r * 180 / math.Pi
}
