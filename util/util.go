// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// util/util.go
package util

import (
	"encoding/json"
	"sort"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// safeString converts a *string to types.String safely.
func SafeString(s *string) types.String {
	if s == nil {
		return types.StringValue("")
	}
	return types.StringValue(*s)
}

func SafeBoolDatasource(b *bool) types.Bool {
	if b == nil {
		return types.BoolNull()
	}
	return types.BoolValue(*b)
}

func SafeStringDatasource(s *string) types.String {
	if s == nil {
		return types.StringNull()
	}
	return types.StringValue(*s)
}

// SafeDeref safely dereferences a *string, returning an empty string if nil.
func SafeDeref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// safeList converts a []string to a Terraform types.List.
func SafeList(items []string) (types.List, diag.Diagnostics) {
	if len(items) == 0 {
		return types.ListValueMust(types.StringType, []attr.Value{}), nil
	}

	var values []attr.Value
	for _, item := range items {
		values = append(values, types.StringValue(item))
	}

	return types.ListValue(types.StringType, values)
}

// ToTypesStringSlice converts a slice of Go strings to a slice of types.String.
func ToTypesStringSlice(items []string) []types.String {
	var result []types.String
	for _, s := range items {
		result = append(result, types.StringValue(s))
	}
	return result
}

func ConvertStringsToTFListString(items []string) types.List {
	var elements []attr.Value
	for _, item := range items {
		elements = append(elements, types.StringValue(item))
	}

	if len(elements) == 0 {
		return types.ListNull(types.StringType)
	}

	return types.ListValueMust(types.StringType, elements)
}

// ConvertStringsToTypesString converts a slice of Go strings to a slice of types.String.
func ConvertStringsToTypesString(items []string) []types.String {
	var result []types.String
	for _, item := range items {
		result = append(result, types.StringValue(item))
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

// marshalDeterministic marshals a map[string]string into a JSON string
// with keys sorted in lexicographical order.
func MarshalDeterministic(m map[string]string) (string, error) {
	// Get the keys and sort them.
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Build an ordered map slice.
	ordered := make(map[string]string, len(m))
	for _, k := range keys {
		ordered[k] = m[k]
	}

	// Marshal the ordered map.
	b, err := json.Marshal(ordered)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
func StringPtr(v string) *string {
	return &v
}

func SafeStringConnector(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func StringPointerOrEmpty(tfStr types.String) *string {
	if tfStr.IsNull() || tfStr.IsUnknown() {
		// Value is null, unknown, or empty — treat it as not set
		return nil
	}
	val := tfStr.ValueString()
	return &val
}

func ConvertTypesStringToStrings(input []string) []types.String {
	var result []types.String
	for _, v := range input {
		if v != "" {
			result = append(result, types.StringValue(v))
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func SanitizeTypesStringList(input []types.String) []types.String {
	var result []types.String
	for _, v := range input {
		if !v.IsNull() && !v.IsUnknown() && v.ValueString() != "" {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		return nil // or you can return []types.String{} if you prefer an empty list
	}
	return result
}

func ConvertTFStringsToGoStrings(input types.List) []string {
	if input.IsNull() || input.IsUnknown() {
		return nil
	}

	var result []string

	for _, val := range input.Elements() {
		strVal, ok := val.(types.String)
		if !ok || strVal.IsNull() || strVal.IsUnknown() {
			continue
		}
		result = append(result, strVal.ValueString())
	}

	if len(result) == 0 {
		return nil
	}

	return result
}

func NormalizeTFListString(list types.List) types.List {
	if list.IsNull() || list.IsUnknown() || len(list.Elements()) == 0 {
		return types.ListNull(types.StringType)
	}
	return list
}

func SafeInt32(ptr *int32) types.Int32 {
	if ptr == nil {
		return types.Int32Null()
	}
	return types.Int32Value(*ptr)
}

func SafeInt64[T int32 | int64](value *T) types.Int64 {
	if value == nil {
		return types.Int64Null()
	}
	return types.Int64Value(int64(*value))
}

func Int32PtrToTFString(val *int32) types.String {
	if val != nil {
		str := strconv.Itoa(int(*val))
		return types.StringValue(str)
	}
	return types.StringNull()
}
