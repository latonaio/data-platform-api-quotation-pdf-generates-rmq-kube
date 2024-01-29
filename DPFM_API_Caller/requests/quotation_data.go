package requests

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
