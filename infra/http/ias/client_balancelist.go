package ias

type BalanceListRequest struct {
	ParList []BalanceInfo `json:"parList,omitempty"`
}

func (i BalanceListRequest) GetQueryString() string {
	//TODO implement me
	return ""
}

func (i BalanceListRequest) GetPathVariable() string {
	//TODO implement me
	return ""
}

func (i BalanceListRequest) IsSkippable() bool {
	//TODO implement me
	return i.ParList == nil || len(i.ParList) == 0
}

func (i *BalanceListRequest) SetParList(parList []string) {
	for _, par := range parList {
		balanceInfo := BalanceInfo{Par: par}
		i.ParList = append(i.ParList, balanceInfo)
	}
}

type BalanceListResponse struct {
	BalanceList []BalanceInfo `json:"balanceList"`
}

func NewIasBalanceListResponse(balanceList []BalanceInfo) *BalanceListResponse {
	return &BalanceListResponse{BalanceList: balanceList}
}

func (i BalanceListResponse) IsFollowingSkippable() bool {
	//TODO implement me
	return false
}
