package controller

import (
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/liuhangyu/trustcc/common"
	"github.com/liuhangyu/trustcc/models"
)

func EntryDataLogic(ctx contractapi.TransactionContextInterface, args []string) error {
	switch args[0] {
	case "insert":
		return InsertData(ctx, args[1])
	case "update":
	case "delete":
	case "query":
	}

	return common.ErrFoundFun.ToError()
}

func InsertData(ctx contractapi.TransactionContextInterface, args string) error {
	//变量定义
	var (
		itemKVs  = make([]models.ItemKV, 0)
		funcName = "InsertData"
	)

	//参数反序列化
	err := json.Unmarshal([]byte(args), &itemKVs)
	if err != nil {
		return common.ErrJsonUnmarshal.FormatErrMsg(funcName, err).ToError()
	}

	for idx := range itemKVs {
		//数据处理－保存
		colsKey, err := ctx.GetStub().CreateCompositeKey(models.TAB_COLID_FSEQ,
			[]string{itemKVs[idx].TabName, ctx.GetStub().GetTxID(), itemKVs[idx].ColSeq})
		if err != nil {
			return common.ErrNewCompKey.FormatErrMsg(funcName, err).ToError()
		}

		valBys, err := json.Marshal(itemKVs[idx].Val)
		if err != nil {
			return common.ErrJsonMarshal.FormatErrMsg(funcName, err).ToError()
		}

		err = ctx.GetStub().PutState(colsKey, valBys)
		if err != nil {
			return common.ErrPutState.FormatErrMsg(funcName, err).ToError()
		}
	}

	//返回事件
	err = ctx.GetStub().SetEvent(models.EvInsertDate, []byte(args))
	if err != nil {
		return common.ErrSetEvent.FormatErrMsg(funcName, err).ToError()
	}

	//返回值
	return nil
}
