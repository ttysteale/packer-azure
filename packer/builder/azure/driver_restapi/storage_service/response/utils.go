package response

import (
	"encoding/xml"
	"fmt"
	"github.com/ttysteale/packer-azure/packer/builder/azure/driver_restapi/settings"
	"io"
	"io/ioutil"
	"log"
)

func readBody(body io.ReadCloser) ([]byte, error) {
	bodyData, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return bodyData, nil
}

func toModel(body io.ReadCloser, model interface{}) (interface{}, error) {

	if body == nil {
		return nil, fmt.Errorf("response body is nil")
	}

	bodyData, err := readBody(body)
	if err != nil {
		return nil, err
	}

	if settings.LogRawResponseBody {
		log.Printf("Response raw body:\n%s\n", string(bodyData))
	}

	err = xml.Unmarshal(bodyData, model)
	if err != nil {
		return nil, err
	}

	return model, nil
}
