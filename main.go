package main

import (
	"github.com/nduson/txn-report/controllers"

	errorMsg "github.com/nduson/txn-report/errors"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {

		errorMsg.LogError(err)
		//log.Fatal("Error loading .env file")
		//log.Fatal().Msg("Error loading .env file")
	}

	controllers.GeneratePaymentTxn("30593125063856129", "30735730862989313")

	// Create a single reader which can be called multiple times
	//reader := bufio.NewReader(os.Stdin)
	// Prompt and read
	//fmt.Print("Enter Starting Cursor: ")
	//text, _ := reader.ReadString('\n')

	// validate Start Cusors
	// GetPaymentHistroy()
	//fmt.Print("Enter Ending Cursor: ")
	//text2, _ := reader.ReadString('\n')

	// validate End Cusors
	// GetPaymentHistroy()

	// Check if starting cursor is less then ending cursor

	// process and generate payment history in cv

	/// GeneratePaymentHistoryCSV()

	// Trim whitespace and print
	//fmt.Printf("Text1: \"%s\", Text2: \"%s\"\n",
	//	strings.TrimSpace(text), strings.TrimSpace(text2))
}
