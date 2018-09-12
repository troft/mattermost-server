// Copyright (c) 2018-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package app

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

func (a *App) GetGroup(id string) (*model.Group, *model.AppError) {
	var result *store.LayeredStoreSupplierResult

	if result := <-a.Srv.Store.Group().Get(id); result.Err != nil {
		return nil, result.Err
	}

	return result.Data.(*model.Group), nil
}

func (a *App) GetGroupsPage(page int, perPage int) ([]*model.Group, *model.AppError) {
	var result *store.LayeredStoreSupplierResult

	if result := <-a.Srv.Store.Group().GetAllPage(page*perPage, perPage); result.Err != nil {
		return nil, result.Err
	}

	return result.Data.([]*model.Group), nil
}

func (a *App) CreateGroup(group *model.Group) (*model.Group, *model.AppError) {
	var result *store.LayeredStoreSupplierResult

	// Clear any user-provided values for trusted properties.
	group.Id = ""
	group.CreateAt = 0
	group.UpdateAt = 0
	group.DeleteAt = 0

	if result := <-a.Srv.Store.Group().Save(group); result.Err != nil {
		return nil, result.Err
	}

	return result.Data.(*model.Group), nil
}

func (a *App) PatchGroup(group *model.Group) (*model.Group, *model.AppError) {
	var result *store.LayeredStoreSupplierResult

	if result := <-a.Srv.Store.Group().Save(group); result.Err != nil {
		return nil, result.Err
	}

	return result.Data.(*model.Group), nil
}

func (a *App) DeleteGroup(groupID string) (*model.Group, *model.AppError) {
	var result *store.LayeredStoreSupplierResult

	if result := <-a.Srv.Store.Group().Delete(groupID); result.Err != nil {
		return nil, result.Err
	}

	return result.Data.(*model.Group), nil
}

func (a *App) CreateGroupMember(groupID string, userID string) (*model.GroupMember, *model.AppError) {
	var result *store.LayeredStoreSupplierResult

	if result := <-a.Srv.Store.Group().CreateMember(groupID, userID); result.Err != nil {
		return nil, result.Err
	}

	return result.Data.(*model.GroupMember), nil
}

func (a *App) DeleteGroupMember(groupID string, userID string) (*model.GroupMember, *model.AppError) {
	var result *store.LayeredStoreSupplierResult

	if result := <-a.Srv.Store.Group().DeleteMember(groupID, userID); result.Err != nil {
		return nil, result.Err
	}

	return result.Data.(*model.GroupMember), nil
}

func (a *App) GetGroupTeam(groupID string, teamID string) (*model.GroupTeam, *model.AppError) {
	var result *store.LayeredStoreSupplierResult

	if result := <-a.Srv.Store.Group().GetGroupTeam(groupID, teamID); result.Err != nil {
		return nil, result.Err
	}

	return result.Data.(*model.GroupTeam), nil
}

func (a *App) CreateGroupTeam(groupTeam *model.GroupTeam) (*model.GroupTeam, *model.AppError) {
	var result *store.LayeredStoreSupplierResult

	// Clear any user-provided values for trusted properties.
	groupTeam.CreateAt = 0
	groupTeam.UpdateAt = 0
	groupTeam.DeleteAt = 0

	if result := <-a.Srv.Store.Group().SaveGroupTeam(groupTeam); result.Err != nil {
		return nil, result.Err
	}

	return result.Data.(*model.GroupTeam), nil
}

func (a *App) DeleteGroupTeam(groupID string, teamID string) (*model.GroupTeam, *model.AppError) {
	var result *store.LayeredStoreSupplierResult

	if result := <-a.Srv.Store.Group().DeleteGroupTeam(groupID, teamID); result.Err != nil {
		return nil, result.Err
	}

	return result.Data.(*model.GroupTeam), nil
}

func (a *App) GetGroupChannel(groupID string, channelID string) (*model.GroupChannel, *model.AppError) {
	var result *store.LayeredStoreSupplierResult

	if result := <-a.Srv.Store.Group().GetGroupChannel(groupID, channelID); result.Err != nil {
		return nil, result.Err
	}

	return result.Data.(*model.GroupChannel), nil
}

func (a *App) CreateGroupChannel(groupChannel *model.GroupChannel) (*model.GroupChannel, *model.AppError) {
	var result *store.LayeredStoreSupplierResult

	// Clear any user-provided values for trusted properties.
	groupChannel.CreateAt = 0
	groupChannel.UpdateAt = 0
	groupChannel.DeleteAt = 0

	if result := <-a.Srv.Store.Group().SaveGroupChannel(groupChannel); result.Err != nil {
		return nil, result.Err
	}

	return result.Data.(*model.GroupChannel), nil
}

func (a *App) DeleteGroupChannel(groupID string, channelID string) (*model.GroupChannel, *model.AppError) {
	var result *store.LayeredStoreSupplierResult

	if result := <-a.Srv.Store.Group().DeleteGroupChannel(groupID, channelID); result.Err != nil {
		return nil, result.Err
	}

	return result.Data.(*model.GroupChannel), nil
}
