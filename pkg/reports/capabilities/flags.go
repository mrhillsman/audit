package capabilities

// BindFlags define the flags used to generate the bundle report
type BindFlags struct {
	IndexImage        string `json:"image"`
	Limit             int32  `json:"limit"`
	HeadOnly          bool   `json:"headOnly"`
	DisableScorecard  bool   `json:"disableScorecard"`
	DisableValidators bool   `json:"disableValidators"`
	ServerMode        bool   `json:"serverMode"`
	Label             string `json:"label"`
	LabelValue        string `json:"labelValue"`
	Filter            string `json:"filter"`
	FilterBundle      string `json:"filterBundle"`
	OutputPath        string `json:"outputPath"`
	OutputFormat      string `json:"outputFormat"`
	ContainerEngine   string `json:"containerEngine"`
	InstallMode       string `json:"installMode"`
	KubeConfig        string `json:"kubeConfig"`
	Namespace         string `json:"namespace"`
	PullSecretName    string `json:"pullSecretName"`
	ServiceAccount    string `json:"serviceAccount"`
	Timeout           string `json:"timeout"`
	SkipTLS           string `json:"skipTls"`
}
