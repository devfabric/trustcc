package router

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/liuhangyu/trustcc/common"
	"github.com/liuhangyu/trustcc/controller"
)

type TrustCC struct {
	contractapi.Contract
}

func (t *TrustCC) Init(ctx contractapi.TransactionContextInterface) error {
	putErr := ctx.GetStub().PutState("ping", []byte("pong"))
	if putErr != nil {
		return common.ErrPutState.FormatErrMsg("Init", putErr).ToError()
	}
	return nil
}

func (t *TrustCC) Invoke(ctx contractapi.TransactionContextInterface) error {
	fn, args := ctx.GetStub().GetFunctionAndParameters()
	fmt.Printf("in route Invoke, module:%v, function:%v, len(args):%d\n", fn, args, len(args))

	switch fn {
	case "tab":
		return controller.EntryTabLogic(ctx, args)
	case "data":
		return controller.EntryDataLogic(ctx, args)
	default:
		return common.ErrFoundMod.ToError()
	}
}

// func (t *TrustCC) Delete(ctx contractapi.TransactionContextInterface) error {
// 	fn, args := ctx.GetStub().GetFunctionAndParameters()
// 	fmt.Println(fn, "args:", args)
// 	switch fn {
// 	default:
// 		return common.ErrFoundMod.ToError()
// 	}
// }

// func (t *TrustCC) Query(ctx contractapi.TransactionContextInterface) (string, error) {
// 	fn, args := ctx.GetStub().GetFunctionAndParameters()
// 	fmt.Println(fn, "args:", args)
// 	switch fn {

// 	default:
// 		return "", common.ErrFoundMod.ToError()
// 	}
// }

func Start() {
	cc, err := contractapi.NewChaincode(new(TrustCC))
	if err != nil {
		fmt.Printf("Error new chaincode: %s", err)
		return
	}
	if err := cc.Start(); err != nil {
		fmt.Printf("Error starting ABstore chaincode: %s", err)
	}
}
