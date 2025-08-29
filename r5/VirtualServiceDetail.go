package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r5/StructureDefinition/VirtualServiceDetail
type VirtualServiceDetail struct {
	Id                           *string                `json:"id,omitempty"`
	Extension                    []Extension            `json:"extension,omitempty"`
	ChannelType                  *Coding                `json:"channelType,omitempty"`
	AddressUrl                   *string                `json:"addressUrl"`
	AddressString                *string                `json:"addressString"`
	AddressContactPoint          *ContactPoint          `json:"addressContactPoint"`
	AddressExtendedContactDetail *ExtendedContactDetail `json:"addressExtendedContactDetail"`
	AdditionalInfo               []string               `json:"additionalInfo,omitempty"`
	MaxParticipants              *int                   `json:"maxParticipants,omitempty"`
	SessionKey                   *string                `json:"sessionKey,omitempty"`
}
