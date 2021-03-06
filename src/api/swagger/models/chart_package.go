package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*ChartPackage chart package

swagger:model chartPackage
*/
type ChartPackage struct {

	/* app version

	Required: true
	Min Length: 1
	*/
	AppVersion *string `json:"appVersion"`

	/* created

	Required: true
	Min Length: 1
	*/
	Created *string `json:"created"`

	/* description

	Required: true
	Min Length: 1
	*/
	Description *string `json:"description"`

	/* digest

	Required: true
	Min Length: 1
	*/
	Digest *string `json:"digest"`

	/* home

	Required: true
	Min Length: 1
	*/
	Home *string `json:"home"`

	/* icon

	Min Length: 1
	*/
	Icon string `json:"icon,omitempty"`

	/* keywords
	 */
	Keywords []string `json:"keywords,omitempty"`

	/* maintainers

	Required: true
	*/
	Maintainers []*Maintainer `json:"maintainers"`

	/* name

	Required: true
	Min Length: 1
	*/
	Name *string `json:"name"`

	/* repo
	 */
	Repo string `json:"repo,omitempty"`

	/* sources

	Required: true
	*/
	Sources []string `json:"sources"`

	/* urls

	Required: true
	*/
	Urls []string `json:"urls"`

	/* version

	Required: true
	Min Length: 1
	*/
	Version *string `json:"version"`
}

// Validate validates this chart package
func (m *ChartPackage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDigest(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHome(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIcon(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKeywords(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMaintainers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSources(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUrls(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChartPackage) validateAppVersion(formats strfmt.Registry) error {

	if err := validate.Required("appVersion", "body", m.AppVersion); err != nil {
		return err
	}

	if err := validate.MinLength("appVersion", "body", string(*m.AppVersion), 1); err != nil {
		return err
	}

	return nil
}

func (m *ChartPackage) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", m.Created); err != nil {
		return err
	}

	if err := validate.MinLength("created", "body", string(*m.Created), 1); err != nil {
		return err
	}

	return nil
}

func (m *ChartPackage) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 1); err != nil {
		return err
	}

	return nil
}

func (m *ChartPackage) validateDigest(formats strfmt.Registry) error {

	if err := validate.Required("digest", "body", m.Digest); err != nil {
		return err
	}

	if err := validate.MinLength("digest", "body", string(*m.Digest), 1); err != nil {
		return err
	}

	return nil
}

func (m *ChartPackage) validateHome(formats strfmt.Registry) error {

	if err := validate.Required("home", "body", m.Home); err != nil {
		return err
	}

	if err := validate.MinLength("home", "body", string(*m.Home), 1); err != nil {
		return err
	}

	return nil
}

func (m *ChartPackage) validateIcon(formats strfmt.Registry) error {

	if swag.IsZero(m.Icon) { // not required
		return nil
	}

	if err := validate.MinLength("icon", "body", string(m.Icon), 1); err != nil {
		return err
	}

	return nil
}

func (m *ChartPackage) validateKeywords(formats strfmt.Registry) error {

	if swag.IsZero(m.Keywords) { // not required
		return nil
	}

	return nil
}

func (m *ChartPackage) validateMaintainers(formats strfmt.Registry) error {

	if err := validate.Required("maintainers", "body", m.Maintainers); err != nil {
		return err
	}

	for i := 0; i < len(m.Maintainers); i++ {

		if swag.IsZero(m.Maintainers[i]) { // not required
			continue
		}

		if m.Maintainers[i] != nil {

			if err := m.Maintainers[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ChartPackage) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	return nil
}

func (m *ChartPackage) validateSources(formats strfmt.Registry) error {

	if err := validate.Required("sources", "body", m.Sources); err != nil {
		return err
	}

	return nil
}

func (m *ChartPackage) validateUrls(formats strfmt.Registry) error {

	if err := validate.Required("urls", "body", m.Urls); err != nil {
		return err
	}

	return nil
}

func (m *ChartPackage) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if err := validate.MinLength("version", "body", string(*m.Version), 1); err != nil {
		return err
	}

	return nil
}
