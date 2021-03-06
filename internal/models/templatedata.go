package models

import "github.com/wandachu/bookings/internal/forms"

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string // cross-site request forgery
	Flash           string
	Warning         string
	Error           string
	Form            *forms.Form
	IsAuthenticated int // 0 : not logged in. 1: logged in
}
