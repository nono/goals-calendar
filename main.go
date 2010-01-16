package main

import (
	"fmt"
	"goldorak"
	"strconv"
	"strings"
)

var goal *goldorak.Model


// Show the calendar
func showCalendar(action *goldorak.Action, cal *Calendar) {
	action.Assign("name", cal.Title())
	action.Assign("year", fmt.Sprint(cal.Year))
	action.Assign("month", cal.MonthAsText())
	action.Assign("prev_url", "/" + cal.Goal.Param + "/" + cal.PrevMonth().String())
	action.Assign("next_url", "/" + cal.Goal.Param + "/" + cal.NextMonth().String())
	action.Assign("current", goldorak.Pluralize(cal.CurrentStreak(), "jour"))
	action.Assign("longest", goldorak.Pluralize(cal.LongestStreak(), "jour"))
	action.Assign("rows", "") // FIXME
	action.Template("calendar")
}

// Form for creating a new goal
func newGoal(action *goldorak.Action, param string, public bool) {
	action.Assign("name", param)
	action.Assign("public", "") // FIXME
	action.Template("new_goal")
}

// Create a goal
func createGoal(action *goldorak.Action, name string, public bool) {
	g := goal.Create(name)
	g.Set("name", name)
	g.Set("current", "0")
	g.Set("longest", "0")
	action.Redirect("/" + g.Param)
}

// Home Page
func homepage(action *goldorak.Action) {
	action.Template("homepage")
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
	//goldorak.Get("/.*(/[0-9]+/[0-9]+)?", func(action *goldorak.Action, params []string) {
	goldorak.Get("/.*", func(action *goldorak.Action, params []string) {
		g := goal.Find(params[0])
		if g != nil {
			cal := NewCal(g)
			if (len(params) > 2) {
				cal.Year, _ = strconv.Atoi(params[1])
				cal.Month,_ = strconv.Atoi(params[2])
			}
			showCalendar(action, cal)
		} else {
			newGoal(action, params[0], true)
		}
	})

	// Show the form for creating a new goal
	goldorak.Get("/objectifs/nouveau", func(action *goldorak.Action, params []string) {
		newGoal(action, "", true)
	})

	// Create a goal
	goldorak.Post("/objectifs", func(action *goldorak.Action, params []string) {
		p_public := action.Param("public")
		public   := len(p_public) > 0 && p_public[0] == "1"
		p_name   := action.Param("name")
		name     := ""
		if len(p_name) > 0 {
			name = strings.TrimSpace(p_name[0])
		}
		if name != "" {
			createGoal(action, name, public)
		} else {
			newGoal(action, name, public)
		}
	})

	// Home Page
	goldorak.Get("/", func(action *goldorak.Action, params []string) {
		homepage(action)
	})

	/************/
	/* Let's go */
	/************/
	goldorak.Start()
}

