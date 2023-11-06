//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azappconfig

import "github.com/Azure/azure-sdk-for-go/sdk/data/azappconfig/internal/generated"

// SettingFields are fields to retrieve from a configuration setting.
type SettingFields = generated.KeyValueFields

const (
	// The primary identifier of a configuration setting.
	SettingFieldsKey SettingFields = generated.KeyValueFieldsKey

	// A label used to group configuration settings.
	SettingFieldsLabel SettingFields = generated.KeyValueFieldsLabel

	// The value of the configuration setting.
	SettingFieldsValue SettingFields = generated.KeyValueFieldsValue

	// The content type of the configuration setting's value.
	SettingFieldsContentType SettingFields = generated.KeyValueFieldsContentType

	// An ETag indicating the version of a configuration setting within a configuration store.
	SettingFieldsETag SettingFields = generated.KeyValueFieldsEtag

	// The last time a modifying operation was performed on the given configuration setting.
	SettingFieldsLastModified SettingFields = generated.KeyValueFieldsLastModified

	// A value indicating whether the configuration setting is read-only.
	SettingFieldsIsReadOnly SettingFields = generated.KeyValueFieldsLocked

	// A list of tags that can help identify what a configuration setting may be applicable for.
	SettingFieldsTags SettingFields = generated.KeyValueFieldsTags
)
