package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/bytes"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/rpc/client/local"
	"github.com/cometbft/cometbft/types"
	"google.golang.org/protobuf/proto"
)

func SendTx[T proto.Message](rpc *local.Local, msg T) error {
	ctx := context.Background()

	tx, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	result, err := rpc.BroadcastTxSync(ctx, tx)
	if err != nil {
		return err
	}

	if result.Code != abcitypes.CodeTypeOK {
		return errors.New(result.Log)
	}

	txChan, err := rpc.Subscribe(ctx, "tx-subscriber", fmt.Sprintf("tm.event = 'Tx' AND tx.hash = '%X'", result.Hash))
	if err != nil {
		return err
	}

	fmt.Println("txhash", result.Hash.String())

	defer func() {
		if err := rpc.Unsubscribe(ctx, "tx-subscriber", fmt.Sprintf("tm.event = 'Tx' AND tx.hash = '%X'", result.Hash)); err != nil {
			// Handle the unsubscribe error if necessary
			fmt.Println("Failed to unsubscribe:", err)
		}
	}()

	select {
	case txRes := <-txChan:
		etx := txRes.Data.(types.EventDataTx)
		if etx.TxResult.Result.Code != abcitypes.CodeTypeOK {
			return fmt.Errorf("tx %s failed to index", result.Hash)
		}
		return nil
	case <-time.After(30 * time.Second):
		return errors.New("tx waiting timeout")
	}
}

func ToTxHash(msg proto.Message) (string, error) {
	b, err := proto.Marshal(msg)
	if err != nil {
		return "", err
	}

	tx := types.Tx(b)
	hash := tx.Hash()
	hexBytes := bytes.HexBytes(hash)
	hashStr := hexBytes.String()
	
	return hashStr, nil
}

type NoopLogger struct{}

func (l NoopLogger) Debug(msg string, keyvals ...any) {}
func (l NoopLogger) Info(msg string, keyvals ...any)  {}
func (l NoopLogger) Error(msg string, keyvals ...any) {}

func (l NoopLogger) With(keyvals ...any) log.Logger {
	return l
}
