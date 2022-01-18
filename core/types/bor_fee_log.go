package types

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

var transferLogSig = common.HexToHash("0xe6497e3ee548a3372136af2fcb0696db31fc6cf20260707645068bd3fe97f3c4")
var transferFeeLogSig = common.HexToHash("0x4dfe1bbbcf077ddc3e01291eea2d5c70c2b422b415d95645b9adcfd678cb1d63")
var feeAddress = common.HexToAddress("0x0000000000000000000000000000000000001010")

func ReceiptsRecover(receipts Receipts) Receipts {
	var result Receipts
	for _, receipt := range receipts {
		var receiptNew = &Receipt{}
		*receiptNew = *receipt
		var logs []*Log
		for _, log := range receipt.Logs {
			if strings.EqualFold(log.Address.String(), feeAddress.String()) && len(log.Topics) > 0 {
				if strings.EqualFold(log.Topics[0].String(), transferLogSig.String()) || strings.EqualFold(log.Topics[0].String(), transferFeeLogSig.String()) {
					continue
				}
			}
			logs = append(logs, log)
		}
		receiptNew.Logs = logs
		result = append(result, receiptNew)
	}
	return result
}
