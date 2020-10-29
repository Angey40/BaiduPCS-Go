// Package converter 格式, 类型转换包
package converter

import (
	"fmt"
	"reflect"
	"strconv"
	"unicode"
	"unsafe"
)

const (
	// B byte
	B = (int64)(1 << (10 * iota))
	// KB kilobyte
	KB
	// MB megabyte
	MB
	// GB gigabyte
	GB
	// TB terabyte
	TB
	// PB petabyte
	PB
)

// ConvertFileSize 文件大小格式化输出
func ConvertFileSize(size int64, precision ...int) string {
	pint := "6"
	if len(precision) == 1 {
		pint = fmt.Sprint(precision[0])
	}
	if size < 0 {
		return "0B"
	}
	if size < KB {
		return fmt.Sprintf("%dB", size)
	}
	if size < MB {
		return fmt.Sprintf("%."+pint+"fKB", float64(size)/float64(KB))
	}
	if size < GB {
		return fmt.Sprintf("%."+pint+"fMB", float64(size)/float64(MB))
	}
	if size < TB {
		return fmt.Sprintf("%."+pint+"fGB", float64(size)/float64(GB))
	}
	if size < PB {
		return fmt.Sprintf("%."+pint+"fTB", float64(size)/float64(TB))
	}
	return fmt.Sprintf("%."+pint+"fPB", float64(size)/float64(PB))
}

// ToString unsafe 转换, 将 []byte 转换为 string
func ToString(p []byte) string {
	return *(*string)(unsafe.Pointer(&p))
}

// ToBytes unsafe 转换, 将 string 转换为 []byte
func ToBytes(str string) []byte {
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: strHeader.Data,
		Len:  strHeader.Len,
		Cap:  strHeader.Len,
	}))
}

// IntToBool int 类型转换为 bool
func IntToBool(i int) bool {
	return i != 0
}

// SliceInt64ToString []int64 转换为 []string
func SliceInt64ToString(si []int64) (ss []string) {
	ss = make([]string, 0, len(si))
	for k := range si {
		ss = append(ss, strconv.FormatInt(si[k], 10))
	}
	return ss
}

// SliceStringToInt64 []string 转换为 []int64
func SliceStringToInt64(ss []string) (si []int64) {
	si = make([]int64, 0, len(ss))
	var (
		i   int64
		err error
	)
	for k := range ss {
		i, err = strconv.ParseInt(ss[k], 10, 64)
		if err != nil {
			continue
		}
		si = append(si, i)
	}
	return
}

// SliceStringToInt []string 转换为 []int
func SliceStringToInt(ss []string) (si []int) {
	si = make([]int, 0, len(ss))
	var (
		i   int
		err error
	)
	for k := range ss {
		i, err = strconv.Atoi(ss[k])
		if err != nil {
			continue
		}
		si = append(si, i)
	}
	return
}

// MustInt 将string转换为int, 忽略错误
func MustInt(s string) (n int) {
	n, _ = strconv.Atoi(s)
	return
}

// MustInt64 将string转换为int64, 忽略错误
func MustInt64(s string) (i int64) {
	i, _ = strconv.ParseInt(s, 10, 64)
	return
}

// ShortDisplay 缩略显示字符串s, 显示长度为num, 缩略的内容用"..."填充
func ShortDisplay(s string, num int) string {
	rs := []rune(s)
	for k := 0; k < len(rs); k++ {
		if unicode.Is(unicode.C, rs[k]) { // 去除无效字符
			rs = append(rs[:k], rs[k+1:]...)
			k--
			continue
		}
		if k >= num {
			return string(rs[:k]) + "..."
		}
	}
	return string(rs)
}
