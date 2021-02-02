package log

// LOGGER STRUCTURE DEFINITION
// LogBuilder is the interface that every concrete implementation should obey
type Detail interface {

	//SETTER DEFINITION
	// SetUUID sets the UUID created for the log
	SetUUID(uuID string)
	// SetIP sets the IP
	SetIP(IP string)
	// SetTimeStamp sets the provided timestamp in the format DD/MM/YYYY HH:MM:SS
	SetTimeStamp(timeStamp string)
	// SetServiceName sets the service name where the log is being generated
	SetServiceName(serviceName string)
	// SetHostname sets the name of the host instance
	SetHostname(hostName string)
	// SetRequestBody sets the request body provided by the client
	SetRequestBody(requestBody string)
	// SetResponseBody sets the response body coming from the redirected service
	SetResponseBody(responseBody string)
	// SetDestinationURL sets the destination URL
	SetDestinationURL(destinationURL string)
	// SetStep sets the step of the service being executed
	SetStep(step string)
	// SetLevel sets the log level: DEBUG, INFO, WARNING, ERROR
	SetLevel(level string)
	// SetProduct sets the product related to the flow where the log is generated
	SetProduct(product string)
	// SetApplication sets the application type where the log is generated
	SetApplication(application string)
	// SetClass sets the class where the log is generated
	SetClass(class string)
	// SetMethod sets the method inside the class where the log is generated
	SetMethod(method string)
	// SetLanguage sets the language
	SetLanguage(language string)
	// SetLogMessage sets the log message providing more specific details
	SetLogMessage(logMessage string)

	//GETTER DEFINITION
	// SetUUID sets the UUID created for the log
	GetUUID() (uuID string)
	// SetIP sets the IP
	GetIP() (IP string)
	// SetTimeStamp sets the provided timestamp in the format DD/MM/YYYY HH:MM:SS
	GetTimeStamp() (timeStamp string)
	// SetServiceName sets the service name where the log is being generated
	GetServiceName() (serviceName string)
	// SetHostname sets the name of the host instance
	GetHostname() (hostName string)
	// SetRequestBody sets the request body provided by the client
	GetRequestBody() (requestBody string)
	// SetResponseBody sets the response body coming from the redirected service
	GetResponseBody() (responseBody string)
	// SetDestinationURL sets the destination URL
	GetDestinationURL() (destinationURL string)
	// SetStep sets the step of the service being executed
	GetStep() (step string)
	// SetLevel sets the log level: DEBUG, INFO, WARNING, ERROR
	GetLevel() (level string)
	// SetProduct sets the product related to the flow where the log is generated
	GetProduct() (product string)
	// SetApplication sets the application type where the log is generated
	GetApplication() (application string)
	// SetClass sets the class where the log is generated
	GetClass() (class string)
	// SetMethod sets the method inside the class where the log is generated
	GetMethod() (method string)
	// SetLanguage sets the language
	GetLanguage() (language string)
	// SetThread sets thread where the request is running
	GetLogMessage() (logMessage string)
}
