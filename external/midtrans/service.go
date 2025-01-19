package midtrans

import (
	"ecommerce/internal/model"
	"fmt"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

func AddTransaction(order *model.Order, product *model.Product) (*snap.Response, error) {
	productId := strconv.FormatInt(int64(order.ProductId), 32)
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  order.Id,
			GrossAmt: int64(order.GrossAmount),
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    productId,
				Name:  product.Name,
				Price: int64(product.Price),
				Qty:   int32(order.Quantity),
			},
		},
	}
	snapReq, err := snap.CreateTransaction(req)
	if err != nil {
		if err.StatusCode == 400 {
			return nil, fmt.Errorf("%s", err.Message)
		} else if err.StatusCode == 401 {
			return nil, fmt.Errorf("%s", err.Message)
		} else if err.StatusCode == 500 {
			return nil, fmt.Errorf("%s", err.Message)
		}
	}
	return snapReq, nil
}
func CheckTransaction(orderId *string) (*coreapi.TransactionStatusResponse, error) {
	response, err := coreapi.CheckTransaction(*orderId)
	if err != nil {
		return nil, err
	}
	return response, nil
}
