package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-google-account-access-token-requests-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-google-account-access-token-requests-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-google-account-access-token-requests-rmq-kube/config"
	"encoding/json"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
	"io"
	"net/http"
)

func (c *DPFMAPICaller) GoogleAccountAccessToken(
	input *dpfm_api_input_reader.SDC,
	errs *[]error,
	log *logger.Logger,
	conf *config.Conf,
) *[]dpfm_api_output_formatter.GoogleAccountAccessToken {
	var googleAccountAccessToken []dpfm_api_output_formatter.GoogleAccountAccessToken

	urlString := input.GoogleAccountAccessToken.URL

	resp, err := http.Post(urlString, "application/x-www-form-urlencoded", nil)
	defer resp.Body.Close()

	if err != nil {
		*errs = append(*errs, xerrors.Errorf("URL does not contain Code"))
		return nil
	}

	if resp.StatusCode != 200 {
		*errs = append(*errs, xerrors.Errorf("Error code: %d", resp.StatusCode))
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		*errs = append(*errs, xerrors.Errorf("Error body: %v", err.Error()))
		return nil
	}

	var tokenData map[string]interface{}
	if err := json.Unmarshal(body, &tokenData); err != nil {
		*errs = append(*errs, xerrors.Errorf("Unmarshal error: %v", err.Error()))
		return nil
	}

	googleAccountAccessToken = append(
		googleAccountAccessToken,
		dpfm_api_output_formatter.GoogleAccountAccessToken{
			AccessToken: tokenData["access_token"].(string),
		},
	)

	return &googleAccountAccessToken
}
