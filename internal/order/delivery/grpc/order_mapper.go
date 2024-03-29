package order

import (
	"github.com/materials-resources/s_prophet/internal/order/domain"
	rpc "github.com/materials-resources/s_prophet/proto/order/v1alpha0"
)

func ToPBGetPickTicketByIdResponse(pickTicket *domain.PickTicket) (*rpc.GetPickTicketByIdResponse, error) {
	var pb = &rpc.GetPickTicketByIdResponse{
		Id:                 pickTicket.ID,
		OrderId:            pickTicket.OrderId,
		OrderPurchaseOrder: pickTicket.OrderPurchaseOrder,
		ShippingAddress: &rpc.Address{
			Name:       pickTicket.ShippingAddress.Name,
			LineOne:    pickTicket.ShippingAddress.LineOne,
			LineTwo:    pickTicket.ShippingAddress.LineTwo,
			City:       pickTicket.ShippingAddress.City,
			State:      pickTicket.ShippingAddress.State,
			PostalCode: pickTicket.ShippingAddress.PostalCode,
			Country:    "",
		},
	}
	return pb, nil
}
func ToPBGetOrderResponse(order *domain.Order) (*rpc.GetOrderResponse, error) {
	var pb = &rpc.GetOrderResponse{Id: order.ID,
		PurchaseOrder: order.PurchaseOrder,
		ShippingAddress: &rpc.Address{
			Name:       order.ShippingAddress.Name,
			LineOne:    order.ShippingAddress.LineOne,
			LineTwo:    order.ShippingAddress.LineTwo,
			City:       order.ShippingAddress.City,
			State:      order.ShippingAddress.State,
			PostalCode: order.ShippingAddress.PostalCode,
		},
		CustomerContact: &rpc.CustomerContact{
			Id:          order.CustomerContact.Id,
			Name:        order.CustomerContact.Name,
			PhoneNumber: order.CustomerContact.PhoneNumber,
		},
	}
	for _, item := range order.Items {
		pb.OrderItems = append(pb.OrderItems, &rpc.GetOrderResponse_OrderItem{
			Id:              item.ID,
			Sn:              item.SN,
			Name:            item.Name,
			QuantityOrdered: item.QuantityOrdered,
			QuantityUnit:    "",
			CostPerUnit:     0,
			TotalPrice:      0,
		},
		)
	}

	return pb, nil
}
