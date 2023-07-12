package utils

import (
	"archive/zip"
	"bufio"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// 获取当前日期时间
func GetCurrentDateTime() string {
	timeObj := time.Now()
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	dateTime := fmt.Sprintf("%d-%d-%d %d:%d:%d", year, month, day, hour, minute, second)
	return dateTime
}

// 获取当前日期
func GetCurrentDate(format string) string {
	timeObj := time.Now()
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	dateTime := fmt.Sprintf(format, year, month, day)
	return dateTime
}

// 计算加密后的密码
func GetEncryptPassword(password string) string {
	secret := "momodaqiq*()"
	adminPass := fmt.Sprintf("%x", md5.Sum([]byte(password+secret)))
	return adminPass
}

// base64加密
func GetBase64Encode(str string) string {
	input := []byte(str)
	encodeString := base64.StdEncoding.EncodeToString(input)
	return encodeString
}

// base64解密
func GetBase64Decode(str string) string {
	res := ""
	decodeBytes, error := base64.StdEncoding.DecodeString(str)
	if error != nil {
		fmt.Println("Base64 decode fail,", error)
	} else {
		res = string(decodeBytes)
	}
	return res
}

// url加密
func GetUrlEncode(str string) string {
	input := []byte(str)
	encodeString := base64.URLEncoding.EncodeToString(input)
	return encodeString
}

// url解密
func GetUrlDecode(str string) string {
	res := ""
	decodeBytes, error := base64.URLEncoding.DecodeString(str)
	if error != nil {
		fmt.Println("Url decode fail,", error)
	} else {
		res = string(decodeBytes)
	}
	return res
}

// 写入文件
func WriteFile(filePath string, content string, isLine bool) {
	file, error := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if error != nil {
		fmt.Println("Open file failed,", error)
		return
	}
	write := bufio.NewWriter(file)
	write.WriteString(content)
	if isLine {
		write.WriteString("\n")
	}
	write.Flush()
	file.Close()
}

// 写入日志文件
func WriteLog(group string, file string, message string, isWrite bool, isShow bool) {
	dateTime := GetCurrentDateTime()
	date := GetCurrentDate("%d-%d-%d")
	filePath := "logs/debug_" + date + ".log"
	logContent := dateTime + " [" + group + "] 发生在文件:" + file + " 消息:" + message
	if isWrite {
		WriteFile(filePath, logContent, true)
	}
	if isShow {
		fmt.Println(logContent)
	}
}

// srcFile could be a single file or a directory
func MakeZip(srcFile []string, destZip string) error {
	zipfile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	for _, s := range srcFile {
		filepath.Walk(s, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}

			header.Name = strings.TrimPrefix(path, filepath.Dir(s)+"/")
			// header.Name = path
			if info.IsDir() {
				header.Name += "/"
			} else {
				header.Method = zip.Deflate
			}

			writer, err := archive.CreateHeader(header)
			if err != nil {
				return err
			}

			if !info.IsDir() {
				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer file.Close()
				_, err = io.Copy(writer, file)
			}
			return err
		})
	}

	return err
}
