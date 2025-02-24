package trongrid

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/gorilla/schema"
	"github.com/rs/zerolog"
	"golang.org/x/time/rate"
	"time"
)

type API interface {
	// ListTransactions
	// Docs: https://developers.tron.network/reference/get-trc20-transaction-info-by-account-address
	ListTransactionsTrc20(ctx context.Context, req *ListTransactionsRequest) (resp *TRC20Response, err error)
	ListTransactions(ctx context.Context, req *ListTransactionsRequest) (resp *ListTransactionsResponse, err error)
}

type api struct {
	encoder *schema.Encoder
	decoder *schema.Decoder
	logger  *zerolog.Logger
	cl      *resty.Client
	token   string
	uri     string
	debug   bool
}

func NewAPI(opts ...Option) API {
	x := &api{
		encoder: NewEncoder(),
		decoder: NewDecoder(),
		logger:  nil,
		cl:      nil,
		token:   "",
		uri:     "",
		debug:   false,
	}
	for _, opt := range opts {
		opt(x)
	}

	if len(x.uri) == 0 {
		x.uri = URI
	}

	cl := resty.New().
		SetBaseURL(x.uri).
		SetDebug(x.debug).
		SetRateLimiter(rate.NewLimiter(rate.Every(time.Second), 1)).
		SetRedirectPolicy(resty.NoRedirectPolicy()).
		SetTimeout(timeout)
	if x.logger != nil {
		cl.SetLogger(NewLogger(x.logger))
	}

	if len(x.token) != 0 {
		cl.SetHeader("TRON-PRO-API-KEY", x.token)
	}

	// Set default retry settings
	cl.SetRetryCount(3). // Default to 3 retries
				SetRetryWaitTime(1 * time.Second).   // Wait 1 second between retries
				SetRetryMaxWaitTime(5 * time.Second) // Maximum wait time of 5 seconds

	// Configure retry conditions
	cl.AddRetryCondition(func(response *resty.Response, err error) bool {
		return err != nil || response.StatusCode() >= 500
	})

	x.cl = cl

	return x
}
