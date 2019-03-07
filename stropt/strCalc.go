package stropt

import (
	"errors"
	"strconv"
	"strings"
)

/*
	用于字符串转数字，运算后，再将结果转为字符串或数字
*/

// 累加float类型数据 prec 精度
func SumStrFloat(fs []string, prec int) (sResult string, result float64, err error) {
	if len(fs) < 1 {
		err = errors.New("参数个数不正确!")
		return
	}

	for _, item := range fs {
		f := 0.0
		f, err = strconv.ParseFloat(item, 64)
		if err != nil {
			return
		}

		result += f
	}

	sResult = strconv.FormatFloat(result, 'f', prec, 64)
	return
}

// 累计uinit类型数据
func SumStrUint(fs []string) (sResult string, result uint64, err error) {
	if len(fs) < 1 {
		err = errors.New("参数个数不正确!")
		return
	}

	for _, item := range fs {
		tmp := uint64(0)

		if strings.HasSuffix(item, ".000000") {
			item = strings.TrimSuffix(item, ".000000")
		}

		tmp, err = strconv.ParseUint(item, 10, 64)
		if err != nil {
			return
		}

		result += tmp
	}

	sResult = strconv.FormatUint(result, 10)
	return
}

// 累计init类型数据
func SumStrInt(fs []string) (sResult string, result int64, err error) {
	if len(fs) < 1 {
		err = errors.New("参数个数不正确!")
		return
	}

	for _, item := range fs {

		if strings.HasSuffix(item, ".000000") {
			item = strings.TrimSuffix(item, ".000000")
		}

		tmp := int64(0)
		tmp, err = strconv.ParseInt(item, 10, 64)
		if err != nil {
			return
		}

		result += tmp
	}

	sResult = strconv.FormatInt(result, 10)
	return
}

// float 除 fs1 / fs2 ; prec： 精度
func DivideStrFloat(fs1 string, fs2 string, prec int) (sResult string, result float64, err error) {
	sResult = "0" //默认值
	if fs1 == "0" {
		return
	}

	f1 := 0.0
	f1, err = strconv.ParseFloat(fs1, 64)
	if err != nil {
		return
	}

	// 返回0
	if f1 > -0.000001 && f1 < 0.000001 {
		return
	}

	f2 := 0.0
	f2, err = strconv.ParseFloat(fs2, 64)
	if err != nil {
		return
	}

	if f2 > -0.000001 && f2 < 0.000001 {
		err = errors.New("分母为0;分子:" + fs1 + ";分母:" + fs2)
		return
	}

	result = f1 / f2
	sResult = strconv.FormatFloat(result, 'f', prec, 64)
	return
}
