package transactions

import (
	"context"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"suck_steal/internal/constants"
	"time"
)

func SendBundleTx(rawTxs []hexutil.Bytes) (string, error) {
	res, err := constants.Client.SendPuissant(
		context.Background(),
		rawTxs,
		uint64(time.Now().Unix()+60),
		nil,
	)

	return res, err
}
