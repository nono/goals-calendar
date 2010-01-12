package main

import (
	"goldorak"
	"strconv"
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

	// Hello world
	goldorak.Get("/hello", func(action *goldorak.Action, params []string) {
		action.Assign("name", "world")
		action.NoLayout()
		action.Template("hello")
	});

	// Show a calendar
	goldorak.Get("/.*(/[0-9]+/[0-9]+)?", func(action *goldorak.Action, params []string) {
		cal     := calendar.Find(params[0])
		if cal != nil {
			// Show the calendar
			year, _ := strconv.Atoi(params[0])
			month,_ := strconv.Atoi(params[1])
			action.Assign("name", cal.Get("title"))
			action.Assign("not_used", string(year + month))
			action.Template("calendar")
		} else {
			// TODO create a new calendar
		}
	});

	/************/
	/* Let's go */
	/************/
	goldorak.Start()
}

