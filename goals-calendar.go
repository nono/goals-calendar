package main

import "goldorak"

func main() {
	/******************/
	/* Initialization */
	/******************/
	goldorak.Initialize("config.json")
	// Calendar := goldorak.NewModel("calendar")

	/***********/
	/* Actions */
	/***********/

	// Hello world
	goldorak.Get("/hello", func(action goldorak.Action) {
		action.Assign("name", "World!")
		//action.Assign("name", Calendar.Find("hello").Get("world"))
		action.Layout("")
		action.Render("hello")
	});

	// Show a calendar
	goldorak.Get("/calendars/(.*)", func(action goldorak.Action) {
		action.Render("calendar")
	});

	/************/
	/* Let's go */
	/************/
	goldorak.Start()
}

