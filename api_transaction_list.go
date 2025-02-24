package trongrid

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
)

type ListTransactionsRequest struct {
	MaxTimestamp  time.Time `url:"max_timestamp,omitempty"`
	MinTimestamp  time.Time `url:"min_timestamp,omitempty"`
	Address       string    `url:"Address,omitempty"`
	Fingerprint   string    `url:"fingerprint,omitempty"`
	OrderBy       string    `url:"order_by,omitempty"`
	Limit         int32     `url:"limit,omitempty"`
	OnlyConfirmed bool      `url:"only_confirmed,omitempty"`
	OnlyFrom      bool      `url:"only_from,omitempty"`
	OnlyTo        bool      `url:"only_to,omitempty"`
}

type ListTransactionsResponse struct {
	Meta    *Meta          `json:"meta"`
	Data    []*Transaction `json:"data"`
	Success bool           `json:"success"`
}

func (api *api) ListTransactions(
	ctx context.Context,
	req *ListTransactionsRequest,
) (resp *ListTransactionsResponse, err error) {
	params := url.Values{}
	if err = api.encoder.Encode(req, params); err != nil {
		api.logger.Error().Err(err).Send()

		return nil, err
	}

	r := api.cl.R().
		ForceContentType("application/json").
		SetContext(ctx).
		SetError(new(Error)).
		SetHeader("TRON-PRO-API-KEY", api.token).
		SetPathParam("address", req.Address).
		SetQueryParamsFromValues(params).
		SetResult(new(ListTransactionsResponse))

	var httpResp *resty.Response
	//https://api.shasta.trongrid.io
	if httpResp, err = r.Get(api.uri + "/v1/accounts/{address}/transactions"); err != nil {
		return nil, err
	}

	if v, ok := httpResp.Error().(*Error); ok {
		err = fmt.Errorf("%w: %s", ErrEmpty, v.Error)
		api.logger.Error().Err(err).Send()

		return nil, err
	}

	if v, ok := httpResp.Result().(*ListTransactionsResponse); ok {
		return v, nil
	}

	return nil, ErrEmpty
}

// 修改返回类型和结果映射
func (api *api) ListTransactionsTrc20(ctx context.Context,
	req *ListTransactionsRequest,
) (resp *TRC20Response, err error) { // 修改返回类型为TRC20Response
	params := url.Values{}
	if err = api.encoder.Encode(req, params); err != nil {
		api.logger.Error().Err(err).Send()
		return nil, err
	}

	r := api.cl.R().
		ForceContentType("application/json").
		SetContext(ctx).
		SetError(new(Error)).
		SetHeader("TRON-PRO-API-KEY", api.token).
		SetPathParam("address", req.Address).
		SetQueryParamsFromValues(params).
		SetResult(new(TRC20Response)) // 修改结果映射

	var httpResp *resty.Response
	if httpResp, err = r.Get(api.uri + "/v1/accounts/{address}/transactions/trc20"); err != nil {
		return nil, err
	}

	if v, ok := httpResp.Error().(*Error); ok {
		err = fmt.Errorf("%w: %s", ErrEmpty, v.Error)
		api.logger.Error().Err(err).Send()
		return nil, err
	}

	if v, ok := httpResp.Result().(*TRC20Response); ok {
		// 添加分页参数处理（如果需要）
		//if req.Fingerprint == "" && v.Meta != nil {
		//	// 可以在这里处理分页指纹
		//}
		return v, nil
	}

	return nil, ErrEmpty
}
