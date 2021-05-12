package sdk

/*
   Copyright 2016 Alexander I.Grafov <grafov@gmail.com>
   Copyright 2016-2019 The Grafana SDK authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

	   http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

   ॐ तारे तुत्तारे तुरे स्व
*/

type Team struct {
	ID          uint   `json:"id"`
	OrgID       uint   `json:"orgId"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	AvatarUrl   string `json:"avatarUrl"`
	MemberCount int    `json:"memberCount"`
}

type TeamMsg struct {
	Message string `json:"message"`
	TeamID  uint   `json:"teamId"`
}

type PageTeams struct {
	TotalCount int    `json:"totalCount"`
	Teams      []Team `json:"teams"`
	Page       int    `json:"page"`
	PerPage    int    `json:"perPage"`
}
