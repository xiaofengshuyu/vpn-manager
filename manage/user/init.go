package user

import (
	"crypto/md5"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/xiaofengshuyu/vpn-manager/manage/common"
)

var (
	// ErrVertifyCodeInvalid veritify code ivalid error
	ErrVertifyCodeInvalid = errors.New("veritify code is invalid")
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

func makeToken(base string) string {
	str := fmt.Sprintf("%s_%d", base, random.Intn(1<<30))
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
