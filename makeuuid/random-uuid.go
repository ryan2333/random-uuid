package makeuuid

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomUuid(uuidLen int) (uuid string) {
	for i := 0; i < uuidLen; i++ {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) % 26
		if n%2 == 0 {
			uuid = fmt.Sprintf("%s%c", uuid, 'a'+n)
		} else {
			uuid = fmt.Sprintf("%s%c", uuid, 'A'+n)
		}
	}
	return
}
