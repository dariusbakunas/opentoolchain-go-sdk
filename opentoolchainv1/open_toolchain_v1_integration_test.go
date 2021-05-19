// +build integration

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
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/dbakuna/opentoolchain-go-sdk/opentoolchainv1"
)

/**
 * This file contains an integration test for the opentoolchainv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`OpenToolchainV1 Integration Tests`, func() {

	const externalConfigFile = "../open_toolchain_v1.env"

	var (
		err          error
		openToolchainService *opentoolchainv1.OpenToolchainV1
		serviceURL   string
		config       map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(opentoolchainv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Printf("Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			openToolchainServiceOptions := &opentoolchainv1.OpenToolchainV1Options{}

			openToolchainService, err = opentoolchainv1.NewOpenToolchainV1UsingExternalConfig(openToolchainServiceOptions)

			Expect(err).To(BeNil())
			Expect(openToolchainService).ToNot(BeNil())
			Expect(openToolchainService.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`GetToolchain - Returns details about a particular toolchain`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetToolchain(getToolchainOptions *GetToolchainOptions)`, func() {

			getToolchainOptions := &opentoolchainv1.GetToolchainOptions{
				GUID: core.StringPtr("testString"),
				EnvID: core.StringPtr("ibm:yp:us-south"),
			}

			toolchain, response, err := openToolchainService.GetToolchain(getToolchainOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(toolchain).ToNot(BeNil())

		})
	})
})

//
// Utility functions are declared in the unit test file
//
