package convert

import "time"

// TimeToString :
func TimeToString(t time.Time) string {
	return t.Format(time.RFC3339)
}
