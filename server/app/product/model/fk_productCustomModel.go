package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"server/app/product/model/vo_fk"
	"server/common"
	"strconv"
	"time"
)

type defaultFkProductCustomModel struct {
	rdb *redis.Client
}

const (
	// FkProductCachePrefix 缓存前缀
	FkProductCachePrefix = "fk_product:"
	// FkProductCacheExpire 缓存过期时间
	FkProductCacheExpire = 10 * time.Minute
)

func (d *defaultFkProductCustomModel) CacheGetByID(ctx context.Context, id string) (fp *FkProduct, ok bool) {
	span, _ := common.Tracer(ctx, "cache-fk_product-get-id")
	defer span.End()

	result, err := d.rdb.Get(ctx, FkProductCachePrefix+id).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			logx.Errorf("[ fk-product ] cache not exist in  %s", id)
			return nil, false
		}
		logx.Errorf("[ fk-product ]get cache error: %s", err.Error())
		return nil, false
	}

	fp = &FkProduct{}
	err = json.Unmarshal([]byte(result), fp)
	if err != nil {
		logx.Errorf("[ fk-product ] cache 反序列化 error: %s", err.Error())
		return nil, false
	}
	logx.Infof("[ fk-product ] cache hit in %s", id)
	return fp, true
}

func (d *defaultFkProductCustomModel) CacheInsert(ctx context.Context, fp *FkProduct) (ok bool) {
	span, _ := common.Tracer(context.Background(), "cache-fk_product-insert")
	defer span.End()

	data, err := json.Marshal(fp)
	if err != nil {
		logx.Errorf("[ fk-product ] cache 序列化 error: %s", err.Error())
		return false
	}

	err = d.rdb.Set(ctx, FkProductCachePrefix+fp.FkProductID, string(data), FkProductCacheExpire).Err()
	if err != nil {
		logx.Errorf("[ fk-product ] cache insert error: %s", err.Error())
		return false
	}
	logx.Infof("[ fk-product ] cache insert success in %s", fp.FkProductID)
	return true
}

func (d *defaultFkProductCustomModel) CacheDelete(ctx context.Context, id string) {
	span, _ := common.Tracer(ctx, "cache-fk_product-delete")
	defer span.End()

	err := d.rdb.Del(ctx, FkProductCachePrefix+id).Err()
	if err != nil {
		logx.Errorf("[ fk-product ] cache delete error: %s", err.Error())
		return
	}
	logx.Infof("[ fk-product ] cache delete success in %s", id)
	return
}

func (d *defaultFkProductCustomModel) CacheGetList(ctx context.Context, key string) (fps []*FkProduct, ids []string, ok bool) {
	span, _ := common.Tracer(ctx, "cache-fk_product-get-list")
	defer span.End()

	result, err := d.rdb.Get(ctx, FkProductCachePrefix+key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			logx.Errorf("[ fk-product ] cache not exist in  %s", key)
			return nil, nil, false
		}
		logx.Errorf("[ fk-product ]get cache error: %s", err.Error())
		return nil, nil, false
	}

	var idList []string
	err = json.Unmarshal([]byte(result), &idList)
	if err != nil {
		logx.Errorf("[ fk-product ] cache 反序列化 error: %s", err.Error())
		return nil, nil, false
	}

	var fpsList []*FkProduct
	for _, id := range idList {
		fp, ok := d.CacheGetByID(ctx, id)
		if !ok {
			ids = append(ids, id)
			continue
		}
		fpsList = append(fpsList, fp)
	}
	logx.Infof("[ fk-product ] cache not exist in %v", ids)

	return fpsList, ids, true
}

func (d *defaultFkProductCustomModel) getCacheKey(ctx context.Context, req *vo_fk.FkProductListReq) (key string) {
	span, _ := common.Tracer(ctx, "cache-fk_product-get-key")
	defer span.End()

	key = "list"
	return fmt.Sprintf("%s:%d:%d:%d:%d", key, req.Page, req.PageSize, req.StartTime, req.EndTime)
}

func (d *defaultFkProductCustomModel) getTotal(ctx context.Context, key string) int64 {
	span, _ := common.Tracer(ctx, "cache-fk_product-get-total")
	defer span.End()

	result, err := d.rdb.Get(ctx, FkProductCachePrefix+key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			logx.Errorf("[ fk-product ] cache not exist in  %s", key)
			return 0
		}
		logx.Errorf("[ fk-product ]get cache error: %s", err.Error())
		return 0
	}

	total, _ := strconv.ParseInt(result, 10, 64)

	return total

}

func (d *defaultFkProductCustomModel) setTotal(ctx context.Context, key string, total int64) {
	span, _ := common.Tracer(ctx, "cache-fk_product-set-total")
	defer span.End()

	err := d.rdb.Set(ctx, FkProductCachePrefix+key, total, FkProductCacheExpire).Err()
	if err != nil {
		logx.Errorf("[ fk-product ] cache insert error: %s", err.Error())
		return
	}
	logx.Infof("[ fk-product ] cache insert success in %s", key)
	return
}
