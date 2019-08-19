package types

type Tick struct {
	Version   int64        `json:"version"`
	Timestamp int64        `json:"ts"`
	Bids      [][2]float64 `json:"bids"`
	Asks      [][2]float64 `json:"asks"`
}
