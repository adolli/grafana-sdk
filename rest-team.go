package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

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

func (r *Client) SearchTeamsWithPaging(ctx context.Context, query *string, perpage, page *int) (PageTeams, error) {
	var (
		raw       []byte
		pageTeams PageTeams
		code      int
		err       error
	)

	var params url.Values = nil
	if perpage != nil && page != nil {
		if params == nil {
			params = url.Values{}
		}
		params["perpage"] = []string{fmt.Sprint(*perpage)}
		params["page"] = []string{fmt.Sprint(*page)}
	}

	if query != nil {
		if params == nil {
			params = url.Values{}
		}
		params["query"] = []string{*query}
	}

	if raw, code, err = r.get(ctx, "api/teams/search", params); err != nil {
		return pageTeams, err
	}

	if code != 200 {
		return pageTeams, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	if err := dec.Decode(&pageTeams); err != nil {
		return pageTeams, fmt.Errorf("unmarshal users: %s\n%s", err, raw)
	}
	return pageTeams, err
}

func (r *Client) GetTeam(ctx context.Context, name string) (Team, error) {
	var (
		raw       []byte
		pageTeams PageTeams
		code      int
		err       error
	)
	params := url.Values{}
	params["name"] = []string{name}
	if raw, code, err = r.get(ctx, "api/teams/search", params); err != nil {
		return Team{}, err
	}

	if code != 200 {
		return Team{}, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	if err := dec.Decode(&pageTeams); err != nil {
		return Team{}, fmt.Errorf("unmarshal users: %s\n%s", err, raw)
	}
	if pageTeams.TotalCount != 1 {
		return Team{}, fmt.Errorf("teams num != 1")
	}
	return pageTeams.Teams[0], err
}

func (r *Client) GetAllTeam(ctx context.Context) ([]Team, error) {
	query := ""
	perpage := 99999
	page := 1
	PageTeams, err := r.SearchTeamsWithPaging(ctx, &query, &perpage, &page)
	if err != nil {
		return nil, err
	}
	return PageTeams.Teams, nil
}

func (r *Client) GetUserTeams(ctx context.Context) {

}

func (r *Client) CreateTeam(ctx context.Context, team Team) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)
	if raw, err = json.Marshal(team); err != nil {
		return StatusMessage{}, err
	}
	if raw, _, err = r.post(ctx, "api/teams", nil, raw); err != nil {
		return StatusMessage{}, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
}

func (r *Client) DeleteTeam(ctx context.Context, id uint) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)
	if raw, _, err = r.delete(ctx, fmt.Sprintf("api/teams/%d", id)); err != nil {
		return StatusMessage{}, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
}

func (r *Client) AddTeamMember(ctx context.Context, tid, uid uint) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)
	if raw, err = json.Marshal(struct {
		UserID uint `json:"userId"`
	}{UserID: uid}); err != nil {
		return StatusMessage{}, err
	}
	if raw, _, err = r.post(ctx, fmt.Sprintf("api/teams/%d/members", tid), nil, raw); err != nil {
		return StatusMessage{}, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
}
