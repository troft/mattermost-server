// Copyright (c) 2016-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package app

import (
	"testing"

	"github.com/mattermost/mattermost-server/model"
)

func TestGetGroup(t *testing.T)      {}
func TestGetGroupsPage(t *testing.T) {}

func TestCreateGroup(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	id := model.NewId()
	group := &model.Group{
		DisplayName: "dn_" + id,
		Name:        "name" + id,
		Type:        model.GroupTypeLdap,
	}

	if _, err := th.App.CreateGroup(group); err != nil {
		t.Log(err)
		t.Fatal("Should create a new group")
	}

	if _, err := th.App.CreateGroup(group); err == nil {
		t.Fatal("Should not create a new group - group already exist")
	}
}

func TestPatchGroup(t *testing.T)         {}
func TestDeleteGroup(t *testing.T)        {}
func TestCreateGroupMember(t *testing.T)  {}
func TestDeleteGroupMember(t *testing.T)  {}
func TestGetGroupTeam(t *testing.T)       {}
func TestCreateGroupTeam(t *testing.T)    {}
func TestDeleteGroupTeam(t *testing.T)    {}
func TestGetGroupChannel(t *testing.T)    {}
func TestCreateGroupChannel(t *testing.T) {}
func TestDeleteGroupChannel(t *testing.T) {}
