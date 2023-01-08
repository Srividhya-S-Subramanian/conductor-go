//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package workflow

import "github.com/conductor-sdk/conductor-go/sdk/model"

type HttpMethod string

const (
	GET     HttpMethod = "GET"
	PUT     HttpMethod = "PUT"
	POST    HttpMethod = "POST"
	DELETE  HttpMethod = "DELETE"
	HEAD    HttpMethod = "HEAD"
	OPTIONS HttpMethod = "OPTIONS"
)

// NewHttpTask Create a new HTTP Task
func NewHttpTask(taskRefName string, input *HttpInput) *Task {
	if len(input.Method) == 0 {
		input.Method = GET
	}
	return &Task{
		model.WorkflowTask{
			Name:              taskRefName,
			TaskReferenceName: taskRefName,
			Description:       "",
			WorkflowTaskType:  string(HTTP),
			Optional:          false,
			InputParameters: map[string]interface{}{
				"http_request": input,
			},
		},
	}
}

// HttpInput Input to the HTTP task
type HttpInput struct {
	Method            HttpMethod          `json:"method"`
	Uri               string              `json:"uri"`
	Headers           map[string][]string `json:"headers,omitempty"`
	Accept            string              `json:"accept,omitempty"`
	ContentType       string              `json:"contentType,omitempty"`
	ConnectionTimeOut int16               `json:"ConnectionTimeOut,omitempty"`
	ReadTimeout       int16               `json:"readTimeout,omitempty"`
	Body              interface{}         `json:"body,omitempty"`
}
