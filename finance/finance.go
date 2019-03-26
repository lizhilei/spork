package finance

import (
	"math"
)

/**
 * 基于固定利率及等额分期付款方式，返回贷款的每期付款额。
 *
 * @param r
 *            必需。贷款利率
 * @param n
 *            必需。该项贷款的付款总数。
 * @param p
 *            必需。现值，或一系列未来付款的当前值的累积和，也称为本金。
 * @param f
 *            必须。未来值，或在最后一次付款后希望得到的现金余额，如果省略 fv，则假设其值为 0（零），也就是一笔贷款的未来值为 0。
 * @param t
 *            必须。数字 0（零）或 1，用以指示各期的付款时间是在期初还是期末。 0 或省略 期末 false 1 期初 true
 * @return
 */
func Pmt(r float64, n float64, p float64, f float64, t bool) (retval float64) {
	if r == 0 {
		retval = -1 * (f + p) / n
	} else {
		var r1 float64 = r + 1
		r2 := r1
		if !t {
			r2 = 1
		}
		retval = (f + p*math.Pow(r1, n)) * r / (r2 * (1 - math.Pow(r1, n)))
	}
	return
}

/**
 * 基于固定利率及等额分期付款方式，返回某项投资的总期数。
 *
 * @param r
 *            必需。各期利率。
 * @param y
 *            必需。各期所应支付的金额，其数值在整个年金期间保持不变。通常，pmt 包括本金和利息，但不包括其他费用或税款。
 * @param p
 *            必需。现值，或一系列未来付款的当前值的累积和。
 * @param f
 *            必需。未来值，或在最后一次付款后希望得到的现金余额。如果省略 fv，则假设其值为 0（例如，一笔贷款的未来值即为 0）。
 * @param t
 *            必需。数字 0 或 1，用以指定各期的付款时间是在期初还是期末。 0 或省略 期末 1 期初
 * @return
 */
func Nper(r float64, y float64, p float64, f float64, t bool) (retval float64) {
	if r == 0 {
		retval = -1 * (f + p) / y
	} else {
		r1 := r + 1
		r2 := r1
		if !t {
			r2 = 1
		}
		ryr := r2 * y / r
		a1 := math.Log(ryr - f)
		a2 := math.Log(p + ryr)
		if ryr-f < 0 {
			a1 = math.Log(f - ryr)
			a2 = math.Log(-p - ryr)
		}
		a3 := math.Log(r1)
		retval = (a1 - a2) / a3

	}
	return
}

func Npv(r float64, cfs []float64) (npv float64) {
	var r1 float64 = r + 1
	var trate float64 = r1
	iSize := len(cfs)
	for i := 0; i < iSize; i += 1 {
		npv += cfs[i] / trate
		trate *= r1
	}
	return
}

/**
 * 基于固定利率及等额分期付款方式，返回某项投资的未来值。
 *
 * @param r
 *            各期利率
 * @param n
 *            总投资期，即该项投资的付款期总数
 * @param y
 *            必需。各期所应支付的金额，其数值在整个年金期间保持不变。通常，pmt 包括本金和利息，但不包括其他费用或税款。
 * @param f
 *            必需。未来值，或在最后一次付款后希望得到的现金余额。如果省略 fv，则假设其值为 0（例如，一笔贷款的未来值即为 0）。
 * @param t
 *            必需。数字 0 或 1，用以指定各期的付款时间是在期初还是期末。 0 或省略 期末 1 期初
 */
func Pv(r float64, n float64, y float64, f float64, t bool) (retval float64) {
	if r == 0 {
		retval = -1 * ((n * y) + f)
	} else {
		r1 := r + 1
		r2 := r1
		if !t {
			r2 = 1
		}
		retval = (((1-math.Pow(r1, n))/r)*r2*y - f) / math.Pow(r1, n)
	}

	return
}

/**
 * 基于固定利率及等额分期付款方式，返回某项投资的未来值。
 *
 * @param r
 *            必需。贷款利率
 * @param n
 *            必需。该项贷款的付款总数。
 * @param y
 *            必需。各期所应支付的金额，其数值在整个年金期间保持不变。通常，pmt 包括本金和利息，但不包括其他费用或税款。
 * @param p
 *            必需。现值，或一系列未来付款的当前值的累积和，也称为本金。
 * @param t
 *            必须。数字 0（零）或 1，用以指示各期的付款时间是在期初还是期末。 0 或省略 期末 1 期初
 *
 */
func Fv(r float64, n float64, y float64, p float64, t bool) (retval float64) {
	if r == 0 {
		retval = -1 * (p + (n + y))
	} else {
		r1 := r + 1
		r2 := r1
		if !t {
			r2 = 1
		}
		retval = ((1-math.Pow(r1, n))*r2*y)/r - p*math.Pow(r1, n)

	}
	return
}

/**
 * 基于固定利率及等额分期付款方式，返回贷款的每期付款额。
 *
 * @param nper   年金的付款总期数。
 * @param pmt    每期的付款金额，在年金周期内不能更改。 通常，pmt 包括本金和利息，但不含其他费用或税金。 如果省略 pmt，则必须包括 fv 参数。
 * @param pv     现值即一系列未来付款当前值的总和。
 * @param fv     未来值，或在最后一次付款后希望得到的现金余额。如果省略 fv，则假定其值为 0（例如，贷款的未来值是 0）。如果省略 fv，则必须包括 pmt 参数。
 * @param type   类型数字 0 或 1，用以指定各期的付款时间是在期初还是期末。
 * @param Guess  预期利率。
 * @return
 */
func Rate(nper, pmt, pv, fv, tp, guess float64) (rate float64) {
	const (
		FINANCIAL_MAX_ITERATIONS int64   = 128
		FINANCIAL_PRECISION      float64 = 0.000000001
	)
	if guess <= 0 || guess >= 1 {
		guess = 0.1
	}

	var y, y0, y1, x0, x1, f, i float64
	rate = guess
	if math.Abs(rate) < FINANCIAL_PRECISION {
		y = pv*(1+nper*rate) + pmt*(1+rate*tp)*nper + fv
	} else {
		f = math.Exp(nper * math.Log(1+rate))
		y = pv*f + pmt*(1/rate+tp)*(f-1) + fv
	}
	y0 = pv + pmt*nper + fv
	y1 = pv*f + pmt*(1/rate+tp)*(f-1) + fv

	x1 = rate
	for {
		if !((math.Abs(y0-y1) > FINANCIAL_PRECISION) && (1 < FINANCIAL_MAX_ITERATIONS)) {
			break
		}
		rate = (y1*x0 - y0*x1) / (y1 - y0)
		x0 = x1
		x1 = rate

		if math.Abs(rate) < FINANCIAL_PRECISION {
			y = pv*(1+nper*rate) + pmt*(1+rate*tp)*nper + fv
		} else {
			f = math.Pow(1+rate, nper)
			y = pv*f + pmt*(1/rate+tp)*(f-1) + fv
		}

		y0 = y1
		y1 = y
		i += 1
	}

	return
}
