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
					fmt.Fprintf(res, "%s", `{"name": "Name", "id": "ID", "toolchainId": "ToolchainID", "pipelineOwner": "PipelineOwner", "enabled": false, "type": "Type", "created": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "updated_at_timestamp": 18, "created_timestamp": 16, "envProperties": [{"name": "Name", "value": "Value", "type": "Type"}], "status": "Status", "url": "URL", "runs_url": "RunsURL", "toolchainCRN": "ToolchainCRN", "pipelineDefinitionId": "PipelineDefinitionID"}`)
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "id": "ID", "toolchainId": "ToolchainID", "pipelineOwner": "PipelineOwner", "enabled": false, "type": "Type", "created": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "updated_at_timestamp": 18, "created_timestamp": 16, "envProperties": [{"name": "Name", "value": "Value", "type": "Type"}], "status": "Status", "url": "URL", "runs_url": "RunsURL", "toolchainCRN": "ToolchainCRN", "pipelineDefinitionId": "PipelineDefinitionID"}`)
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

				// Construct an instance of the EnvProperty model
				envPropertyModel := new(opentoolchainv1.EnvProperty)
				envPropertyModel.Name = core.StringPtr("testString")
				envPropertyModel.Value = core.StringPtr("testString")
				envPropertyModel.Type = core.StringPtr("testString")

				// Construct an instance of the PatchTektonPipelineOptions model
				patchTektonPipelineOptionsModel := new(opentoolchainv1.PatchTektonPipelineOptions)
				patchTektonPipelineOptionsModel.GUID = core.StringPtr("testString")
				patchTektonPipelineOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				patchTektonPipelineOptionsModel.EnvProperties = []opentoolchainv1.EnvProperty{*envPropertyModel}
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "id": "ID", "toolchainId": "ToolchainID", "pipelineOwner": "PipelineOwner", "enabled": false, "type": "Type", "created": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "updated_at_timestamp": 18, "created_timestamp": 16, "envProperties": [{"name": "Name", "value": "Value", "type": "Type"}], "status": "Status", "url": "URL", "runs_url": "RunsURL", "toolchainCRN": "ToolchainCRN", "pipelineDefinitionId": "PipelineDefinitionID"}`)
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

				// Construct an instance of the EnvProperty model
				envPropertyModel := new(opentoolchainv1.EnvProperty)
				envPropertyModel.Name = core.StringPtr("testString")
				envPropertyModel.Value = core.StringPtr("testString")
				envPropertyModel.Type = core.StringPtr("testString")

				// Construct an instance of the PatchTektonPipelineOptions model
				patchTektonPipelineOptionsModel := new(opentoolchainv1.PatchTektonPipelineOptions)
				patchTektonPipelineOptionsModel.GUID = core.StringPtr("testString")
				patchTektonPipelineOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				patchTektonPipelineOptionsModel.EnvProperties = []opentoolchainv1.EnvProperty{*envPropertyModel}
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "id": "ID", "toolchainId": "ToolchainID", "pipelineOwner": "PipelineOwner", "enabled": false, "type": "Type", "created": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "updated_at_timestamp": 18, "created_timestamp": 16, "envProperties": [{"name": "Name", "value": "Value", "type": "Type"}], "status": "Status", "url": "URL", "runs_url": "RunsURL", "toolchainCRN": "ToolchainCRN", "pipelineDefinitionId": "PipelineDefinitionID"}`)
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

				// Construct an instance of the EnvProperty model
				envPropertyModel := new(opentoolchainv1.EnvProperty)
				envPropertyModel.Name = core.StringPtr("testString")
				envPropertyModel.Value = core.StringPtr("testString")
				envPropertyModel.Type = core.StringPtr("testString")

				// Construct an instance of the PatchTektonPipelineOptions model
				patchTektonPipelineOptionsModel := new(opentoolchainv1.PatchTektonPipelineOptions)
				patchTektonPipelineOptionsModel.GUID = core.StringPtr("testString")
				patchTektonPipelineOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				patchTektonPipelineOptionsModel.EnvProperties = []opentoolchainv1.EnvProperty{*envPropertyModel}
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

				// Construct an instance of the EnvProperty model
				envPropertyModel := new(opentoolchainv1.EnvProperty)
				envPropertyModel.Name = core.StringPtr("testString")
				envPropertyModel.Value = core.StringPtr("testString")
				envPropertyModel.Type = core.StringPtr("testString")

				// Construct an instance of the PatchTektonPipelineOptions model
				patchTektonPipelineOptionsModel := new(opentoolchainv1.PatchTektonPipelineOptions)
				patchTektonPipelineOptionsModel.GUID = core.StringPtr("testString")
				patchTektonPipelineOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				patchTektonPipelineOptionsModel.EnvProperties = []opentoolchainv1.EnvProperty{*envPropertyModel}
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

				// Construct an instance of the EnvProperty model
				envPropertyModel := new(opentoolchainv1.EnvProperty)
				envPropertyModel.Name = core.StringPtr("testString")
				envPropertyModel.Value = core.StringPtr("testString")
				envPropertyModel.Type = core.StringPtr("testString")

				// Construct an instance of the PatchTektonPipelineOptions model
				patchTektonPipelineOptionsModel := new(opentoolchainv1.PatchTektonPipelineOptions)
				patchTektonPipelineOptionsModel.GUID = core.StringPtr("testString")
				patchTektonPipelineOptionsModel.EnvID = core.StringPtr("ibm:yp:us-south")
				patchTektonPipelineOptionsModel.EnvProperties = []opentoolchainv1.EnvProperty{*envPropertyModel}
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
				createServiceInstanceParamsParametersModel.Name = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.Type = core.StringPtr("testString")
				createServiceInstanceParamsParametersModel.UIPipeline = core.BoolPtr(true)
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
			It(`Invoke NewCreateServiceInstanceParamsParameters successfully`, func() {
				name := "testString"
				model, err := openToolchainService.NewCreateServiceInstanceParamsParameters(name)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
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
			It(`Invoke NewPatchTektonPipelineOptions successfully`, func() {
				// Construct an instance of the EnvProperty model
				envPropertyModel := new(opentoolchainv1.EnvProperty)
				Expect(envPropertyModel).ToNot(BeNil())
				envPropertyModel.Name = core.StringPtr("testString")
				envPropertyModel.Value = core.StringPtr("testString")
				envPropertyModel.Type = core.StringPtr("testString")
				Expect(envPropertyModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(envPropertyModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(envPropertyModel.Type).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the PatchTektonPipelineOptions model
				guid := "testString"
				envID := "ibm:yp:us-south"
				patchTektonPipelineOptionsModel := openToolchainService.NewPatchTektonPipelineOptions(guid, envID)
				patchTektonPipelineOptionsModel.SetGUID("testString")
				patchTektonPipelineOptionsModel.SetEnvID("ibm:yp:us-south")
				patchTektonPipelineOptionsModel.SetEnvProperties([]opentoolchainv1.EnvProperty{*envPropertyModel})
				patchTektonPipelineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(patchTektonPipelineOptionsModel).ToNot(BeNil())
				Expect(patchTektonPipelineOptionsModel.GUID).To(Equal(core.StringPtr("testString")))
				Expect(patchTektonPipelineOptionsModel.EnvID).To(Equal(core.StringPtr("ibm:yp:us-south")))
				Expect(patchTektonPipelineOptionsModel.EnvProperties).To(Equal([]opentoolchainv1.EnvProperty{*envPropertyModel}))
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
