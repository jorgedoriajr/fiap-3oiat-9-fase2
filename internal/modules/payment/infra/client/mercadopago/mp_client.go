package mercadopago

import (
	"bytes"
	"context"
	"encoding/json"
	"hamburgueria/internal/modules/payment/domain"
	"hamburgueria/internal/modules/payment/infra/client/mercadopago/request"
	"hamburgueria/internal/modules/payment/infra/client/mercadopago/response"
	"hamburgueria/internal/modules/payment/port/output"
	"hamburgueria/internal/modules/payment/usecase/command"
	"net/http"
	"sync"
)

type MercadoPagoClient struct {
	host   string
	bearer string
	client http.Client
}

func (mpc MercadoPagoClient) CreatePayment(ctx context.Context, command command.CreatePaymentCommand) (*domain.Payment, error) {
	return mpc.post(ctx, command)
}

func (mpc MercadoPagoClient) post(ctx context.Context, command command.CreatePaymentCommand) (*domain.Payment, error) {
	// url mp https://api.mercadopago.com/instore/orders/qr/seller/collectors/{user_id}/pos/{external_pos_id}/qrs
	qr := request.MapToMPQrCodePaymentRequest(command)
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(qr)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "URL", &buf)
	req.Header.Set("Authorization", mpc.bearer)
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

	paymentEntity := mp_qr_code_response.MpQrCodeResponseToPaymentEntity()

	if err != nil {
		return nil, err
	}

	return paymentEntity, nil

}

var (
	mercadoPagoClient     MercadoPagoClient
	mercadoPagoClientOnce sync.Once
)

func GetCreateMercadoPagoClient(host string, bearer string, client http.Client) output.PaymentClient {
	mercadoPagoClientOnce.Do(func() {
		mercadoPagoClient = MercadoPagoClient{host: host, bearer: bearer, client: client}

	})
	return mercadoPagoClient
}
