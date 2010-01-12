package main

import (
	"fmt"
	"goldorak"
	"strconv"
	"time"
)


func main() {
	/******************/
	/* Initialization */
	/******************/
	goldorak.Initialize("config.json")
	conn := goldorak.Connect()
	calendar := conn.NewModel("calendar")

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
		var year int64
		var month int
		cal := calendar.Find(params[0])
		if cal != nil {
			// Show the calendar
			if (len(params) > 2) {
				year, _ = strconv.Atoi64(params[0])
				month,_ = strconv.Atoi(params[1])
			} else {
				now := time.LocalTime()
				year, month = now.Year, now.Month
			}
			action.Assign("name", cal.Get("title"))
			action.Assign("year", fmt.Sprint("%04d", year))
			action.Assign("month", fmt.Sprint("%02d", month))
			action.Template("calendar")
		} else {
			action.Assign("name", params[0])
			action.Template("new_calendar")
		}
	});

	// Create a calendar
	goldorak.Post("/calendars", func(action *goldorak.Action, params []string) {
		slug := goldorak.Parameterize(params[0])
		cal := calendar.Create(slug)
		action.Redirect("/" + cal.Param)
	});

	/************/
	/* Let's go */
	/************/
	goldorak.Start()
}

