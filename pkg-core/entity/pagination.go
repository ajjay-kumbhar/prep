package entity

type Pagination struct {
	Total    int `json:"total"`
	PageSize int `json:"page_size"`
	PageNum  int `json:"page_num"`
}

func (p *Pagination) Validate() {

	if p.PageNum == 0 {
		p.PageNum = 1
	}

	if p.PageSize == 0 {
		p.PageSize = 2
	}

}
