package sdk

import (
	"encoding/json"
)

type TargetCommonInfo struct {
	RefID      string `json:"refId"`
	Datasource string `json:"datasource,omitempty"`
	Hide       bool   `json:"hide,omitempty"`
}

type OpentsdbTargetInfo struct {
	Aggregator           string           `json:"aggregator"`
	Alias                string           `json:"alias,omitempty"`
	CurrentFilterGroupBy bool             `json:"currentFilterGroupBy,omitempty"`
	CurrentFilterKey     string           `json:"currentFilterKey,omitempty"`
	CurrentFilterType    string           `json:"currentFilterType,omitempty"`
	CurrentFilterValue   string           `json:"currentFilterValue,omitempty"`
	CurrentTagKey        string           `json:"currentTagKey,omitempty"`
	CurrentTagValue      string           `json:"currentTagValue,omitempty"`
	DownsampleAggregator string           `json:"downsampleAggregator,omitempty"`
	DownsampleFillPolicy string           `json:"downsampleFillPolicy,omitempty"`
	DownsampleInterval   string           `json:"downsampleInterval,omitempty"`
	Filters              []OpentsdbFilter `json:"filters,omitempty"`
	Metric               string           `json:"metric"`
}

type MixedTargetInfo struct {
	// For PostgreSQL
	Table        string `json:"table,omitempty"`
	TimeColumn   string `json:"timeColumn,omitempty"`
	MetricColumn string `json:"metricColumn,omitempty"`
	RawSql       string `json:"rawSql,omitempty"`
	Select       [][]struct {
		Params []string `json:"params,omitempty"`
		Type   string   `json:"type,omitempty"`
	} `json:"select,omitempty"`
	Where []struct {
		Type     string   `json:"type,omitempty"`
		Name     string   `json:"name,omitempty"`
		Params   []string `json:"params,omitempty"`
		Datatype string   `json:"datatype,omitempty"`
	} `json:"where,omitempty"`
	Group []struct {
		Type   string   `json:"type,omitempty"`
		Params []string `json:"params,omitempty"`
	} `json:"group,omitempty"`

	// For Prometheus
	Expr           string `json:"expr,omitempty"`
	IntervalFactor int    `json:"intervalFactor,omitempty"`
	Interval       string `json:"interval,omitempty"`
	Step           int    `json:"step,omitempty"`
	LegendFormat   string `json:"legendFormat,omitempty"`
	Instant        bool   `json:"instant,omitempty"`
	Format         string `json:"format,omitempty"`

	// For InfluxDB
	Measurement string `json:"measurement,omitempty"`

	// For Elasticsearch
	DsType  *string `json:"dsType,omitempty"`
	Metrics []struct {
		ID    string `json:"id"`
		Field string `json:"field"`
		Type  string `json:"type"`
	} `json:"metrics,omitempty"`
	Query      string `json:"query,omitempty"`
	Alias      string `json:"alias,omitempty"`
	RawQuery   bool   `json:"rawQuery,omitempty"`
	TimeField  string `json:"timeField,omitempty"`
	BucketAggs []struct {
		ID       string `json:"id"`
		Field    string `json:"field"`
		Type     string `json:"type"`
		Settings struct {
			Interval    string `json:"interval,omitempty"`
			MinDocCount int    `json:"min_doc_count"`
			Order       string `json:"order,omitempty"`
			OrderBy     string `json:"orderBy,omitempty"`
			Size        string `json:"size,omitempty"`
		} `json:"settings"`
	} `json:"bucketAggs,omitempty"`

	// For Graphite
	Target string `json:"target,omitempty"`

	// For CloudWatch
	Namespace  string            `json:"namespace,omitempty"`
	MetricName string            `json:"metricName,omitempty"`
	Statistics []string          `json:"statistics,omitempty"`
	Dimensions map[string]string `json:"dimensions,omitempty"`
	Period     string            `json:"period,omitempty"`
	Region     string            `json:"region,omitempty"`

	// For the Stackdriver data source. Find out more information at
	// https:/grafana.com/docs/grafana/v6.0/features/datasources/stackdriver/
	ProjectName        string                    `json:"projectName,omitempty"`
	AlignOptions       []StackdriverAlignOptions `json:"alignOptions,omitempty"`
	AliasBy            string                    `json:"aliasBy,omitempty"`
	MetricType         string                    `json:"metricType,omitempty"`
	MetricKind         string                    `json:"metricKind,omitempty"`
	Filters            []string                  `json:"filters,omitempty"`
	AlignmentPeriod    string                    `json:"alignmentPeriod,omitempty"`
	CrossSeriesReducer string                    `json:"crossSeriesReducer,omitempty"`
	PerSeriesAligner   string                    `json:"perSeriesAligner,omitempty"`
	ValueType          string                    `json:"valueType,omitempty"`
	GroupBys           []string                  `json:"groupBys,omitempty"`
}

type OpentsdbFilter struct {
	Filter  string `json:"filter"`
	GroupBy bool   `json:"groupBy"`
	TagK    string `json:"tagk"`
	Type    string `json:"type"`
}

func (t *Target) UnmarshalJSON(b []byte) (err error) {
	var common TargetCommonInfo
	err = json.Unmarshal(b, &common)
	if err != nil {
		return
	}

	// try specific datasource first
	var opentsdbTarget OpentsdbTargetInfo
	err = json.Unmarshal(b, &opentsdbTarget)
	if err == nil {
		if opentsdbTarget.Metric != "" && opentsdbTarget.Aggregator != "" {
			t.OpentsdbTargetInfo = &opentsdbTarget
			return
		}
	}

	var mixed MixedTargetInfo
	err = json.Unmarshal(b, &mixed)
	if err == nil {
		t.MixedTargetInfo = mixed
	}
	return
}

func (t *Target) MarshalJSON() ([]byte, error) {
	// try specific datasource first
	if t.OpentsdbTargetInfo != nil {
		var out = struct {
			TargetCommonInfo
			OpentsdbTargetInfo
		}{t.TargetCommonInfo, *t.OpentsdbTargetInfo}
		return json.Marshal(out)
	}

	var out = struct {
		TargetCommonInfo
		MixedTargetInfo
	}{t.TargetCommonInfo, t.MixedTargetInfo}
	return json.Marshal(out)
}
