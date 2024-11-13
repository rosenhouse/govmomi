/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
Package vapi/crypto provides access to the Crypto Manager REST APIs that are not available in the SOAP API.
Currently for creating and deleting native providers only.
See the top-level package crypto for getting provider details via crypto.ManagerKmip.
See also: https://blogs.vmware.com/code/2023/07/30/automate-vsphere-native-key-providers/
*/
package crypto