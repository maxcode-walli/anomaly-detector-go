package main

type Transaction struct {
	UserID                            string       `json:"userID,omitempty"`
	ExternalAccountID                 string       `json:"externalAccountID,omitempty"`
	TransactionStatus                 string       `json:"transactionStatus,omitempty" firestore:"transactionStatus,omitempty"`
	TransactionId                     string       `json:"transactionId,omitempty" firestore:"transactionId,omitempty"`
	Uuid                              string       `json:"uuid,omitempty" firestore:"uuid,omitempty"`
	EndToEndId                        string       `json:"endToEndId,omitempty" firestore:"endToEndId,omitempty"`
	TransactionAmount                 *Money       `json:"transactionAmount,omitempty" firestore:"transactionAmount,omitempty"`
	BookingDate                       string       `json:"bookingDate,omitempty" firestore:"bookingDate,omitempty"`
	ValueDate                         string       `json:"valueDate,omitempty" firestore:"valueDate,omitempty"`
	BalanceAfterTransaction           string       `json:"balanceAfterTransaction,omitempty" firestore:"balanceAfterTransaction,omitempty"`
	RemittanceInformationUnstructured string       `json:"remittanceInformationUnstructured,omitempty" firestore:"remittanceInformationUnstructured,omitempty"`
	RemittanceInformationStructured   string       `json:"remittanceInformationStructured,omitempty" firestore:"remittanceInformationStructured,omitempty"`
	CreditorAccount                   *CounterPart `json:"creditorAccount,omitempty" firestore:"creditorAccount,omitempty"`
	DebtorAccount                     *CounterPart `json:"debtorAccount,omitempty" firestore:"debtorAccount,omitempty"`
}

type Money struct {
	Amount   int64  `json:"amount" firestore:"amount"`
	Currency string `json:"currency" firestore:"currency"`
}

type CounterPart struct {
	Name     string                 `json:"name,omitempty"`
	IBAN     string                 `json:"iban,omitempty"`
	MetaInfo map[string]interface{} `json:"metaInfo,omitempty"`
}

type AccountUser struct {
	Email         string `firestore:"email,omitempty"`
	DisplayName   string `firestore:"displayName,omitempty"`
	PhotoUrl      string `firestore:"photoURL,omitempty"`
	Uid           string `firestore:"uid,omitempty"`
	EmailVerified bool   `firestore:"emailVerified"`
	Job           string `firestore:"job"`
}
