// Copyright (c) Microsoft Open Technologies, Inc.
// All Rights Reserved.
// Licensed under the Apache License, Version 2.0.
// See License.txt in the project root for license information.
package storage_tests

import (
	"github.com/ttysteale/packer-azure/packer/builder/azure/driver_restapi/storage_service/request"
	"testing"
)

func _TestPutBlob(t *testing.T) {

	errMassage := "TestPutBlob error: %s\n"

	sa := request.NewStorageServiceDriver(g_accountName, g_secret)

	filePath := "d:\\Packer.io\\example\\Azure\\srcFolder\\npp.6.6.3.Installer.exe"

	containerName := "packer-provision"

	_, err := sa.PutBlob(containerName, filePath)

	if err != nil {
		t.Errorf(errMassage, err.Error())
	}

	t.Error("eom")
}
