package main

import (
	"fmt"
	"goldorak"
	"strconv"
	"strings"
)

var goal *goldorak.Model


// Show the calendar
func showCalendar(action *goldorak.Action, goal *goldorak.Instance, params []string) {
	cal := NewCal()
	if (len(params) > 2) {
		cal.Year, _ = strconv.Atoi(params[1])
		cal.Month,_ = strconv.Atoi(params[2])
	}
	prev := cal.PrevMonth()
	next := cal.NextMonth()
	current, _ := strconv.Atoi(goal.Get("current"))
	longest, _ := strconv.Atoi(goal.Get("longest"))
	action.Assign("name", goal.Get("title"))
	action.Assign("year", fmt.Sprint("%04d", cal.Year))
	action.Assign("month", fmt.Sprint("%02d", cal.Month))
	action.Assign("prev_url", "/" + params[0] + "/" + prev.String())
	action.Assign("next_url", "/" + params[0] + "/" + next.String())
	action.Assign("current", goldorak.Pluralize(current, "jour"))
	action.Assign("longest", goldorak.Pluralize(longest, "jour"))
	action.Assign("rows", "") // FIXME
	action.Template("calendar")
}

// Form for creating a new goal
func newGoal(action *goldorak.Action, param string) {
	action.Assign("name", param)
	action.Template("new_calendar")
}

// Create a goal
func createGoal(action *goldorak.Action, name string) {
	s := goldorak.Parameterize(name)
	g := goal.Create(s)
	g.Set("name", name)
	g.Set("current", "0")
	g.Set("longest", "0")
	action.Redirect("/" + s)
}

func main() {
	/******************/
	/* Initialization */
	/******************/
	goldorak.Initialize("config.json")
	conn := goldorak.Connect()
	goal  = conn.NewModel("goal")

	/***********/
	/* Actions */
	/***********/

	// Layout
	goldorak.DefaultLayout(func(action *goldorak.Action) {
		action.Assign("favicon", goldorak.StaticUrl("favicon.png"))
		action.Assign("stylesheet", goldorak.StaticUrl("styles.css"))
		action.Template("layout")
	})

	// Show a calendar
	goldorak.Get("/.*(/[0-9]+/[0-9]+)?", func(action *goldorak.Action, params []string) {
		g := goal.Find(params[0])
		if g != nil {
			showCalendar(action, g, params)
		} else {
			newGoal(action, params[0])
		}
	});

	// Create a calendar
	goldorak.Post("/calendars", func(action *goldorak.Action, params []string) {
		name := strings.TrimSpace(params[0])
		if name != "" {
			createGoal(action, name)
		} else {
			newGoal(action, "")
		}
	});

	/************/
	/* Let's go */
	/************/
	goldorak.Start()
}

