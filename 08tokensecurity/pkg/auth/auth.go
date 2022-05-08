package auth

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"time"
)

/*
@author RandySun
@create 2022-03-29-9:15
*/

// Token token认证
type Token struct {
	AppID     string
	AppSecret string
}

// GetRequestMetadata 获取当前请求认证所需的元数据（metadata）
func (t *Token) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	// 设置一个种子
	rand.Seed(time.Now().UnixNano())
	// Intn返回一个取值范围在[0,n)的伪随机int值
	num := rand.Intn(100) + 1 // 随机1-100
	rangeSeed := strconv.Itoa(num)
	log.Println("GetRequestMetadata 每次访问服务端方法都会被调用 添加自定义认证", rangeSeed)

	return map[string]string{"app_id": t.AppID, "app_secret": t.AppSecret, "range_seed": rangeSeed}, nil
}

// RequireTransportSecurity 是否需要基于 TLS 认证进行安全传输,返回false不进行TLS验证
func (t *Token) RequireTransportSecurity() bool {
	return false
}
