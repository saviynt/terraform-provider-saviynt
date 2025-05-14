// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package endpointsutil

var OwnerTypeMap = map[string]string{
	"":"",
	"0": "None",
	"1": "User", 
	"2": "Usergroup",
}

var RoleTypeMap = map[string]string{
	"":"",
	"0": "Node",
	"1": "Enabler",
	"2": "Transactional",
	"3": "Emergency Access",
	"4": "Enterprise",
	"5": "Application",
	"6": "Entitlement",
}

var RequestOptionMap = map[string]string{
	"0": "None",
	"1": "DropdownSingle",
	"2": "Table",
	"3": "TableOnlyAdd",
}

var RequiredMap = map[string]bool{
	"":false,
	"0": false,
	"1": true,
}

var ShowOnMap = map[string]string {
	"":"",
	"-1": "All",
	"0": "ShowOnServiceAccountRequest",
	"1": "ShowOnApplicationRequest",
}

func TranslateValue(input string, valueMap map[string]string) string {
	if input == "" {
		return ""
	}
	if val, ok := valueMap[input]; ok {
		return val
	}
	return input 
}