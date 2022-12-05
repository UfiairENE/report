var StellarSdk = require('stellar-sdk')
const axios = require('axios');
let fs = require("fs")

var server = new StellarSdk.Server('https://expansion.bantu.network');
StellarSdk.Networks.PUBLIC = "Public Bantu Network"

let networkURL = "https://expansion.bantu.network/"
let csvHeader = "Cursor,Source,Destination,Amount,Date,Hash,TransactionUrl";
let fileName = "Payment-Txn.csv"




 const GeneratePaymentTxn = async (startCursor, endCursor) => {

   // let nextCursor = startCursor

    let  endCursorInt = Number(endCursor)
     let nextCursorInt = Number(startCursor)

     fs.writeFileSync(fileName, csvHeader);
 
     for (;nextCursorInt < endCursorInt;){
 
       const  ops = await GenerateNextBatch(startCursor , endCursor)

         if(ops.length > 0){
       //  console.log("csvrow: " + ops.length)
         let nextCursor =  ProcessBatch(ops)
          //   console.log("Next cursor: " + nextCursor)
             nextCursorInt = Number(startCursor)
             startCursor  = nextCursor
         }else{
            break
         }
 
     }
  }

const GenerateNextBatch = async (startCursor, endCursor) => {
    console.log("Start Cursor is: " + startCursor + " End Cursor " + endCursor)
    let newRecord =  []

    let result

   // await axios.get('https://expansion.bantu.network/payments?cursor='+startCursor+'&limit=200'

    await axios.get( networkURL + 'payments?cursor='+startCursor+'&limit=200').then(resp => {

   // console.log(resp.data);

    let ops = resp.data._embedded.records
    result = resp.data._embedded.records

   console.log(ops.length)

    if(ops.length > 0){
        let  lastOpsCursor = ops[ops.length - 1]

        let c = lastOpsCursor

       let endCursorInt = Number(endCursor)
        let lastOpsCursorInt =  Number(c.paging_token)

        if (Number(lastOpsCursorInt) < Number(endCursorInt)) {
            ops.forEach(function(record) { 
                //console.log(record)
                if(record.type === "payment"){
                    let	currentRecordCursorInt = Number(record.paging_token)
                    if (currentRecordCursorInt < endCursorInt) {
                        let newRow = {}
                        newRecord.push(record)
                    }
                }
                
              });
        }
        return
        }

}).catch(function (err) {
    console.log(err);
    return []
})

    return newRecord
  }

const ProcessBatch =  (records) => {
    let endCursor = 0

    records.forEach(function(record) {  
        endCursor = record.paging_token
      let newRow = `\r\n${record.paging_token},${record.from},${record.to},${record.amount},${record.created_at},${record.transaction_hash},${record._links.transaction.href}`
        WriteToCSV(newRow)
      //  console.log("Process Batch......... " + record.paging_token +  " " +  record.created_at +  " " + record.source_account
      });

    return endCursor
  }

  const WriteToCSV =  (newRow) => {

   
    fs.appendFileSync(fileName, newRow);

  }
  
GeneratePaymentTxn("33369031146737665","33465247004106753")
