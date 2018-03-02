package planfix

import (
	"encoding/xml"
)

// in all requests except auth.login
type XmlRequester interface {
	SetSid(sid string)
	SetAccount(account string)
	GetMethod() string
}
type XmlRequestAuth struct {
	XMLName xml.Name `xml:"request"`
	Method  string   `xml:"method,attr"`
	Account string   `xml:"account"`
	Sid     string   `xml:"sid"`
}

func (a *XmlRequestAuth) SetSid(sid string) {
	a.Sid = sid
}
func (a *XmlRequestAuth) SetAccount(account string) {
	a.Account = account
}
func (a *XmlRequestAuth) GetMethod() string {
	return a.Method
}

// in all responses
type XmlResponseStatus struct {
	XMLName xml.Name `xml:"response"`
	Status  string   `xml:"status,attr"`
	Code    string   `xml:"code"`
	Message string   `xml:"message"`
}

type XmlResponseFile struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type XmlResponseActionUser struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type XmlResponseActionAnalitic struct {
	ID   int    `xml:"id"`
	Key  int    `xml:"key"`
	Name string `xml:"name"`
}

type XmlResponseAnaliticOptions struct {
	ID      int                               `xml:"id"`
	Name    string                            `xml:"name"`
	GroupID int                               `xml:"group>id"`
	Fields  []XmlResponseAnaliticOptionsField `xml:"fields>field"`
}

type XmlResponseAnaliticOptionsField struct {
	ID         int      `xml:"id"`
	Num        int      `xml:"num"`
	Name       string   `xml:"name"`
	Type       string   `xml:"type"`
	ListValues []string `xml:"list>value"`
	HandbookID int      `xml:"handbook>id"`
}

type XmlResponseAnaliticHandbookRecord struct {
	Key       int                                       `xml:"key"`
	ParentKey int                                       `xml:"parentKey"`
	IsGroup   int                                       `xml:"isGroup"`
	Values    []XmlResponseAnaliticHandbookRecordValues `xml:"value"`
	ValuesMap map[string]string
}

type XmlResponseAnaliticHandbookRecordValues struct {
	Name        string `xml:"name,attr"`
	Value       string `xml:"value,attr"`
	IsDisplayed int    `xml:"isDisplayed,attr"`
}

type XmlResponseAction struct {
	ID                           int                         `xml:"id"`
	Description                  string                      `xml:"description"`
	OldStatus                    int                         `xml:"statusChange>oldStatus,omitempty"`
	NewStatus                    int                         `xml:"statusChange>newStatus,omitempty"`
	IsNotRead                    bool                        `xml:"isNotRead"`
	FromEmail                    bool                        `xml:"fromEmail"`
	DateTime                     string                      `xml:"dateTime"`
	TaskID                       int                         `xml:"task>id"`
	TaskTitle                    string                      `xml:"task>title"`
	ContactGeneral               int                         `xml:"contact>general"`
	ContactName                  string                      `xml:"contact>name"`
	Owner                        XmlResponseActionUser       `xml:"owner"`
	ProjectID                    int                         `xml:"project>id"`
	ProjectTitle                 string                      `xml:"project>title"`
	TaskExpectDateChangedOldDate string                      `xml:"taskExpectDateChanged>oldDate"`
	TaskExpectDateChangedNewDate string                      `xml:"taskExpectDateChanged>newDate"`
	TaskStartTimeChangedOldDate  string                      `xml:"taskStartTimeChanged>oldDate"`
	TaskStartTimeChangedNewDate  string                      `xml:"taskStartTimeChanged>newDate"`
	Files                        []XmlResponseFile           `xml:"files>file"`
	NotifiedList                 []XmlResponseActionUser     `xml:"notifiedList>user"`
	Analitics                    []XmlResponseActionAnalitic `xml:"analitics>analitic"`
}

// TODO: добавить все поля из https://planfix.ru/docs/ПланФикс_API_task.get
type XmlResponseTask struct {
	ID           int    `xml:"id"`
	Title        string `xml:"title"`
	Description  string `xml:"description"`
	General      int    `xml:"general"`
	ProjectID    int    `xml:"project>id"`
	ProjectTitle string `xml:"project>title"`
}

type XmlResponseAnalitic struct {
	ID        int    `xml:"id"`
	Name      string `xml:"name"`
	GroupID   int    `xml:"group>id"`
	GroupName string `xml:"group>name"`
}

type XmlRequestAnalitic struct {
	ID       int                       `xml:"id"`
	ItemData []XmlRequestAnaliticField `xml:"analiticData>itemData"`
}

type XmlRequestAnaliticField struct {
	FieldID int         `xml:"fieldId"`
	Value   interface{} `xml:"value"`
}

// TODO: добавить все поля из https://planfix.ru/docs/ПланФикс_API_user.get
type XmlResponseUser struct {
	ID       int    `xml:"id"`
	Name     string `xml:"name"`
	LastName string `xml:"lastName"`
	Login    string `xml:"login"`
	Email    string `xml:"email"`
}

// auth.login
type XmlRequestAuthLogin struct {
	XMLName xml.Name `xml:"request"`
	Method  string   `xml:"method,attr"`

	Account  string `xml:"account"`
	Login    string `xml:"login"`
	Password string `xml:"password"`
}

func (a *XmlRequestAuthLogin) SetSid(sid string) {}
func (a *XmlRequestAuthLogin) SetAccount(account string) {
	a.Account = account
}
func (a *XmlRequestAuthLogin) GetMethod() string {
	return a.Method
}

// auth.login response
type XmlResponseAuth struct {
	XMLName xml.Name `xml:"response"`
	Sid     string   `xml:"sid"`
}

// action.get
type XmlRequestActionGet struct {
	XmlRequestAuth
	XMLName  xml.Name `xml:"request"`
	ActionID int      `xml:"action>id"`
}

// action.get response
type XmlResponseActionGet struct {
	XMLName xml.Name          `xml:"response"`
	Action  XmlResponseAction `xml:"action"`
}

// action.getList
type XmlRequestActionGetList struct {
	XmlRequestAuth
	XMLName xml.Name `xml:"request"`

	TaskID         int    `xml:"task>id,omitempty"`
	TaskGeneral    int    `xml:"task>general,omitempty"`
	ContactGeneral int    `xml:"contact>general,omitempty"`
	PageCurrent    int    `xml:"pageCurrent"`
	PageSize       int    `xml:"pageSize"`
	Sort           string `xml:"sort"`
}

// action.getList response
type XmlResponseActionGetList struct {
	XMLName xml.Name `xml:"response"`
	Actions struct {
		ActionsCount      int                 `xml:"count,attr"`
		ActionsTotalCount int                 `xml:"totalCount,attr"`
		Actions           []XmlResponseAction `xml:"action"`
	} `xml:"actions"`
}

// action.add
type XmlRequestActionAdd struct {
	XmlRequestAuth
	XMLName xml.Name `xml:"request"`

	Description    string               `xml:"action>description"`
	TaskID         int                  `xml:"action>task>id,omitempty"`
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
	XMLName  xml.Name `xml:"response"`
	ActionID int      `xml:"action>id"`
}

// analitic.getList
type XmlRequestAnaliticGetList struct {
	XmlRequestAuth
	XMLName xml.Name `xml:"request"`

	AnaliticGroupID int `xml:"analiticGroupId,omitempty"`
}

// analitic.getList response
type XmlResponseAnaliticGetList struct {
	XMLName   xml.Name `xml:"response"`
	Analitics struct {
		AnaliticsCount      int                   `xml:"count,attr"`
		AnaliticsTotalCount int                   `xml:"totalCount,attr"`
		Analitics           []XmlResponseAnalitic `xml:"analitic"`
	} `xml:"analitics"`
}

// analitic.getOptions
type XmlRequestAnaliticGetOptions struct {
	XmlRequestAuth
	XMLName xml.Name `xml:"request"`

	AnaliticID int `xml:"analitic>id"`
}

// analitic.getOptions response
type XmlResponseAnaliticGetOptions struct {
	XMLName  xml.Name                   `xml:"response"`
	Analitic XmlResponseAnaliticOptions `xml:"analitic"`
}

// analitic.getHandbook
type XmlRequestAnaliticGetHandbook struct {
	XmlRequestAuth
	XMLName xml.Name `xml:"request"`

	HandbookID int `xml:"handbook>id"`
}

// analitic.getHandbook response
type XmlResponseAnaliticGetHandbook struct {
	XMLName xml.Name                            `xml:"response"`
	Records []XmlResponseAnaliticHandbookRecord `xml:"records>record"`
}

// task.get
type XmlRequestTaskGet struct {
	XmlRequestAuth
	XMLName xml.Name `xml:"request"`

	TaskID      int `xml:"task>id,omitempty"`
	TaskGeneral int `xml:"task>general,omitempty"`
}

// task.get response
type XmlResponseTaskGet struct {
	XMLName xml.Name        `xml:"response"`
	Task    XmlResponseTask `xml:"task"`
}

// user.get
type XmlRequestUserGet struct {
	XmlRequestAuth
	XMLName xml.Name `xml:"request"`

	UserID int `xml:"user>id,omitempty"`
}

// user.get response
type XmlResponseUserGet struct {
	XMLName xml.Name        `xml:"response"`
	User    XmlResponseUser `xml:"user"`
}
