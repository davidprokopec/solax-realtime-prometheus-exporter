package inverter

import (
	"github.com/davidprokopec/solax-realtime-prometheus-exporter/solax/inverter/fields"
	"github.com/davidprokopec/solax-realtime-prometheus-exporter/solax/inverter/units"
)

type X3HybridG4 struct {
	Type        int32 `json:"type"`
	SN          string    `json:"SN"`
	Ver         string    `json:"ver"`
	Data        []float64 `json:"Data"`
	Information []any     `json:"Information"`
}

type IndexUnit struct {
	Index int
	Unit  string
}

type Decoder map[string]IndexUnit

func (x X3HybridG4) Field(field string) float64 {
	f, ok := x.decode()[field]
	if !ok {
		return 0
	}
	return x.Data[f.Index]
}

func (x X3HybridG4) decode() Decoder {
	return Decoder{
		fields.G1_Voltage:   {0, units.V},
		fields.G2_Voltage:   {1, units.V},
		fields.G3_Voltage:   {2, units.V},
		fields.G1_Current:   {3, units.A},
		fields.G2_Current:   {4, units.A},
		fields.G3_Current:   {5, units.A},
		fields.G1_Power:     {6, units.W},
		fields.G2_Power:     {7, units.W},
		fields.G3_Power:     {8, units.W},
		fields.PV1_Current:  {10, units.A},
		fields.PV2_Current:  {11, units.A},
		fields.PV1_Voltage:  {12, units.V},
		fields.PV2_Voltage:  {13, units.V},
		fields.PV1_Power:    {14, units.W},
		fields.PV2_Power:    {15, units.W},
		fields.G1_Frequency: {16, units.HZ},
		fields.G2_Frequency: {17, units.HZ},
		fields.G3_Frequency: {18, units.HZ},
		fields.Run_Mode:     {19, ""},
	}
}

// func (x X3HybridG2) decode() Decoder {
// 	return Decoder{
// 		fields.PV1_Current:          {0, units.A},
// 		fields.PV2_Current:          {1, units.A},
// 		fields.PV1_Voltage:          {2, units.V},
// 		fields.PV2_Voltage:          {3, units.V},
// 		fields.Output_Current:       {4, units.A},
// 		fields.Network_Voltage:      {5, units.V},
// 		fields.AC_Power:             {6, units.W},
// 		fields.Inverter_Temperature: {7, units.C},
// 		fields.Todays_Energy:        {8, units.KWH},
// 		fields.Total_Energy:         {9, units.KWH},
// 		fields.Exported_Power:       {10, units.W},
// 		fields.PV1_Power:            {11, units.W},
// 		fields.PV2_Power:            {12, units.W},
// 		fields.Total_FeedIn_Energy:  {41, units.KWH},
// 		fields.Total_Consumption:    {42, units.KWH},
// 		fields.Power_Now:            {43, units.W},
// 		fields.Grid_Frequency:       {50, units.HZ},
// 	}
// }
