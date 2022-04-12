// Copyright 1999-2021 Tencent Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package integration

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func TestNewSignature(t *testing.T) {

	signMethod := "TC3-HMAC-SHA256"
	reqMethodList := []string{"GET", "POST"}
	newSignature(signMethod, reqMethodList, t)
}

func newSignature(signMethod string, reqMethodList []string, t *testing.T) {

	for _, reqMethod := range reqMethodList {
		credential := common.NewCredential(
			os.Getenv("TENCENTCLOUD_SECRET_ID"),
			os.Getenv("TENCENTCLOUD_SECRET_KEY"),
		)
		cpf := profile.NewClientProfile()
		cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
		cpf.HttpProfile.ReqMethod = reqMethod
		cpf.SignMethod = signMethod
		client, _ := cvm.NewClient(credential, "ap-guangzhou", cpf)

		request := cvm.NewDescribeZonesRequest()

		response, err := client.DescribeZones(request)
		if _, ok := err.(*errors.TencentCloudSDKError); ok {
			t.Errorf(fmt.Sprintf("An API error has returned: %s", err))

		}
		if err != nil {
			t.Errorf(fmt.Sprintf("fail to init client: %v", err))
		}
		fmt.Printf("%s\n", response.ToJsonString())
	}

	for _, reqMethod := range reqMethodList {
		fmt.Printf("request: %s，use %s sign method.\n", reqMethod, signMethod)
		credential := common.NewCredential(
			os.Getenv("TENCENTCLOUD_SECRET_ID"),
			os.Getenv("TENCENTCLOUD_SECRET_KEY"),
		)
		cpf := profile.NewClientProfile()
		cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
		client, _ := cvm.NewClient(credential, "ap-guangzhou", cpf)

		request := cvm.NewDescribeInstancesOperationLimitRequest()

		request.InstanceIds = common.StringPtrs([]string{""})
		request.Operation = common.StringPtr("INSTANCE_DEGRADE")

		response, err := client.DescribeInstancesOperationLimit(request)
		if _, ok := err.(*errors.TencentCloudSDKError); ok {
			fmt.Printf("An API error has returned: %s\n", err)
			if strings.Index(err.Error(), "InvalidInstanceId.Malformed") == -1 {
				t.Errorf(fmt.Sprintf("failt to request!"))
			}
		} else {
			if err != nil {
				t.Errorf(fmt.Sprintf("fail to init client: %v", err))
			}
			fmt.Printf("%s\n", response.ToJsonString())
			t.Errorf(fmt.Sprintf("failt to request!"))
		}
	}
}
