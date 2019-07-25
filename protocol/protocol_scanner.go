package protocol

import (
	"github.com/gugronnier/onionscan/config"
	"github.com/gugronnier/onionscan/report"
)

type Scanner interface {
	ScanProtocol(hiddenService string, onionscanConfig *config.OnionScanConfig, report *report.OnionScanReport)
}
