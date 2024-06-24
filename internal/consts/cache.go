package consts

// CachePrefix Redis缓存统一前缀
const (
	CachePrefix               = "Cache::Novel::"
	HomeBookCache             = "homeBookCache"
	LatestNewsCache           = "latestNewsCache"
	BookVisitRankCache        = "bookVisitRankCache"
	BookNewestRankCache       = "bookNewestRankCache"
	BookUpdateRankCache       = "bookUpdateRankCache"
	HomeFriendLinkCache       = "homeFriendLinkCache"
	BookCategoryListCache     = "bookCategoryListCache"
	BookInfoCache             = "bookInfoCache"
	BookChapterCache          = "bookChapterCache"
	BookContentCache          = "bookContentCache"
	LastUpdateBookIDListCache = "lastUpdateBookIdListCache"
	UserInfoCache             = "userInfoCache"
	AuthorInfoCache           = "authorInfoCache"
)

type CacheConfig struct {
	Name    string
	TTL     int
	MaxSize int
}

// 创建了一个新的缓存实例

func NewCacheConfig(name string, ttl, maxSize int) *CacheConfig {
	return &CacheConfig{
		Name:    CachePrefix + name,
		TTL:     ttl,
		MaxSize: maxSize,
	}
}
