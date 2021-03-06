// Copyright (c) Microsoft Open Technologies, Inc.
// All Rights Reserved.
// Licensed under the Apache License, Version 2.0.
// See License.txt in the project root for license information.
package storage_tests

import (
	"fmt"
	"github.com/ttysteale/packer-azure/packer/builder/azure/driver_restapi/storage_service/request"
	"github.com/ttysteale/packer-azure/packer/builder/azure/driver_restapi/storage_service/response"
	"testing"
)

func _TestListContainers(t *testing.T) {

	errMassage := "TestListContainers: %s\n"

	sa := request.NewStorageServiceDriver(g_accountName, g_secret)

	resp, err := sa.ListContainers()

	if err != nil {
		t.Errorf(errMassage, err.Error())
	}

	list, err := response.ParseContainersList(resp.Body)
	if err != nil {
		t.Errorf(errMassage, err.Error())
	}

	fmt.Printf("list: %v\n", list)

	t.Error("eom")
}
