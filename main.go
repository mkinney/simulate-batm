package main

import (
  "net/http"
  "log"
  "fmt"
)

func calculate_crypto_amount(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(`{"BTC":{"cryptoAmount":0.01017675,"cryptoDiscount":0E-8,"discountCode":null,"discountQuotient":null,"feeDiscount":null,"fiatDiscount":0.00,"fixedTransactionFee":5.0000000000}}`))
}

func sell_crypto(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(`{"cashAmount":100.00,"cashCurrency":"USD","cryptoAddress":"1KnLwMzVv3zwj1oqe9F4HHAjkGZeisKRCs","cryptoAmount":0.011258,"cryptoCurrency":"BTC","customData":{},"fixedTransactionFee":5.0000000000,"localTransactionId":"EATZUZ","remoteTransactionId":"RIMMVU","status":0,"transactionUUID":"UD78U4YKKS7BSCBA","validityInMinutes":30}`))
}


const PORT = "8000"

func main() {

  mux := http.NewServeMux()
  mux.HandleFunc("/calculate_crypto_amount", calculate_crypto_amount)
  mux.HandleFunc("/sell_crypto", sell_crypto)

  fmt.Printf("Open http://localhost:%s\n", PORT)
  err := http.ListenAndServe(":" + PORT, mux)
  log.Fatal(err)
}
