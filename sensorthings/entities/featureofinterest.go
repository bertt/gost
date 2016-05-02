package entities

import "encoding/json"

// FeatureOfInterest in SensorThings represents the phenomena an Observation is detecting. In some cases a FeatureOfInterest
// can be the Location of the Sensor and therefore of the Observation. A FeatureOfInterest is linked to a single Observation
type FeatureOfInterest struct {
	ID              string         `json:"@iot.id"`
	NavSelf         string         `json:"@iot.selfLink"`
	Description     string         `json:"description"`
	EncodingType    string         `json:"encodingtype"`
	Feature         string         `json:"feature"`
	NavObservations string         `json:"Observations@iot.navigationLink,omitempty"`
	Observations    []*Observation `json:"Observations,omitempty"`
}

// GetEntityType returns the EntityType for FeatureOfInterest
func (f *FeatureOfInterest) GetEntityType() EntityType {
	return EntityTypeFeatureOfInterest
}

// ParseEntity tries to parse the given json byte array into the current entity
func (f *FeatureOfInterest) ParseEntity(data []byte) error {
	foi := &f
	err := json.Unmarshal(data, foi)
	if err != nil {
		return err
	}

	return nil
}

// ContainsMandatoryParams checks if all mandatory params for a FeatureOfInterest are available before posting
func (f *FeatureOfInterest) ContainsMandatoryParams() (bool, []error) {
	err := []error{}
	CheckMandatoryParam(&err, f.Description, f.GetEntityType(), "description")
	CheckMandatoryParam(&err, f.EncodingType, f.GetEntityType(), "encodingtype")
	CheckMandatoryParam(&err, f.Feature, f.GetEntityType(), "feature")

	if len(err) != 0 {
		return false, err
	}

	return true, nil
}

// SetLinks sets the entity specific navigation links if needed
func (f *FeatureOfInterest) SetLinks(externalURL string) {
	f.NavSelf = CreateEntitySefLink(externalURL, EntityLinkFeatureOfInterests.ToString(), f.ID)
	f.NavObservations = CreateEntityLink(f.Observations == nil, EntityLinkFeatureOfInterests.ToString(), EntityLinkObservations.ToString(), f.ID)
}