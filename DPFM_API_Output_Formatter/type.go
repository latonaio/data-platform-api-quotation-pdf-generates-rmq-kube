package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
	QuotationPdf        string      `json:"quotation_pdf"`
}

type Message struct {
	Header     []Header   `json:"Header"`
}

type Header struct {
	Quotation							int			`json:"Quotation"`
	QuotationDate						string		`json:"QuotationDate"`
	QuotationType						string		`json:"QuotationType"`
	QuotationStatus						string		`json:"QuotationStatus"`
	Buyer								int 		`json:"Buyer"`
	BuyerName							string		`json:"BuyerName"`
	Seller								int			`json:"Seller"`
	SellerName							string		`json:"SellerName"`
	BindingPeriodValidityStartDate		*string		`json:"BindingPeriodValidityStartDate"`
	BindingPeriodValidityEndDate		*string		`json:"BindingPeriodValidityEndDate"`
	TotalGrossAmount					float32		`json:"TotalGrossAmount"`
	TransactionCurrency					string		`json:"TransactionCurrency"`
	PricingDate							string		`json:"PricingDate"`
	RequestedDeliveryDate				string		`json:"RequestedDeliveryDate"`
	Incoterms							*string		`json:"Incoterms"`
	PaymentTerms						string		`json:"PaymentTerms"`
	PaymentMethod						string		`json:"PaymentMethod"`
	Contract		                 	*int     	`json:"Contract"`
	ContractItem	                 	*int     	`json:"ContractItem"`
	Items				  				[]Items		`json:"Items"`
}

type Items struct {
	Quotation								int		`json:"Quotation"`
	QuotationItem							int		`json:"QuotationItem"`
	QuotationItemCategory					string	`json:"QuotationItemCategory"`
	QuotationItemText						string	`json:"QuotationItemText"`
	QuotationItemTextByBuyer				string	`json:"QuotationItemTextByBuyer"`
	QuotationItemTextBySeller				string	`json:"QuotationItemTextBySeller"`
	Product									string	`json:"Product"`
	SizeOrDimensionText                     *string `json:"SizeOrDimensionText"`
	BaseUnit								string	`json:"BaseUnit"`
	DeliveryUnit							float32	`json:"DeliveryUnit"`
	RequestedDeliveryDate					string	`json:"RequestedDeliveryDate"`
	QuotationQuantityInBaseUnit				string	`json:"QuotationQuantityInBaseUnit"`
	QuotationQuantityInDeliveryUnit			string	`json:"QuotationQuantityInDeliveryUnit"`
	NetAmount								float32 `json:"NetAmount"`
	TaxAmount								float32 `json:"TaxAmount"`
	GrossAmount								float32 `json:"GrossAmount"`
}
