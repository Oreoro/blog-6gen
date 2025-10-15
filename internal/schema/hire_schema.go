/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package schema

// PublicHireInfoResp represents public hire info
type PublicHireInfoResp struct {
	Enabled bool         `json:"enabled"`
	Rate    *HireRate    `json:"rate,omitempty"`
	Contact *HireContact `json:"contact,omitempty"`
	Note    string       `json:"note,omitempty"`
}

type HireRate struct {
	Currency string `json:"currency" validate:"omitempty,gt=0,lte=16"`
	Amount   int64  `json:"amount" validate:"omitempty,min=0"`
	Unit     string `json:"unit" validate:"omitempty,oneof=hour day project" enums:"hour,day,project"`
}

type HireContact struct {
	Email string `json:"email,omitempty" validate:"omitempty,email,lte=256"`
	URL   string `json:"url,omitempty" validate:"omitempty,url,lte=1024"`
}

// UpdateHireInfoReq for owner to update their hire info
type UpdateHireInfoReq struct {
	Enabled bool         `json:"enabled"`
	Rate    *HireRate    `json:"rate"`
	Contact *HireContact `json:"contact"`
	Note    string       `json:"note" validate:"omitempty,lte=2000,sanitizer"`
	UserID  string       `json:"-"`
}

// GetUserHireInfoReq for querying by username
type GetUserHireInfoReq struct {
	Username string `form:"username" validate:"required,gte=2,lte=30"`
}
