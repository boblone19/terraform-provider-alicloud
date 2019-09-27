package emr

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyClusterServiceConfigForAdmin invokes the emr.ModifyClusterServiceConfigForAdmin API synchronously
// api document: https://help.aliyun.com/api/emr/modifyclusterserviceconfigforadmin.html
func (client *Client) ModifyClusterServiceConfigForAdmin(request *ModifyClusterServiceConfigForAdminRequest) (response *ModifyClusterServiceConfigForAdminResponse, err error) {
	response = CreateModifyClusterServiceConfigForAdminResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyClusterServiceConfigForAdminWithChan invokes the emr.ModifyClusterServiceConfigForAdmin API asynchronously
// api document: https://help.aliyun.com/api/emr/modifyclusterserviceconfigforadmin.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyClusterServiceConfigForAdminWithChan(request *ModifyClusterServiceConfigForAdminRequest) (<-chan *ModifyClusterServiceConfigForAdminResponse, <-chan error) {
	responseChan := make(chan *ModifyClusterServiceConfigForAdminResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyClusterServiceConfigForAdmin(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyClusterServiceConfigForAdminWithCallback invokes the emr.ModifyClusterServiceConfigForAdmin API asynchronously
// api document: https://help.aliyun.com/api/emr/modifyclusterserviceconfigforadmin.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyClusterServiceConfigForAdminWithCallback(request *ModifyClusterServiceConfigForAdminRequest, callback func(response *ModifyClusterServiceConfigForAdminResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyClusterServiceConfigForAdminResponse
		var err error
		defer close(result)
		response, err = client.ModifyClusterServiceConfigForAdmin(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyClusterServiceConfigForAdminRequest is the request struct for api ModifyClusterServiceConfigForAdmin
type ModifyClusterServiceConfigForAdminRequest struct {
	*requests.RpcRequest
	RefreshHostConfig    requests.Boolean `position:"Query" name:"RefreshHostConfig"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ConfigType           string           `position:"Query" name:"ConfigType"`
	HostInstanceId       string           `position:"Query" name:"HostInstanceId"`
	Author               string           `position:"Query" name:"Author"`
	GroupId              string           `position:"Query" name:"GroupId"`
	ClusterId            string           `position:"Query" name:"ClusterId"`
	UserId               string           `position:"Query" name:"UserId"`
	CustomConfigParams   string           `position:"Query" name:"CustomConfigParams"`
	ServiceName          string           `position:"Query" name:"ServiceName"`
	Comment              string           `position:"Query" name:"Comment"`
	GatewayClusterIdList *[]string        `position:"Query" name:"GatewayClusterIdList"  type:"Repeated"`
	ConfigParams         string           `position:"Query" name:"ConfigParams"`
}

// ModifyClusterServiceConfigForAdminResponse is the response struct for api ModifyClusterServiceConfigForAdmin
type ModifyClusterServiceConfigForAdminResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyClusterServiceConfigForAdminRequest creates a request to invoke ModifyClusterServiceConfigForAdmin API
func CreateModifyClusterServiceConfigForAdminRequest() (request *ModifyClusterServiceConfigForAdminRequest) {
	request = &ModifyClusterServiceConfigForAdminRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ModifyClusterServiceConfigForAdmin", "emr", "openAPI")
	return
}

// CreateModifyClusterServiceConfigForAdminResponse creates a response to parse from ModifyClusterServiceConfigForAdmin response
func CreateModifyClusterServiceConfigForAdminResponse() (response *ModifyClusterServiceConfigForAdminResponse) {
	response = &ModifyClusterServiceConfigForAdminResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}