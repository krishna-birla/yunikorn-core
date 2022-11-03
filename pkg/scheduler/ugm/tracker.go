/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package ugm

import (
	"github.com/apache/yunikorn-core/pkg/common/resources"
	"github.com/apache/yunikorn-core/pkg/common/security"
)

// Tracker Defines a set of interfaces to track and retrieve the user group resource usage
type Tracker interface {
	GetUserResources(user security.UserGroup) *resources.Resource
	GetGroupResources(group string) *resources.Resource

	GetUsersResources() map[string]*UserTracker
	GetGroupsResources() map[string]*GroupTracker

	IncreaseTrackedResource(queuePath, applicationID string, usage *resources.Resource, user security.UserGroup) error
	DecreaseTrackedResource(queuePath, applicationID string, usage *resources.Resource, user security.UserGroup, removeApp bool) error
}
