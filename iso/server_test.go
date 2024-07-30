package iso_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/herudins/sharedisolib/iso"
	"github.com/herudins/sharedisolib/tool"
)

type handler struct{}

func (h *handler) ExecuteTransaction(msg *iso.Message) (string, error) {
	fmt.Printf("Receiving iso message: [%s] \n", tool.AsJSON(msg))

	if msg.ResponseCode == iso.RcFail {
		msg.ResponseCode = iso.RcSuccess
		msg.ResponseMessage = "Transaksi berhasil"
		msg.Amount = "40000"
		return iso.RcSuccess, nil
	}

	return iso.RcFail, errors.New("Sengaja")
}

func OffTestServer(t *testing.T) {
	var server iso.TCPServer
	server.Handler = &handler{}
	fmt.Println("Starting server @ 5000 ...")
	server.Serve(":", 5000)
}

func TestIsoString(t *testing.T) {
	var isoLib iso.Message
	isoLib.MTI = "2200"
	isoLib.ProcessingCode = iso.PcInquiry
	isoLib.ResponseCode = iso.RcFail
	isoLib.ResponseMessage = "This is from client"
	isoLib.SetAmount(50000)

	fmt.Println(isoLib.String())
}

// func TestClient(t *testing.T) {
// 	var iso Message
// 	iso.MTI = "2200"
// 	iso.ProcessingCode = PcInquiry
// 	iso.ResponseCode = RcFail
// 	iso.ResponseMessage = "This is from client"

// 	fmt.Println("Sending Request")
// 	fmt.Println(iso.String())

// 	if err := iso.Execute("localhost", 5000); err != nil {
// 		t.Error(err)
// 	}

// 	// Equal(t, iso.ResponseCode, RcSuccess)
// 	fmt.Println(iso.String())
// }
