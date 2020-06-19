package translate

import (
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

func DeserializeSimulationBegins(data map[string]interface{}, v *SimulationBeginsData) error {
	var err error
	if err = mapstructure.Decode(data, v); err != nil {
		return errors.Wrap(err, "error while decoding simulation begins data")
	}

	// Deserialize profiles
	for wlId, profilesData := range data["profiles"].(map[string]interface{}) {
		for profileName, profileData := range profilesData.(map[string]interface{}) {
			profile := v.Profiles[wlId][profileName]
			profileDataMap := profileData.(map[string]interface{})

			// Profile specs
			if err := mapstructure.Decode(profileDataMap, &profile.Specs); err != nil {
				return errors.Wrap(err, "could not deserialize profile specs")
			}
			delete(profile.Specs, "type")
			delete(profile.Specs, "ret")
			// TODO : how to modify profile in place?
			v.Profiles[wlId][profileName] = profile
		}
	}
	return nil
}

func DeserializeJobSubmitted(data map[string]interface{}, v *Job) error {
	// TODO : handle redis
	if data["job"] == nil {
		return errors.Errorf("[translate] Could not find job data. Perhaps you are trying to launch a simulation with redis enabled, which is not currently supported by the simulator.")
	}

	if err := mapstructure.Decode(data["job"], v); err != nil {
		return errors.Wrap(err, "[translate] failed to deserialize a job")
	}

	if err := mapstructure.Decode(data["job"], &(v.CustomFields)); err != nil {
		return errors.Wrap(err, "[translate] failed to deserialize a job's custom fields")
	}
	delete(v.CustomFields, "id")
	delete(v.CustomFields, "subtime")
	delete(v.CustomFields, "res")
	delete(v.CustomFields, "profile")

	return nil
}

func DeserializeJobCompleted(data map[string]interface{}, v *JobCompletedData) error {
	if err := mapstructure.Decode(data, v); err != nil {
		return errors.Wrap(err, "[translate] failed to deserialize a JOB_COMPLETED message")
	}
	return nil
}
