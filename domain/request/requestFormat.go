package request

type Notify struct {
	Phone        string `bson:"phone" json:"phone"`
	Organization string `bson:"organization" json:"organization"`
}

type IoMTStatus struct {
	Suspend bool `bson:"suspend" json:"suspend"`
}

// const (
// 	Running  = "Running"
// 	Pending  = "Pending"
// 	Finished = "Finished"
// 	Killed   = "Killed"
// )

// type DeviceInfo struct {
// 	IsBusy  bool `bson:"isBusy" json:"isBusy"`
// 	IsAlive bool `bson:"isAlive" json:"isAlive"`
// }

// type Task struct {
// 	TaskID    string `bson:"TaskID" json:"TaskID"`
// 	JobStatus string `bson:"JobStatus" json:"JobStatus"`
// 	Msg       string `bson:"Msg" json:"Msg"`
// 	Code      int    `bson:"Code" json:"Code"`
// }

// type Result struct {
// 	TaskID         string `bson:"TaskID" json:"TaskID"`
// 	JobStatus      string `bson:"JobStatus" json:"JobStatus"`
// 	Msg            string `bson:"Msg" json:"Msg"`
// 	Code           int    `bson:"Code" json:"Code"`
// 	AnalysisResult string `bson:"AnalysisResult" json:"AnalysisResult"`
// }
