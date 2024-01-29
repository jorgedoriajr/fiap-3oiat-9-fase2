package request

import (
	"hamburgueria/internal/modules/payment/usecase/command"
	"time"
)

type QrCodePaymentRequest struct {
	Title             string `json:"title"`
	Description       string `json:"description"`
	ExternalReference string `json:"external_reference"`
	NotificationUrl   string `json:"notification_url"`
	TotalAmount       int    `json:"total_amount"`
	ExpirationDate    string `json:"expiration_date"`
	Items             []Item `json:"items"`
}

type Item struct {
	Title       string `json:"title"`
	UnitPrice   int    `json:"unit_price"`
	Quantity    int    `json:"quantity"`
	UnitMeasure string `json:"unit_measure"`
	TotalAmount int    `json:"total_amount"`
}

func MapToMPQrCodePaymentRequest(command command.CreatePaymentCommand, callBackURl string) QrCodePaymentRequest {
	var items []Item
	for _, orderItem := range command.OrderItems {
		items = append(items, Item{
			Title:       orderItem.Name,
			UnitPrice:   orderItem.Amount,
			Quantity:    orderItem.Quantity,
			UnitMeasure: "unit",
			TotalAmount: orderItem.TotalAmount,
		})
	}
	return QrCodePaymentRequest{
		Title:             "Pedido Hamburgueria",
		Description:       "Pedido Hamburgueria",
		ExternalReference: command.OrderId.String(),
		NotificationUrl:   callBackURl,
		TotalAmount:       command.Amount,
		ExpirationDate:    setExperionDateTime(),
		Items:             items,
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
