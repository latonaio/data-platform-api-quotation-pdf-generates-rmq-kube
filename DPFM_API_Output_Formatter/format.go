package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-quotation-pdf-generates-rmq-kube/DPFM_API_Input_Formatter"
)

func SetToPdfData(
	headerData *dpfm_api_input_reader.Header,
	inputItems []dpfm_api_input_reader.Items,
) *Header {
	pm := &Header{}

	var items []Items
	for _, dataItems := range inputItems {
		if headerData.Quotation == dataItems.Quotation {
			items = append(items, Items{
				Quotation:							dataItems.Quotation,
				QuotationItem:						dataItems.QuotationItem,
				Product:							dataItems.Product,
				SizeOrDimensionText:				dataItems.SizeOrDimensionText,
				QuotationItemText:        			dataItems.QuotationItemText,
				QuotationQuantityInBaseUnit:		dataItems.QuotationQuantityInBaseUnit,
				QuotationQuantityInDeliveryUnit:	dataItems.QuotationQuantityInDeliveryUnit,
				BaseUnit:							dataItems.BaseUnit,
				DeliveryUnit:        				dataItems.DeliveryUnit,
				RequestedDeliveryDate:        		dataItems.RequestedDeliveryDate,
				NetAmount:							dataItems.NetAmount,
				TaxAmount:							dataItems.TaxAmount,
				GrossAmount:						dataItems.GrossAmount,
			})
		}
	}

	pm = &Header{
				Quotation:   					headerData.Quotation,
				QuotationStatus:   				headerData.QuotationStatus,
				QuotationDate:   				headerData.QuotationDate,
				QuotationType:   				headerData.QuotationType,
				Buyer: 							headerData.Buyer,
				BuyerName: 						headerData.BuyerName,
				Seller: 						headerData.Seller,
				SellerName:              		headerData.SellerName,
				RequestedDeliveryDate: 			headerData.RequestedDeliveryDate,
				TotalGrossAmount: 				headerData.TotalGrossAmount,
				Contract: 						headerData.Contract,
				ContractItem: 					headerData.ContractItem,
				Incoterms: 						headerData.Incoterms,
				PricingDate: 					headerData.PricingDate,
				PaymentTerms: 					headerData.PaymentTerms,
				PaymentMethod: 					headerData.PaymentMethod,
				TransactionCurrency: 			headerData.TransactionCurrency,
				Items:     						items,
	}

	return pm
}
