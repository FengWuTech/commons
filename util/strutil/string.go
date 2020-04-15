package strutil

import (
	"bytes"
	"crypto/rand"
	"github.com/unknwon/com"
	"math/big"
	mrand "math/rand"
	"strconv"
	"strings"
	"time"
)

func IntSliceToStringSlice(ints []int) []string {
	var strSlice []string
	for _, v := range ints {
		strSlice = append(strSlice, strconv.Itoa(v))
	}
	return strSlice
}

func StringToIntSlice(str string, separator string) []int {
	var arrIntValues []int
	arrStrValues := strings.Split(str, separator)
	for _, v := range arrStrValues {
		arrIntValues = append(arrIntValues, com.StrTo(v).MustInt())
	}
	return arrIntValues
}

func IntSliceToString(ints []int, split string) string {
	var strs []string
	for _, v := range ints {
		strs = append(strs, strconv.Itoa(v))
	}
	return strings.Join(strs, split)
}

func CreateRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

func GenNonceStr(len int) string {
	var nonce = make([]byte, len)
	var r = mrand.New(mrand.NewSource(time.Now().Unix()))
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		nonce[i] = byte(b)
	}
	return string(nonce)
}
