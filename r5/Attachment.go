//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

// http://hl7.org/fhir/r5/StructureDefinition/Attachment
type Attachment struct {
	Id          *string     `json:"id,omitempty"`
	Extension   []Extension `json:"extension,omitempty"`
	ContentType *string     `json:"contentType,omitempty"`
	Language    *string     `json:"language,omitempty"`
	Data        *string     `json:"data,omitempty"`
	Url         *string     `json:"url,omitempty"`
	Size        *int64      `json:"size,omitempty"`
	Hash        *string     `json:"hash,omitempty"`
	Title       *string     `json:"title,omitempty"`
	Creation    *string     `json:"creation,omitempty"`
	Height      *int        `json:"height,omitempty"`
	Width       *int        `json:"width,omitempty"`
	Frames      *int        `json:"frames,omitempty"`
	Duration    *float64    `json:"duration,omitempty"`
	Pages       *int        `json:"pages,omitempty"`
}
