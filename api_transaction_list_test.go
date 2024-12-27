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
