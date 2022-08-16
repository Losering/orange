### 1. "首页Banner"

1. route definition

- Url: /v1/home/banner
- Method: GET
- Request: `-`
- Response: `HomeBannerResponse`

2. request definition



3. response definition



```golang
type HomeBannerResponse struct {
	Banners []*Banner `json:"banners"`
}
```

### 2. "限时抢购"

1. route definition

- Url: /v1/flashsale
- Method: GET
- Request: `-`
- Response: `FlashSaleResponse`

2. request definition



3. response definition



```golang
type FlashSaleResponse struct {
	StartTime int64 `json:"start_time"` // 抢购开始时间
	Products []*Product `json:"products"`
}
```

### 3. "推荐商品列表"

1. route definition

- Url: /v1/recommend
- Method: GET
- Request: `RecommendRequest`
- Response: `RecommendResponse`

2. request definition



```golang
type RecommendRequest struct {
	Cursor int64 `json:"cursor"`
	Ps int64 `form:"ps,default=20"` // 每页大小
}
```


3. response definition



```golang
type RecommendResponse struct {
	Products []*Product `json:"products"`
	IsEnd bool `json:"is_end"` // 是否最后一页
	RecommendTime int64 `json:"recommend_time"` // 商品列表最后一个商品的推荐时间
}
```

### 4. "分类商品列表"

1. route definition

- Url: /v1/category/list
- Method: GET
- Request: `CategoryListRequest`
- Response: `CategoryListResponse`

2. request definition



```golang
type CategoryListRequest struct {
	Cursor int64 `form:"cursor"` // 分页游标
	Ps int64 `form:"ps,default=20"` // 每页大小
	Category string `form:"category"` // 分类
	Sort string `form:"sort"` // 排序
}
```


3. response definition



```golang
type CategoryListResponse struct {
	Products []*Product `json:"products"`
	IsEnd bool `json:"is_end"`
	LastVal int64 `json:"last_val"`
}
```

### 5. "购物车列表"

1. route definition

- Url: /v1/cart/list
- Method: GET
- Request: `CartListRequest`
- Response: `CartListResponse`

2. request definition



```golang
type CartListRequest struct {
	UID int64 `form:"uid"`
}
```


3. response definition



```golang
type CartListResponse struct {
	Products []*CartProduct `json:"products"`
}
```

### 6. "商品评论列表"

1. route definition

- Url: /v1/product/comment
- Method: GET
- Request: `ProductCommentRequest`
- Response: `ProductCommentResponse`

2. request definition



```golang
type ProductCommentRequest struct {
	ProductID int64 `form:"product_id"`
	Cursor int64 `form:"cursor"`
	Ps int64 `form:"ps,default=20"`
}
```


3. response definition



```golang
type ProductCommentResponse struct {
	Comments []*Comment `json:"comments"`
	IsEnd bool `json:"is_end"` // 是否最后一页
	CommentTime int64 `json:"comment_time"` // 评论列表最后一个评论的时间
}
```

### 7. "订单列表"

1. route definition

- Url: /v1/order/list
- Method: GET
- Request: `OrderListRequest`
- Response: `OrderListResponse`

2. request definition



```golang
type OrderListRequest struct {
	UID int64 `form:"uid"`
	Status int32 `form:"status,optional"`
	Cursor int64 `form:"cursor,optional"`
	Ps int64 `form:"ps,default=20"`
}
```


3. response definition



```golang
type OrderListResponse struct {
	Orders []*Order `json:"orders"`
	IsEnd bool `json:"is_end"` // 是否最后一页
	OrderTime int64 `json:"order_time"`
}
```

### 8. "商品详情"

1. route definition

- Url: /v1/product/detail
- Method: GET
- Request: `ProductDetailRequest`
- Response: `ProductDetailResponse`

2. request definition



```golang
type ProductDetailRequest struct {
	ProductID int64 `form:"product_id"`
}
```


3. response definition



```golang
type ProductDetailResponse struct {
	Product *Product `json:"product"`
	Comments []*Comment `json:"comments"`
}
```

