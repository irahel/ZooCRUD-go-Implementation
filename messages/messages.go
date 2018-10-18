package messages

const(
	//Prefix
	LOG_PREFIX string = "[LOG]"
	CLI_PREFIX string = "[Client] "
	ERROR_PREFIX string = "[ERROR ] "
	SERVER_PREFIX string = "[Server] "

	//Message
		//Connection
	MESSAGE_CONNECTED string = "Connection successfully ESTABLISHED "
	MESSAGE_DISCONNECTED string = "Connection successfully STOPPED "
	MESSAGE_CONNECT_ERROR string = "Connection error "
		//Commands
	MESSAGE_COMMAND_EMPTY string = "Empty command"
	MESSAGE_COMMAND_MANY_ARGUMMENTS string = "Command have many argumments"
	MESSAGE_COMMAND_FEW_ARGUMMENTS string = "Command have few argumments"
			//Create    
	MESSAGE_CREATE_SUCESS string = "Node created successfully"
	MESSAGE_CREATE_FAILED string = "Node created failed"
			//Read
	MESSAGE_EXISTS_TRUE string = "Node exists"
	MESSAGE_EXISTS_FALSE string = "Node NOT exists"
			//List
	MESSAGE_LIST string = " childrens path: "
			//Update
	MESSAGE_UPDATE_SUCESS string = "Node updated successfully"
	MESSAGE_UPDATE_FAILED string = "Node NOT updated"
			//Delete
	MESSAGE_DELETE_SUCESS string = "Node deleted successfully"
	MESSAGE_DELETE_FAILED string = "Node NOT deleted"
		//Usage
			//Incorrect
	MESSAGE_INCORRECT_USAGE_LINE_COMMAND string = "Incorrect calling line command"
	MESSAGE_INCORRECT_USAGE_COMMAND string = "Incorrect calling command"
			//Best Usage
	MESSAGE_USAGE_LINE_COMMAND string = "Usage:'''go run Cli.go <Server IP> <Server Port>'''"
	MESSAGE_USAGE_CREATE_COMMAND string = "Usage:'''CREATE/UPDATE <path> <value>'''"
	MESSAGE_USAGE_READ_COMMAND string = "Usage:'''READ/EXISTS/LIST <path> <value>'''"
		//Erros
	MESSAGE_BAD_VERSION string = "Version doesn’t match"
)