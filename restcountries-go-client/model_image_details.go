/*
OpenAPI Example
API version: 1.0.0
*/

package openapi

import (
	"bytes"
	"encoding/json"
)

var _ MappedNullable = &ImageDetails{}

type ImageDetails struct {
	Png string  `json:"png,omitempty"`
	Svg string  `json:"svg,omitempty"`
	Alt *string `json:"alt,omitempty"`
}

type _ImageDetails ImageDetails

func NewImageDetails(png string, svg string) *ImageDetails {
	return &ImageDetails{Png: png, Svg: svg}
}

func NewImageDetailsWithDefaults() *ImageDetails {
	return &ImageDetails{}
}

func (o *ImageDetails) GetPng() string {
	if o == nil {
		return ""
	}
	return o.Png
}
func (o *ImageDetails) GetPngOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Png, true
}
func (o *ImageDetails) SetPng(v string) { o.Png = v }

func (o *ImageDetails) GetSvg() string {
	if o == nil {
		return ""
	}
	return o.Svg
}
func (o *ImageDetails) GetSvgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Svg, true
}
func (o *ImageDetails) SetSvg(v string) { o.Svg = v }

func (o *ImageDetails) GetAlt() string {
	if o == nil || IsNil(o.Alt) {
		return ""
	}
	return *o.Alt
}
func (o *ImageDetails) GetAltOk() (*string, bool) {
	if o == nil || IsNil(o.Alt) {
		return nil, false
	}
	return o.Alt, true
}
func (o *ImageDetails) HasAlt() bool {
	return o != nil && !IsNil(o.Alt)
}
func (o *ImageDetails) SetAlt(v string) { o.Alt = &v }

func (o ImageDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Png != "" {
		toSerialize["png"] = o.Png
	}
	if o.Svg != "" {
		toSerialize["svg"] = o.Svg
	}
	if !IsNil(o.Alt) {
		toSerialize["alt"] = o.Alt
	}
	return toSerialize, nil
}

func (o *ImageDetails) UnmarshalJSON(data []byte) (err error) {
	// All fields optional — no required property validation
	varImageDetails := _ImageDetails{}
	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varImageDetails)
	if err != nil {
		return err
	}
	*o = ImageDetails(varImageDetails)
	return nil
}

type NullableImageDetails struct {
	value *ImageDetails
	isSet bool
}

func (v NullableImageDetails) Get() *ImageDetails     { return v.value }
func (v *NullableImageDetails) Set(val *ImageDetails) { v.value = val; v.isSet = true }
func (v NullableImageDetails) IsSet() bool            { return v.isSet }
func (v *NullableImageDetails) Unset()                { v.value = nil; v.isSet = false }
func NewNullableImageDetails(val *ImageDetails) *NullableImageDetails {
	return &NullableImageDetails{value: val, isSet: true}
}
func (v NullableImageDetails) MarshalJSON() ([]byte, error) { return json.Marshal(v.value) }
func (v *NullableImageDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
