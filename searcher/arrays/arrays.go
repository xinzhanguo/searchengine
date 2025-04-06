package arrays

const (
	LOW  = 0
	HIGH = 1
)

// BinarySearch 二分查找
func BinarySearch(arr []uint64, target uint64) bool {
	low := 0
	high := len(arr) - 1
	for low < high {
		mid := (low + high) >> 1
		if arr[mid] >= target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return arr != nil && arr[low] == target
}

func ArrayUint32Exists(arr []uint32, target uint32) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func ArrayUint64Exists(arr []uint64, target uint64) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func ArrayStringExists(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

// MergeArrayUint32 合并两个数组
func MergeArrayUint32(target []uint64, source []uint64) []uint64 {

	for _, val := range source {
		if !BinarySearch(target, val) {
			target = append(target, val)
		}
	}
	return target
}

func Find(arr []uint64, target uint64) int {
	for index, v := range arr {
		if v == target {
			return index
		}
	}
	return -1
}
