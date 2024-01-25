package request

import (
	"hamburgueria/internal/modules/payment/usecase/command"
	"time"
)

type QrCodePaymentRequest struct {
	ExternalReference string `json:"external_reference"`
	NotificationUrl   string `json:"notification_url"`
	TotalAmount       int    `json:"total_amount"`
	ExpirationDate    string `json:"expiration_date"`
}

func MapToMPQrCodePaymentRequest(command command.CreatePaymentCommand, callBackURl string) QrCodePaymentRequest {
	return QrCodePaymentRequest{
		ExternalReference: command.OrderId.String(),
		NotificationUrl:   callBackURl,
		TotalAmount:       command.Amount,
		ExpirationDate:    setExperionDateTime(),
	}

}

func setExperionDateTime() string {
	// Obtém a hora atual
	currentTime := time.Now()
	// Adiciona 5 minutos ao horário atual
	newTime := currentTime.Add(5 * time.Minute)
	// Define o layout desejado para a formatação
	layout := "2006-01-02T15:04:05.999-04:00"
	// Formata a nova data conforme o layout especificado
	formattedTime := newTime.Format(layout)
	// Imprime a nova data formatada
	return formattedTime
}
