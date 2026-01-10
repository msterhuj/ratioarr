package trackers

type Ratio struct {
	Uploaded   int64
	Downloaded int64
	Ratio      float64
}

type Tracker interface {
	Name() string
	FetchRatio() (*Ratio, error)
}
