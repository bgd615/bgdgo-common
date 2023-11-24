package strc

import (
	"crypto/md5"
	"encoding/json"
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func MapInterface2String(inputData map[string]interface{}) map[string]string {
	outputData := map[string]string{}
	for key, value := range inputData {
		switch value.(type) {
		case string:
			outputData[key] = value.(string)
		case int:
			tmp := value.(int)
			outputData[key] = strconv.Itoa(tmp)
		case int64:
			tmp := value.(int64)
			outputData[key] = strconv.FormatInt(tmp, 10)
		}
	}
	return outputData
}

func JobParseCtl(ctl string) (string, string, string, error) {
	arr := strings.Split(ctl, ".")
	if len(arr) < 3 {
		return "", "", "", errors.New("Err: parse ctl err.")
	}
	sys := arr[0][0:5]
	return sys, arr[0], fmt.Sprintf("%v.%v", arr[1], arr[2]), nil
}

func Md5s(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func DisableEscapeHtml(dat interface{}) (string, error) {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	if err := jsonEncoder.Encode(dat); err != nil {
		return "", err
	}
	return bf.String(), nil
}
