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

package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateTieringPolicy invokes the nas.CreateTieringPolicy API synchronously
// api document: https://help.aliyun.com/api/nas/createtieringpolicy.html
func (client *Client) CreateTieringPolicy(request *CreateTieringPolicyRequest) (response *CreateTieringPolicyResponse, err error) {
	response = CreateCreateTieringPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTieringPolicyWithChan invokes the nas.CreateTieringPolicy API asynchronously
// api document: https://help.aliyun.com/api/nas/createtieringpolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTieringPolicyWithChan(request *CreateTieringPolicyRequest) (<-chan *CreateTieringPolicyResponse, <-chan error) {
	responseChan := make(chan *CreateTieringPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTieringPolicy(request)
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

// CreateTieringPolicyWithCallback invokes the nas.CreateTieringPolicy API asynchronously
// api document: https://help.aliyun.com/api/nas/createtieringpolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTieringPolicyWithCallback(request *CreateTieringPolicyRequest, callback func(response *CreateTieringPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTieringPolicyResponse
		var err error
		defer close(result)
		response, err = client.CreateTieringPolicy(request)
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

// CreateTieringPolicyRequest is the request struct for api CreateTieringPolicy
type CreateTieringPolicyRequest struct {
	*requests.RpcRequest
	Name        string           `position:"Query" name:"Name"`
	Description string           `position:"Query" name:"Description"`
	Mtime       requests.Integer `position:"Query" name:"Mtime"`
	Atime       requests.Integer `position:"Query" name:"Atime"`
	Ctime       requests.Integer `position:"Query" name:"Ctime"`
	Size        requests.Integer `position:"Query" name:"Size"`
	FileName    string           `position:"Query" name:"FileName"`
	RecallTime  requests.Integer `position:"Query" name:"RecallTime"`
	CheckLimit  requests.Boolean `position:"Query" name:"CheckLimit"`
}

// CreateTieringPolicyResponse is the response struct for api CreateTieringPolicy
type CreateTieringPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateCreateTieringPolicyRequest creates a request to invoke CreateTieringPolicy API
func CreateCreateTieringPolicyRequest() (request *CreateTieringPolicyRequest) {
	request = &CreateTieringPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "CreateTieringPolicy", "nas", "openAPI")
	return
}

// CreateCreateTieringPolicyResponse creates a response to parse from CreateTieringPolicy response
func CreateCreateTieringPolicyResponse() (response *CreateTieringPolicyResponse) {
	response = &CreateTieringPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}