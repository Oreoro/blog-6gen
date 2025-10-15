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

package migrations

import (
	"context"
	"time"

	"xorm.io/xorm"
)

// addHireFieldsToUser adds hire-related fields to user table
func addHireFieldsToUser(ctx context.Context, x *xorm.Engine) error {
	type User struct {
		HireMeEnabled     bool       `xorm:"not null default false BOOL hire_me_enabled"`
		HireRateCurrency  string     `xorm:"VARCHAR(16) hire_rate_currency"`
		HireRateAmount    int64      `xorm:"BIGINT(20) hire_rate_amount"`
		HireRateUnit      string     `xorm:"VARCHAR(16) hire_rate_unit"`
		HireContactEmail  string     `xorm:"VARCHAR(256) hire_contact_email"`
		HireContactURL    string     `xorm:"VARCHAR(1024) hire_contact_url"`
		HireNote          string     `xorm:"TEXT hire_note"`
		HireLastUpdatedAt *time.Time `xorm:"DATETIME hire_last_updated_at"`
	}
	return x.Context(ctx).Sync(new(User))
}

func addHireFields(ctx context.Context, x *xorm.Engine) error {
	return addHireFieldsToUser(ctx, x)
}
