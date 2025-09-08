package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r5/StructureDefinition/VirtualServiceDetail
type VirtualServiceDetail struct {
	Id                           *string                `json:"id,omitempty"`
	Extension                    []Extension            `json:"extension,omitempty"`
	ChannelType                  *Coding                `json:"channelType,omitempty"`
	AddressUrl                   *string                `json:"addressUrl,omitempty"`
	AddressString                *string                `json:"addressString,omitempty"`
	AddressContactPoint          *ContactPoint          `json:"addressContactPoint,omitempty"`
	AddressExtendedContactDetail *ExtendedContactDetail `json:"addressExtendedContactDetail,omitempty"`
	AdditionalInfo               []string               `json:"additionalInfo,omitempty"`
	MaxParticipants              *int                   `json:"maxParticipants,omitempty"`
	SessionKey                   *string                `json:"sessionKey,omitempty"`
}
