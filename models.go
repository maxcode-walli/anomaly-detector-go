package p

type Transaction struct {
	UserID                            string       `json:"userID,omitempty"`
	ExternalAccountID                 string       `json:"externalAccountID,omitempty"`
	TransactionStatus                 string       `json:"transactionStatus,omitempty"`
	TransactionId                     string       `json:"transactionId,omitempty"`
	Uuid                              string       `json:"uuid,omitempty"`
	EndToEndId                        string       `json:"endToEndId,omitempty"`
	TransactionAmount                 *Money       `json:"transactionAmount,omitempty"`
	BookingDate                       string       `json:"bookingDate,omitempty"`
	ValueDate                         string       `json:"valueDate,omitempty"`
	BalanceAfterTransaction           string       `json:"balanceAfterTransaction,omitempty"`
	RemittanceInformationUnstructured string       `json:"remittanceInformationUnstructured,omitempty"`
	RemittanceInformationStructured   string       `json:"remittanceInformationStructured,omitempty"`
	CreditorAccount                   *CounterPart `json:"creditorAccount,omitempty"`
	DebtorAccount                     *CounterPart `json:"debtorAccount,omitempty"`
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

type User struct {
	Email         string `json:"email,omitempty"`
	DisplayName   string `json:"displayName,omitempty"`
	PhotoUrl      string `json:"photoUrl,omitempty"`
	Uid           string `json:"uid,omitempty"`
	EmailVerified string `json:"emailVerified,omitempty"`
	Job           string `json:"job,omitempty"`
}
