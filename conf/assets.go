package conf

import (
	vb "github.com/giskook/vav-common/base"
	"github.com/giskook/vav-ms/base"
)

const (
	ASSETS_FORMAT_FILE        string = "./assets/formats.line"
	ASSETS_SAMPLING_RATE_FILE string = "./assets/sampling_rates.line"
	NOTSET                    string = "not set"
)

func read_format_file() ([]string, error) {
	return vb.ReadLineFile(ASSETS_FORMAT_FILE)
}

func read_sampling_rate_file() ([]string, error) {
	return vb.ReadLineFile(ASSETS_SAMPLING_RATE_FILE)
}

func (cnf *Conf) CheckFormat(format_index int) (string, error) {
	if format_index >= len(cnf.Formats) {
		return "", base.ERROR_BAD_REQUEST_ASSETS_OVER_RANGE
	}

	v := cnf.Formats[format_index]
	if v == NOTSET {
		return "", base.ERROR_INTERNAL_SERVER_ERROR_FORMATS_NOT_SUPPORT
	}

	return v, nil
}

func (cnf *Conf) CheckSamplingRate(sampling_rate_index int) (string, error) {
	if sampling_rate_index >= len(cnf.SamplingRates) {
		return "", base.ERROR_BAD_REQUEST_ASSETS_OVER_RANGE
	}

	return cnf.SamplingRates[sampling_rate_index], nil
}
