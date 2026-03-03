/*
OpenAPI Example
API version: 1.0.0
*/

package openapi

import (
	"encoding/json"
	"bytes"
)

var _ MappedNullable = &AllresponsebodyInnerFlags{}

type AllresponsebodyInnerFlags struct {
	Png string `json:"png,omitempty"`
	Svg string `json:"svg,omitempty"`
	Alt string `json:"alt,omitempty"`
}

type _AllresponsebodyInnerFlags AllresponsebodyInnerFlags

func NewAllresponsebodyInnerFlags(png string, svg string, alt string) *AllresponsebodyInnerFlags {
	return &AllresponsebodyInnerFlags{Png: png, Svg: svg, Alt: alt}
}

func NewAllresponsebodyInnerFlagsWithDefaults() *AllresponsebodyInnerFlags {
	return &AllresponsebodyInnerFlags{}
}

func (o *AllresponsebodyInnerFlags) GetPng() string {
	if o == nil { return "" }
	return o.Png
}
func (o *AllresponsebodyInnerFlags) GetPngOk() (*string, bool) {
	if o == nil { return nil, false }
	return &o.Png, true
}
func (o *AllresponsebodyInnerFlags) SetPng(v string) { o.Png = v }

func (o *AllresponsebodyInnerFlags) GetSvg() string {
	if o == nil { return "" }
	return o.Svg
}
func (o *AllresponsebodyInnerFlags) GetSvgOk() (*string, bool) {
	if o == nil { return nil, false }
	return &o.Svg, true
}
func (o *AllresponsebodyInnerFlags) SetSvg(v string) { o.Svg = v }

func (o *AllresponsebodyInnerFlags) GetAlt() string {
	if o == nil { return "" }
	return o.Alt
}
func (o *AllresponsebodyInnerFlags) GetAltOk() (*string, bool) {
	if o == nil { return nil, false }
	return &o.Alt, true
}
func (o *AllresponsebodyInnerFlags) SetAlt(v string) { o.Alt = v }

func (o AllresponsebodyInnerFlags) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil { return []byte{}, err }
	return json.Marshal(toSerialize)
}

func (o AllresponsebodyInnerFlags) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Png != "" { toSerialize["png"] = o.Png }
	if o.Svg != "" { toSerialize["svg"] = o.Svg }
	if o.Alt != "" { toSerialize["alt"] = o.Alt }
	return toSerialize, nil
}

func (o *AllresponsebodyInnerFlags) UnmarshalJSON(data []byte) (err error) {
	varFlags := _AllresponsebodyInnerFlags{}
	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varFlags)
	if err != nil { return err }
	*o = AllresponsebodyInnerFlags(varFlags)
	return nil
}

type NullableAllresponsebodyInnerFlags struct {
	value *AllresponsebodyInnerFlags
	isSet bool
}
func (v NullableAllresponsebodyInnerFlags) Get() *AllresponsebodyInnerFlags { return v.value }
func (v *NullableAllresponsebodyInnerFlags) Set(val *AllresponsebodyInnerFlags) { v.value = val; v.isSet = true }
func (v NullableAllresponsebodyInnerFlags) IsSet() bool { return v.isSet }
func (v *NullableAllresponsebodyInnerFlags) Unset() { v.value = nil; v.isSet = false }
func NewNullableAllresponsebodyInnerFlags(val *AllresponsebodyInnerFlags) *NullableAllresponsebodyInnerFlags {
	return &NullableAllresponsebodyInnerFlags{value: val, isSet: true}
}
func (v NullableAllresponsebodyInnerFlags) MarshalJSON() ([]byte, error) { return json.Marshal(v.value) }
func (v *NullableAllresponsebodyInnerFlags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
