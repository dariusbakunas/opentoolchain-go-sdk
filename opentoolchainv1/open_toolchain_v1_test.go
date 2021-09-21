/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package opentoolchainv1_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/dariusbakunas/opentoolchain-go-sdk/opentoolchainv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`OpenToolchainV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(openToolchainService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(openToolchainService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
				URL: "https://opentoolchainv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(openToolchainService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_TOOLCHAIN_URL":       "https://opentoolchainv1/api",
				"OPEN_TOOLCHAIN_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1UsingExternalConfig(&opentoolchainv1.OpenToolchainV1Options{})
				Expect(openToolchainService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := openToolchainService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != openToolchainService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(openToolchainService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(openToolchainService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1UsingExternalConfig(&opentoolchainv1.OpenToolchainV1Options{
					URL: "https://testService/api",
				})
				Expect(openToolchainService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := openToolchainService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != openToolchainService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(openToolchainService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(openToolchainService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1UsingExternalConfig(&opentoolchainv1.OpenToolchainV1Options{})
				err := openToolchainService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := openToolchainService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != openToolchainService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(openToolchainService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(openToolchainService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_TOOLCHAIN_URL":       "https://opentoolchainv1/api",
				"OPEN_TOOLCHAIN_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1UsingExternalConfig(&opentoolchainv1.OpenToolchainV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(openToolchainService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_TOOLCHAIN_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1UsingExternalConfig(&opentoolchainv1.OpenToolchainV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(openToolchainService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = opentoolchainv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`PatchToolchain(patchToolchainOptions *PatchToolchainOptions)`, func() {
		patchToolchainPath := "/devops/toolchains/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(patchToolchainPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env_id"]).To(Equal([]string{"ibm:yp:us-south"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke PatchToolchain successfully`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := openToolchainService.PatchToolchain(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the PatchToolchainOptions model
				patchToolchainOptionsModel := new(opentoolchainv1.PatchToolchainOptions)
				patchToolchainOptionsModel.GUID = core.StringPtr("testString")
				patchToolchainOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				patchToolchainOptionsModel.Name = core.StringPtr("testString")
				patchToolchainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = openToolchainService.PatchToolchain(patchToolchainOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke PatchToolchain with error: Operation validation and request error`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Construct an instance of the PatchToolchainOptions model
				patchToolchainOptionsModel := new(opentoolchainv1.PatchToolchainOptions)
				patchToolchainOptionsModel.GUID = core.StringPtr("testString")
				patchToolchainOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				patchToolchainOptionsModel.Name = core.StringPtr("testString")
				patchToolchainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := openToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := openToolchainService.PatchToolchain(patchToolchainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the PatchToolchainOptions model with no property values
				patchToolchainOptionsModelNew := new(opentoolchainv1.PatchToolchainOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = openToolchainService.PatchToolchain(patchToolchainOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteToolchain(deleteToolchainOptions *DeleteToolchainOptions)`, func() {
		deleteToolchainPath := "/devops/toolchains/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteToolchainPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["env_id"]).To(Equal([]string{"ibm:yp:us-south"}))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteToolchain successfully`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := openToolchainService.DeleteToolchain(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteToolchainOptions model
				deleteToolchainOptionsModel := new(opentoolchainv1.DeleteToolchainOptions)
				deleteToolchainOptionsModel.GUID = core.StringPtr("testString")
				deleteToolchainOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				deleteToolchainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = openToolchainService.DeleteToolchain(deleteToolchainOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteToolchain with error: Operation validation and request error`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Construct an instance of the DeleteToolchainOptions model
				deleteToolchainOptionsModel := new(opentoolchainv1.DeleteToolchainOptions)
				deleteToolchainOptionsModel.GUID = core.StringPtr("testString")
				deleteToolchainOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				deleteToolchainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := openToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := openToolchainService.DeleteToolchain(deleteToolchainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteToolchainOptions model with no property values
				deleteToolchainOptionsModelNew := new(opentoolchainv1.DeleteToolchainOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = openToolchainService.DeleteToolchain(deleteToolchainOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateToolchain(createToolchainOptions *CreateToolchainOptions)`, func() {
		createToolchainPath := "/devops/setup/deploy"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createToolchainPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["env_id"]).To(Equal([]string{"ibm:yp:us-south"}))
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateToolchain successfully`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := openToolchainService.CreateToolchain(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CreateToolchainOptions model
				createToolchainOptionsModel := new(opentoolchainv1.CreateToolchainOptions)
				createToolchainOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				createToolchainOptionsModel.Repository = core.StringPtr("testString")
				createToolchainOptionsModel.Autocreate = core.BoolPtr(true)
				createToolchainOptionsModel.ResourceGroupID = core.StringPtr("testString")
				createToolchainOptionsModel.RepositoryToken = core.StringPtr("testString")
				createToolchainOptionsModel.Branch = core.StringPtr("testString")
				createToolchainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = openToolchainService.CreateToolchain(createToolchainOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke CreateToolchain with error: Operation validation and request error`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Construct an instance of the CreateToolchainOptions model
				createToolchainOptionsModel := new(opentoolchainv1.CreateToolchainOptions)
				createToolchainOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				createToolchainOptionsModel.Repository = core.StringPtr("testString")
				createToolchainOptionsModel.Autocreate = core.BoolPtr(true)
				createToolchainOptionsModel.ResourceGroupID = core.StringPtr("testString")
				createToolchainOptionsModel.RepositoryToken = core.StringPtr("testString")
				createToolchainOptionsModel.Branch = core.StringPtr("testString")
				createToolchainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := openToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := openToolchainService.CreateToolchain(createToolchainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the CreateToolchainOptions model with no property values
				createToolchainOptionsModelNew := new(opentoolchainv1.CreateToolchainOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = openToolchainService.CreateToolchain(createToolchainOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateServiceInstance(createServiceInstanceOptions *CreateServiceInstanceOptions) - Operation response error`, func() {
		createServiceInstancePath := "/devops/service_instances"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createServiceInstancePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["env_id"]).To(Equal([]string{"ibm:yp:us-south"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateServiceInstance with error: Operation response processing error`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Construct an instance of the CreateServiceInstanceParamsParameters model
				createServiceInstanceParamsParametersModel := new(opentoolchainv1.CreateServiceInstanceParamsParameters)
				createServiceInstanceParamsParametersModel.APIKey = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ServiceKey = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.KeyType = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ServiceID = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ServiceName = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ServiceURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.UserEmail = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.UserPhone = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.Authorized = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ChannelName = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.TeamURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.PipelineStart = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.PipelineSuccess = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.PipelineFail = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.ToolchainBind = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.ToolchainUnbind = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.APIToken = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.GitID = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.APIRootURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.Legal = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.RepoURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.TokenURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.PrivateRepo = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.HasIssues = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.InstanceName = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.IntegrationStatus = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.Region = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ResourceGroup = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.EnableTraceability = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.Name = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.Type = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.UIPipeline = core.BoolPtr(true)

				// Construct an instance of the CreateServiceInstanceOptions model
				createServiceInstanceOptionsModel := new(opentoolchainv1.CreateServiceInstanceOptions)
				createServiceInstanceOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				createServiceInstanceOptionsModel.ToolchainID = core.StringPtr("testString")
				createServiceInstanceOptionsModel.ServiceID = core.StringPtr("testString")
				createServiceInstanceOptionsModel.Parameters = createServiceInstanceParamsParametersModel
				createServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := openToolchainService.CreateServiceInstance(createServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				openToolchainService.EnableRetries(0, 0)
				result, response, operationErr = openToolchainService.CreateServiceInstance(createServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateServiceInstance(createServiceInstanceOptions *CreateServiceInstanceOptions)`, func() {
		createServiceInstancePath := "/devops/service_instances"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createServiceInstancePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env_id"]).To(Equal([]string{"ibm:yp:us-south"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": "Status"}`)
				}))
			})
			It(`Invoke CreateServiceInstance successfully with retries`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())
				openToolchainService.EnableRetries(0, 0)

				// Construct an instance of the CreateServiceInstanceParamsParameters model
				createServiceInstanceParamsParametersModel := new(opentoolchainv1.CreateServiceInstanceParamsParameters)
				createServiceInstanceParamsParametersModel.APIKey = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ServiceKey = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.KeyType = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ServiceID = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ServiceName = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ServiceURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.UserEmail = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.UserPhone = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.Authorized = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ChannelName = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.TeamURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.PipelineStart = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.PipelineSuccess = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.PipelineFail = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.ToolchainBind = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.ToolchainUnbind = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.APIToken = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.GitID = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.APIRootURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.Legal = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.RepoURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.TokenURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.PrivateRepo = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.HasIssues = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.InstanceName = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.IntegrationStatus = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.Region = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ResourceGroup = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.EnableTraceability = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.Name = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.Type = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.UIPipeline = core.BoolPtr(true)

				// Construct an instance of the CreateServiceInstanceOptions model
				createServiceInstanceOptionsModel := new(opentoolchainv1.CreateServiceInstanceOptions)
				createServiceInstanceOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				createServiceInstanceOptionsModel.ToolchainID = core.StringPtr("testString")
				createServiceInstanceOptionsModel.ServiceID = core.StringPtr("testString")
				createServiceInstanceOptionsModel.Parameters = createServiceInstanceParamsParametersModel
				createServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := openToolchainService.CreateServiceInstanceWithContext(ctx, createServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				openToolchainService.DisableRetries()
				result, response, operationErr := openToolchainService.CreateServiceInstance(createServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = openToolchainService.CreateServiceInstanceWithContext(ctx, createServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createServiceInstancePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env_id"]).To(Equal([]string{"ibm:yp:us-south"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": "Status"}`)
				}))
			})
			It(`Invoke CreateServiceInstance successfully`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := openToolchainService.CreateServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateServiceInstanceParamsParameters model
				createServiceInstanceParamsParametersModel := new(opentoolchainv1.CreateServiceInstanceParamsParameters)
				createServiceInstanceParamsParametersModel.APIKey = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ServiceKey = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.KeyType = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ServiceID = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ServiceName = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ServiceURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.UserEmail = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.UserPhone = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.Authorized = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ChannelName = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.TeamURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.PipelineStart = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.PipelineSuccess = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.PipelineFail = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.ToolchainBind = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.ToolchainUnbind = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.APIToken = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.GitID = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.APIRootURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.Legal = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.RepoURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.TokenURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.PrivateRepo = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.HasIssues = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.InstanceName = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.IntegrationStatus = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.Region = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ResourceGroup = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.EnableTraceability = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.Name = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.Type = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.UIPipeline = core.BoolPtr(true)

				// Construct an instance of the CreateServiceInstanceOptions model
				createServiceInstanceOptionsModel := new(opentoolchainv1.CreateServiceInstanceOptions)
				createServiceInstanceOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				createServiceInstanceOptionsModel.ToolchainID = core.StringPtr("testString")
				createServiceInstanceOptionsModel.ServiceID = core.StringPtr("testString")
				createServiceInstanceOptionsModel.Parameters = createServiceInstanceParamsParametersModel
				createServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = openToolchainService.CreateServiceInstance(createServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateServiceInstance with error: Operation validation and request error`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Construct an instance of the CreateServiceInstanceParamsParameters model
				createServiceInstanceParamsParametersModel := new(opentoolchainv1.CreateServiceInstanceParamsParameters)
				createServiceInstanceParamsParametersModel.APIKey = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ServiceKey = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.KeyType = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ServiceID = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ServiceName = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ServiceURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.UserEmail = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.UserPhone = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.Authorized = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ChannelName = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.TeamURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.PipelineStart = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.PipelineSuccess = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.PipelineFail = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.ToolchainBind = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.ToolchainUnbind = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.APIToken = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.GitID = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.APIRootURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.Legal = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.RepoURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.TokenURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.PrivateRepo = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.HasIssues = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.InstanceName = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.IntegrationStatus = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.Region = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ResourceGroup = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.EnableTraceability = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.Name = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.Type = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.UIPipeline = core.BoolPtr(true)

				// Construct an instance of the CreateServiceInstanceOptions model
				createServiceInstanceOptionsModel := new(opentoolchainv1.CreateServiceInstanceOptions)
				createServiceInstanceOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				createServiceInstanceOptionsModel.ToolchainID = core.StringPtr("testString")
				createServiceInstanceOptionsModel.ServiceID = core.StringPtr("testString")
				createServiceInstanceOptionsModel.Parameters = createServiceInstanceParamsParametersModel
				createServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := openToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := openToolchainService.CreateServiceInstance(createServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateServiceInstanceOptions model with no property values
				createServiceInstanceOptionsModelNew := new(opentoolchainv1.CreateServiceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = openToolchainService.CreateServiceInstance(createServiceInstanceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke CreateServiceInstance successfully`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Construct an instance of the CreateServiceInstanceParamsParameters model
				createServiceInstanceParamsParametersModel := new(opentoolchainv1.CreateServiceInstanceParamsParameters)
				createServiceInstanceParamsParametersModel.APIKey = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ServiceKey = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.KeyType = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ServiceID = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ServiceName = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ServiceURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.UserEmail = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.UserPhone = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.Authorized = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ChannelName = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.TeamURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.PipelineStart = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.PipelineSuccess = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.PipelineFail = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.ToolchainBind = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.ToolchainUnbind = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.APIToken = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.GitID = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.APIRootURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.Legal = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.RepoURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.TokenURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.PrivateRepo = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.HasIssues = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.InstanceName = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.IntegrationStatus = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.Region = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ResourceGroup = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.EnableTraceability = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.Name = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.Type = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.UIPipeline = core.BoolPtr(true)

				// Construct an instance of the CreateServiceInstanceOptions model
				createServiceInstanceOptionsModel := new(opentoolchainv1.CreateServiceInstanceOptions)
				createServiceInstanceOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				createServiceInstanceOptionsModel.ToolchainID = core.StringPtr("testString")
				createServiceInstanceOptionsModel.ServiceID = core.StringPtr("testString")
				createServiceInstanceOptionsModel.Parameters = createServiceInstanceParamsParametersModel
				createServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := openToolchainService.CreateServiceInstance(createServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteServiceInstance(deleteServiceInstanceOptions *DeleteServiceInstanceOptions)`, func() {
		deleteServiceInstancePath := "/devops/service_instances/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteServiceInstancePath))
					Expect(req.Method).To(Equal("DELETE"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env_id"]).To(Equal([]string{"ibm:yp:us-south"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteServiceInstance successfully`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := openToolchainService.DeleteServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteServiceInstanceOptions model
				deleteServiceInstanceOptionsModel := new(opentoolchainv1.DeleteServiceInstanceOptions)
				deleteServiceInstanceOptionsModel.GUID = core.StringPtr("testString")
				deleteServiceInstanceOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				deleteServiceInstanceOptionsModel.ToolchainID = core.StringPtr("testString")
				deleteServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = openToolchainService.DeleteServiceInstance(deleteServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteServiceInstance with error: Operation validation and request error`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Construct an instance of the DeleteServiceInstanceOptions model
				deleteServiceInstanceOptionsModel := new(opentoolchainv1.DeleteServiceInstanceOptions)
				deleteServiceInstanceOptionsModel.GUID = core.StringPtr("testString")
				deleteServiceInstanceOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				deleteServiceInstanceOptionsModel.ToolchainID = core.StringPtr("testString")
				deleteServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := openToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := openToolchainService.DeleteServiceInstance(deleteServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteServiceInstanceOptions model with no property values
				deleteServiceInstanceOptionsModelNew := new(opentoolchainv1.DeleteServiceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = openToolchainService.DeleteServiceInstance(deleteServiceInstanceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PatchServiceInstance(patchServiceInstanceOptions *PatchServiceInstanceOptions)`, func() {
		patchServiceInstancePath := "/devops/service_instances/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(patchServiceInstancePath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env_id"]).To(Equal([]string{"ibm:yp:us-south"}))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke PatchServiceInstance successfully`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := openToolchainService.PatchServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the PatchServiceInstanceParamsParameters model
				patchServiceInstanceParamsParametersModel := new(opentoolchainv1.PatchServiceInstanceParamsParameters)
				patchServiceInstanceParamsParametersModel.APIKey = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.ServiceKey = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.KeyType = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.ServiceID = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.ServiceName = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.ServiceURL = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.UserEmail = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.UserPhone = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.Name = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.Type = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.UIPipeline = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.RepoURL = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.PrivateRepo = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.HasIssues = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.EnableTraceability = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.InstanceName = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.IntegrationStatus = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.Region = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.ResourceGroup = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.ChannelName = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.TeamURL = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.PipelineStart = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.PipelineSuccess = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.PipelineFail = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.ToolchainBind = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.ToolchainUnbind = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.APIToken = core.StringPtr("testString")

				// Construct an instance of the PatchServiceInstanceOptions model
				patchServiceInstanceOptionsModel := new(opentoolchainv1.PatchServiceInstanceOptions)
				patchServiceInstanceOptionsModel.GUID = core.StringPtr("testString")
				patchServiceInstanceOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				patchServiceInstanceOptionsModel.ToolchainID = core.StringPtr("testString")
				patchServiceInstanceOptionsModel.ServiceID = core.StringPtr("testString")
				patchServiceInstanceOptionsModel.Parameters = patchServiceInstanceParamsParametersModel
				patchServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = openToolchainService.PatchServiceInstance(patchServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke PatchServiceInstance with error: Operation validation and request error`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Construct an instance of the PatchServiceInstanceParamsParameters model
				patchServiceInstanceParamsParametersModel := new(opentoolchainv1.PatchServiceInstanceParamsParameters)
				patchServiceInstanceParamsParametersModel.APIKey = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.ServiceKey = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.KeyType = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.ServiceID = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.ServiceName = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.ServiceURL = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.UserEmail = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.UserPhone = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.Name = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.Type = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.UIPipeline = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.RepoURL = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.PrivateRepo = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.HasIssues = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.EnableTraceability = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.InstanceName = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.IntegrationStatus = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.Region = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.ResourceGroup = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.ChannelName = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.TeamURL = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.PipelineStart = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.PipelineSuccess = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.PipelineFail = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.ToolchainBind = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.ToolchainUnbind = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.APIToken = core.StringPtr("testString")

				// Construct an instance of the PatchServiceInstanceOptions model
				patchServiceInstanceOptionsModel := new(opentoolchainv1.PatchServiceInstanceOptions)
				patchServiceInstanceOptionsModel.GUID = core.StringPtr("testString")
				patchServiceInstanceOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				patchServiceInstanceOptionsModel.ToolchainID = core.StringPtr("testString")
				patchServiceInstanceOptionsModel.ServiceID = core.StringPtr("testString")
				patchServiceInstanceOptionsModel.Parameters = patchServiceInstanceParamsParametersModel
				patchServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := openToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := openToolchainService.PatchServiceInstance(patchServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the PatchServiceInstanceOptions model with no property values
				patchServiceInstanceOptionsModelNew := new(opentoolchainv1.PatchServiceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = openToolchainService.PatchServiceInstance(patchServiceInstanceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetServiceInstance(getServiceInstanceOptions *GetServiceInstanceOptions) - Operation response error`, func() {
		getServiceInstancePath := "/devops/service_instances/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getServiceInstancePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["env_id"]).To(Equal([]string{"ibm:yp:us-south"}))
					Expect(req.URL.Query()["toolchainId"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetServiceInstance with error: Operation response processing error`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Construct an instance of the GetServiceInstanceOptions model
				getServiceInstanceOptionsModel := new(opentoolchainv1.GetServiceInstanceOptions)
				getServiceInstanceOptionsModel.GUID = core.StringPtr("testString")
				getServiceInstanceOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				getServiceInstanceOptionsModel.ToolchainID = core.StringPtr("testString")
				getServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := openToolchainService.GetServiceInstance(getServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				openToolchainService.EnableRetries(0, 0)
				result, response, operationErr = openToolchainService.GetServiceInstance(getServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetServiceInstance(getServiceInstanceOptions *GetServiceInstanceOptions)`, func() {
		getServiceInstancePath := "/devops/service_instances/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getServiceInstancePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["env_id"]).To(Equal([]string{"ibm:yp:us-south"}))
					Expect(req.URL.Query()["toolchainId"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"serviceInstance": {"instance_id": "InstanceID", "dashboard_url": "DashboardURL", "service_id": "ServiceID", "parameters": {"mapKey": "anyValue"}}}`)
				}))
			})
			It(`Invoke GetServiceInstance successfully with retries`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())
				openToolchainService.EnableRetries(0, 0)

				// Construct an instance of the GetServiceInstanceOptions model
				getServiceInstanceOptionsModel := new(opentoolchainv1.GetServiceInstanceOptions)
				getServiceInstanceOptionsModel.GUID = core.StringPtr("testString")
				getServiceInstanceOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				getServiceInstanceOptionsModel.ToolchainID = core.StringPtr("testString")
				getServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := openToolchainService.GetServiceInstanceWithContext(ctx, getServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				openToolchainService.DisableRetries()
				result, response, operationErr := openToolchainService.GetServiceInstance(getServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = openToolchainService.GetServiceInstanceWithContext(ctx, getServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getServiceInstancePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["env_id"]).To(Equal([]string{"ibm:yp:us-south"}))
					Expect(req.URL.Query()["toolchainId"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"serviceInstance": {"instance_id": "InstanceID", "dashboard_url": "DashboardURL", "service_id": "ServiceID", "parameters": {"mapKey": "anyValue"}}}`)
				}))
			})
			It(`Invoke GetServiceInstance successfully`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := openToolchainService.GetServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetServiceInstanceOptions model
				getServiceInstanceOptionsModel := new(opentoolchainv1.GetServiceInstanceOptions)
				getServiceInstanceOptionsModel.GUID = core.StringPtr("testString")
				getServiceInstanceOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				getServiceInstanceOptionsModel.ToolchainID = core.StringPtr("testString")
				getServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = openToolchainService.GetServiceInstance(getServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetServiceInstance with error: Operation validation and request error`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Construct an instance of the GetServiceInstanceOptions model
				getServiceInstanceOptionsModel := new(opentoolchainv1.GetServiceInstanceOptions)
				getServiceInstanceOptionsModel.GUID = core.StringPtr("testString")
				getServiceInstanceOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				getServiceInstanceOptionsModel.ToolchainID = core.StringPtr("testString")
				getServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := openToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := openToolchainService.GetServiceInstance(getServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetServiceInstanceOptions model with no property values
				getServiceInstanceOptionsModelNew := new(opentoolchainv1.GetServiceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = openToolchainService.GetServiceInstance(getServiceInstanceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetServiceInstance successfully`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Construct an instance of the GetServiceInstanceOptions model
				getServiceInstanceOptionsModel := new(opentoolchainv1.GetServiceInstanceOptions)
				getServiceInstanceOptionsModel.GUID = core.StringPtr("testString")
				getServiceInstanceOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				getServiceInstanceOptionsModel.ToolchainID = core.StringPtr("testString")
				getServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := openToolchainService.GetServiceInstance(getServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTektonPipeline(getTektonPipelineOptions *GetTektonPipelineOptions) - Operation response error`, func() {
		getTektonPipelinePath := "/devops/pipelines/tekton/api/v1/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelinePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["env_id"]).To(Equal([]string{"ibm:yp:us-south"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTektonPipeline with error: Operation response processing error`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineOptions model
				getTektonPipelineOptionsModel := new(opentoolchainv1.GetTektonPipelineOptions)
				getTektonPipelineOptionsModel.GUID = core.StringPtr("testString")
				getTektonPipelineOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				getTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := openToolchainService.GetTektonPipeline(getTektonPipelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				openToolchainService.EnableRetries(0, 0)
				result, response, operationErr = openToolchainService.GetTektonPipeline(getTektonPipelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTektonPipeline(getTektonPipelineOptions *GetTektonPipelineOptions)`, func() {
		getTektonPipelinePath := "/devops/pipelines/tekton/api/v1/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelinePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["env_id"]).To(Equal([]string{"ibm:yp:us-south"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "dashboard_url": "DashboardURL", "resourceGroupId": "ResourceGroupID", "id": "ID", "toolchainId": "ToolchainID", "pipelineOwner": "PipelineOwner", "enabled": false, "type": "Type", "created": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "updated_at_timestamp": 18, "created_timestamp": 16, "envProperties": [{"name": "Name", "value": "Value", "type": "Type"}], "inputs": [{"type": "Type", "serviceInstanceId": "ServiceInstanceID", "shardDefinitionId": "ShardDefinitionID", "scmSource": {"path": "Path", "url": "URL", "type": "Type", "blindConnection": false, "branch": "Branch"}}], "triggers": [{"id": "ID", "name": "Name", "eventListener": "EventListener", "disabled": true, "scmSource": {"url": "URL", "type": "Type", "branch": "Branch", "pattern": "Pattern"}, "type": "Type", "events": {"push": true, "pull_request": false, "pull_request_closed": false}, "serviceInstanceId": "ServiceInstanceID"}], "status": "Status", "url": "URL", "runs_url": "RunsURL", "toolchainCRN": "ToolchainCRN", "pipelineDefinitionId": "PipelineDefinitionID"}`)
				}))
			})
			It(`Invoke GetTektonPipeline successfully with retries`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())
				openToolchainService.EnableRetries(0, 0)

				// Construct an instance of the GetTektonPipelineOptions model
				getTektonPipelineOptionsModel := new(opentoolchainv1.GetTektonPipelineOptions)
				getTektonPipelineOptionsModel.GUID = core.StringPtr("testString")
				getTektonPipelineOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				getTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := openToolchainService.GetTektonPipelineWithContext(ctx, getTektonPipelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				openToolchainService.DisableRetries()
				result, response, operationErr := openToolchainService.GetTektonPipeline(getTektonPipelineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = openToolchainService.GetTektonPipelineWithContext(ctx, getTektonPipelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelinePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["env_id"]).To(Equal([]string{"ibm:yp:us-south"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "dashboard_url": "DashboardURL", "resourceGroupId": "ResourceGroupID", "id": "ID", "toolchainId": "ToolchainID", "pipelineOwner": "PipelineOwner", "enabled": false, "type": "Type", "created": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "updated_at_timestamp": 18, "created_timestamp": 16, "envProperties": [{"name": "Name", "value": "Value", "type": "Type"}], "inputs": [{"type": "Type", "serviceInstanceId": "ServiceInstanceID", "shardDefinitionId": "ShardDefinitionID", "scmSource": {"path": "Path", "url": "URL", "type": "Type", "blindConnection": false, "branch": "Branch"}}], "triggers": [{"id": "ID", "name": "Name", "eventListener": "EventListener", "disabled": true, "scmSource": {"url": "URL", "type": "Type", "branch": "Branch", "pattern": "Pattern"}, "type": "Type", "events": {"push": true, "pull_request": false, "pull_request_closed": false}, "serviceInstanceId": "ServiceInstanceID"}], "status": "Status", "url": "URL", "runs_url": "RunsURL", "toolchainCRN": "ToolchainCRN", "pipelineDefinitionId": "PipelineDefinitionID"}`)
				}))
			})
			It(`Invoke GetTektonPipeline successfully`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := openToolchainService.GetTektonPipeline(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTektonPipelineOptions model
				getTektonPipelineOptionsModel := new(opentoolchainv1.GetTektonPipelineOptions)
				getTektonPipelineOptionsModel.GUID = core.StringPtr("testString")
				getTektonPipelineOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				getTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = openToolchainService.GetTektonPipeline(getTektonPipelineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTektonPipeline with error: Operation validation and request error`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineOptions model
				getTektonPipelineOptionsModel := new(opentoolchainv1.GetTektonPipelineOptions)
				getTektonPipelineOptionsModel.GUID = core.StringPtr("testString")
				getTektonPipelineOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				getTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := openToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := openToolchainService.GetTektonPipeline(getTektonPipelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTektonPipelineOptions model with no property values
				getTektonPipelineOptionsModelNew := new(opentoolchainv1.GetTektonPipelineOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = openToolchainService.GetTektonPipeline(getTektonPipelineOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetTektonPipeline successfully`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineOptions model
				getTektonPipelineOptionsModel := new(opentoolchainv1.GetTektonPipelineOptions)
				getTektonPipelineOptionsModel.GUID = core.StringPtr("testString")
				getTektonPipelineOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				getTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := openToolchainService.GetTektonPipeline(getTektonPipelineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PatchTektonPipeline(patchTektonPipelineOptions *PatchTektonPipelineOptions) - Operation response error`, func() {
		patchTektonPipelinePath := "/devops/pipelines/tekton/api/v1/testString/config"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(patchTektonPipelinePath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["env_id"]).To(Equal([]string{"ibm:yp:us-south"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PatchTektonPipeline with error: Operation response processing error`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Construct an instance of the PatchTektonPipelineParamsWorker model
				patchTektonPipelineParamsWorkerModel := new(opentoolchainv1.PatchTektonPipelineParamsWorker)
				patchTektonPipelineParamsWorkerModel.WorkerID = core.StringPtr("testString")
				patchTektonPipelineParamsWorkerModel.WorkerName = core.StringPtr("testString")
				patchTektonPipelineParamsWorkerModel.WorkerType = core.StringPtr("testString")

				// Construct an instance of the EnvProperty model
				envPropertyModel := new(opentoolchainv1.EnvProperty)
				envPropertyModel.Name = core.StringPtr("testString")
				envPropertyModel.Value = core.StringPtr("testString")
				envPropertyModel.Type = core.StringPtr("testString")

				// Construct an instance of the TektonPipelineInputScmSource model
				tektonPipelineInputScmSourceModel := new(opentoolchainv1.TektonPipelineInputScmSource)
				tektonPipelineInputScmSourceModel.Path = core.StringPtr("testString")
				tektonPipelineInputScmSourceModel.URL = core.StringPtr("testString")
				tektonPipelineInputScmSourceModel.Type = core.StringPtr("testString")
				tektonPipelineInputScmSourceModel.BlindConnection = core.BoolPtr(true)
				tektonPipelineInputScmSourceModel.Branch = core.StringPtr("testString")

				// Construct an instance of the TektonPipelineInput model
				tektonPipelineInputModel := new(opentoolchainv1.TektonPipelineInput)
				tektonPipelineInputModel.Type = core.StringPtr("testString")
				tektonPipelineInputModel.ServiceInstanceID = core.StringPtr("testString")
				tektonPipelineInputModel.ShardDefinitionID = core.StringPtr("testString")
				tektonPipelineInputModel.ScmSource = tektonPipelineInputScmSourceModel

				// Construct an instance of the TektonPipelineTriggerScmSource model
				tektonPipelineTriggerScmSourceModel := new(opentoolchainv1.TektonPipelineTriggerScmSource)
				tektonPipelineTriggerScmSourceModel.URL = core.StringPtr("testString")
				tektonPipelineTriggerScmSourceModel.Type = core.StringPtr("testString")
				tektonPipelineTriggerScmSourceModel.Branch = core.StringPtr("testString")
				tektonPipelineTriggerScmSourceModel.Pattern = core.StringPtr("testString")

				// Construct an instance of the TektonPipelineTriggerEvents model
				tektonPipelineTriggerEventsModel := new(opentoolchainv1.TektonPipelineTriggerEvents)
				tektonPipelineTriggerEventsModel.Push = core.BoolPtr(true)
				tektonPipelineTriggerEventsModel.PullRequest = core.BoolPtr(true)
				tektonPipelineTriggerEventsModel.PullRequestClosed = core.BoolPtr(true)

				// Construct an instance of the TektonPipelineTrigger model
				tektonPipelineTriggerModel := new(opentoolchainv1.TektonPipelineTrigger)
				tektonPipelineTriggerModel.ID = core.StringPtr("testString")
				tektonPipelineTriggerModel.Name = core.StringPtr("testString")
				tektonPipelineTriggerModel.EventListener = core.StringPtr("testString")
				tektonPipelineTriggerModel.Disabled = core.BoolPtr(true)
				tektonPipelineTriggerModel.ScmSource = tektonPipelineTriggerScmSourceModel
				tektonPipelineTriggerModel.Type = core.StringPtr("testString")
				tektonPipelineTriggerModel.Events = tektonPipelineTriggerEventsModel
				tektonPipelineTriggerModel.ServiceInstanceID = core.StringPtr("testString")

				// Construct an instance of the PatchTektonPipelineOptions model
				patchTektonPipelineOptionsModel := new(opentoolchainv1.PatchTektonPipelineOptions)
				patchTektonPipelineOptionsModel.GUID = core.StringPtr("testString")
				patchTektonPipelineOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				patchTektonPipelineOptionsModel.Worker = patchTektonPipelineParamsWorkerModel
				patchTektonPipelineOptionsModel.EnvProperties = []opentoolchainv1.EnvProperty{*envPropertyModel}
				patchTektonPipelineOptionsModel.Inputs = []opentoolchainv1.TektonPipelineInput{*tektonPipelineInputModel}
				patchTektonPipelineOptionsModel.Triggers = []opentoolchainv1.TektonPipelineTrigger{*tektonPipelineTriggerModel}
				patchTektonPipelineOptionsModel.PipelineDefinitionID = core.StringPtr("testString")
				patchTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := openToolchainService.PatchTektonPipeline(patchTektonPipelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				openToolchainService.EnableRetries(0, 0)
				result, response, operationErr = openToolchainService.PatchTektonPipeline(patchTektonPipelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PatchTektonPipeline(patchTektonPipelineOptions *PatchTektonPipelineOptions)`, func() {
		patchTektonPipelinePath := "/devops/pipelines/tekton/api/v1/testString/config"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(patchTektonPipelinePath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env_id"]).To(Equal([]string{"ibm:yp:us-south"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "dashboard_url": "DashboardURL", "resourceGroupId": "ResourceGroupID", "id": "ID", "toolchainId": "ToolchainID", "pipelineOwner": "PipelineOwner", "enabled": false, "type": "Type", "created": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "updated_at_timestamp": 18, "created_timestamp": 16, "envProperties": [{"name": "Name", "value": "Value", "type": "Type"}], "inputs": [{"type": "Type", "serviceInstanceId": "ServiceInstanceID", "shardDefinitionId": "ShardDefinitionID", "scmSource": {"path": "Path", "url": "URL", "type": "Type", "blindConnection": false, "branch": "Branch"}}], "triggers": [{"id": "ID", "name": "Name", "eventListener": "EventListener", "disabled": true, "scmSource": {"url": "URL", "type": "Type", "branch": "Branch", "pattern": "Pattern"}, "type": "Type", "events": {"push": true, "pull_request": false, "pull_request_closed": false}, "serviceInstanceId": "ServiceInstanceID"}], "status": "Status", "url": "URL", "runs_url": "RunsURL", "toolchainCRN": "ToolchainCRN", "pipelineDefinitionId": "PipelineDefinitionID"}`)
				}))
			})
			It(`Invoke PatchTektonPipeline successfully with retries`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())
				openToolchainService.EnableRetries(0, 0)

				// Construct an instance of the PatchTektonPipelineParamsWorker model
				patchTektonPipelineParamsWorkerModel := new(opentoolchainv1.PatchTektonPipelineParamsWorker)
				patchTektonPipelineParamsWorkerModel.WorkerID = core.StringPtr("testString")
				patchTektonPipelineParamsWorkerModel.WorkerName = core.StringPtr("testString")
				patchTektonPipelineParamsWorkerModel.WorkerType = core.StringPtr("testString")

				// Construct an instance of the EnvProperty model
				envPropertyModel := new(opentoolchainv1.EnvProperty)
				envPropertyModel.Name = core.StringPtr("testString")
				envPropertyModel.Value = core.StringPtr("testString")
				envPropertyModel.Type = core.StringPtr("testString")

				// Construct an instance of the TektonPipelineInputScmSource model
				tektonPipelineInputScmSourceModel := new(opentoolchainv1.TektonPipelineInputScmSource)
				tektonPipelineInputScmSourceModel.Path = core.StringPtr("testString")
				tektonPipelineInputScmSourceModel.URL = core.StringPtr("testString")
				tektonPipelineInputScmSourceModel.Type = core.StringPtr("testString")
				tektonPipelineInputScmSourceModel.BlindConnection = core.BoolPtr(true)
				tektonPipelineInputScmSourceModel.Branch = core.StringPtr("testString")

				// Construct an instance of the TektonPipelineInput model
				tektonPipelineInputModel := new(opentoolchainv1.TektonPipelineInput)
				tektonPipelineInputModel.Type = core.StringPtr("testString")
				tektonPipelineInputModel.ServiceInstanceID = core.StringPtr("testString")
				tektonPipelineInputModel.ShardDefinitionID = core.StringPtr("testString")
				tektonPipelineInputModel.ScmSource = tektonPipelineInputScmSourceModel

				// Construct an instance of the TektonPipelineTriggerScmSource model
				tektonPipelineTriggerScmSourceModel := new(opentoolchainv1.TektonPipelineTriggerScmSource)
				tektonPipelineTriggerScmSourceModel.URL = core.StringPtr("testString")
				tektonPipelineTriggerScmSourceModel.Type = core.StringPtr("testString")
				tektonPipelineTriggerScmSourceModel.Branch = core.StringPtr("testString")
				tektonPipelineTriggerScmSourceModel.Pattern = core.StringPtr("testString")

				// Construct an instance of the TektonPipelineTriggerEvents model
				tektonPipelineTriggerEventsModel := new(opentoolchainv1.TektonPipelineTriggerEvents)
				tektonPipelineTriggerEventsModel.Push = core.BoolPtr(true)
				tektonPipelineTriggerEventsModel.PullRequest = core.BoolPtr(true)
				tektonPipelineTriggerEventsModel.PullRequestClosed = core.BoolPtr(true)

				// Construct an instance of the TektonPipelineTrigger model
				tektonPipelineTriggerModel := new(opentoolchainv1.TektonPipelineTrigger)
				tektonPipelineTriggerModel.ID = core.StringPtr("testString")
				tektonPipelineTriggerModel.Name = core.StringPtr("testString")
				tektonPipelineTriggerModel.EventListener = core.StringPtr("testString")
				tektonPipelineTriggerModel.Disabled = core.BoolPtr(true)
				tektonPipelineTriggerModel.ScmSource = tektonPipelineTriggerScmSourceModel
				tektonPipelineTriggerModel.Type = core.StringPtr("testString")
				tektonPipelineTriggerModel.Events = tektonPipelineTriggerEventsModel
				tektonPipelineTriggerModel.ServiceInstanceID = core.StringPtr("testString")

				// Construct an instance of the PatchTektonPipelineOptions model
				patchTektonPipelineOptionsModel := new(opentoolchainv1.PatchTektonPipelineOptions)
				patchTektonPipelineOptionsModel.GUID = core.StringPtr("testString")
				patchTektonPipelineOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				patchTektonPipelineOptionsModel.Worker = patchTektonPipelineParamsWorkerModel
				patchTektonPipelineOptionsModel.EnvProperties = []opentoolchainv1.EnvProperty{*envPropertyModel}
				patchTektonPipelineOptionsModel.Inputs = []opentoolchainv1.TektonPipelineInput{*tektonPipelineInputModel}
				patchTektonPipelineOptionsModel.Triggers = []opentoolchainv1.TektonPipelineTrigger{*tektonPipelineTriggerModel}
				patchTektonPipelineOptionsModel.PipelineDefinitionID = core.StringPtr("testString")
				patchTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := openToolchainService.PatchTektonPipelineWithContext(ctx, patchTektonPipelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				openToolchainService.DisableRetries()
				result, response, operationErr := openToolchainService.PatchTektonPipeline(patchTektonPipelineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = openToolchainService.PatchTektonPipelineWithContext(ctx, patchTektonPipelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(patchTektonPipelinePath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env_id"]).To(Equal([]string{"ibm:yp:us-south"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "dashboard_url": "DashboardURL", "resourceGroupId": "ResourceGroupID", "id": "ID", "toolchainId": "ToolchainID", "pipelineOwner": "PipelineOwner", "enabled": false, "type": "Type", "created": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "updated_at_timestamp": 18, "created_timestamp": 16, "envProperties": [{"name": "Name", "value": "Value", "type": "Type"}], "inputs": [{"type": "Type", "serviceInstanceId": "ServiceInstanceID", "shardDefinitionId": "ShardDefinitionID", "scmSource": {"path": "Path", "url": "URL", "type": "Type", "blindConnection": false, "branch": "Branch"}}], "triggers": [{"id": "ID", "name": "Name", "eventListener": "EventListener", "disabled": true, "scmSource": {"url": "URL", "type": "Type", "branch": "Branch", "pattern": "Pattern"}, "type": "Type", "events": {"push": true, "pull_request": false, "pull_request_closed": false}, "serviceInstanceId": "ServiceInstanceID"}], "status": "Status", "url": "URL", "runs_url": "RunsURL", "toolchainCRN": "ToolchainCRN", "pipelineDefinitionId": "PipelineDefinitionID"}`)
				}))
			})
			It(`Invoke PatchTektonPipeline successfully`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := openToolchainService.PatchTektonPipeline(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PatchTektonPipelineParamsWorker model
				patchTektonPipelineParamsWorkerModel := new(opentoolchainv1.PatchTektonPipelineParamsWorker)
				patchTektonPipelineParamsWorkerModel.WorkerID = core.StringPtr("testString")
				patchTektonPipelineParamsWorkerModel.WorkerName = core.StringPtr("testString")
				patchTektonPipelineParamsWorkerModel.WorkerType = core.StringPtr("testString")

				// Construct an instance of the EnvProperty model
				envPropertyModel := new(opentoolchainv1.EnvProperty)
				envPropertyModel.Name = core.StringPtr("testString")
				envPropertyModel.Value = core.StringPtr("testString")
				envPropertyModel.Type = core.StringPtr("testString")

				// Construct an instance of the TektonPipelineInputScmSource model
				tektonPipelineInputScmSourceModel := new(opentoolchainv1.TektonPipelineInputScmSource)
				tektonPipelineInputScmSourceModel.Path = core.StringPtr("testString")
				tektonPipelineInputScmSourceModel.URL = core.StringPtr("testString")
				tektonPipelineInputScmSourceModel.Type = core.StringPtr("testString")
				tektonPipelineInputScmSourceModel.BlindConnection = core.BoolPtr(true)
				tektonPipelineInputScmSourceModel.Branch = core.StringPtr("testString")

				// Construct an instance of the TektonPipelineInput model
				tektonPipelineInputModel := new(opentoolchainv1.TektonPipelineInput)
				tektonPipelineInputModel.Type = core.StringPtr("testString")
				tektonPipelineInputModel.ServiceInstanceID = core.StringPtr("testString")
				tektonPipelineInputModel.ShardDefinitionID = core.StringPtr("testString")
				tektonPipelineInputModel.ScmSource = tektonPipelineInputScmSourceModel

				// Construct an instance of the TektonPipelineTriggerScmSource model
				tektonPipelineTriggerScmSourceModel := new(opentoolchainv1.TektonPipelineTriggerScmSource)
				tektonPipelineTriggerScmSourceModel.URL = core.StringPtr("testString")
				tektonPipelineTriggerScmSourceModel.Type = core.StringPtr("testString")
				tektonPipelineTriggerScmSourceModel.Branch = core.StringPtr("testString")
				tektonPipelineTriggerScmSourceModel.Pattern = core.StringPtr("testString")

				// Construct an instance of the TektonPipelineTriggerEvents model
				tektonPipelineTriggerEventsModel := new(opentoolchainv1.TektonPipelineTriggerEvents)
				tektonPipelineTriggerEventsModel.Push = core.BoolPtr(true)
				tektonPipelineTriggerEventsModel.PullRequest = core.BoolPtr(true)
				tektonPipelineTriggerEventsModel.PullRequestClosed = core.BoolPtr(true)

				// Construct an instance of the TektonPipelineTrigger model
				tektonPipelineTriggerModel := new(opentoolchainv1.TektonPipelineTrigger)
				tektonPipelineTriggerModel.ID = core.StringPtr("testString")
				tektonPipelineTriggerModel.Name = core.StringPtr("testString")
				tektonPipelineTriggerModel.EventListener = core.StringPtr("testString")
				tektonPipelineTriggerModel.Disabled = core.BoolPtr(true)
				tektonPipelineTriggerModel.ScmSource = tektonPipelineTriggerScmSourceModel
				tektonPipelineTriggerModel.Type = core.StringPtr("testString")
				tektonPipelineTriggerModel.Events = tektonPipelineTriggerEventsModel
				tektonPipelineTriggerModel.ServiceInstanceID = core.StringPtr("testString")

				// Construct an instance of the PatchTektonPipelineOptions model
				patchTektonPipelineOptionsModel := new(opentoolchainv1.PatchTektonPipelineOptions)
				patchTektonPipelineOptionsModel.GUID = core.StringPtr("testString")
				patchTektonPipelineOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				patchTektonPipelineOptionsModel.Worker = patchTektonPipelineParamsWorkerModel
				patchTektonPipelineOptionsModel.EnvProperties = []opentoolchainv1.EnvProperty{*envPropertyModel}
				patchTektonPipelineOptionsModel.Inputs = []opentoolchainv1.TektonPipelineInput{*tektonPipelineInputModel}
				patchTektonPipelineOptionsModel.Triggers = []opentoolchainv1.TektonPipelineTrigger{*tektonPipelineTriggerModel}
				patchTektonPipelineOptionsModel.PipelineDefinitionID = core.StringPtr("testString")
				patchTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = openToolchainService.PatchTektonPipeline(patchTektonPipelineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PatchTektonPipeline with error: Operation validation and request error`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Construct an instance of the PatchTektonPipelineParamsWorker model
				patchTektonPipelineParamsWorkerModel := new(opentoolchainv1.PatchTektonPipelineParamsWorker)
				patchTektonPipelineParamsWorkerModel.WorkerID = core.StringPtr("testString")
				patchTektonPipelineParamsWorkerModel.WorkerName = core.StringPtr("testString")
				patchTektonPipelineParamsWorkerModel.WorkerType = core.StringPtr("testString")

				// Construct an instance of the EnvProperty model
				envPropertyModel := new(opentoolchainv1.EnvProperty)
				envPropertyModel.Name = core.StringPtr("testString")
				envPropertyModel.Value = core.StringPtr("testString")
				envPropertyModel.Type = core.StringPtr("testString")

				// Construct an instance of the TektonPipelineInputScmSource model
				tektonPipelineInputScmSourceModel := new(opentoolchainv1.TektonPipelineInputScmSource)
				tektonPipelineInputScmSourceModel.Path = core.StringPtr("testString")
				tektonPipelineInputScmSourceModel.URL = core.StringPtr("testString")
				tektonPipelineInputScmSourceModel.Type = core.StringPtr("testString")
				tektonPipelineInputScmSourceModel.BlindConnection = core.BoolPtr(true)
				tektonPipelineInputScmSourceModel.Branch = core.StringPtr("testString")

				// Construct an instance of the TektonPipelineInput model
				tektonPipelineInputModel := new(opentoolchainv1.TektonPipelineInput)
				tektonPipelineInputModel.Type = core.StringPtr("testString")
				tektonPipelineInputModel.ServiceInstanceID = core.StringPtr("testString")
				tektonPipelineInputModel.ShardDefinitionID = core.StringPtr("testString")
				tektonPipelineInputModel.ScmSource = tektonPipelineInputScmSourceModel

				// Construct an instance of the TektonPipelineTriggerScmSource model
				tektonPipelineTriggerScmSourceModel := new(opentoolchainv1.TektonPipelineTriggerScmSource)
				tektonPipelineTriggerScmSourceModel.URL = core.StringPtr("testString")
				tektonPipelineTriggerScmSourceModel.Type = core.StringPtr("testString")
				tektonPipelineTriggerScmSourceModel.Branch = core.StringPtr("testString")
				tektonPipelineTriggerScmSourceModel.Pattern = core.StringPtr("testString")

				// Construct an instance of the TektonPipelineTriggerEvents model
				tektonPipelineTriggerEventsModel := new(opentoolchainv1.TektonPipelineTriggerEvents)
				tektonPipelineTriggerEventsModel.Push = core.BoolPtr(true)
				tektonPipelineTriggerEventsModel.PullRequest = core.BoolPtr(true)
				tektonPipelineTriggerEventsModel.PullRequestClosed = core.BoolPtr(true)

				// Construct an instance of the TektonPipelineTrigger model
				tektonPipelineTriggerModel := new(opentoolchainv1.TektonPipelineTrigger)
				tektonPipelineTriggerModel.ID = core.StringPtr("testString")
				tektonPipelineTriggerModel.Name = core.StringPtr("testString")
				tektonPipelineTriggerModel.EventListener = core.StringPtr("testString")
				tektonPipelineTriggerModel.Disabled = core.BoolPtr(true)
				tektonPipelineTriggerModel.ScmSource = tektonPipelineTriggerScmSourceModel
				tektonPipelineTriggerModel.Type = core.StringPtr("testString")
				tektonPipelineTriggerModel.Events = tektonPipelineTriggerEventsModel
				tektonPipelineTriggerModel.ServiceInstanceID = core.StringPtr("testString")

				// Construct an instance of the PatchTektonPipelineOptions model
				patchTektonPipelineOptionsModel := new(opentoolchainv1.PatchTektonPipelineOptions)
				patchTektonPipelineOptionsModel.GUID = core.StringPtr("testString")
				patchTektonPipelineOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				patchTektonPipelineOptionsModel.Worker = patchTektonPipelineParamsWorkerModel
				patchTektonPipelineOptionsModel.EnvProperties = []opentoolchainv1.EnvProperty{*envPropertyModel}
				patchTektonPipelineOptionsModel.Inputs = []opentoolchainv1.TektonPipelineInput{*tektonPipelineInputModel}
				patchTektonPipelineOptionsModel.Triggers = []opentoolchainv1.TektonPipelineTrigger{*tektonPipelineTriggerModel}
				patchTektonPipelineOptionsModel.PipelineDefinitionID = core.StringPtr("testString")
				patchTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := openToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := openToolchainService.PatchTektonPipeline(patchTektonPipelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PatchTektonPipelineOptions model with no property values
				patchTektonPipelineOptionsModelNew := new(opentoolchainv1.PatchTektonPipelineOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = openToolchainService.PatchTektonPipeline(patchTektonPipelineOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke PatchTektonPipeline successfully`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Construct an instance of the PatchTektonPipelineParamsWorker model
				patchTektonPipelineParamsWorkerModel := new(opentoolchainv1.PatchTektonPipelineParamsWorker)
				patchTektonPipelineParamsWorkerModel.WorkerID = core.StringPtr("testString")
				patchTektonPipelineParamsWorkerModel.WorkerName = core.StringPtr("testString")
				patchTektonPipelineParamsWorkerModel.WorkerType = core.StringPtr("testString")

				// Construct an instance of the EnvProperty model
				envPropertyModel := new(opentoolchainv1.EnvProperty)
				envPropertyModel.Name = core.StringPtr("testString")
				envPropertyModel.Value = core.StringPtr("testString")
				envPropertyModel.Type = core.StringPtr("testString")

				// Construct an instance of the TektonPipelineInputScmSource model
				tektonPipelineInputScmSourceModel := new(opentoolchainv1.TektonPipelineInputScmSource)
				tektonPipelineInputScmSourceModel.Path = core.StringPtr("testString")
				tektonPipelineInputScmSourceModel.URL = core.StringPtr("testString")
				tektonPipelineInputScmSourceModel.Type = core.StringPtr("testString")
				tektonPipelineInputScmSourceModel.BlindConnection = core.BoolPtr(true)
				tektonPipelineInputScmSourceModel.Branch = core.StringPtr("testString")

				// Construct an instance of the TektonPipelineInput model
				tektonPipelineInputModel := new(opentoolchainv1.TektonPipelineInput)
				tektonPipelineInputModel.Type = core.StringPtr("testString")
				tektonPipelineInputModel.ServiceInstanceID = core.StringPtr("testString")
				tektonPipelineInputModel.ShardDefinitionID = core.StringPtr("testString")
				tektonPipelineInputModel.ScmSource = tektonPipelineInputScmSourceModel

				// Construct an instance of the TektonPipelineTriggerScmSource model
				tektonPipelineTriggerScmSourceModel := new(opentoolchainv1.TektonPipelineTriggerScmSource)
				tektonPipelineTriggerScmSourceModel.URL = core.StringPtr("testString")
				tektonPipelineTriggerScmSourceModel.Type = core.StringPtr("testString")
				tektonPipelineTriggerScmSourceModel.Branch = core.StringPtr("testString")
				tektonPipelineTriggerScmSourceModel.Pattern = core.StringPtr("testString")

				// Construct an instance of the TektonPipelineTriggerEvents model
				tektonPipelineTriggerEventsModel := new(opentoolchainv1.TektonPipelineTriggerEvents)
				tektonPipelineTriggerEventsModel.Push = core.BoolPtr(true)
				tektonPipelineTriggerEventsModel.PullRequest = core.BoolPtr(true)
				tektonPipelineTriggerEventsModel.PullRequestClosed = core.BoolPtr(true)

				// Construct an instance of the TektonPipelineTrigger model
				tektonPipelineTriggerModel := new(opentoolchainv1.TektonPipelineTrigger)
				tektonPipelineTriggerModel.ID = core.StringPtr("testString")
				tektonPipelineTriggerModel.Name = core.StringPtr("testString")
				tektonPipelineTriggerModel.EventListener = core.StringPtr("testString")
				tektonPipelineTriggerModel.Disabled = core.BoolPtr(true)
				tektonPipelineTriggerModel.ScmSource = tektonPipelineTriggerScmSourceModel
				tektonPipelineTriggerModel.Type = core.StringPtr("testString")
				tektonPipelineTriggerModel.Events = tektonPipelineTriggerEventsModel
				tektonPipelineTriggerModel.ServiceInstanceID = core.StringPtr("testString")

				// Construct an instance of the PatchTektonPipelineOptions model
				patchTektonPipelineOptionsModel := new(opentoolchainv1.PatchTektonPipelineOptions)
				patchTektonPipelineOptionsModel.GUID = core.StringPtr("testString")
				patchTektonPipelineOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				patchTektonPipelineOptionsModel.Worker = patchTektonPipelineParamsWorkerModel
				patchTektonPipelineOptionsModel.EnvProperties = []opentoolchainv1.EnvProperty{*envPropertyModel}
				patchTektonPipelineOptionsModel.Inputs = []opentoolchainv1.TektonPipelineInput{*tektonPipelineInputModel}
				patchTektonPipelineOptionsModel.Triggers = []opentoolchainv1.TektonPipelineTrigger{*tektonPipelineTriggerModel}
				patchTektonPipelineOptionsModel.PipelineDefinitionID = core.StringPtr("testString")
				patchTektonPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := openToolchainService.PatchTektonPipeline(patchTektonPipelineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTektonPipelineDefinition(getTektonPipelineDefinitionOptions *GetTektonPipelineDefinitionOptions) - Operation response error`, func() {
		getTektonPipelineDefinitionPath := "/devops/pipelines/tekton/api/v1/testString/definition"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelineDefinitionPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["env_id"]).To(Equal([]string{"ibm:yp:us-south"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTektonPipelineDefinition with error: Operation response processing error`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineDefinitionOptions model
				getTektonPipelineDefinitionOptionsModel := new(opentoolchainv1.GetTektonPipelineDefinitionOptions)
				getTektonPipelineDefinitionOptionsModel.GUID = core.StringPtr("testString")
				getTektonPipelineDefinitionOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				getTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := openToolchainService.GetTektonPipelineDefinition(getTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				openToolchainService.EnableRetries(0, 0)
				result, response, operationErr = openToolchainService.GetTektonPipelineDefinition(getTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTektonPipelineDefinition(getTektonPipelineDefinitionOptions *GetTektonPipelineDefinitionOptions)`, func() {
		getTektonPipelineDefinitionPath := "/devops/pipelines/tekton/api/v1/testString/definition"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelineDefinitionPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["env_id"]).To(Equal([]string{"ibm:yp:us-south"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"pipelineId": "PipelineID", "repoUrl": "RepoURL", "branch": "Branch", "path": "Path", "sha": "Sha", "type": "Type", "id": "ID", "shardRepos": [{"sha": "Sha", "shardDefinitionId": "ShardDefinitionID", "repoUrl": "RepoURL", "path": "Path"}]}`)
				}))
			})
			It(`Invoke GetTektonPipelineDefinition successfully with retries`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())
				openToolchainService.EnableRetries(0, 0)

				// Construct an instance of the GetTektonPipelineDefinitionOptions model
				getTektonPipelineDefinitionOptionsModel := new(opentoolchainv1.GetTektonPipelineDefinitionOptions)
				getTektonPipelineDefinitionOptionsModel.GUID = core.StringPtr("testString")
				getTektonPipelineDefinitionOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				getTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := openToolchainService.GetTektonPipelineDefinitionWithContext(ctx, getTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				openToolchainService.DisableRetries()
				result, response, operationErr := openToolchainService.GetTektonPipelineDefinition(getTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = openToolchainService.GetTektonPipelineDefinitionWithContext(ctx, getTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTektonPipelineDefinitionPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["env_id"]).To(Equal([]string{"ibm:yp:us-south"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"pipelineId": "PipelineID", "repoUrl": "RepoURL", "branch": "Branch", "path": "Path", "sha": "Sha", "type": "Type", "id": "ID", "shardRepos": [{"sha": "Sha", "shardDefinitionId": "ShardDefinitionID", "repoUrl": "RepoURL", "path": "Path"}]}`)
				}))
			})
			It(`Invoke GetTektonPipelineDefinition successfully`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := openToolchainService.GetTektonPipelineDefinition(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTektonPipelineDefinitionOptions model
				getTektonPipelineDefinitionOptionsModel := new(opentoolchainv1.GetTektonPipelineDefinitionOptions)
				getTektonPipelineDefinitionOptionsModel.GUID = core.StringPtr("testString")
				getTektonPipelineDefinitionOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				getTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = openToolchainService.GetTektonPipelineDefinition(getTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTektonPipelineDefinition with error: Operation validation and request error`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineDefinitionOptions model
				getTektonPipelineDefinitionOptionsModel := new(opentoolchainv1.GetTektonPipelineDefinitionOptions)
				getTektonPipelineDefinitionOptionsModel.GUID = core.StringPtr("testString")
				getTektonPipelineDefinitionOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				getTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := openToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := openToolchainService.GetTektonPipelineDefinition(getTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTektonPipelineDefinitionOptions model with no property values
				getTektonPipelineDefinitionOptionsModelNew := new(opentoolchainv1.GetTektonPipelineDefinitionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = openToolchainService.GetTektonPipelineDefinition(getTektonPipelineDefinitionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetTektonPipelineDefinition successfully`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Construct an instance of the GetTektonPipelineDefinitionOptions model
				getTektonPipelineDefinitionOptionsModel := new(opentoolchainv1.GetTektonPipelineDefinitionOptions)
				getTektonPipelineDefinitionOptionsModel.GUID = core.StringPtr("testString")
				getTektonPipelineDefinitionOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				getTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := openToolchainService.GetTektonPipelineDefinition(getTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTektonPipelineDefinition(createTektonPipelineDefinitionOptions *CreateTektonPipelineDefinitionOptions) - Operation response error`, func() {
		createTektonPipelineDefinitionPath := "/devops/pipelines/tekton/api/v1/testString/definition"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTektonPipelineDefinitionPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["env_id"]).To(Equal([]string{"ibm:yp:us-south"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTektonPipelineDefinition with error: Operation response processing error`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Construct an instance of the CreateTektonPipelineDefinitionParamsInputsItemScmSource model
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel := new(opentoolchainv1.CreateTektonPipelineDefinitionParamsInputsItemScmSource)
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.Path = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.URL = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.Type = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.BlindConnection = core.BoolPtr(true)
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.Branch = core.StringPtr("testString")

				// Construct an instance of the CreateTektonPipelineDefinitionParamsInputsItem model
				createTektonPipelineDefinitionParamsInputsItemModel := new(opentoolchainv1.CreateTektonPipelineDefinitionParamsInputsItem)
				createTektonPipelineDefinitionParamsInputsItemModel.ScmSource = createTektonPipelineDefinitionParamsInputsItemScmSourceModel
				createTektonPipelineDefinitionParamsInputsItemModel.Type = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemModel.ServiceInstanceID = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemModel.ShardDefinitionID = core.StringPtr("testString")

				// Construct an instance of the CreateTektonPipelineDefinitionOptions model
				createTektonPipelineDefinitionOptionsModel := new(opentoolchainv1.CreateTektonPipelineDefinitionOptions)
				createTektonPipelineDefinitionOptionsModel.GUID = core.StringPtr("testString")
				createTektonPipelineDefinitionOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				createTektonPipelineDefinitionOptionsModel.Inputs = []opentoolchainv1.CreateTektonPipelineDefinitionParamsInputsItem{*createTektonPipelineDefinitionParamsInputsItemModel}
				createTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := openToolchainService.CreateTektonPipelineDefinition(createTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				openToolchainService.EnableRetries(0, 0)
				result, response, operationErr = openToolchainService.CreateTektonPipelineDefinition(createTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTektonPipelineDefinition(createTektonPipelineDefinitionOptions *CreateTektonPipelineDefinitionOptions)`, func() {
		createTektonPipelineDefinitionPath := "/devops/pipelines/tekton/api/v1/testString/definition"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTektonPipelineDefinitionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env_id"]).To(Equal([]string{"ibm:yp:us-south"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"definition": {"pipelineId": "PipelineID", "repoUrl": "RepoURL", "branch": "Branch", "path": "Path", "sha": "Sha", "id": "ID"}, "inputs": [{"type": "Type", "serviceInstanceId": "ServiceInstanceID", "shardDefinitionId": "ShardDefinitionID", "scmSource": {"path": "Path", "url": "URL", "type": "Type", "blindConnection": false, "branch": "Branch"}}]}`)
				}))
			})
			It(`Invoke CreateTektonPipelineDefinition successfully with retries`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())
				openToolchainService.EnableRetries(0, 0)

				// Construct an instance of the CreateTektonPipelineDefinitionParamsInputsItemScmSource model
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel := new(opentoolchainv1.CreateTektonPipelineDefinitionParamsInputsItemScmSource)
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.Path = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.URL = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.Type = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.BlindConnection = core.BoolPtr(true)
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.Branch = core.StringPtr("testString")

				// Construct an instance of the CreateTektonPipelineDefinitionParamsInputsItem model
				createTektonPipelineDefinitionParamsInputsItemModel := new(opentoolchainv1.CreateTektonPipelineDefinitionParamsInputsItem)
				createTektonPipelineDefinitionParamsInputsItemModel.ScmSource = createTektonPipelineDefinitionParamsInputsItemScmSourceModel
				createTektonPipelineDefinitionParamsInputsItemModel.Type = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemModel.ServiceInstanceID = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemModel.ShardDefinitionID = core.StringPtr("testString")

				// Construct an instance of the CreateTektonPipelineDefinitionOptions model
				createTektonPipelineDefinitionOptionsModel := new(opentoolchainv1.CreateTektonPipelineDefinitionOptions)
				createTektonPipelineDefinitionOptionsModel.GUID = core.StringPtr("testString")
				createTektonPipelineDefinitionOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				createTektonPipelineDefinitionOptionsModel.Inputs = []opentoolchainv1.CreateTektonPipelineDefinitionParamsInputsItem{*createTektonPipelineDefinitionParamsInputsItemModel}
				createTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := openToolchainService.CreateTektonPipelineDefinitionWithContext(ctx, createTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				openToolchainService.DisableRetries()
				result, response, operationErr := openToolchainService.CreateTektonPipelineDefinition(createTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = openToolchainService.CreateTektonPipelineDefinitionWithContext(ctx, createTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTektonPipelineDefinitionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env_id"]).To(Equal([]string{"ibm:yp:us-south"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"definition": {"pipelineId": "PipelineID", "repoUrl": "RepoURL", "branch": "Branch", "path": "Path", "sha": "Sha", "id": "ID"}, "inputs": [{"type": "Type", "serviceInstanceId": "ServiceInstanceID", "shardDefinitionId": "ShardDefinitionID", "scmSource": {"path": "Path", "url": "URL", "type": "Type", "blindConnection": false, "branch": "Branch"}}]}`)
				}))
			})
			It(`Invoke CreateTektonPipelineDefinition successfully`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := openToolchainService.CreateTektonPipelineDefinition(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateTektonPipelineDefinitionParamsInputsItemScmSource model
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel := new(opentoolchainv1.CreateTektonPipelineDefinitionParamsInputsItemScmSource)
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.Path = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.URL = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.Type = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.BlindConnection = core.BoolPtr(true)
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.Branch = core.StringPtr("testString")

				// Construct an instance of the CreateTektonPipelineDefinitionParamsInputsItem model
				createTektonPipelineDefinitionParamsInputsItemModel := new(opentoolchainv1.CreateTektonPipelineDefinitionParamsInputsItem)
				createTektonPipelineDefinitionParamsInputsItemModel.ScmSource = createTektonPipelineDefinitionParamsInputsItemScmSourceModel
				createTektonPipelineDefinitionParamsInputsItemModel.Type = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemModel.ServiceInstanceID = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemModel.ShardDefinitionID = core.StringPtr("testString")

				// Construct an instance of the CreateTektonPipelineDefinitionOptions model
				createTektonPipelineDefinitionOptionsModel := new(opentoolchainv1.CreateTektonPipelineDefinitionOptions)
				createTektonPipelineDefinitionOptionsModel.GUID = core.StringPtr("testString")
				createTektonPipelineDefinitionOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				createTektonPipelineDefinitionOptionsModel.Inputs = []opentoolchainv1.CreateTektonPipelineDefinitionParamsInputsItem{*createTektonPipelineDefinitionParamsInputsItemModel}
				createTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = openToolchainService.CreateTektonPipelineDefinition(createTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTektonPipelineDefinition with error: Operation validation and request error`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Construct an instance of the CreateTektonPipelineDefinitionParamsInputsItemScmSource model
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel := new(opentoolchainv1.CreateTektonPipelineDefinitionParamsInputsItemScmSource)
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.Path = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.URL = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.Type = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.BlindConnection = core.BoolPtr(true)
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.Branch = core.StringPtr("testString")

				// Construct an instance of the CreateTektonPipelineDefinitionParamsInputsItem model
				createTektonPipelineDefinitionParamsInputsItemModel := new(opentoolchainv1.CreateTektonPipelineDefinitionParamsInputsItem)
				createTektonPipelineDefinitionParamsInputsItemModel.ScmSource = createTektonPipelineDefinitionParamsInputsItemScmSourceModel
				createTektonPipelineDefinitionParamsInputsItemModel.Type = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemModel.ServiceInstanceID = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemModel.ShardDefinitionID = core.StringPtr("testString")

				// Construct an instance of the CreateTektonPipelineDefinitionOptions model
				createTektonPipelineDefinitionOptionsModel := new(opentoolchainv1.CreateTektonPipelineDefinitionOptions)
				createTektonPipelineDefinitionOptionsModel.GUID = core.StringPtr("testString")
				createTektonPipelineDefinitionOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				createTektonPipelineDefinitionOptionsModel.Inputs = []opentoolchainv1.CreateTektonPipelineDefinitionParamsInputsItem{*createTektonPipelineDefinitionParamsInputsItemModel}
				createTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := openToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := openToolchainService.CreateTektonPipelineDefinition(createTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTektonPipelineDefinitionOptions model with no property values
				createTektonPipelineDefinitionOptionsModelNew := new(opentoolchainv1.CreateTektonPipelineDefinitionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = openToolchainService.CreateTektonPipelineDefinition(createTektonPipelineDefinitionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke CreateTektonPipelineDefinition successfully`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Construct an instance of the CreateTektonPipelineDefinitionParamsInputsItemScmSource model
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel := new(opentoolchainv1.CreateTektonPipelineDefinitionParamsInputsItemScmSource)
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.Path = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.URL = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.Type = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.BlindConnection = core.BoolPtr(true)
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.Branch = core.StringPtr("testString")

				// Construct an instance of the CreateTektonPipelineDefinitionParamsInputsItem model
				createTektonPipelineDefinitionParamsInputsItemModel := new(opentoolchainv1.CreateTektonPipelineDefinitionParamsInputsItem)
				createTektonPipelineDefinitionParamsInputsItemModel.ScmSource = createTektonPipelineDefinitionParamsInputsItemScmSourceModel
				createTektonPipelineDefinitionParamsInputsItemModel.Type = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemModel.ServiceInstanceID = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemModel.ShardDefinitionID = core.StringPtr("testString")

				// Construct an instance of the CreateTektonPipelineDefinitionOptions model
				createTektonPipelineDefinitionOptionsModel := new(opentoolchainv1.CreateTektonPipelineDefinitionOptions)
				createTektonPipelineDefinitionOptionsModel.GUID = core.StringPtr("testString")
				createTektonPipelineDefinitionOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				createTektonPipelineDefinitionOptionsModel.Inputs = []opentoolchainv1.CreateTektonPipelineDefinitionParamsInputsItem{*createTektonPipelineDefinitionParamsInputsItemModel}
				createTektonPipelineDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := openToolchainService.CreateTektonPipelineDefinition(createTektonPipelineDefinitionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetToolchain(getToolchainOptions *GetToolchainOptions) - Operation response error`, func() {
		getToolchainPath := "/devops/toolchains/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getToolchainPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["env_id"]).To(Equal([]string{"ibm:yp:us-south"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetToolchain with error: Operation response processing error`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Construct an instance of the GetToolchainOptions model
				getToolchainOptionsModel := new(opentoolchainv1.GetToolchainOptions)
				getToolchainOptionsModel.GUID = core.StringPtr("testString")
				getToolchainOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				getToolchainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := openToolchainService.GetToolchain(getToolchainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				openToolchainService.EnableRetries(0, 0)
				result, response, operationErr = openToolchainService.GetToolchain(getToolchainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetToolchain(getToolchainOptions *GetToolchainOptions)`, func() {
		getToolchainPath := "/devops/toolchains/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getToolchainPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["env_id"]).To(Equal([]string{"ibm:yp:us-south"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"toolchain_guid": "ToolchainGUID", "name": "Name", "description": "Description", "key": "Key", "container": {"guid": "GUID", "type": "Type"}, "crn": "CRN", "created": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "creator": "Creator", "generator": "Generator", "template": {"getting_started": "GettingStarted", "services_total": 13, "name": "Name", "type": "Type", "url": "URL", "source": "Source", "locale": "Locale"}, "tags": ["Tags"], "lifecycle_messaging_webhook_id": "LifecycleMessagingWebhookID", "region_id": "RegionID", "services": [{"broker_id": "BrokerID", "service_id": "ServiceID", "container": {"guid": "GUID", "type": "Type"}, "updated_at": "2019-01-01T12:00:00.000Z", "parameters": {"mapKey": "anyValue"}, "status": {"state": "State"}, "dashboard_url": "DashboardURL", "region_id": "RegionID", "instance_id": "InstanceID", "description": "Description", "tags": ["Tags"], "url": "URL", "toolchain_binding": {"status": {"state": "State"}, "name": "Name", "webhook_id": "WebhookID"}}]}`)
				}))
			})
			It(`Invoke GetToolchain successfully with retries`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())
				openToolchainService.EnableRetries(0, 0)

				// Construct an instance of the GetToolchainOptions model
				getToolchainOptionsModel := new(opentoolchainv1.GetToolchainOptions)
				getToolchainOptionsModel.GUID = core.StringPtr("testString")
				getToolchainOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				getToolchainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := openToolchainService.GetToolchainWithContext(ctx, getToolchainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				openToolchainService.DisableRetries()
				result, response, operationErr := openToolchainService.GetToolchain(getToolchainOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = openToolchainService.GetToolchainWithContext(ctx, getToolchainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getToolchainPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["env_id"]).To(Equal([]string{"ibm:yp:us-south"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"toolchain_guid": "ToolchainGUID", "name": "Name", "description": "Description", "key": "Key", "container": {"guid": "GUID", "type": "Type"}, "crn": "CRN", "created": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "creator": "Creator", "generator": "Generator", "template": {"getting_started": "GettingStarted", "services_total": 13, "name": "Name", "type": "Type", "url": "URL", "source": "Source", "locale": "Locale"}, "tags": ["Tags"], "lifecycle_messaging_webhook_id": "LifecycleMessagingWebhookID", "region_id": "RegionID", "services": [{"broker_id": "BrokerID", "service_id": "ServiceID", "container": {"guid": "GUID", "type": "Type"}, "updated_at": "2019-01-01T12:00:00.000Z", "parameters": {"mapKey": "anyValue"}, "status": {"state": "State"}, "dashboard_url": "DashboardURL", "region_id": "RegionID", "instance_id": "InstanceID", "description": "Description", "tags": ["Tags"], "url": "URL", "toolchain_binding": {"status": {"state": "State"}, "name": "Name", "webhook_id": "WebhookID"}}]}`)
				}))
			})
			It(`Invoke GetToolchain successfully`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := openToolchainService.GetToolchain(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetToolchainOptions model
				getToolchainOptionsModel := new(opentoolchainv1.GetToolchainOptions)
				getToolchainOptionsModel.GUID = core.StringPtr("testString")
				getToolchainOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				getToolchainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = openToolchainService.GetToolchain(getToolchainOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetToolchain with error: Operation validation and request error`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Construct an instance of the GetToolchainOptions model
				getToolchainOptionsModel := new(opentoolchainv1.GetToolchainOptions)
				getToolchainOptionsModel.GUID = core.StringPtr("testString")
				getToolchainOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				getToolchainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := openToolchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := openToolchainService.GetToolchain(getToolchainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetToolchainOptions model with no property values
				getToolchainOptionsModelNew := new(opentoolchainv1.GetToolchainOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = openToolchainService.GetToolchain(getToolchainOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetToolchain successfully`, func() {
				openToolchainService, serviceErr := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openToolchainService).ToNot(BeNil())

				// Construct an instance of the GetToolchainOptions model
				getToolchainOptionsModel := new(opentoolchainv1.GetToolchainOptions)
				getToolchainOptionsModel.GUID = core.StringPtr("testString")
				getToolchainOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				getToolchainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := openToolchainService.GetToolchain(getToolchainOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			openToolchainService, _ := opentoolchainv1.NewOpenToolchainV1(&opentoolchainv1.OpenToolchainV1Options{
				URL:           "http://opentoolchainv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateServiceInstanceOptions successfully`, func() {
				// Construct an instance of the CreateServiceInstanceParamsParameters model
				createServiceInstanceParamsParametersModel := new(opentoolchainv1.CreateServiceInstanceParamsParameters)
				Expect(createServiceInstanceParamsParametersModel).ToNot(BeNil())
				createServiceInstanceParamsParametersModel.APIKey = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ServiceKey = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.KeyType = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ServiceID = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ServiceName = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ServiceURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.UserEmail = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.UserPhone = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.Authorized = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ChannelName = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.TeamURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.PipelineStart = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.PipelineSuccess = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.PipelineFail = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.ToolchainBind = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.ToolchainUnbind = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.APIToken = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.GitID = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.APIRootURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.Legal = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.RepoURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.TokenURL = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.PrivateRepo = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.HasIssues = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.InstanceName = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.IntegrationStatus = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.Region = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.ResourceGroup = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.EnableTraceability = core.BoolPtr(true)
				createServiceInstanceParamsParametersModel.Name = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.Type = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.UIPipeline = core.BoolPtr(true)
				Expect(createServiceInstanceParamsParametersModel.APIKey).To(Equal(core.StringPtr("testString")))
				Expect(createServiceInstanceParamsParametersModel.ServiceKey).To(Equal(core.StringPtr("testString")))
				Expect(createServiceInstanceParamsParametersModel.KeyType).To(Equal(core.StringPtr("testString")))
				Expect(createServiceInstanceParamsParametersModel.ServiceID).To(Equal(core.StringPtr("testString")))
				Expect(createServiceInstanceParamsParametersModel.ServiceName).To(Equal(core.StringPtr("testString")))
				Expect(createServiceInstanceParamsParametersModel.ServiceURL).To(Equal(core.StringPtr("testString")))
				Expect(createServiceInstanceParamsParametersModel.UserEmail).To(Equal(core.StringPtr("testString")))
				Expect(createServiceInstanceParamsParametersModel.UserPhone).To(Equal(core.StringPtr("testString")))
				Expect(createServiceInstanceParamsParametersModel.Authorized).To(Equal(core.StringPtr("testString")))
				Expect(createServiceInstanceParamsParametersModel.ChannelName).To(Equal(core.StringPtr("testString")))
				Expect(createServiceInstanceParamsParametersModel.TeamURL).To(Equal(core.StringPtr("testString")))
				Expect(createServiceInstanceParamsParametersModel.PipelineStart).To(Equal(core.BoolPtr(true)))
				Expect(createServiceInstanceParamsParametersModel.PipelineSuccess).To(Equal(core.BoolPtr(true)))
				Expect(createServiceInstanceParamsParametersModel.PipelineFail).To(Equal(core.BoolPtr(true)))
				Expect(createServiceInstanceParamsParametersModel.ToolchainBind).To(Equal(core.BoolPtr(true)))
				Expect(createServiceInstanceParamsParametersModel.ToolchainUnbind).To(Equal(core.BoolPtr(true)))
				Expect(createServiceInstanceParamsParametersModel.APIToken).To(Equal(core.StringPtr("testString")))
				Expect(createServiceInstanceParamsParametersModel.GitID).To(Equal(core.StringPtr("testString")))
				Expect(createServiceInstanceParamsParametersModel.APIRootURL).To(Equal(core.StringPtr("testString")))
				Expect(createServiceInstanceParamsParametersModel.Legal).To(Equal(core.BoolPtr(true)))
				Expect(createServiceInstanceParamsParametersModel.RepoURL).To(Equal(core.StringPtr("testString")))
				Expect(createServiceInstanceParamsParametersModel.TokenURL).To(Equal(core.StringPtr("testString")))
				Expect(createServiceInstanceParamsParametersModel.PrivateRepo).To(Equal(core.BoolPtr(true)))
				Expect(createServiceInstanceParamsParametersModel.HasIssues).To(Equal(core.BoolPtr(true)))
				Expect(createServiceInstanceParamsParametersModel.InstanceName).To(Equal(core.StringPtr("testString")))
				Expect(createServiceInstanceParamsParametersModel.IntegrationStatus).To(Equal(core.StringPtr("testString")))
				Expect(createServiceInstanceParamsParametersModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(createServiceInstanceParamsParametersModel.ResourceGroup).To(Equal(core.StringPtr("testString")))
				Expect(createServiceInstanceParamsParametersModel.EnableTraceability).To(Equal(core.BoolPtr(true)))
				Expect(createServiceInstanceParamsParametersModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createServiceInstanceParamsParametersModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(createServiceInstanceParamsParametersModel.UIPipeline).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the CreateServiceInstanceOptions model
				envID := "ibm:yp:us-south"
				createServiceInstanceOptionsModel := openToolchainService.NewCreateServiceInstanceOptions(envID)
				createServiceInstanceOptionsModel.SetEnvID("ibm:yp:us-south")
				createServiceInstanceOptionsModel.SetToolchainID("testString")
				createServiceInstanceOptionsModel.SetServiceID("testString")
				createServiceInstanceOptionsModel.SetParameters(createServiceInstanceParamsParametersModel)
				createServiceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createServiceInstanceOptionsModel).ToNot(BeNil())
				Expect(createServiceInstanceOptionsModel.EnvID).To(Equal(core.StringPtr("ibm:yp:us-south")))
				Expect(createServiceInstanceOptionsModel.ToolchainID).To(Equal(core.StringPtr("testString")))
				Expect(createServiceInstanceOptionsModel.ServiceID).To(Equal(core.StringPtr("testString")))
				Expect(createServiceInstanceOptionsModel.Parameters).To(Equal(createServiceInstanceParamsParametersModel))
				Expect(createServiceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTektonPipelineDefinitionOptions successfully`, func() {
				// Construct an instance of the CreateTektonPipelineDefinitionParamsInputsItemScmSource model
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel := new(opentoolchainv1.CreateTektonPipelineDefinitionParamsInputsItemScmSource)
				Expect(createTektonPipelineDefinitionParamsInputsItemScmSourceModel).ToNot(BeNil())
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.Path = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.URL = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.Type = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.BlindConnection = core.BoolPtr(true)
				createTektonPipelineDefinitionParamsInputsItemScmSourceModel.Branch = core.StringPtr("testString")
				Expect(createTektonPipelineDefinitionParamsInputsItemScmSourceModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(createTektonPipelineDefinitionParamsInputsItemScmSourceModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(createTektonPipelineDefinitionParamsInputsItemScmSourceModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(createTektonPipelineDefinitionParamsInputsItemScmSourceModel.BlindConnection).To(Equal(core.BoolPtr(true)))
				Expect(createTektonPipelineDefinitionParamsInputsItemScmSourceModel.Branch).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateTektonPipelineDefinitionParamsInputsItem model
				createTektonPipelineDefinitionParamsInputsItemModel := new(opentoolchainv1.CreateTektonPipelineDefinitionParamsInputsItem)
				Expect(createTektonPipelineDefinitionParamsInputsItemModel).ToNot(BeNil())
				createTektonPipelineDefinitionParamsInputsItemModel.ScmSource = createTektonPipelineDefinitionParamsInputsItemScmSourceModel
				createTektonPipelineDefinitionParamsInputsItemModel.Type = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemModel.ServiceInstanceID = core.StringPtr("testString")
				createTektonPipelineDefinitionParamsInputsItemModel.ShardDefinitionID = core.StringPtr("testString")
				Expect(createTektonPipelineDefinitionParamsInputsItemModel.ScmSource).To(Equal(createTektonPipelineDefinitionParamsInputsItemScmSourceModel))
				Expect(createTektonPipelineDefinitionParamsInputsItemModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(createTektonPipelineDefinitionParamsInputsItemModel.ServiceInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createTektonPipelineDefinitionParamsInputsItemModel.ShardDefinitionID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateTektonPipelineDefinitionOptions model
				guid := "testString"
				envID := "ibm:yp:us-south"
				createTektonPipelineDefinitionOptionsModel := openToolchainService.NewCreateTektonPipelineDefinitionOptions(guid, envID)
				createTektonPipelineDefinitionOptionsModel.SetGUID("testString")
				createTektonPipelineDefinitionOptionsModel.SetEnvID("ibm:yp:us-south")
				createTektonPipelineDefinitionOptionsModel.SetInputs([]opentoolchainv1.CreateTektonPipelineDefinitionParamsInputsItem{*createTektonPipelineDefinitionParamsInputsItemModel})
				createTektonPipelineDefinitionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTektonPipelineDefinitionOptionsModel).ToNot(BeNil())
				Expect(createTektonPipelineDefinitionOptionsModel.GUID).To(Equal(core.StringPtr("testString")))
				Expect(createTektonPipelineDefinitionOptionsModel.EnvID).To(Equal(core.StringPtr("ibm:yp:us-south")))
				Expect(createTektonPipelineDefinitionOptionsModel.Inputs).To(Equal([]opentoolchainv1.CreateTektonPipelineDefinitionParamsInputsItem{*createTektonPipelineDefinitionParamsInputsItemModel}))
				Expect(createTektonPipelineDefinitionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateToolchainOptions successfully`, func() {
				// Construct an instance of the CreateToolchainOptions model
				envID := "ibm:yp:us-south"
				repository := "testString"
				createToolchainOptionsModel := openToolchainService.NewCreateToolchainOptions(envID, repository)
				createToolchainOptionsModel.SetEnvID("ibm:yp:us-south")
				createToolchainOptionsModel.SetRepository("testString")
				createToolchainOptionsModel.SetAutocreate(true)
				createToolchainOptionsModel.SetResourceGroupID("testString")
				createToolchainOptionsModel.SetRepositoryToken("testString")
				createToolchainOptionsModel.SetBranch("testString")
				createToolchainOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createToolchainOptionsModel).ToNot(BeNil())
				Expect(createToolchainOptionsModel.EnvID).To(Equal(core.StringPtr("ibm:yp:us-south")))
				Expect(createToolchainOptionsModel.Repository).To(Equal(core.StringPtr("testString")))
				Expect(createToolchainOptionsModel.Autocreate).To(Equal(core.BoolPtr(true)))
				Expect(createToolchainOptionsModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(createToolchainOptionsModel.RepositoryToken).To(Equal(core.StringPtr("testString")))
				Expect(createToolchainOptionsModel.Branch).To(Equal(core.StringPtr("testString")))
				Expect(createToolchainOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteServiceInstanceOptions successfully`, func() {
				// Construct an instance of the DeleteServiceInstanceOptions model
				guid := "testString"
				envID := "ibm:yp:us-south"
				deleteServiceInstanceOptionsModel := openToolchainService.NewDeleteServiceInstanceOptions(guid, envID)
				deleteServiceInstanceOptionsModel.SetGUID("testString")
				deleteServiceInstanceOptionsModel.SetEnvID("ibm:yp:us-south")
				deleteServiceInstanceOptionsModel.SetToolchainID("testString")
				deleteServiceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteServiceInstanceOptionsModel).ToNot(BeNil())
				Expect(deleteServiceInstanceOptionsModel.GUID).To(Equal(core.StringPtr("testString")))
				Expect(deleteServiceInstanceOptionsModel.EnvID).To(Equal(core.StringPtr("ibm:yp:us-south")))
				Expect(deleteServiceInstanceOptionsModel.ToolchainID).To(Equal(core.StringPtr("testString")))
				Expect(deleteServiceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteToolchainOptions successfully`, func() {
				// Construct an instance of the DeleteToolchainOptions model
				guid := "testString"
				envID := "ibm:yp:us-south"
				deleteToolchainOptionsModel := openToolchainService.NewDeleteToolchainOptions(guid, envID)
				deleteToolchainOptionsModel.SetGUID("testString")
				deleteToolchainOptionsModel.SetEnvID("ibm:yp:us-south")
				deleteToolchainOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteToolchainOptionsModel).ToNot(BeNil())
				Expect(deleteToolchainOptionsModel.GUID).To(Equal(core.StringPtr("testString")))
				Expect(deleteToolchainOptionsModel.EnvID).To(Equal(core.StringPtr("ibm:yp:us-south")))
				Expect(deleteToolchainOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEnvProperty successfully`, func() {
				name := "testString"
				value := "testString"
				typeVar := "testString"
				model, err := openToolchainService.NewEnvProperty(name, value, typeVar)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGetServiceInstanceOptions successfully`, func() {
				// Construct an instance of the GetServiceInstanceOptions model
				guid := "testString"
				envID := "ibm:yp:us-south"
				toolchainID := "testString"
				getServiceInstanceOptionsModel := openToolchainService.NewGetServiceInstanceOptions(guid, envID, toolchainID)
				getServiceInstanceOptionsModel.SetGUID("testString")
				getServiceInstanceOptionsModel.SetEnvID("ibm:yp:us-south")
				getServiceInstanceOptionsModel.SetToolchainID("testString")
				getServiceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getServiceInstanceOptionsModel).ToNot(BeNil())
				Expect(getServiceInstanceOptionsModel.GUID).To(Equal(core.StringPtr("testString")))
				Expect(getServiceInstanceOptionsModel.EnvID).To(Equal(core.StringPtr("ibm:yp:us-south")))
				Expect(getServiceInstanceOptionsModel.ToolchainID).To(Equal(core.StringPtr("testString")))
				Expect(getServiceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTektonPipelineDefinitionOptions successfully`, func() {
				// Construct an instance of the GetTektonPipelineDefinitionOptions model
				guid := "testString"
				envID := "ibm:yp:us-south"
				getTektonPipelineDefinitionOptionsModel := openToolchainService.NewGetTektonPipelineDefinitionOptions(guid, envID)
				getTektonPipelineDefinitionOptionsModel.SetGUID("testString")
				getTektonPipelineDefinitionOptionsModel.SetEnvID("ibm:yp:us-south")
				getTektonPipelineDefinitionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTektonPipelineDefinitionOptionsModel).ToNot(BeNil())
				Expect(getTektonPipelineDefinitionOptionsModel.GUID).To(Equal(core.StringPtr("testString")))
				Expect(getTektonPipelineDefinitionOptionsModel.EnvID).To(Equal(core.StringPtr("ibm:yp:us-south")))
				Expect(getTektonPipelineDefinitionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTektonPipelineOptions successfully`, func() {
				// Construct an instance of the GetTektonPipelineOptions model
				guid := "testString"
				envID := "ibm:yp:us-south"
				getTektonPipelineOptionsModel := openToolchainService.NewGetTektonPipelineOptions(guid, envID)
				getTektonPipelineOptionsModel.SetGUID("testString")
				getTektonPipelineOptionsModel.SetEnvID("ibm:yp:us-south")
				getTektonPipelineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTektonPipelineOptionsModel).ToNot(BeNil())
				Expect(getTektonPipelineOptionsModel.GUID).To(Equal(core.StringPtr("testString")))
				Expect(getTektonPipelineOptionsModel.EnvID).To(Equal(core.StringPtr("ibm:yp:us-south")))
				Expect(getTektonPipelineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetToolchainOptions successfully`, func() {
				// Construct an instance of the GetToolchainOptions model
				guid := "testString"
				envID := "ibm:yp:us-south"
				getToolchainOptionsModel := openToolchainService.NewGetToolchainOptions(guid, envID)
				getToolchainOptionsModel.SetGUID("testString")
				getToolchainOptionsModel.SetEnvID("ibm:yp:us-south")
				getToolchainOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getToolchainOptionsModel).ToNot(BeNil())
				Expect(getToolchainOptionsModel.GUID).To(Equal(core.StringPtr("testString")))
				Expect(getToolchainOptionsModel.EnvID).To(Equal(core.StringPtr("ibm:yp:us-south")))
				Expect(getToolchainOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPatchServiceInstanceOptions successfully`, func() {
				// Construct an instance of the PatchServiceInstanceParamsParameters model
				patchServiceInstanceParamsParametersModel := new(opentoolchainv1.PatchServiceInstanceParamsParameters)
				Expect(patchServiceInstanceParamsParametersModel).ToNot(BeNil())
				patchServiceInstanceParamsParametersModel.APIKey = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.ServiceKey = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.KeyType = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.ServiceID = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.ServiceName = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.ServiceURL = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.UserEmail = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.UserPhone = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.Name = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.Type = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.UIPipeline = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.RepoURL = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.PrivateRepo = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.HasIssues = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.EnableTraceability = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.InstanceName = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.IntegrationStatus = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.Region = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.ResourceGroup = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.ChannelName = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.TeamURL = core.StringPtr("testString")
				patchServiceInstanceParamsParametersModel.PipelineStart = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.PipelineSuccess = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.PipelineFail = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.ToolchainBind = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.ToolchainUnbind = core.BoolPtr(true)
				patchServiceInstanceParamsParametersModel.APIToken = core.StringPtr("testString")
				Expect(patchServiceInstanceParamsParametersModel.APIKey).To(Equal(core.StringPtr("testString")))
				Expect(patchServiceInstanceParamsParametersModel.ServiceKey).To(Equal(core.StringPtr("testString")))
				Expect(patchServiceInstanceParamsParametersModel.KeyType).To(Equal(core.StringPtr("testString")))
				Expect(patchServiceInstanceParamsParametersModel.ServiceID).To(Equal(core.StringPtr("testString")))
				Expect(patchServiceInstanceParamsParametersModel.ServiceName).To(Equal(core.StringPtr("testString")))
				Expect(patchServiceInstanceParamsParametersModel.ServiceURL).To(Equal(core.StringPtr("testString")))
				Expect(patchServiceInstanceParamsParametersModel.UserEmail).To(Equal(core.StringPtr("testString")))
				Expect(patchServiceInstanceParamsParametersModel.UserPhone).To(Equal(core.StringPtr("testString")))
				Expect(patchServiceInstanceParamsParametersModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(patchServiceInstanceParamsParametersModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(patchServiceInstanceParamsParametersModel.UIPipeline).To(Equal(core.BoolPtr(true)))
				Expect(patchServiceInstanceParamsParametersModel.RepoURL).To(Equal(core.StringPtr("testString")))
				Expect(patchServiceInstanceParamsParametersModel.PrivateRepo).To(Equal(core.BoolPtr(true)))
				Expect(patchServiceInstanceParamsParametersModel.HasIssues).To(Equal(core.BoolPtr(true)))
				Expect(patchServiceInstanceParamsParametersModel.EnableTraceability).To(Equal(core.BoolPtr(true)))
				Expect(patchServiceInstanceParamsParametersModel.InstanceName).To(Equal(core.StringPtr("testString")))
				Expect(patchServiceInstanceParamsParametersModel.IntegrationStatus).To(Equal(core.StringPtr("testString")))
				Expect(patchServiceInstanceParamsParametersModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(patchServiceInstanceParamsParametersModel.ResourceGroup).To(Equal(core.StringPtr("testString")))
				Expect(patchServiceInstanceParamsParametersModel.ChannelName).To(Equal(core.StringPtr("testString")))
				Expect(patchServiceInstanceParamsParametersModel.TeamURL).To(Equal(core.StringPtr("testString")))
				Expect(patchServiceInstanceParamsParametersModel.PipelineStart).To(Equal(core.BoolPtr(true)))
				Expect(patchServiceInstanceParamsParametersModel.PipelineSuccess).To(Equal(core.BoolPtr(true)))
				Expect(patchServiceInstanceParamsParametersModel.PipelineFail).To(Equal(core.BoolPtr(true)))
				Expect(patchServiceInstanceParamsParametersModel.ToolchainBind).To(Equal(core.BoolPtr(true)))
				Expect(patchServiceInstanceParamsParametersModel.ToolchainUnbind).To(Equal(core.BoolPtr(true)))
				Expect(patchServiceInstanceParamsParametersModel.APIToken).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the PatchServiceInstanceOptions model
				guid := "testString"
				envID := "ibm:yp:us-south"
				patchServiceInstanceOptionsModel := openToolchainService.NewPatchServiceInstanceOptions(guid, envID)
				patchServiceInstanceOptionsModel.SetGUID("testString")
				patchServiceInstanceOptionsModel.SetEnvID("ibm:yp:us-south")
				patchServiceInstanceOptionsModel.SetToolchainID("testString")
				patchServiceInstanceOptionsModel.SetServiceID("testString")
				patchServiceInstanceOptionsModel.SetParameters(patchServiceInstanceParamsParametersModel)
				patchServiceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(patchServiceInstanceOptionsModel).ToNot(BeNil())
				Expect(patchServiceInstanceOptionsModel.GUID).To(Equal(core.StringPtr("testString")))
				Expect(patchServiceInstanceOptionsModel.EnvID).To(Equal(core.StringPtr("ibm:yp:us-south")))
				Expect(patchServiceInstanceOptionsModel.ToolchainID).To(Equal(core.StringPtr("testString")))
				Expect(patchServiceInstanceOptionsModel.ServiceID).To(Equal(core.StringPtr("testString")))
				Expect(patchServiceInstanceOptionsModel.Parameters).To(Equal(patchServiceInstanceParamsParametersModel))
				Expect(patchServiceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPatchTektonPipelineOptions successfully`, func() {
				// Construct an instance of the PatchTektonPipelineParamsWorker model
				patchTektonPipelineParamsWorkerModel := new(opentoolchainv1.PatchTektonPipelineParamsWorker)
				Expect(patchTektonPipelineParamsWorkerModel).ToNot(BeNil())
				patchTektonPipelineParamsWorkerModel.WorkerID = core.StringPtr("testString")
				patchTektonPipelineParamsWorkerModel.WorkerName = core.StringPtr("testString")
				patchTektonPipelineParamsWorkerModel.WorkerType = core.StringPtr("testString")
				Expect(patchTektonPipelineParamsWorkerModel.WorkerID).To(Equal(core.StringPtr("testString")))
				Expect(patchTektonPipelineParamsWorkerModel.WorkerName).To(Equal(core.StringPtr("testString")))
				Expect(patchTektonPipelineParamsWorkerModel.WorkerType).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the EnvProperty model
				envPropertyModel := new(opentoolchainv1.EnvProperty)
				Expect(envPropertyModel).ToNot(BeNil())
				envPropertyModel.Name = core.StringPtr("testString")
				envPropertyModel.Value = core.StringPtr("testString")
				envPropertyModel.Type = core.StringPtr("testString")
				Expect(envPropertyModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(envPropertyModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(envPropertyModel.Type).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the TektonPipelineInputScmSource model
				tektonPipelineInputScmSourceModel := new(opentoolchainv1.TektonPipelineInputScmSource)
				Expect(tektonPipelineInputScmSourceModel).ToNot(BeNil())
				tektonPipelineInputScmSourceModel.Path = core.StringPtr("testString")
				tektonPipelineInputScmSourceModel.URL = core.StringPtr("testString")
				tektonPipelineInputScmSourceModel.Type = core.StringPtr("testString")
				tektonPipelineInputScmSourceModel.BlindConnection = core.BoolPtr(true)
				tektonPipelineInputScmSourceModel.Branch = core.StringPtr("testString")
				Expect(tektonPipelineInputScmSourceModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(tektonPipelineInputScmSourceModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(tektonPipelineInputScmSourceModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(tektonPipelineInputScmSourceModel.BlindConnection).To(Equal(core.BoolPtr(true)))
				Expect(tektonPipelineInputScmSourceModel.Branch).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the TektonPipelineInput model
				tektonPipelineInputModel := new(opentoolchainv1.TektonPipelineInput)
				Expect(tektonPipelineInputModel).ToNot(BeNil())
				tektonPipelineInputModel.Type = core.StringPtr("testString")
				tektonPipelineInputModel.ServiceInstanceID = core.StringPtr("testString")
				tektonPipelineInputModel.ShardDefinitionID = core.StringPtr("testString")
				tektonPipelineInputModel.ScmSource = tektonPipelineInputScmSourceModel
				Expect(tektonPipelineInputModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(tektonPipelineInputModel.ServiceInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(tektonPipelineInputModel.ShardDefinitionID).To(Equal(core.StringPtr("testString")))
				Expect(tektonPipelineInputModel.ScmSource).To(Equal(tektonPipelineInputScmSourceModel))

				// Construct an instance of the TektonPipelineTriggerScmSource model
				tektonPipelineTriggerScmSourceModel := new(opentoolchainv1.TektonPipelineTriggerScmSource)
				Expect(tektonPipelineTriggerScmSourceModel).ToNot(BeNil())
				tektonPipelineTriggerScmSourceModel.URL = core.StringPtr("testString")
				tektonPipelineTriggerScmSourceModel.Type = core.StringPtr("testString")
				tektonPipelineTriggerScmSourceModel.Branch = core.StringPtr("testString")
				tektonPipelineTriggerScmSourceModel.Pattern = core.StringPtr("testString")
				Expect(tektonPipelineTriggerScmSourceModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(tektonPipelineTriggerScmSourceModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(tektonPipelineTriggerScmSourceModel.Branch).To(Equal(core.StringPtr("testString")))
				Expect(tektonPipelineTriggerScmSourceModel.Pattern).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the TektonPipelineTriggerEvents model
				tektonPipelineTriggerEventsModel := new(opentoolchainv1.TektonPipelineTriggerEvents)
				Expect(tektonPipelineTriggerEventsModel).ToNot(BeNil())
				tektonPipelineTriggerEventsModel.Push = core.BoolPtr(true)
				tektonPipelineTriggerEventsModel.PullRequest = core.BoolPtr(true)
				tektonPipelineTriggerEventsModel.PullRequestClosed = core.BoolPtr(true)
				Expect(tektonPipelineTriggerEventsModel.Push).To(Equal(core.BoolPtr(true)))
				Expect(tektonPipelineTriggerEventsModel.PullRequest).To(Equal(core.BoolPtr(true)))
				Expect(tektonPipelineTriggerEventsModel.PullRequestClosed).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the TektonPipelineTrigger model
				tektonPipelineTriggerModel := new(opentoolchainv1.TektonPipelineTrigger)
				Expect(tektonPipelineTriggerModel).ToNot(BeNil())
				tektonPipelineTriggerModel.ID = core.StringPtr("testString")
				tektonPipelineTriggerModel.Name = core.StringPtr("testString")
				tektonPipelineTriggerModel.EventListener = core.StringPtr("testString")
				tektonPipelineTriggerModel.Disabled = core.BoolPtr(true)
				tektonPipelineTriggerModel.ScmSource = tektonPipelineTriggerScmSourceModel
				tektonPipelineTriggerModel.Type = core.StringPtr("testString")
				tektonPipelineTriggerModel.Events = tektonPipelineTriggerEventsModel
				tektonPipelineTriggerModel.ServiceInstanceID = core.StringPtr("testString")
				Expect(tektonPipelineTriggerModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(tektonPipelineTriggerModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(tektonPipelineTriggerModel.EventListener).To(Equal(core.StringPtr("testString")))
				Expect(tektonPipelineTriggerModel.Disabled).To(Equal(core.BoolPtr(true)))
				Expect(tektonPipelineTriggerModel.ScmSource).To(Equal(tektonPipelineTriggerScmSourceModel))
				Expect(tektonPipelineTriggerModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(tektonPipelineTriggerModel.Events).To(Equal(tektonPipelineTriggerEventsModel))
				Expect(tektonPipelineTriggerModel.ServiceInstanceID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the PatchTektonPipelineOptions model
				guid := "testString"
				envID := "ibm:yp:us-south"
				patchTektonPipelineOptionsModel := openToolchainService.NewPatchTektonPipelineOptions(guid, envID)
				patchTektonPipelineOptionsModel.SetGUID("testString")
				patchTektonPipelineOptionsModel.SetEnvID("ibm:yp:us-south")
				patchTektonPipelineOptionsModel.SetWorker(patchTektonPipelineParamsWorkerModel)
				patchTektonPipelineOptionsModel.SetEnvProperties([]opentoolchainv1.EnvProperty{*envPropertyModel})
				patchTektonPipelineOptionsModel.SetInputs([]opentoolchainv1.TektonPipelineInput{*tektonPipelineInputModel})
				patchTektonPipelineOptionsModel.SetTriggers([]opentoolchainv1.TektonPipelineTrigger{*tektonPipelineTriggerModel})
				patchTektonPipelineOptionsModel.SetPipelineDefinitionID("testString")
				patchTektonPipelineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(patchTektonPipelineOptionsModel).ToNot(BeNil())
				Expect(patchTektonPipelineOptionsModel.GUID).To(Equal(core.StringPtr("testString")))
				Expect(patchTektonPipelineOptionsModel.EnvID).To(Equal(core.StringPtr("ibm:yp:us-south")))
				Expect(patchTektonPipelineOptionsModel.Worker).To(Equal(patchTektonPipelineParamsWorkerModel))
				Expect(patchTektonPipelineOptionsModel.EnvProperties).To(Equal([]opentoolchainv1.EnvProperty{*envPropertyModel}))
				Expect(patchTektonPipelineOptionsModel.Inputs).To(Equal([]opentoolchainv1.TektonPipelineInput{*tektonPipelineInputModel}))
				Expect(patchTektonPipelineOptionsModel.Triggers).To(Equal([]opentoolchainv1.TektonPipelineTrigger{*tektonPipelineTriggerModel}))
				Expect(patchTektonPipelineOptionsModel.PipelineDefinitionID).To(Equal(core.StringPtr("testString")))
				Expect(patchTektonPipelineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPatchToolchainOptions successfully`, func() {
				// Construct an instance of the PatchToolchainOptions model
				guid := "testString"
				envID := "ibm:yp:us-south"
				patchToolchainOptionsModel := openToolchainService.NewPatchToolchainOptions(guid, envID)
				patchToolchainOptionsModel.SetGUID("testString")
				patchToolchainOptionsModel.SetEnvID("ibm:yp:us-south")
				patchToolchainOptionsModel.SetName("testString")
				patchToolchainOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(patchToolchainOptionsModel).ToNot(BeNil())
				Expect(patchToolchainOptionsModel.GUID).To(Equal(core.StringPtr("testString")))
				Expect(patchToolchainOptionsModel.EnvID).To(Equal(core.StringPtr("ibm:yp:us-south")))
				Expect(patchToolchainOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(patchToolchainOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewTektonPipelineTrigger successfully`, func() {
				eventListener := "testString"
				typeVar := "testString"
				model, err := openToolchainService.NewTektonPipelineTrigger(eventListener, typeVar)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
