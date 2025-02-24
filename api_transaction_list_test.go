package trongrid_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"

	"github.com/eliohn/go-trongrid"
)

func TestApi_ListTransactions(t *testing.T) {
	t.Parallel()

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	api := trongrid.NewAPI(
		trongrid.WithURI("https://api.shasta.trongrid.io"),
		trongrid.WithDebug(),
		trongrid.WithLogger(&logger),
		trongrid.WithToken("622ec85e-7406-431d-9caf-0a19501469a4"),
	)

	ctx := context.Background()
	now := time.Now()

	modelListTransactionsRequest, err := api.ListTransactions(ctx, &trongrid.ListTransactionsRequest{
		MaxTimestamp:  now.Add(-(time.Hour * 24)),
		MinTimestamp:  now,
		Address:       "TDuzLK9vBRuSdhLovyy5gCD2bGp4fjecHk",
		Fingerprint:   "",
		OrderBy:       "block_timestamp,desc",
		Limit:         200,
		OnlyConfirmed: true,
		OnlyFrom:      false,
		OnlyTo:        false,
	})
	require.NoError(t, err)
	require.NotNil(t, modelListTransactionsRequest)
}
func TestApi_ListTransactionsTrc20(t *testing.T) {
	t.Parallel()

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	api := trongrid.NewAPI(
		trongrid.WithURI("https://api.shasta.trongrid.io"),
		trongrid.WithDebug(),
		trongrid.WithLogger(&logger),
		trongrid.WithToken("622ec85e-7406-431d-9caf-0a19501469a4"),
	)

	ctx := context.Background()
	now := time.Now()
	// TDkHqdvt6ZRnBCbhj3ytYdWTgJkE6LHNfH

	modelListTransactionsRequest, err := api.ListTransactionsTrc20(ctx, &trongrid.ListTransactionsRequest{
		MaxTimestamp:  now.Add(-(time.Hour * 24)),
		MinTimestamp:  now,
		Address:       "TDuzLK9vBRuSdhLovyy5gCD2bGp4fjecHk",
		Fingerprint:   "",
		OrderBy:       "block_timestamp,desc",
		Limit:         200,
		OnlyConfirmed: true,
		OnlyFrom:      false,
		OnlyTo:        false,
	})
	require.NoError(t, err)
	require.NotNil(t, modelListTransactionsRequest)
}

func TestApi_ListTransactionsTrc20Main(t *testing.T) {
	t.Parallel()

	//logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	api := trongrid.NewAPI(
		trongrid.WithURI("https://api.trongrid.io"),
		//trongrid.WithDebug(),
		//trongrid.WithLogger(&logger),
		trongrid.WithToken("622ec85e-7406-431d-9caf-0a19501469a4"),
	)

	ctx := context.Background()
	now := time.Now()
	// TDkHqdvt6ZRnBCbhj3ytYdWTgJkE6LHNfH

	modelListTransactionsRequest, err := api.ListTransactionsTrc20(ctx, &trongrid.ListTransactionsRequest{
		MaxTimestamp:  now.Add(-(time.Hour * 1)),
		MinTimestamp:  now,
		Address:       "TDkHqdvt6ZRnBCbhj3ytYdWTgJkE6LHNfH",
		Fingerprint:   "",
		OrderBy:       "block_timestamp,desc",
		Limit:         10,
		OnlyConfirmed: true,
		OnlyFrom:      false,
		OnlyTo:        true,
	})
	require.NoError(t, err)
	//require.NotNil(t, modelListTransactionsRequest)
	for i := 0; i < len(modelListTransactionsRequest.Data); i++ {
		item := modelListTransactionsRequest.Data[i]
		t.Logf("交易ID: %v", item)
		formattedTime := time.Unix(item.BlockTimestamp/1000, 0).Format("2006-01-02 15:04:05")
		t.Logf("交易时间: %s %v", formattedTime, trongrid.ParseValue(item.Value, item.TokenInfo.Decimals))

	}
}
