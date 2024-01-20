package mercadopago

import (
	"bytes"
	"context"
	"encoding/json"
	"hamburgueria/internal/modules/payment/infra/client/mercadopago/request"
	"hamburgueria/internal/modules/payment/infra/client/mercadopago/response"
	"hamburgueria/internal/modules/payment/usecase/command"
	"hamburgueria/internal/modules/payment/usecase/result"
	"net/http"
)

type mercadoPagoClient struct {
	client http.Client
}

func (mpc mercadoPagoClient) CreatePayment(ctx context.Context, command command.CreatePaymentCommand) (*result.PaymentProcessed, error) {
	return mpc.post(ctx, command)
}

func (mpc mercadoPagoClient) post(ctx context.Context, command command.CreatePaymentCommand) (*result.PaymentProcessed, error) {
	// url mp https://api.mercadopago.com/instore/orders/qr/seller/collectors/{user_id}/pos/{external_pos_id}/qrs
	qr := request.MapToMPQrCodePaymentRequest(command)
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(qr)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "URL", &buf)
	req.Header.Set("Authorization", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := mpc.client.Do(req)
	if err != nil {
		return nil, err
	}
	mp_qr_code_response := response.QrCodePaymentResponse{}

	err = json.NewDecoder(resp.Body).Decode(&mp_qr_code_response)
	if err != nil {
		return nil, err
	}
	return nil, nil

}
