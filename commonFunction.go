package kbutils

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// GeneratePIDFile 生成一个pID文件，格式：<name>.<ID>.pID
func GeneratePIDFile(name string, id int) {
	var filename string
	if id == 0 {
		filename = fmt.Sprintf("%s.pid", name)
	} else {
		filename = fmt.Sprintf("%s.%d.pid", name, id)
	}

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	pid := os.Getpid()
	fmt.Fprintf(f, "%d", pid)
	//	Debug("new %s[%d], pID: %d", name, id, pid)
}

var r *rand.Rand

// GenRandomString 生成随机字符串
func GenRandomString(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	if r == nil {
		r = rand.New(rand.NewSource(time.Now().Unix()))
	}
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// GetRandomRange 生成一个1-n的随机数
func GetRandomRange(_min, _max int) int {
	if r == nil {
		r = rand.New(rand.NewSource(time.Now().Unix()))
	}
	return r.Intn(_max-_min) + _min
}

// SliceRandList 生成一个随机数序列
func SliceRandList(min, max int) []int {
	if max < min {
		min, max = max, min
	}
	length := max - min + 1
	if r == nil {
		r = rand.New(rand.NewSource(time.Now().Unix()))
	}
	list := r.Perm(length)
	for index := range list {
		list[index] += min
	}
	return list
}

// GetRandom 生成一个1-n的随机数
func GetRandom(_max int) int {
	if r == nil {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	return r.Intn(_max) + 1
}

// TimeToStr 格式化时间戳
func TimeToStr(t int64) string {
	return time.Unix(t, 0).Format("2006-01-02 15:04:05")
}

// StrToTime 时间字符串转时间戳
func StrToTime(str string) int64 {
	loc, _ := time.LoadLocation("Local")
	theTime, err := time.ParseInLocation("2006-01-02 15:04:05", str, loc)
	if err == nil {
		return theTime.Unix()
	} else {
		return 0
	}
}

// NextTime 当前时间下一个指定时间
func NextTime(h, m, s int) int64 {
	now := time.Now()
	year, mon, day := now.Date()
	hour, min, sec := now.Clock()

	ts := StrToTime(fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", year, mon, day, 0, 0, 0))

	if (3600*h + 60*m + s) <= (3600*hour + 60*min + sec) {
		ts += int64(3600*24 + 3600*h + 60*m + s)
	} else {
		ts += int64(3600*h + 60*m + s)
	}
	return ts
}

// LowerCasedName 小写字符串
func LowerCasedName(name string) string {
	newstr := make([]rune, 0)
	firstTime := true

	for _, chr := range name {
		if isUpper := 'A' <= chr && chr <= 'Z'; isUpper {
			if firstTime == true {
				firstTime = false
			} else {
				newstr = append(newstr, '_')
			}
			chr -= ('A' - 'a')
		}
		newstr = append(newstr, chr)
	}
	return string(newstr)
}

// UpperCasedName 大写字符串
func UpperCasedName(name string) string {
	newstr := make([]rune, 0)
	upNextChar := true
	for _, chr := range name {
		switch {
		case upNextChar:
			upNextChar = false
			chr -= ('a' - 'A')
		case chr == '_':
			upNextChar = true
			continue
		}
		newstr = append(newstr, chr)
	}
	return string(newstr)
}

/**
	SendEmail 发送邮件
	param：to 发送给谁，比如：example@example.com;example1@163.com;example2@sina.com.cn;...
	param：user : example@example.com login smtp server user
	password: xxxxx login smtp server password
	host: smtp.example.com:port   smtp.163.com:25
	to: example@example.com;example1@163.com;example2@sina.com.cn;...
	subject:The subject of mail
	body: The content of mail
	mailtyoe: mail type html or text
	result：error 错误对象
**/
func SendEmail(to string, user string, password string, host string, subject string, body string, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var contentType string
	if mailtype == "html" {
		contentType = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	return err
}

// Uint32ToBytes 从int32转化为[]byte
func Uint32ToBytes(i uint32) []byte {
	return []byte{byte((i >> 24) & 0xff), byte((i >> 16) & 0xff),
		byte((i >> 8) & 0xff), byte(i & 0xff)}
}

// BytesToUint32 从[]byte转化为int32
func BytesToUint32(buf []byte) uint32 {
	return uint32(buf[0])<<24 + uint32(buf[1])<<16 + uint32(buf[2])<<8 +
		uint32(buf[3])
}

// ListDir 获取指定目录下的所有文件/文件夹，不进入下一级目录搜索，可以匹配后缀过滤。
func ListDir(dirPth string, suffix string, fileOrDir bool) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fileOrDir {
			// 列出文件
			if fi.IsDir() { // 忽略目录
				continue
			}
		} else {
			// 列出文件夹
			if !fi.IsDir() { // 忽略目录
				continue
			}
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}

// WalkDir 获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix)                                                     //忽略后缀匹配的大小写
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}

// FileExist 判断文件是否存在。
func FileExist(path string) bool {
	_, err := os.Stat(path)
	return (err == nil)
}

// GetExternalIP 获取外部ip。
func GetExternalIP() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return (string(body))
}

// GetInternalIP 获取内部ip。
func GetInternalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// WhereAmI return a string containing the file name, function name
// and the line number of a specified entry on the call stack
func WhereAmI(depthList ...int) string {
	var depth int
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}
	function, file, line, _ := runtime.Caller(depth)
	return fmt.Sprintf("File: %s  Function: %s Line: %d", chopPath(file), runtime.FuncForPC(function).Name(), line)
}

// return the source filename after the last slash
func chopPath(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	} else {
		return original[i+1:]
	}
}

// CheckMapKeyAndInit 检验MAP里键是否存在,不存在就赋值,python map.get(key,0)
func CheckMapKeyAndInit(v map[int]int, key int, initValue int) {
	if _, ok := v[key]; !ok {
		v[key] = initValue
	}
}
