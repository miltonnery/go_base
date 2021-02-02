package json

// JSON LOG BUILDER IMPLEMENTATION

//Struct for logs formatted in JSON
type LogDetailsJSON struct {
	UUID           string `json:"uuid"`           //Unique tracking identifier
	IP             string `json:"ip"`             //Remote client IP
	TimeStamp      string `json:"timestamp"`      //Log creation time: DD/MM/YYYY HH:MM:SS
	ServiceName    string `json:"service"`        //Microservice name
	Hostname       string `json:"hostname"`       //Name of pod
	RequestBody    string `json:"request"`        //Client request
	ResponseBody   string `json:"response"`       //Redirected service response
	DestinationURL string `json:"destinationUrl"` //Requested path before redirection
	Step           string `json:"step"`           //Paso o contexto que se está ejecutando (cotización, pago, envío de correo, etc.)
	Level          string `json:"level"`          //Log level: Debug, Info, Warn or Error
	Product        string `json:"product"`        // LifeMiles product
	Application    string `json:"application"`    //Application name
	Class          string `json:"class"`          //requested service name
	Method         string `json:"method"`         // response status
	Language       string `json:"language"`       // reason of response
	LogMessage     string `json:"logMessage"`     // extra details message
}

func NewLogDetailsJSON() *LogDetailsJSON {
	return &LogDetailsJSON{
		UUID:           "",
		IP:             "",
		TimeStamp:      "",
		ServiceName:    "",
		Hostname:       "",
		RequestBody:    "",
		ResponseBody:   "",
		DestinationURL: "",
		Step:           "",
		Level:          "",
		Product:        "",
		Application:    "",
		Class:          "",
		Method:         "",
		Language:       "",
		LogMessage:     "",
	}
}

func (lmlJSON *LogDetailsJSON) SetUUID(uuID string) {
	lmlJSON.UUID = uuID
}

func (lmlJSON *LogDetailsJSON) SetIP(IP string) {
	lmlJSON.IP = IP
}

func (lmlJSON *LogDetailsJSON) SetTimeStamp(timeStamp string) {
	lmlJSON.TimeStamp = timeStamp
}

func (lmlJSON *LogDetailsJSON) SetServiceName(serviceName string) {
	lmlJSON.ServiceName = serviceName
}

func (lmlJSON *LogDetailsJSON) SetHostname(hostName string) {
	lmlJSON.Hostname = hostName
}

func (lmlJSON *LogDetailsJSON) SetRequestBody(requestBody string) {
	lmlJSON.RequestBody = requestBody
}

func (lmlJSON *LogDetailsJSON) SetResponseBody(responseBody string) {
	lmlJSON.ResponseBody = responseBody
}

func (lmlJSON *LogDetailsJSON) SetDestinationURL(destinationURL string) {
	lmlJSON.DestinationURL = destinationURL
}

func (lmlJSON *LogDetailsJSON) SetStep(step string) {
	lmlJSON.Step = step
}

func (lmlJSON *LogDetailsJSON) SetLevel(level string) {
	lmlJSON.Level = level
}

func (lmlJSON *LogDetailsJSON) SetProduct(product string) {
	lmlJSON.Product = product
}

func (lmlJSON *LogDetailsJSON) SetApplication(application string) {
	lmlJSON.Application = application
}

func (lmlJSON *LogDetailsJSON) SetClass(class string) {
	lmlJSON.Class = class
}

func (lmlJSON *LogDetailsJSON) SetMethod(method string) {
	lmlJSON.Method = method
}

func (lmlJSON *LogDetailsJSON) SetLanguage(language string) {
	lmlJSON.Language = language
}

func (lmlJSON *LogDetailsJSON) SetLogMessage(logMessage string) {
	lmlJSON.LogMessage = logMessage
}

func (lmlJSON *LogDetailsJSON) GetUUID() (uuID string) {
	uuID = lmlJSON.UUID
	return
}

func (lmlJSON *LogDetailsJSON) GetIP() (IP string) {
	IP = lmlJSON.IP
	return
}

func (lmlJSON *LogDetailsJSON) GetTimeStamp() (timeStamp string) {
	timeStamp = lmlJSON.TimeStamp
	return
}

func (lmlJSON *LogDetailsJSON) GetServiceName() (serviceName string) {
	serviceName = lmlJSON.ServiceName
	return
}

func (lmlJSON *LogDetailsJSON) GetHostname() (hostName string) {
	hostName = lmlJSON.Hostname
	return
}

func (lmlJSON *LogDetailsJSON) GetRequestBody() (requestBody string) {
	return lmlJSON.RequestBody
}

func (lmlJSON *LogDetailsJSON) GetResponseBody() (responseBody string) {
	return lmlJSON.ResponseBody
}

func (lmlJSON *LogDetailsJSON) GetDestinationURL() (destinationURL string) {
	return lmlJSON.DestinationURL
}

func (lmlJSON *LogDetailsJSON) GetStep() (step string) {
	return lmlJSON.Step
}

func (lmlJSON *LogDetailsJSON) GetLevel() (level string) {
	return lmlJSON.Level
}

func (lmlJSON *LogDetailsJSON) GetProduct() (product string) {
	return lmlJSON.Product
}

func (lmlJSON *LogDetailsJSON) GetApplication() (application string) {
	return lmlJSON.Application
}

func (lmlJSON *LogDetailsJSON) GetClass() (class string) {
	return lmlJSON.Class
}

func (lmlJSON *LogDetailsJSON) GetMethod() (method string) {
	return lmlJSON.Method
}

func (lmlJSON *LogDetailsJSON) GetLanguage() (language string) {
	return lmlJSON.Language
}

func (lmlJSON *LogDetailsJSON) GetLogMessage() (logMessage string) {
	return lmlJSON.LogMessage
}
