package sdk

import (
	"encoding/json"
	"testing"
)

func TestTarget_Unmarshal1(t *testing.T) {
	// raw data with some residue fields
	// such as the following `dimensions`,`group`
	raw := []byte(`
[
  {
    "describe": "",
    "dimensions": [
      "$rds_instance"
    ],
    "expr": "topk(3,acs_rds_dashboard_CpuUsage_Maximum__{app_id=\"$appid\"})",
    "group": "",
    "interval": "",
    "legendFormat": "{{instance_name}}",
    "metric": "CpuUsage",
    "period": "60",
    "project": "acs_rds_dashboard",
    "refId": "A",
    "target": [
      "Maximum"
    ],
    "type": "timeserie",
    "xcol": "timestamp",
    "ycol": [
      "Maximum"
    ]
  },
  {
    "expr": "topk(3,aws_rds_CPUUtilization_Maximum_s{app_id=~\"$appid\"})",
    "hide": false,
    "interval": "",
    "legendFormat": "{{DBInstanceIdentifier}}",
    "refId": "B"
  }
]
`)
	var targets []Target
	err := json.Unmarshal(raw, &targets)
	if err != nil {
		t.Error(err.Error())
	}
}
