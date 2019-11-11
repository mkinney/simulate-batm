Simulate some BATM web services (for development/testing)

1. calculate_crypto_amount

On the server:
  curl -k 'https://localhost:7743/extensions/example/calculate_crypto_amount?serial_number=BT102781&fiat_currency=USD&crypto_currency=BTC&fiat_amount=100'

returns:
  {"BTC":{"cryptoAmount":0.01017675,"cryptoDiscount":0E-8,"discountCode":null,"discountQuotient":null,"feeDiscount":null,"fiatDiscount":0.00,"fixedTransactionFee":5.0000000000}}

locally:
  curl 'http://localhost:8000/calculate_crypto_amount?serial_number=BT102781&fiat_currency=USD&crypto_currency=BTC&fiat_amount=100'


2. sell_crypto

On the server:
  curl -k -H "Content-Type: application/json" 'https://localhost:7743/extensions/example/sell_crypto?serial_number=BT102781&fiat_amount=100.00&fiat_currency=USD&crypto_currency=BTC&crypto_amount=0.01017534'

returns:
  {"cashAmount":100.00,"cashCurrency":"USD","cryptoAddress":"1KnLwMzVv3zwj1oqe9F4HHAjkGZeisKRCs","cryptoAmount":0.011258,"cryptoCurrency":"BTC","customData":{},"fixedTransactionFee":5.0000000000,"localTransactionId":"EATZUZ","remoteTransactionId":"RIMMVU","status":0,"transactionUUID":"UD78U4YKKS7BSCBA","validityInMinutes":30}

locally:
  curl 'http://localhost:8000/calculate_crypto_amount?serial_number=BT102781&fiat_currency=USD&crypto_currency=BTC&fiat_amount=100'
