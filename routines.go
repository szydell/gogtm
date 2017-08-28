package gogtm

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/szydell/mstools"
)

func createRoutines(workDir string) {
	// prepare paths
	path, pathO, pathR := generatePaths(workDir)
	// create directory tree for routines
	err := os.MkdirAll(pathR, os.ModePerm)
	mstools.ErrCheck(err) // drop log and fatal close on error
	// create path for objects
	err = os.Mkdir(pathO, os.ModePerm)
	mstools.ErrCheck(err) // drop log and fatal close on error

	// get already configured 'gtmroutines' from environment
	routines := os.Getenv("gtmroutines")
	// concatenate old value with path created for this session
	routines += " " + pathO + "(" + pathR + ")"

	// create file with routines
	generateRoutineFile(pathR)

	// set 'gtmroutines' env variable to access internal gogtm file with routines
	os.Setenv("gtmroutines", routines)

	// prepare path for gtmaccess.ci
	ciPath := filepath.Join(path, "gtmaccess.ci")
	// generate gtmaccess.ci
	generateCiFile(ciPath)
	// set 'GTMCI' env variable to access interface file needed by gt.m api
	os.Setenv("GTMCI", ciPath)

}

func generatePaths(workDir string) (path string, pathO string, pathR string) {
	// add unique directory name for this session (do not mix routines between sessions)
	path = filepath.Join(workDir, "gogtm/"+goSessionID)
	// create directories 'o' for objects, 'r' for routines
	pathR = path + "/r"
	pathO = path + "/o"
	return
}

func cleanRoutines(workDir string) {
	path, _, _ := generatePaths(workDir)

	os.RemoveAll(path)
}

func generateCiFile(path string) {
	data := []byte(`gtminit   : void init^%gtmaccess( O:gtm_char_t* )
gtmdata   : void data^%gtmaccess( I:gtm_char_t*, O:gtm_uint_t*, O:gtm_char_t* )
gtmget    : void get^%gtmaccess( I:gtm_char_t*, I:gtm_string_t*, O:gtm_char_t*, O:gtm_char_t* )
gtmkill   : void kill^%gtmaccess( I:gtm_char_t*, O:gtm_char_t* )
gtmorder  : void order^%gtmaccess( I:gtm_char_t*, I:gtm_char_t*, O:gtm_char_t*, O:gtm_char_t* )
gtmquery  : void query^%gtmaccess( I:gtm_char_t*, O:gtm_char_t*, O:gtm_char_t* )
gtmset    : void set^%gtmaccess( I:gtm_char_t*, I:gtm_string_t*, O:gtm_char_t*)
gtmxecute : void xecute^%gtmaccess( I:gtm_char_t*, O:gtm_char_t*, O:gtm_char_t* )
gtmzkill  : void zkill^%gtmaccess( I:gtm_char_t*, O:gtm_char_t* )
gvstat    : void gvstat^%gtmaccess( O:gtm_char_t*, O:gtm_char_t* )
`)

	err := ioutil.WriteFile(path, data, 0400)
	mstools.ErrCheck(err)
}

func generateRoutineFile(path string) {
	// routines internally used by gogtm. M language.
	data := []byte(`%gtmaccess    ; entry points to access GT.M
    quit
    ;
init(error)
    set $ztrap="new tmp set error=$ecode set tmp=$piece($ecode,"","",2) quit:$quit $extract(tmp,2,$length(tmp)) quit"
    quit:$quit 0 quit
    ;
data(var,value,error)
	set value=$data(@var)
	quit:$quit 0 quit
	;
get(var,opt,value,error)
    set value=$GET(@var,opt)
    quit:$quit 0 quit
    ;
gvstat(stats,error)
    N RET
    S REGION=$V("GVFIRST") S RET=REGION_"->"_$V("GVSTAT",REGION)
    F I=1:1 S REGION=$V("GVNEXT",REGION) Q:REGION=""  S RET=RET_"|"_REGION_"->"_$V("GVSTAT",REGION)
    set stats=RET
    ;
kill(var,error)
    kill @var
    quit:$quit 0 quit
    ;
order(var,dir,value,error)
    set value=$order(@var,dir) 
    quit:$quit 0 quit
    ;
query(var,value,error)
    set value=$query(@var)
    quit:$quit 0 quit
    ;
set(var,value,error)
    set @var=value
    quit:$quit 0 quit
    ;
xecute(code,value,error)
    xecute code
    quit:$quit 0 quit
    ;
zkill(var,error)
    zkill @var
    quit:$quit 0 quit
    ;
`)
	path = filepath.Join(path, "_gtmaccess.m")
	err := ioutil.WriteFile(path, data, 0400)
	mstools.ErrCheck(err)
}
