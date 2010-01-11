package main

import "goldorak"

func main() {
	/******************/
	/* Initialization */
	/******************/
	goldorak.Initialize("config.json")
	//calendar := goldorak.NewModel("calendar")

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
		action.Assign("name", params[0])
		//action.Assign("name", calendar.Find("hello").Get("world"))
		//action.NoLayout()
		action.Template("hello")
	});

	// Show a calendar
	goldorak.Get("/calendars/.*", func(action *goldorak.Action, params []string) {
		action.Template("calendar")
	});

	/************/
	/* Let's go */
	/************/
	goldorak.Start()
}

