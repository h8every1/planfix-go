package planfix

import "encoding/xml"

/*type XmlRequest struct {
	XMLName xml.Name `xml:"request"`
	Method  string   `xml:"method,attr"`
}

func (r xmlRequest) getMethod() string {
	return r.Method
}
func (r *xmlRequest) setMethod(m string) {
	r.Method = m
}

type xmlResponse struct {
	XMLName xml.Name `xml:"response"`
	Status  string   `xml:"status,attr"`
}*/

type XmlResponseFile struct {
	Id   int    `xml:"id"`
	Name string `xml:"name"`
}

type XmlResponseUser struct {
	Id   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type XmlResponseActionAnalitic struct {
	Id   int    `xml:"id"`
	Key  int    `xml:"key"`
	Name string `xml:"name"`
}

type XmlResponseAnaliticOptions struct {
	Id      int                               `xml:"id"`
	Name    string                            `xml:"name"`
	GroupId int                               `xml:"group>id"`
	Fields  []XmlResponseAnaliticOptionsField `xml:"fields>field"`
}

type XmlResponseAnaliticOptionsField struct {
	Id         int      `xml:"id"`
	Num        int      `xml:"num"`
	Name       string   `xml:"name"`
	Type       string   `xml:"type"`
	ListValues []string `xml:"list>value"`
	HandbookId int      `xml:"handbook>id"`
}

type XmlResponseAction struct {
	Id                           int                         `xml:"id"`
	Description                  string                      `xml:"description"`
	OldStatus                    int                         `xml:"statusChange>oldStatus,omitempty"`
	NewStatus                    int                         `xml:"statusChange>newStatus,omitempty"`
	IsNotRead                    bool                        `xml:"isNotRead"`
	FromEmail                    bool                        `xml:"fromEmail"`
	DateTime                     string                      `xml:"dateTime"`
	TaskId                       int                         `xml:"task>id"`
	TaskTitle                    string                      `xml:"task>title"`
	ContactGeneral               int                         `xml:"contact>general"`
	ContactName                  string                      `xml:"contact>name"`
	Owner                        XmlResponseUser             `xml:"owner"`
	ProjectId                    int                         `xml:"project>id"`
	ProjectTitle                 string                      `xml:"project>title"`
	TaskExpectDateChangedOldDate string                      `xml:"taskExpectDateChanged>oldDate"`
	TaskExpectDateChangedNewDate string                      `xml:"taskExpectDateChanged>newDate"`
	TaskStartTimeChangedOldDate  string                      `xml:"taskStartTimeChanged>oldDate"`
	TaskStartTimeChangedNewDate  string                      `xml:"taskStartTimeChanged>newDate"`
	Files                        []XmlResponseFile           `xml:"files>file"`
	NotifiedList                 []XmlResponseUser           `xml:"notifiedList>user"`
	Analitics                    []XmlResponseActionAnalitic `xml:"analitics>analitic"`
}

// TODO: добавить все поля из https://planfix.ru/docs/ПланФикс_API_task.get
type XmlResponseTask struct {
	Id           int    `xml:"id"`
	Title        string `xml:"title"`
	Description  string `xml:"description"`
	General      int    `xml:"general"`
	ProjectId    int    `xml:"project>id"`
	ProjectTitle string `xml:"project>title"`
}

type XmlResponseAnalitic struct {
	Id        int    `xml:"id"`
	Name      string `xml:"name"`
	GroupId   int    `xml:"group>id"`
	GroupName string `xml:"group>name"`
}

type XmlRequestAnalitic struct {
	Id       int                       `xml:"id"`
	ItemData []XmlRequestAnaliticField `xml:"analiticData>itemData"`
}

type XmlRequestAnaliticField struct {
	FieldId int         `xml:"fieldId"`
	Value   interface{} `xml:"value"`
}

// auth.login
type XmlRequestAuth struct {
	XMLName xml.Name `xml:"request"`
	Method  string   `xml:"method,attr"`

	Account  string `xml:"account"`
	Login    string `xml:"login"`
	Password string `xml:"password"`
}

// auth.login response
type XmlResponseAuth struct {
	XMLName xml.Name `xml:"response"`
	Status  string   `xml:"status,attr"`
	Code    string   `xml:"code"`
	Message string   `xml:"message"`

	Sid string `xml:"sid"`
}

// action.get
type XmlRequestActionGet struct {
	XMLName xml.Name `xml:"request"`
	Method  string   `xml:"method,attr"`
	Account string   `xml:"account"`
	Sid     string   `xml:"sid"`

	ActionId int `xml:"action>id"`
}

// action.get response
type XmlResponseActionGet struct {
	XMLName xml.Name `xml:"response"`
	Status  string   `xml:"status,attr"`
	Code    string   `xml:"code"`
	Message string   `xml:"message"`

	Action XmlResponseAction `xml:"action"`
}

// action.getList
type XmlRequestActionGetList struct {
	XMLName xml.Name `xml:"request"`
	Method  string   `xml:"method,attr"`
	Account string   `xml:"account"`
	Sid     string   `xml:"sid"`

	TaskId         int    `xml:"task>id,omitempty"`
	TaskGeneral    int    `xml:"task>general,omitempty"`
	ContactGeneral int    `xml:"contact>general,omitempty"`
	PageCurrent    int    `xml:"pageCurrent"`
	PageSize       int    `xml:"pageSize"`
	Sort           string `xml:"sort"`
}

// action.getList response
type XmlResponseActionGetList struct {
	XMLName xml.Name `xml:"response"`
	Status  string   `xml:"status,attr"`
	Code    string   `xml:"code"`
	Message string   `xml:"message"`

	Actions struct {
		ActionsCount      int                 `xml:"count,attr"`
		ActionsTotalCount int                 `xml:"totalCount,attr"`
		Actions           []XmlResponseAction `xml:"action"`
	} `xml:"actions"`
}

// action.add
type XmlRequestActionAdd struct {
	XMLName xml.Name `xml:"request"`
	Method  string   `xml:"method,attr"`
	Account string   `xml:"account"`
	Sid     string   `xml:"sid"`

	Description    string               `xml:"action>description"`
	TaskId         int                  `xml:"action>task>id,omitempty"`
	TaskGeneral    int                  `xml:"action>task>general,omitempty"`
	ContactGeneral int                  `xml:"action>contact>general,omitempty"`
	TaskNewStatus  int                  `xml:"action>taskNewStatus,omitempty"`
	NotifiedList   []XmlResponseUser    `xml:"action>notifiedList>user,omitempty"`
	IsHidden       int                  `xml:"action>isHidden"`
	Owner          XmlResponseUser      `xml:"action>owner,omitempty"`
	DateTime       string               `xml:"action>dateTime,omitempty"`
	Analitics      []XmlRequestAnalitic `xml:"action>analitics>analitic,omitempty"`
}

// action.add response
type XmlResponseActionAdd struct {
	XMLName xml.Name `xml:"response"`
	Status  string   `xml:"status,attr"`
	Code    string   `xml:"code"`
	Message string   `xml:"message"`

	ActionId int `xml:"action>id"`
}

// analitic.getList
type XmlRequestAnaliticGetList struct {
	XMLName xml.Name `xml:"request"`
	Method  string   `xml:"method,attr"`
	Account string   `xml:"account"`
	Sid     string   `xml:"sid"`

	AnaliticGroupId int `xml:"analiticGroupId,omitempty"`
}

// analitic.getList response
type XmlResponseAnaliticGetList struct {
	XMLName xml.Name `xml:"response"`
	Status  string   `xml:"status,attr"`
	Code    string   `xml:"code"`
	Message string   `xml:"message"`

	Analitics struct {
		AnaliticsCount      int                   `xml:"count,attr"`
		AnaliticsTotalCount int                   `xml:"totalCount,attr"`
		Analitics           []XmlResponseAnalitic `xml:"analitic"`
	} `xml:"analitics"`
}

// analitic.getOptions
type XmlRequestAnaliticGetOptions struct {
	XMLName xml.Name `xml:"request"`
	Method  string   `xml:"method,attr"`
	Account string   `xml:"account"`
	Sid     string   `xml:"sid"`

	AnaliticId int `xml:"analitic>id"`
}

// analitic.getOptions response
type XmlResponseAnaliticGetOptions struct {
	XMLName xml.Name `xml:"response"`
	Status  string   `xml:"status,attr"`
	Code    string   `xml:"code"`
	Message string   `xml:"message"`

	Analitic XmlResponseAnaliticOptions `xml:"analitic"`
}

// task.get
type XmlRequestTaskGet struct {
	XMLName xml.Name `xml:"request"`
	Method  string   `xml:"method,attr"`
	Account string   `xml:"account"`
	Sid     string   `xml:"sid"`

	TaskId      int `xml:"task>id,omitempty"`
	TaskGeneral int `xml:"task>general,omitempty"`
}

// task.get response
type XmlResponseTaskGet struct {
	XMLName xml.Name `xml:"response"`
	Status  string   `xml:"status,attr"`
	Code    string   `xml:"code"`
	Message string   `xml:"message"`

	Task XmlResponseTask `xml:"task"`
}
