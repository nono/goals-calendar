package main

import (
	"fmt"
	"goldorak"
	"strconv"
	"strings"
	"time"
)

// TODO create a type calendar
var goal *goldorak.Model


// Show the calendar
func showCalendar(action *goldorak.Action, goal *goldorak.Instance, params []string) {
	var year, month int
	if (len(params) > 2) {
		year, _ = strconv.Atoi(params[1])
		month,_ = strconv.Atoi(params[2])
	} else {
		now := time.LocalTime()
		year, month = int(now.Year), now.Month
	}
	action.Assign("name", goal.Get("title"))
	action.Assign("year", fmt.Sprint("%04d", year))
	action.Assign("month", fmt.Sprint("%02d", month))
	action.Assign("prev_url", "/" + params[0]) // FIXME
	action.Assign("next_url", "/" + params[0]) // FIXME
	action.Assign("rows", "") // FIXME
	action.Assign("current", goal.Get("current") + " jours") // TODO Pluralize
	action.Assign("longest", goal.Get("longest") + " jours") // TODO Pluralize
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

