// util/util.go
package util

import (
	"encoding/json"
	"sort"

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

// ConvertTypesStringToStrings converts a slice of types.String to a slice of Go strings.
func ConvertTypesStringToStrings(input []types.String) []string {
	var result []string
	for _, s := range input {
		if !s.IsNull() && !s.IsUnknown() {
			result = append(result, s.ValueString())
		} else {
			result = append(result, "")
		}
	}
	return result
}

// ToTypesStringSlice converts a slice of Go strings to a slice of types.String.
func ToTypesStringSlice(items []string) []types.String {
	var result []types.String
	for _, s := range items {
		result = append(result, types.StringValue(s))
	}
	return result
}

// SafeDeref safely dereferences a *string, returning an empty string if nil.
func SafeDeref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// ConvertStringsToTypesString converts a slice of Go strings to a slice of types.String.
func ConvertStringsToTypesString(items []string) []types.String {
	var result []types.String
	for _, item := range items {
		result = append(result, types.StringValue(item))
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
