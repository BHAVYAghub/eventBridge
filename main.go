package main

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eventbridge"
	"log"
	"time"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1")}))
	ebride := eventbridge.New(sess, aws.NewConfig().WithRegion("us-east-1"))

	//createEventBus(ebride)
	//createRule(ebride)
	//
	//createTarget(ebride)
	sendEvent(ebride)
}

func sendEvent(ebride *eventbridge.EventBridge) {
	cust := struct {
		Message string `json:"message"`
	}{
		Message: "Sample",
	}

	byt, _ := json.Marshal(cust)
	putInput := &eventbridge.PutEventsInput{
		Entries: []*eventbridge.PutEventsRequestEntry{
			{
				Detail:       aws.String(string(byt)),
				DetailType:   aws.String("myDetailType"),
				EventBusName: aws.String("default"),
				Time:         aws.Time(time.Now()),
				Source:       aws.String("application"),
			},
		},
	}
	log.Println(ebride.PutEvents(putInput))
}

func createEventBus(ebride *eventbridge.EventBridge) {
	in := &eventbridge.CreateEventBusInput{
		EventSourceName: nil,
		Name:            aws.String("Gofr-Bhavya"),
		Tags:            nil,
	}

	log.Println(ebride.CreateEventBus(in))
}

func createRule(ebride *eventbridge.EventBridge) {
	xy := struct {
		Source []string `json:"source"`
	}{
		Source: []string{"application"},
	}

	byt,_ := json.Marshal(xy)
	str := string(byt)
	log.Println(str)
	//str := "{\n\t\"source\":\t[\"application\"],\n\t\"detail-type\":\t[\"myDetailType\"]\n}"
	// create rule
	log.Println(ebride.PutRule(&eventbridge.PutRuleInput{
		EventBusName: aws.String("default"),
		EventPattern: aws.String("{ \"source\": [\"application\"] }"),
		Name:         aws.String("testRule-Bhavya"),
		ScheduleExpression: aws.String("rate(5 minutes)"),
	}))

	//createTarget(ebride)
}

func createTarget(ebride *eventbridge.EventBridge) {
	log.Println(ebride.PutTargets(&eventbridge.PutTargetsInput{
		EventBusName: aws.String("default"),
		Rule:         aws.String("testRule-Bhavya"),
		Targets: []*eventbridge.Target{
			{
				Arn: aws.String("<lambda arn>"),
				Id:  aws.String("1"),
			},
		},
	}))
}
//{
//"source": ["application"],
//"detail-type": ["myDetailType"]
//}
//{
//"source": ["application"],
//"detail-type": ["myDetailType"]
//}
