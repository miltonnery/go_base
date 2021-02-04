package errorhandling

/*
Copyright 2020 LifeMiles integration team
Use of this source code is governed by internal error code standards made by DEV team

Error categories:
1. Technical errors: from 0   to 500
2. Business  errors: from 501 to 1000

1. Technical errors:
	Intended to describe errors whose origin comes from sources such as:
		- I/O : File reading, file modifications and so on.
		- Persistence: Database persistence connection failures, CRUD functions.
		- Use of specific libraries: All types of errors coming from libraries used in the solution.
		- Connections: Consumption of other services, and so on.
		- All time generic errors: Unknown errors,
		- Test errors: Any type of error defined for unit tests.
		- Technical security: Error reading a security header because it is malformed.

2. Business errors:
	Intended to express all the logical or business errors defined or found for the required solution.
	Example:
		- JSON contract is incorrect.
		- The resulting value from a request is not correct.
		- Logical security: the security header content is incorrect or invalid.

All these errors are open to redefinition and they can vary from service to service.
*/

// INTERNAL ERROR CODES
const (
	// TECHNICAL ERRORS ------------------------------------------------------------------------------------------------|

	// Basic data type
	BasicParsingProblem = 10
	BasicJSONUnMarshall = 15
	BasicJSONMarshall   = 20
	BasicEmptyParameter = 30

	// IO -> File
	IOFileBaseConfigurationNotFound        = 25
	IOFileEnvironmentConfigurationNotFound = 26

	// IO -> Viper
	IOViperUnmarshalProblem = 50

	// Persistence
	PersistenceConnectionLost = 100
	PersistenceRecordNotFound = 110

	// Gateway
	GatewayTimeoutDeadlineExceeded = 200
	GatewayRedirectionError        = 201

	// Security

	// Generic errors
	GenericUnknownError = 400
	GenericTestError    = 450

	// BUSINESS ERRORS -------------------------------------------------------------------------------------------------|

	GenericBusinessTestError = 552
)

// INTERNAL ERROR DESCRIPTIONS
var statusText = map[int]string{

	// Technical Errors ------------------------------------------------------------------------------------------------|

	// Basic data type
	BasicParsingProblem: "basic: there is a problem while trying to parse the data type",
	BasicJSONUnMarshall: "basic: problems trying to un-marshall json into struct",
	BasicJSONMarshall:   "basic: problems trying to marshall struct into json",
	BasicEmptyParameter: "basic: the required parameter is empty",

	//IO -> File
	IOFileBaseConfigurationNotFound:        "file: base configuration not found",
	IOFileEnvironmentConfigurationNotFound: "file: configuration configuration not found",

	//IO -> Viper
	IOViperUnmarshalProblem: "viper: problems detected while trying to unmarshal properties into specified struct",

	//Persistence
	PersistenceConnectionLost: "persistence: connection lost",
	PersistenceRecordNotFound: "persistence: the requested record does not exists",

	//Gateway
	GatewayTimeoutDeadlineExceeded: "gateway: Waiting time deadline exceeded",
	GatewayRedirectionError:        "gateway: Waiting time deadline exceeded",

	//Security

	//Generic
	GenericTestError:    "unit-test: there is something wrong with the unit test :c",
	GenericUnknownError: "generic-errors: unknown error",

	// Business Errors -------------------------------------------------------------------------------------------------|
	GenericBusinessTestError: "generic-business: generic business test error",
}

// ErrorDescription returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func ErrorDescription(code int) string {
	return statusText[code]
}
