package main

func FormText(fhirVersion string) string {
	ret := `package ` + fhirVersion + `

import "fmt"

//html Select for code, from valueset list
templ CodeSelect(fieldname string, current *string, valueset []Coding) {
		<select name={ fieldname }>
			<option value="">--</option>
			for _, c := range valueset {
				if c.Code != nil {
					<option
						value={ *c.Code }
						if current != nil && *c.Code == *current {
							selected
						}
					>
						if c.Display == nil {
							{ *c.Code }
						} else {
							{ *c.Display }
						}
					</option>
				}
			}
		</select>
}

templ CodingSelect(fieldname string, current *Coding, valueset []Coding) {
		<select
			name={ fieldname + ".display" }
			onblur="setCodingFromOptions(this)"
		>
			<option value="">--</option>
			for _, c := range valueset {
				if c.Code != nil {
					<option
						fhir-code={ *c.Code }
						fhir-system={ *c.System }
						fhir-display={ *c.Display }
						if current != nil && *c.Code == *current.Code {
							selected
						}
					>
						if c.Display == nil {
							{ *c.Code }
						} else {
							{ *c.Display }
						}
					</option>
				}
			}
		</select>
		if current != nil && current.Code != nil {
			<input name={ fieldname + ".code" } type="hidden" value={ *current.Code }/>
		} else {
			<input name={ fieldname + ".code" } type="hidden" value=""/>
		}
		if current != nil && current.System != nil {
			<input name={ fieldname + ".system" } type="hidden" value={ *current.System }/>
		} else {
			<input name={ fieldname + ".system" } type="hidden" value=""/>
		}
		if current != nil && current.Display != nil {
			<input name={ fieldname + ".display" } type="hidden" value={ *current.Display }/>
		} else {
			<input name={ fieldname + ".display" } type="hidden" value=""/>
		}
}

templ CodeableConceptSelect(fieldname string, current *CodeableConcept, valueset []Coding) {
	{{ hasOneCoding := current != nil && len(current.Coding) != 0 }}
	<select
		onblur="setCodingFromOptions(this)"
	>
		<option>--</option>
		for _, c := range valueset {
			if c.Code != nil && c.System != nil && c.Display != nil {
				<option
					fhir-code={ *c.Code }
					fhir-system={ *c.System }
					fhir-display={ *c.Display }
					if hasOneCoding && *c.Code == *current.Coding[0].Code {
						selected
					}
				>
					if c.Display == nil {
						{ *c.Code }
					} else {
						{ *c.Display }
					}
				</option>
			}
		}
	</select>
	if hasOneCoding && current.Coding[0].Code != nil {
		<input name={ fieldname + ".coding[0].code" } type="hidden" value={ *current.Coding[0].Code }/>
	} else {
		<input name={ fieldname + ".coding[0].code" } type="hidden" value=""/>
	}
	if hasOneCoding && current.Coding[0].System != nil {
		<input name={ fieldname + ".coding[0].system" } type="hidden" value={ *current.Coding[0].System }/>
	} else {
		<input name={ fieldname + ".coding[0].system" } type="hidden" value=""/>
	}
	if hasOneCoding && current.Coding[0].Display != nil {
		<input name={ fieldname + ".coding[0].display" } type="hidden" value={ *current.Coding[0].Display }/>
	} else {
		<input name={ fieldname + ".coding[0].display" } type="hidden" value=""/>
	}
}

templ AnnotationInput(fieldname string, current *Annotation) {
	if current != nil && current.Time != nil {
		<span>{ *current.Time } </span>
	}
	if current == nil {
		<textarea name={ fieldname } value=""></textarea>
	} else {
		<textarea name={ fieldname } value={ current.Text }></textarea>
	}
}

templ StringInput(fieldname string, current *string) {
	if current == nil {
		<input name={fieldname} value=""/>
	} else {
		<input name={fieldname} value={*current}/>
	}
}

templ BoolInput(fieldname string, current *bool) {
	<select name={fieldname}>
		<option value="" selected={current == nil}>--</option>
		<option value="true" selected={current != nil && *current}>Yes</option>
		<option value="false" selected={current != nil && !*current}>No</option>
	</select>
}

templ Float64Input(fieldname string, current *float64) {
	if current == nil {
		<input type="number" step="any" name={fieldname} value=""/>
	} else {
		<input type="number" step="any" name={fieldname} value={fmt.Sprintf("%f", *current)}/>
	}
}

templ IntInput(fieldname string, current *int) {
	if current == nil {
		<input type="number" name={fieldname} value=""/>
	} else {
		<input type="number" name={fieldname} value={fmt.Sprintf("%d", *current)}/>
	}
}

templ Int64Input(fieldname string, current *int64) {
	if current == nil {
		<input type="number" name={fieldname} value=""/>
	} else {
		<input type="number" name={fieldname} value={fmt.Sprintf("%d", *current)}/>
	}
}

`
	return ret
}
