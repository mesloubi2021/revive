package fixtures

import "time"

type decodeAndValidateRequest struct {
	// BEAWRE : the flag of URLParam should match the const string URLParam
	URLParam    string `json:"-" path:"url_param" validate:"numeric"`
	Text        string `json:"text" validate:"max=10"`
	DefaultInt  int    `json:"defaultInt" default:"10.0"` // MATCH /field's type and default value's type mismatch/
	DefaultInt2 int    `json:"defaultInt2" default:"10"`
	// MATCH:12 /unknown option 'inline' in JSON tag/
	DefaultInt3      int             `json:"defaultInt2,inline" default:"11"` // MATCH /duplicate tag name: 'defaultInt2'/
	DefaultString    string          `json:"defaultString" default:"foo"`
	DefaultBool      bool            `json:"defaultBool" default:"trues"` // MATCH /field's type and default value's type mismatch/
	DefaultBool2     bool            `json:"defaultBool2" default:"true"`
	DefaultBool3     bool            `json:"defaultBool3" default:"false"`
	DefaultFloat     float64         `json:"defaultFloat" default:"f10.0"` // MATCH /field's type and default value's type mismatch/
	DefaultFloat2    float64         `json:"defaultFloat2" default:"10.0"`
	MandatoryStruct  mandatoryStruct `json:"mandatoryStruct" required:"trues"` // MATCH /required should be 'true' or 'false'/
	MandatoryStruct2 mandatoryStruct `json:"mandatoryStruct2" required:"true"`
	MandatoryStruct4 mandatoryStruct `json:"mandatoryStruct4" required:"false"`
	OptionalStruct   *optionalStruct `json:"optionalStruct,omitempty"`
	OptionalQuery    string          `json:"-" querystring:"queryfoo"`
	optionalQuery    string          `json:"-" querystring:"queryfoo"` // MATCH /tag on not-exported field optionalQuery/
	// No-reg test for bug https://github.com/deepsourcelabs/revive/issues/208
	Tiret    string `json:"-,"`
	BadTiret string `json:"other,"` // MATCH /option can not be empty in JSON tag/
}

type RangeAllocation struct {
	metav1.TypeMeta   `json:",inline"` // MATCH /unknown option 'inline' in JSON tag/
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Range             string `json:"range,flow"`  // MATCH /unknown option 'flow' in JSON tag/
	Data              []byte `json:"data,inline"` // MATCH /unknown option 'inline' in JSON tag/
}

type RangeAllocation struct {
	metav1.TypeMeta   `bson:",minsize"`
	metav1.ObjectMeta `bson:"metadata,omitempty"`
	Range             string `bson:"range,flow"` // MATCH /unknown option 'flow' in BSON tag/
	Data              []byte `bson:"data,inline"`
}

type TestContextSpecificTags2 struct {
	A       int       `asn1:"explicit,tag:1"`
	B       int       `asn1:"tag:2"`
	S       string    `asn1:"tag:0,utf8"`
	Ints    []int     `asn1:"set"`
	Version int       `asn1:"optional,explicit,default:0,tag:0"` // MATCH /duplicated tag number 0/
	Time    time.Time `asn1:"explicit,tag:4,other"`              // MATCH /unknown option 'other' in ASN1 tag/
}

type VirtualMachineRelocateSpecDiskLocator struct {
	DynamicData

	DiskId          int32                           `xml:"diskId,attr,cdata"`
	Datastore       ManagedObjectReference          `xml:"datastore,chardata,innerxml"`
	DiskMoveType    string                          `xml:"diskMoveType,omitempty,comment"`
	DiskBackingInfo BaseVirtualDeviceBackingInfo    `xml:"diskBackingInfo,omitempty,any"`
	Profile         []BaseVirtualMachineProfileSpec `xml:"profile,omitempty,other"` // MATCH /unknown option 'other' in XML tag/
}

type TestDuplicatedXMLTags struct {
	A int `xml:"a"`
	B int `xml:"a"` // MATCH /duplicate tag name: 'a'/
	C int `xml:"c"`
}

type TestDuplicatedBSONTags struct {
	A int `bson:"b"`
	B int `bson:"b"` // MATCH /duplicate tag name: 'b'/
	C int `bson:"c"`
}

type TestDuplicatedYAMLTags struct {
	A int `yaml:"b"`
	B int `yaml:"c"`
	C int `yaml:"c"` // MATCH /duplicate tag name: 'c'/
}

type TestDuplicatedProtobufTags struct {
	A int `protobuf:"varint,name=b"`
	B int `protobuf:"varint,name=c"`
	C int `protobuf:"varint,name=c"` // MATCH /duplicate tag name: 'c'/
}

// test case from
// sigs.k8s.io/kustomize/api/types/helmchartargs.go

type HelmChartArgs struct {
	ChartName        string                 `json:"chartName,omitempty" yaml:"chartName,omitempty"`
	ChartVersion     string                 `json:"chartVersion,omitempty" yaml:"chartVersion,omitempty"`
	ChartRepoURL     string                 `json:"chartRepoUrl,omitempty" yaml:"chartRepoUrl,omitempty"`
	ChartHome        string                 `json:"chartHome,omitempty" yaml:"chartHome,omitempty"`
	ChartRepoName    string                 `json:"chartRepoName,omitempty" yaml:"chartRepoName,omitempty"`
	HelmBin          string                 `json:"helmBin,omitempty" yaml:"helmBin,omitempty"`
	HelmHome         string                 `json:"helmHome,omitempty" yaml:"helmHome,omitempty"`
	Values           string                 `json:"values,omitempty" yaml:"values,omitempty"`
	ValuesLocal      map[string]interface{} `json:"valuesLocal,omitempty" yaml:"valuesLocal,omitempty"`
	ValuesMerge      string                 `json:"valuesMerge,omitempty" yaml:"valuesMerge,omitempty"`
	ReleaseName      string                 `json:"releaseName,omitempty" yaml:"releaseName,omitempty"`
	ReleaseNamespace string                 `json:"releaseNamespace,omitempty" yaml:"releaseNamespace,omitempty"`
	ExtraArgs        []string               `json:"extraArgs,omitempty" yaml:"extraArgs,omitempty"`
}
