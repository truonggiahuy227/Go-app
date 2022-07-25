package constant

const (
	LogBackEndPrefix                       string = "[BACKEND]"
	LogBackEndModulePrefix                 string = "[arrow-service]"
	LogBackEndMainInfoPrefix               string = "[Main_Info]"
	LogBackEndMainErrorPrefix              string = "[Main_Error]"
	LogBackEndWorkerInitNewUserInfoPrefix  string = "[WorkerInitNewUser_Info]"
	LogBackEndWorkerInitNewUserErrorPrefix string = "[WorkerInitNewUser_Error]"
)

const (
	LogInfoPrefix                   = LogBackEndPrefix + LogBackEndModulePrefix + LogBackEndMainInfoPrefix
	LogErrorPrefix                  = LogBackEndPrefix + LogBackEndModulePrefix + LogBackEndMainErrorPrefix
	LogWorkerInitNewUserInfoPrefix  = LogBackEndPrefix + LogBackEndModulePrefix + LogBackEndWorkerInitNewUserInfoPrefix
	LogWorkerInitNewUserErrorPrefix = LogBackEndPrefix + LogBackEndModulePrefix + LogBackEndWorkerInitNewUserErrorPrefix
)
