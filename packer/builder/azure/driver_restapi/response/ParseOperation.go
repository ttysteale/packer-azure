// Copyright (c) Microsoft Open Technologies, Inc.
// All Rights Reserved.
// Licensed under the Apache License, Version 2.0.
// See License.txt in the project root for license information.
package response

import (
	"github.com/ttysteale/packer-azure/packer/builder/azure/driver_restapi/response/model"
	"io"
)

func ParseOperation(body io.ReadCloser) (*model.Operation, error) {
	data, err := toModel(body, &model.Operation{})

	if err != nil {
		return nil, err
	}

	m := data.(*model.Operation)

	return m, nil
}
