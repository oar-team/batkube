package translate

import (
	"k8s.io/apimachinery/pkg/watch"
)

/*
Batsim protocol
*/

type BatMessage struct {
	Now    float64 `json:"now"`
	Events []Event `json:"events"`
}

type Event struct {
	Timestamp float64                `json:"timestamp"`
	Type      string                 `json:"type"`
	Data      map[string]interface{} `json:"data"`
}

/*
Batsim message data
*/

type SimulationBeginsData struct {
	NbResources         int  `mapstructure:"nb_resources"`
	NbComputeResources  int  `mapstructure:"nb_compute_resources"`
	NbStorageResources  int  `mapstructure:"nb_storage_resources"`
	AllowComputeSharing bool `mapstructure:"allow_compute_sharing"`
	AllowStorageSharing bool `mapstructure:"allow_storage_sharing"`
	Config              `mapstructure:"config"`
	ComputeResources    []ComputeResource             `mapstructure:"compute_resources"`
	StorageResources    []map[string]interface{}      `mapstructure:"storage_resources"`
	Workloads           map[string]string             `mapstructure:"workloads"`
	Profiles            map[string]map[string]Profile `mapstructure:"profiles"`
}

type Workload struct {
	NbRes    int
	Jobs     []Job
	Profiles map[string]Profile
}

// TODO : config might be properly defined somewhere
// Also, it has default values
type Config map[string]interface{}

type ComputeResource struct {
	Id   int    `mapstructure:"id"`
	Name string `mapstructure:"name"`
	// state in {sleeping, idle, computing, switching_on, switching_off}
	State      string                 `mapstructure:"state"`
	Properties map[string]interface{} `mapstructure:"properties"`
}

type Job struct {
	// Sometimes job id are written as integers
	Id           string                 `mapstructure:"id"`
	Subtime      float64                `mapstructure:"subtime"`
	Res          int                    `mapstructure:"res"`
	Profile      string                 `mapstructure:"profile"`
	CustomFields map[string]interface{} `mapstructure:"-"`
}

type Profile struct {
	Type  string                 `mapstructure:"type"`
	Ret   int                    `mapstructure:"ret"`
	Specs map[string]interface{} `mapstructure:"-"`
}

type ExecuteJobData struct {
	JobId string `mapstructure:"job_id"`
	Alloc string `mapstructure:"alloc"`
}

type JobCompletedData struct {
	JobId      string `mapstructure:"job_id"`
	JobState   string `mapstructure:"job_state"`
	ReturnCode int    `mapstructure:"return_code"`
	Alloc      string `mapstructure:"alloc"`
}

type CallMeLaterData struct {
	Timestamp float64 `mapstructure:"timestamp"`
}

/*
Watch event types
*/
var Added = string(watch.Added)
var Modified = string(watch.Modified)
var Deleted = string(watch.Deleted)
var Error = string(watch.Error)
