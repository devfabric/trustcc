package controller

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/liuhangyu/trustcc/common"
	"github.com/liuhangyu/trustcc/models"
)

func EntryTabLogic(ctx contractapi.TransactionContextInterface, args []string) error {
	switch args[0] {
	case "create": //新建表
		if len(args) >= 2 {
			return CreateTab(ctx, args[1])
		}
		return common.ErrArgsNum.FormatErrMsg("create", fmt.Errorf("input len(agrs) %d, expect input 2 param", len(args))).ToError()
	case "changecol": //更新列名
	case "freeinds": //删除索引
	case "newinds": //新建索引
	case "drop": //删除表
		return nil
	}
	return common.ErrFoundFun.ToError()
}

func CreateTab(ctx contractapi.TransactionContextInterface, args string) error {
	//变量定义
	var (
		metaTabs = make([]models.MetaTab, 0)
		funcName = "CreateTab"
	)

	//参数反序列化
	err := json.Unmarshal([]byte(args), &metaTabs)
	if err != nil {
		return common.ErrJsonUnmarshal.FormatErrMsg(funcName, err).ToError()
	}

	for idx := range metaTabs {
		//数据校验
		Bvalbytes, err := ctx.GetStub().GetState(metaTabs[idx].TabName)
		if err != nil {
			return common.ErrGetState.FormatErrMsg(funcName, err).ToError()
		}
		if Bvalbytes != nil {
			return common.ErrStateKeyExist.FormatErrMsg(funcName, nil).ToError()
		}

		//数据处理－保存
		err = ctx.GetStub().PutState(metaTabs[idx].TabName, []byte(args))
		if err != nil {
			return common.ErrPutState.FormatErrMsg(funcName, err).ToError()
		}
	}

	//返回事件
	err = ctx.GetStub().SetEvent(models.EvCreateTab, []byte(args))
	if err != nil {
		return common.ErrSetEvent.FormatErrMsg(funcName, err).ToError()
	}

	//返回值
	return nil
}
