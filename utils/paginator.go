package utils

type Paginator struct {
	CurrentPage     int64 `json:"currentPage"`     //当前页
	NextPage        int64 `json:"nextPage"`        //下一页
	PrePage         int64 `json:"prePage"`         //上一页
	PageSize        int64 `json:"pageSize"`        //每页数量
	CurrentPageSize int64 `json:"currentPageSize"` //当前页数量
	TotalPage       int64 `json:"totalPage"`       //总页数
	TotalCount      int64 `json:"totalCount"`      //总数量
	FirstPage       bool  `json:"firstPage"`       //为第一页
	LastPage        bool  `json:"lastPage"`        //为最后一页
	// PageList        []int64 `json:"pageList"`        //显示的页
	Max int64
}

func GenPaginator(limit, offset, count int64) Paginator {
	var paginator Paginator
	paginator.TotalCount = count
	paginator.TotalPage = (count + limit - 1) / limit
	paginator.PageSize = limit
	if offset == 0 {
		paginator.FirstPage = true
	} else {
		paginator.FirstPage = false
	}
	if offset == paginator.TotalPage {
		paginator.LastPage = true
	} else {
		paginator.LastPage = false
	}
	if paginator.TotalCount > 0 && paginator.CurrentPage > 0 {
		paginator.Max = paginator.TotalCount / paginator.CurrentPage
	} else {
		paginator.Max = 0
	}
	return paginator

}
