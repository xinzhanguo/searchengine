package pagination

import (
	"fmt"
	"testing"
)

func TestPagination_GetPage(t *testing.T) {
	pagination := new(Pagination)

	pagination.Init(10, 100)

	for i := 1; i <= 10; i++ {
		start, end := pagination.GetPage(i)
		fmt.Println(start, end)
	}
}
