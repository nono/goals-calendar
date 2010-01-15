package main

import (
	"fmt"
	"time"
	"goldorak"
	"strconv"
)

type Calendar struct {
	Year   int
	Month  int
	Goal   *goldorak.Instance
}

func NewCal(goal *goldorak.Instance) *Calendar {
	now := time.LocalTime()
	year, month := int(now.Year), now.Month
	return &Calendar{year, month, goal}
}

func (this *Calendar) String() string {
	return fmt.Sprintf("%04d/%02d", this.Year, this.Month)
}

func (this *Calendar) PrevMonth() *Calendar {
	if this.Month == 1 {
		return &Calendar{this.Year - 1, 12, this.Goal}
	}
	return &Calendar{this.Year, this.Month - 1, this.Goal}
}

func (this *Calendar) NextMonth() *Calendar {
	if this.Month == 12 {
		return &Calendar{this.Year + 1, 1, this.Goal}
	}
	return &Calendar{this.Year, this.Month + 1, this.Goal}
}

func (this *Calendar) CurrentStreak() int {
	ret, _ := strconv.Atoi(this.Goal.Get("current"))
	return ret
}

func (this *Calendar) LongestStreak() int {
	ret, _ := strconv.Atoi(this.Goal.Get("longest"))
	return ret
}

func (this *Calendar) Title() string {
	return this.Goal.Get("title")
}

func (this *Calendar) MonthAsText() string {
	switch this.Month {
		case 1:  return "Janvier"
		case 2:  return "Février"
		case 3:  return "Mars"
		case 4:  return "Avril"
		case 5:  return "Mai"
		case 6:  return "Juin"
		case 7:  return "Juillet"
		case 8:  return "Août"
		case 9:  return "Septembre"
		case 10: return "Octobre"
		case 11: return "Novembre"
		case 12: return "Décembre"
	}
	return "Inconnu"
}

