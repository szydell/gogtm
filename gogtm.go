//Package gogtm enables multithreaded access to GT.M database
package gogtm

import "errors"
import "github.com/szydell/mstools"

type gtmenv struct {
	gbldir  string
	icu     string
	chset   string
	log     string
	tmp     string
	dist    string
	hugeTbl bool
}

func (env *gtmenv) checkParams() error {
	if env.gbldir == "" {
		return errors.New("location of database gld file not set")
	}
	if tmp, err := mstools.FileOrDir(env.gbldir); tmp != 'f' || err != nil {
		return errors.New("gbldir does not point to a file")
	}
	switch env.chset {
	case "M":
	case "UTF-8":
		if env.icu == "" {
			return errors.New("if using UTF-8 libicu must be set")
		}
	default:
		return errors.New("set proper chset (M or UTF-8)")
	}
	if env.chset == "UTF-8" && env.icu == "" {

	}
	if tmp, err := mstools.FileOrDir(env.log); tmp != 'd' || err != nil {
		return errors.New("location for db logs not set or is not a directory")
	}
	if tmp, err := mstools.FileOrDir(env.tmp); tmp != 'd' || err != nil {
		return errors.New("location for db temporary files not set or is not a directory")
	}
	if tmp, err := mstools.FileOrDir(env.dist); tmp != 'd' || err != nil {
		return errors.New("location of gt.m engine not set or is not a directory")
	}

	return nil
}

//Start do all needed checks and initiates connection to GT.M database
func Start(db gtmenv) {

}
