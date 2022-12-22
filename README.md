###Building the project
cd to the go workspace.

`make build`

###Running the application
on the same terminal path run 

`./out/cron_parser <cron expression>"`

###Running tests
run command 

`make test`

###Project Setup
project is divided into 5 packages and subpackages:
1. cmd - houses the main func - the entry point of the application
2. expression_parser - the entry point of the parsing logic. this maps the expression fields to the actual types like minute, hour, month etc. and invokes the respective field parsers
3. field_parsers - this houses the logic of how to parse each type of time field like hour, month etc.
   1. common - common code to parse the field for multiple types and invokes special char handlers for each field type
   2. special char handlers - contains the speacial char handling within the application for each field type hour, month etc.
   3. time unit - all supported time unit in this application [hour, minutes, month, dayof week, day of month] 
4. output_formatter - this is responsible for handling the formatting for the output returned by the expression parser package
5. types - contains all the custom defined types (structs,interfaces) in this project. this ensures ni cyclic dependencies are formed as the project development evolves and types are used across packages.  
