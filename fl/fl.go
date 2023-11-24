package fl

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func IsFile(f string) bool {
	fi, e := os.Stat(f)
	if e != nil {
		return false
	}
	return !fi.IsDir()
}

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func FileExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func Copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}
	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not regular file", src)
	}
	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	distination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer distination.Close()

	nBytes, err := io.Copy(distination, source)
	return nBytes, err

}

func HashCode(key string) int {
	var index int = 0
	index = int(key[0])
	for k := 0; k < len(key); k++ {
		index *= (1103515245 + int(key[k]))
	}
	index >>= 27
	index &= 16 - 1
	return index
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func IsExpired(st string, et string, expt int64) (bool, error) {
	loc, _ := time.LoadLocation("Local")
	timeLayout := "2006-01-02 15:04:05"
	stheTime, _ := time.ParseInLocation(timeLayout, st, loc)
	sst := stheTime.Unix()
	etheTime, _ := time.ParseInLocation(timeLayout, et, loc)
	est := etheTime.Unix()
	if est-sst > expt {
		return true, nil
	}
	return false, nil
}

func CalcFileMD5(filename string) (string, error) {
	f, err := os.Open(filename) //打开文件
	if nil != err {
		fmt.Println(err)
		return "", err
	}
	defer f.Close()
	md5Handle := md5.New()         //创建 md5 句柄
	_, err = io.Copy(md5Handle, f) //将文件内容拷贝到 md5 句柄中
	if nil != err {
		fmt.Println(err)
		return "", err
	}
	md := md5Handle.Sum(nil)        //计算 MD5 值，返回 []byte
	md5str := fmt.Sprintf("%x", md) //将 []byte 转为 string
	return md5str, nil
}
func ReadFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func ReadFile1(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
