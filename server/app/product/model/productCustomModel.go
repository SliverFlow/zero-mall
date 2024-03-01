package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	"k8s.io/apimachinery/pkg/util/json"
	"server/common"
	"time"
)

// 需要缓存层的时候再把相关的东西分离出来
// 公共方法 以及 缓存层操作
type defaultProductCustomModel struct {
	c *redis.Client
}

const ProductCacheKey = "product:"
const ExpireTime = 360 * time.Second

// CacheGetProductByID 从缓存中获取商品信息
func (d *defaultProductCustomModel) CacheGetProductByID(ctx context.Context, id string) (product *Product, ok bool) {
	span, _ := common.Tracer(ctx, "cache-get-product-by-id")
	defer span.End()

	// 从缓存中获取商品信息
	cacheKey := ProductCacheKey + id
	val, err := d.c.Get(ctx, cacheKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, false
		}
		logx.Errorf("redis get cache error: %s", err)
		return nil, false
	}

	// 反序列化
	product = &Product{}
	_ = json.Unmarshal([]byte(val), product)

	return product, true
}

func (d *defaultProductCustomModel) CacheInsertProduct(ctx context.Context, product *Product) (ok bool) {
	span, _ := common.Tracer(ctx, "cache-insert-product")
	defer span.End()

	fmt.Println("cache insert product")

	// 序列化
	data, err := json.Marshal(product)
	if err != nil {
		logx.Errorf("json marshal error: %s", err)
		return false
	}

	// 插入缓存
	cacheKey := ProductCacheKey + product.ProductID
	err = d.c.Set(ctx, cacheKey, data, ExpireTime).Err()
	if err != nil {
		logx.Errorf("redis set cache error: %s", err)
		return false
	}

	return true

}
