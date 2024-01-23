package mercadopago

import (
	"context"
	"github.com/rs/zerolog"
	"hamburgueria/config"
	"hamburgueria/internal/modules/payment/domain"
	"hamburgueria/internal/modules/payment/infra/client/mercadopago/request"
	"hamburgueria/internal/modules/payment/infra/client/mercadopago/response"
	"hamburgueria/internal/modules/payment/port/output"
	"hamburgueria/internal/modules/payment/usecase/command"
	"hamburgueria/pkg/httpclient"
	"sync"
)

type ClientGateway struct {
	userId        string
	externalPosId string
	bearer        string
	client        httpclient.Client
	logger        zerolog.Logger
}

func (mpc ClientGateway) CreatePayment(ctx context.Context, command command.CreatePaymentCommand) (*domain.Payment, error) {
	return mpc.post(ctx, command)
}

func (mpc ClientGateway) post(ctx context.Context, command command.CreatePaymentCommand) (*domain.Payment, error) {
	// url mp https://api.mercadopago.com/instore/orders/qr/seller/collectors/{user_id}/pos/{external_pos_id}/qrs

	httpRequest := httpclient.NewRequest[response.QrCodePaymentResponse](
		ctx, mpc.client, "instore/orders/qr/seller/collectors/{user_id}/pos/{external_pos_id}/qrs",
	).
		WithPathParams(map[string]string{
			"user_id":         mpc.userId,
			"external_pos_id": mpc.externalPosId,
		}).
		WithHeaders(map[string]string{
			"Authorization": mpc.bearer,
		})

	responseMP, err := httpRequest.Post(request.MapToMPQrCodePaymentRequest(command))

	if err != nil {
		return nil, err
	}

	paymentEntity := responseMP.Result.MpQrCodeResponseToPaymentEntity()

	if err != nil {
		return nil, err
	}

	return paymentEntity, nil

}

var (
	mercadoPagoClient     ClientGateway
	mercadoPagoClientOnce sync.Once
)

func GetCreateMercadoPagoClient(client httpclient.Client, mercadoPagoConfig config.MercadoPago, logger zerolog.Logger) output.PaymentClient {
	mercadoPagoClientOnce.Do(func() {
		mercadoPagoClient = ClientGateway{
			userId:        mercadoPagoConfig.UserId,
			externalPosId: mercadoPagoConfig.ExternalPosId,
			bearer:        mercadoPagoConfig.Bearer,
			client:        client,
			logger:        logger,
		}

	})
	return mercadoPagoClient
}
