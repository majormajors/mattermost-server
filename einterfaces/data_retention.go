// Copyright (c) 2016-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

// Code generated by build/interface_gen.go; DO NOT EDIT.

package einterfaces

import (
	"github.com/mattermost/mattermost-server/model"
)

type DataRetentionInterface interface {
	GetPolicy() (*model.DataRetentionPolicy, *model.AppError)
}
