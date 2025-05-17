package model

import (
	"fmt"
	"strings"
	"time"
)

type DateOnly time.Time

const dateLayout = "1999-01-01"

func (date *DateOnly) UnmarshalJSON(byte []byte) error {
	strn := strings.Trim(string(byte), `"`)
	timeParse, err := time.Parse(dateLayout, strn)
	if err != nil {
		return err
	}
	*date = DateOnly(timeParse)
	return nil
}

func (date DateOnly) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, time.Time(date).Format(dateLayout))), nil
}

func (date DateOnly) toTime() time.Time {
	return time.Time(date)
}
