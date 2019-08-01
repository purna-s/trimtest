package trimtest

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// ActivityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-trimtest")

// MyActivity is a stub for your Activity implementation
type XMLParserActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &XMLParserActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *XMLParserActivity) Metadata() *activity.Metadata {
	return a.metadata
}


// Eval implements activity.Activity.Eval
func (a *XMLParserActivity) Eval(ctx activity.Context) (done bool, err error) {

	XMLString := ctx.GetInput("xmlString").(string)

	activityLog.Debugf("XML String is : [%s]", XMLString)
	//fmt.Println("XML String is : ", XMLString)

	if len(XMLString) == 0 {
		activityLog.Debugf("value in the field is empty ")
		//fmt.Println("value in  the field is empty ")

	}
	//	XMLString = (string(XMLString))

	trimstring := strings.TrimSpace(xmlString)

	//fmt.Println(" JSON String ")
	//fmt.Println(string(jsondata))

	// Set the output as part of the context
	activityLog.Debugf("Activity has trimmed spaces Successfully")
	fmt.Println("Activity has trimmed spaces Successfully")

	ctx.SetOutput("output", string(trimstring))

	return true, nil
}
