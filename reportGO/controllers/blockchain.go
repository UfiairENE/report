package controllers

import (
	"encoding/csv"
	"fmt"
	"strconv"

	errorLog "github.com/nduson/txn-report/errors"
	helpers "github.com/nduson/txn-report/helpers"

	//"log"
	"net/http"
	"os"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/protocols/horizon/operations"
)

var configuredHorizonClient *horizonclient.Client

var f *os.File

var csvHead = [][]string{
	{"Cursor", "Source", "Destination", "Amount", "Date", "Hash", "TransactionUrl"},
}

func GetBlockchainClient() *horizonclient.Client {
	if configuredHorizonClient != nil {
		return configuredHorizonClient
	}

	configuredHorizonClient = &horizonclient.Client{
		HorizonURL: os.Getenv("HORIZON_ENDPOINT"),
		HTTP:       http.DefaultClient,
	}

	return configuredHorizonClient
}

func GenerateNextBatch(startCursor , endCursor string) ([]operations.Operation, error) {
	errorLog.LogInfo("Generating next batch from : " + startCursor)

	client := GetBlockchainClient()

	// payments for an account
	opRequest := horizonclient.OperationRequest{Cursor: startCursor, Limit: 200}
	ops, err := client.Payments(opRequest)
	if err != nil {
		return []operations.Operation{}, err
	}
	if len(ops.Embedded.Records) > 0 {
		lastOpsCursor := ops.Embedded.Records[len(ops.Embedded.Records) -1]
		c, ok := lastOpsCursor.(operations.Payment)
			if ok {
		fmt.Println("generate next batch", len(ops.Embedded.Records), startCursor, endCursor, c.PT)
		endCursorInt, _ := strconv.ParseInt(endCursor, 10, 64)
		lastOpsCursorInt, _ := strconv.ParseInt(c.PT, 10, 64)
		if lastOpsCursorInt > endCursorInt  {
			newRecord := []operations.Operation {}
			for _, record := range ops.Embedded.Records {
				c, ok := record.(operations.Payment)
				if ok {
					currentRecordCursorInt, _ := strconv.ParseInt(c.PT, 10, 64)
					if currentRecordCursorInt < endCursorInt {
						newRecord = append(newRecord, record)
					}
				}

			}
			ops.Embedded.Records = newRecord
		}
			} 		

	}
	return ops.Embedded.Records, nil
}

func ProcessBatch(ops []operations.Operation) string {
	records := ops
	var endCursor string
	w := csv.NewWriter(f)
	defer w.Flush()

	for _, value := range records {
		if value.GetType() == "payment" {
			c, ok := value.(operations.Payment)
			endCursor = c.PT

			if ok {
				txnDate := helpers.ConvertDatetime(c.LedgerCloseTime.Format("Jan 02, 2006 3:04:05 PM"))
				csvRow := []string{c.PT, c.SourceAccount, c.To, c.Amount, txnDate, c.GetTransactionHash(), string(c.Links.Transaction.Href)}
				if err := w.Write(csvRow); err != nil {
					errorLog.LogError(err)
				}
			}
		}
	}

	return endCursor
}

func GeneratePaymentTxn(startCursor, endCursor string) {
	openFile, err := os.Create("Payment-Txn.csv")
	if err != nil {
		errorLog.LogError(err)
	}
	f = openFile
	csv.NewWriter(f).WriteAll(csvHead)

	nextCursor := startCursor
	endCursorInt, _ := strconv.ParseInt(endCursor, 10, 64)
	nextCursorInt, _ := strconv.ParseInt(nextCursor, 10, 64)

	for nextCursorInt < endCursorInt {
		ops, err := GenerateNextBatch(nextCursor , endCursor)
		if err != nil {
			errorLog.LogError(err)
		}
		if len(ops) > 0 {
			fmt.Println("csvrow", len(ops))
			nextCursor = ProcessBatch(ops)
			fmt.Println("Next cursor: ", nextCursor)
			nextCursorInt, _ = strconv.ParseInt(nextCursor, 10, 64)
		}else {
			break
		}
		

	}
}
