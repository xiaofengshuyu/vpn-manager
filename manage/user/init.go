package user

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/xiaofengshuyu/vpn-manager/manage/common"
)

var (
	logger = common.Logger
)
var (
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func makeVertifyCode() string {
	return fmt.Sprintf("%06d", random.Intn(1000000))
}
