package urlinspection

import "encoding/xml"

type ES struct {
	Name        string    `bson:"name" json:"name"`
	ClusterName string    `bson:"cluster_name" json:"cluster_name"`
	Version     esVersion `bson:"version" json:"version"`
	Tagline     string    `bson:"tagline" json:"tagline"`
}

type esVersion struct {
	Number                           string `bson:"number" json:"number"`
	BuildFlavor                      string `bson:"build_flavor" json:"build_flavor"`
	BuildHash                        string `bson:"build_hash" json:"build_hash"`
	BuildDate                        string `bson:"build_date" json:"build_date"`
	BuildSnapshot                    bool   `bson:"build_snapshot" json:"build_snapshot"`
	LuceneVersion                    string `bson:"lucene_version" json:"lucene_version"`
	MinimumWireCompatibilityVersion  string `bson:"minimum_wire_compatibility_version" json:"minimum_wire_compatibility_version"`
	MinimumIndexCompatibilityVersion string `bson:"minimum_index_compatibility_version" json:"minimum_index_compatibility_version"`
}

type HTMLMeta struct {
	Title               string `json:"title"`
	Description         string `json:"description"`
	SiteName            string `json:"site_name"`
	ApplicationName     string `json:"application_name"`
	ApplicationVersion  string `json:"application_version"`
	ApplicationTitle    string `json:"application_title"`
	ApplicationRedirect string `json:"application_redirect"`
}

type WSDLMeta struct {
	XMLName         xml.Name `xml:"definitions"`
	Name            string   `xml:"name,attr"`
	TargetNamespace string   `xml:"targetNamespace,attr"`
	Service         Service  `xml:"service"`
}

type Service struct {
	Doc   string  `xml:"documentation"`
	Ports []*Port `xml:"port"`
}

type Port struct {
	XMLName xml.Name `xml:"port"`
	Name    string   `xml:"name,attr"`
	Binding string   `xml:"binding,attr"`
	Address Address  `xml:"address"`
}

type Address struct {
	XMLName  xml.Name `xml:"address"`
	Location string   `xml:"location,attr"`
}
