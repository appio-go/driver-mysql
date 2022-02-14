# Driver-MySQL
Helps you connect to MySQL DB, with a credentials provided in .env file


## Usage
Add APP_NAME="Some name" to project .env file

```
package main
import "github.com/appio-go/logger"

func main(){
	var l logger.Logger
	l.Print("Logname", "info", "/your/request/url", "192.168.0.1", "User identifier", "Hello world, im logger")
}
```
