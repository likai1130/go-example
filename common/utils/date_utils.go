package utils

import (
	"strconv"
	"time"
)

/**
 * 获取当前 8 位字符长度的日期
 */
func GetCurrentDate() (dateLen8 string) {
	currentDate := time.Now().Format("20060102150405")[:8]
	return currentDate
}

/**
	字符类型时间戳
 */
func GetCurrentTimeUnix() string {
	//当前时间戳
	t1 := time.Now().Unix()
	timeContent := strconv.FormatInt(t1, 10)
	return timeContent
}

/* 获取当前 8 位字符长度的日期
 */
func GetCurrentDate8() (dateLen8 string) {
	currentDate := time.Now().Format("20060102")[:8]
	return currentDate
}

/**
当前时间
*/
func CurrentTime() string {
	timeStr := time.Now().Format("2006-01-02 15:04:05") //当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
	return timeStr
}
