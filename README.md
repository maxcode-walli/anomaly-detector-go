# anomaly-detector-go

responds to: **pigeon.eventType=walli.TransactionUpdatedEventV1**
message body structure:
{ 
  "userID": "xT1Luq9BHMM9p9UMBNIMbGzcK4e2", 
  "externalAccountID": "{GUID}", 
  "transactionStatus": "pending", 
  "transactionId": "89534", 
  "uuid": "{GUID}", 
  "endToEndId": "13847", 
  "creditorAccount": { 
    "name": "xxxxx xxxx", "iban": "{IBAN}" 
  }, 
  "transactionAmount": { 
    "amount": 11, 
    "currency": "RON" 
  }, 
  "bookingDate": "yyyy-MM-dd", 
  "valueDate": "yyyy-MM-dd", 
  "balanceAfterTransaction": "99" 
}
