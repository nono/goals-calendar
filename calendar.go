package main

import (
	"fmt"
	"time"
)

type Calendar struct {
	Year  int
	Month int
}

func NewCal() *Calendar {
	now := time.LocalTime()
	year, month := int(now.Year), now.Month
	return &Calendar{year, month}
}

func (this *Calendar) String() string {
	return fmt.Sprintf("%04d/%02d", this.Year, this.Month)
}

func (this *Calendar) PrevMonth() *Calendar {
	if this.Month == 1 {
		return &Calendar{this.Year - 1, 12}
	}
	return &Calendar{this.Year, this.Month - 1}
}

func (this *Calendar) NextMonth() *Calendar {
	if this.Month == 12 {
		return &Calendar{this.Year + 1, 1}
	}
	return &Calendar{this.Year, this.Month + 1}
}

