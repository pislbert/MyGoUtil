// 参考 github.com/mohae
package mem

import (
	"github.com/mohae/deepcopy"
)

func Copy(src interface{}) interface{} {
	return deepcopy.Copy(src)
}
