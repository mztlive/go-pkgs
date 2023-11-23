package pay

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

type Amount struct {
	Total       int `json:"total"`
	Refund      int `json:"refund"`
	PayerTotal  int `json:"payer_total"`
	PayerRefund int `json:"payer_refund"`
}

type RefundTransaction struct {
	MchID               string    `json:"mchid"`
	TransactionID       string    `json:"transaction_id"`
	OutTradeNo          string    `json:"out_trade_no"`
	RefundID            string    `json:"refund_id"`
	OutRefundNo         string    `json:"out_refund_no"`
	RefundStatus        string    `json:"refund_status"`
	SuccessTime         time.Time `json:"success_time"`
	UserReceivedAccount string    `json:"user_received_account"`
	Amount              Amount    `json:"amount"`
}

type NotifyHandlerCallback func(ctx context.Context, notifyRequest *notify.Request, transaction *payments.Transaction) error

type RefundNotifyHandlerCallback func(ctx context.Context, notifyRequest *notify.Request, transaction *RefundTransaction) error

type WechatPayCfg struct {
	AppId    string
	MchId    string
	V3Key    string
	KeyFile  string
	CertFile string
}

// WechatJsPay wechat jsapi pay
func WechatJsPay(ctx context.Context, request *jsapi.PrepayRequest, cfg *WechatPayCfg) (*jsapi.PrepayWithRequestPaymentResponse, error) {
	var (
		prepayResponse *jsapi.PrepayWithRequestPaymentResponse
		prepayResult   *core.APIResult
		client         *core.Client
		err            error
	)

	if client, err = wechatPayClient(ctx, cfg); err != nil {
		return nil, fmt.Errorf("create wechat pay client failed. %w", err)
	}

	svc := jsapi.JsapiApiService{Client: client}

	if prepayResponse, prepayResult, err = svc.PrepayWithRequestPayment(ctx, *request); err != nil {
		return nil, fmt.Errorf("prepay failed. %w", err)
	}

	if prepayResult.Response.StatusCode != http.StatusOK {
		defer prepayResult.Response.Body.Close()
		bodyBytes, _ := io.ReadAll(prepayResult.Response.Body)

		return nil, fmt.Errorf("prepay failed. status not is 200. %s", string(bodyBytes))
	}

	return prepayResponse, nil
}

// Refund wechat pay refund
func Refund(ctx context.Context, request *refunddomestic.CreateRequest, cfg *WechatPayCfg) (*refunddomestic.Refund, error) {
	var (
		client       *core.Client
		apiResult    *core.APIResult
		refundResult *refunddomestic.Refund
		err          error
	)

	if client, err = wechatPayClient(ctx, cfg); err != nil {
		return nil, fmt.Errorf("create wechat pay client failed. %w", err)
	}

	svr := refunddomestic.RefundsApiService{Client: client}
	if refundResult, apiResult, err = svr.Create(ctx, *request); err != nil {
		return nil, fmt.Errorf("refund failed. %w", err)
	}

	if apiResult.Response.StatusCode != http.StatusOK {
		defer apiResult.Response.Body.Close()
		bodyBytes, _ := io.ReadAll(apiResult.Response.Body)

		return nil, fmt.Errorf("refund failed. status not is 200. %s", string(bodyBytes))
	}

	return refundResult, nil
}

// RefundNotifyHandler process on wechat refund notify
// callback receives the parsed notify parameters
//
// if callback func return error. the func also return error
func RefundNotifyHandler(ctx context.Context, request *http.Request, cfg *WechatPayCfg, callback RefundNotifyHandlerCallback) error {
	var (
		notifyData  *notify.Request
		transaction RefundTransaction
		handler     *notify.Handler
		err         error
	)

	if handler, err = createNotifyHandler(ctx, cfg); err != nil {
		return fmt.Errorf("failed to create notify handler. %w", err)
	}

	if notifyData, err = handler.ParseNotifyRequest(ctx, request, &transaction); err != nil {
		return fmt.Errorf("parse notify request failed. %w", err)
	}

	return callback(ctx, notifyData, &transaction)
}

// NotifyHandler process on wechat paid notify
// callback receives the parsed notify parameters
//
// if callback func return error. the func also return error
func NotifyHandler(ctx context.Context, request *http.Request, cfg *WechatPayCfg, callback NotifyHandlerCallback) error {
	var (
		notifyData  *notify.Request
		transaction payments.Transaction
		handler     *notify.Handler
		err         error
	)

	if handler, err = createNotifyHandler(ctx, cfg); err != nil {
		return fmt.Errorf("failed to create notify handler. %w", err)
	}

	if notifyData, err = handler.ParseNotifyRequest(ctx, request, &transaction); err != nil {
		return fmt.Errorf("parse notify request failed. %w", err)
	}

	return callback(ctx, notifyData, &transaction)
}

// returns a wechat notify handler.
// the handler can be used all notify request
func createNotifyHandler(ctx context.Context, cfg *WechatPayCfg) (*notify.Handler, error) {
	var (
		client *core.Client
		err    error
	)

	if client, err = wechatPayClient(ctx, cfg); err != nil {
		return nil, fmt.Errorf("create wechat pay client failed. %w", err)
	}

	err = downloader.MgrInstance().RegisterDownloaderWithClient(ctx, client, cfg.MchId, cfg.V3Key)
	if err != nil {
		return nil, fmt.Errorf("register downloader failed. %w", err)
	}

	certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(cfg.MchId)

	handler := notify.NewNotifyHandler(cfg.V3Key, verifiers.NewSHA256WithRSAVerifier(certificateVisitor))
	return handler, nil
}

// wechatPayClient returns new wechat pay client instance
func wechatPayClient(ctx context.Context, cfg *WechatPayCfg) (*core.Client, error) {
	var (
		privKey      *rsa.PrivateKey
		err          error
		pubKeyBytes  []byte
		serialNumber string
	)

	if pubKeyBytes, err = os.ReadFile(cfg.CertFile); err != nil {
		return nil, fmt.Errorf("read public key file failed. %w", err)
	}

	serialNumber = getSerialNumber(pubKeyBytes)

	if privKey, err = utils.LoadPrivateKeyWithPath(cfg.KeyFile); err != nil {
		return nil, fmt.Errorf("load private key failed. %w", err)
	}

	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(cfg.MchId, serialNumber, privKey, cfg.V3Key),
	}

	return core.NewClient(ctx, opts...)
}

// getSerialNumber 从证书中获取序列号
func getSerialNumber(certPem []byte) string {
	certDERBlock, _ := pem.Decode(certPem)
	x509Cert, _ := x509.ParseCertificate(certDERBlock.Bytes)
	serialNumberBytes := x509Cert.SerialNumber.Bytes()
	serialNumber := strings.ToUpper(hex.EncodeToString(serialNumberBytes))

	return serialNumber
}
