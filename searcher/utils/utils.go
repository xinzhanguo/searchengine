package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/spaolacci/murmur3"
)

func ExecTime(fn func()) float64 {
	start := time.Now()
	fn()
	tc := float64(time.Since(start).Nanoseconds())
	return tc / 1e6
}

func ExecTimeWithError(fn func() error) (float64, error) {
	start := time.Now()
	err := fn()
	tc := float64(time.Since(start).Nanoseconds())
	return tc / 1e6, err
}

func Encoder(data interface{}) []byte {
	if data == nil {
		return nil
	}
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	return buffer.Bytes()
}

func Decoder(data []byte, v interface{}) {
	if data == nil {
		return
	}
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	err := decoder.Decode(v)
	if err != nil {
		panic(err)
	}
}

func Murmur3(key []byte) uint64 {
	hasher := murmur3.New64WithSeed(1)
	hasher.Write(key)
	return hasher.Sum64()
}

// StringToInt 字符串转整数
func StringToInt(value string) uint64 {
	return Murmur3([]byte(value))
}

// StringToInt 字符串转整数
func StringToInt64(value string) uint64 {
	return Murmur3([]byte(value))
}

func Uint32Comparator(a, b interface{}) int {
	aAsserted := a.(uint32)
	bAsserted := b.(uint32)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

func Uint32ToBytes(i uint64) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint64(buf, i)
	return buf
}

func Uint64ToBytes(i uint64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, i)
	return buf
}

// QuickSortAsc 快速排序
func QuickSortAsc(arr []int, start, end int, cmp func(int, int)) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				if cmp != nil {
					cmp(i, j)
				}
				i++
				j--
			}
		}

		if start < j {
			QuickSortAsc(arr, start, j, cmp)
		}
		if end > i {
			QuickSortAsc(arr, i, end, cmp)
		}
	}
}
func DeleteArray(array []uint64, index int) []uint64 {
	return append(array[:index], array[index+1:]...)
}

func ReleaseAssets(file fs.File, out string) {
	if file == nil {
		return
	}

	if out == "" {
		panic("out is empty")
	}

	//判断out文件是否存在
	if _, err := os.Stat(out); os.IsNotExist(err) {
		//读取文件信息
		fileInfo, err := file.Stat()
		if err != nil {
			panic(err)
		}
		buffer := make([]byte, fileInfo.Size())
		_, err = file.Read(buffer)
		if err != nil {
			panic(err)
		}

		// 读取输出文件目录
		outDir := filepath.Dir(out)
		err = os.MkdirAll(outDir, os.ModePerm)
		if err != nil {
			panic(err)
		}

		//创建文件
		outFile, _ := os.Create(out)
		defer func(outFile *os.File) {
			err := outFile.Close()
			if err != nil {
				panic(err)
			}
		}(outFile)

		err = os.WriteFile(out, buffer, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

}

// DirSizeB DirSizeMB getFileSize get file size by path(B)
func DirSizeB(path string) int64 {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	if err != nil {
		return 0
	}

	return size
}

// RemovePunctuation 移除所有的标点符号
func RemovePunctuation(str string) string {
	reg := regexp.MustCompile(`\p{P}+`)
	return reg.ReplaceAllString(str, "")
}

// RemoveSpace 移除所有的空格
func RemoveSpace(str string) string {
	reg := regexp.MustCompile(`\s+`)
	return reg.ReplaceAllString(str, "")
}

// init 注册数据类型
// 防止 gob: type not registered for interface: map[string]interface {}
func init() {
	gob.Register(map[string]interface{}{})
	gob.Register([]interface{}{})
}
