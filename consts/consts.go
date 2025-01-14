/*
   Copyright 2020 The Compose Specification Authors.

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

package consts

const (
	ComposeProjectName   = "COMPOSE_PROJECT_NAME"
	ComposePathSeparator = "COMPOSE_PATH_SEPARATOR"
	ComposeFilePath      = "COMPOSE_FILE"
	ComposeProfiles      = "COMPOSE_PROFILES"
)

const Extensions = "#extensions" // Using # prefix, we prevent risk to conflict with an actual yaml key

type ComposeFileKey struct{}
