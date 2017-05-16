//Package gogtm enables multithreaded access to GT.M database
package gogtm

type gtmenv struct {
	gtmgbldir     string
	gtmIcuVersion string
	gtmChset      string
	gtmLog        string
	gtmTmp        string
	gtmDist       string
	hugeTbl       bool
}

type gogtmenv struct {
	db gtmenv
}

//Start initiates connection to GT.M database
func Start() {

}
