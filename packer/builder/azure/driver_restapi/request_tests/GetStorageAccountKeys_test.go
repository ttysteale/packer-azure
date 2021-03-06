// Copyright (c) Microsoft Open Technologies, Inc.
// All Rights Reserved.
// Licensed under the Apache License, Version 2.0.
// See License.txt in the project root for license information.

package request_tests

import (
	"fmt"
	"github.com/ttysteale/packer-azure/packer/builder/azure/driver_restapi/response"
	"testing"
)

func _TestGetStorageAccountKeys(t *testing.T) {

	errMassage := "TestGetStorageAccountKeys: %s\n"

	reqManager, err := getRequestManager(t)
	if err != nil {
		t.Errorf(errMassage, err.Error())
	}

	storageAccountName := "packervhds"
	requestData := reqManager.GetStorageAccountKeys(storageAccountName)
	resp, err := reqManager.Execute(requestData)

	if err != nil {
		t.Errorf(errMassage, err.Error())
	}

	storageService, err := response.ParseStorageService(resp.Body)

	if err != nil {
		t.Errorf(errMassage, err.Error())
	}

	fmt.Printf("storageService: %v\n\n", storageService)

	t.Error("eom")
}
