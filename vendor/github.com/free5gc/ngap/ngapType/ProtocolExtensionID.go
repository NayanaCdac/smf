package ngapType

// Need to import "github.com/free5gc/aper" if it uses "aper"

type ProtocolExtensionID struct {
	Value int64 `aper:"valueLB:0,valueUB:65535"`
}
