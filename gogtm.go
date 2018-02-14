//Package gogtm enables access to gt.m database
package gogtm

/*
#cgo CFLAGS: -I/opt/fis-gtm/engine
#cgo LDFLAGS: -L/opt/fis-gtm/engine -lgtmshr
#include <gtmxc_types.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#define maxstr 1048576

#ifndef NULL
#define NULL ((void *) 0)
#endif

#define CALLGTM(xyz) status = xyz ;		\
	if (0 != status ) {				\
		gtm_zstatus( msg, 2048 );			\
		snprintf(errmsg, 2048, "Failure of %s with error: %s\n", #xyz, msg); \
		return (int) status; \
	}


int cip_init(char *errmsg, int maxmsglen) {
	gtm_char_t msg[maxmsglen], err[maxmsglen];
	gtm_string_t gtminit_str;
	ci_name_descriptor gtminit;
	gtm_status_t status;


	gtminit_str.address = "gtminit";
	gtminit_str.length = sizeof("gtminit")-1;
	gtminit.rtn_name=gtminit_str;
	gtminit.handle = NULL;

	errmsg[0] = '\0';
	err[0] = '\0';

	CALLGTM (gtm_cip( &gtminit, &err));

	if (0 != strlen( err )){
		snprintf(errmsg, maxmsglen, "cip_init error: [%s]\n", err);
		return 100;
	}
	return 0;
} // end of cip_init

int cip_set(char *s_global, char *s_value, char *errmsg, int maxmsglen) {
	gtm_char_t err[maxmsglen], msg[maxmsglen];
	gtm_string_t gtmset_str, p_value;
	ci_name_descriptor gtmset;
	gtm_status_t status;

	gtmset_str.address = "gtmset";
	gtmset_str.length = sizeof("gtmset")-1;
	gtmset.rtn_name=gtmset_str;
	gtmset.handle = NULL;

	err[0] = '\0';

	p_value.address = ( gtm_char_t *) s_value; p_value.length = strlen(s_value);
	CALLGTM( gtm_cip( &gtmset, s_global, &p_value, &err));

	if (0 != strlen( err )){
		snprintf(errmsg, maxmsglen, "cip_set error: [%s]\n", err);
		fprintf( stderr, "error set: %s", err);
		return 100;
	}
	return 0;
} // end of cip_set

int cip_get(char *s_global, char *s_opt, char *s_ret, char *errmsg, int maxmsglen, int maxretlen) {
	gtm_char_t err[maxmsglen], msg[maxmsglen];
	gtm_string_t gtmget_str, p_opt;
	ci_name_descriptor gtmget;
	gtm_status_t status;

	gtmget_str.address = "gtmget";
	gtmget_str.length = sizeof("gtmget")-1;
	gtmget.rtn_name=gtmget_str;
	gtmget.handle = NULL;

	err[0] = '\0';

	p_opt.address = ( gtm_char_t *) s_opt; p_opt.length = strlen(s_opt);

	CALLGTM( gtm_cip( &gtmget, s_global, &p_opt, s_ret, &err));

	if (0 != strlen( err )){
		snprintf(errmsg, maxmsglen, "cip_get error: [%s]\n", err);
		fprintf( stderr, "error set: %s", err);
		return 100;
	}
	return 0;
} // end of cip_get

int cip_gvstat(char *errmsg, char *s_ret, int maxmsglen) {
	gtm_char_t msg[maxmsglen], err[maxmsglen];
	gtm_string_t gtmgvstat_str;
	ci_name_descriptor gtmgvstat;
	gtm_status_t status;


	gtmgvstat_str.address = "gvstat";
	gtmgvstat_str.length = sizeof("gtmgvstat")-1;
	gtmgvstat.rtn_name=gtmgvstat_str;
	gtmgvstat.handle = NULL;

	errmsg[0] = '\0';
	err[0] = '\0';

	CALLGTM (gtm_cip( &gtmgvstat, s_ret, &err));

	if (0 != strlen( err )){
		snprintf(errmsg, maxmsglen, "cip_gvstat error: [%s]\n", err);
		return 100;
	}
	return 0;
} // end of cip_gvstat

int cip_kill(char *s_global, char *errmsg, int maxmsglen) {
	gtm_char_t err[maxmsglen], msg[maxmsglen];
	gtm_string_t gtmkill_str;
	ci_name_descriptor gtmkill;
	gtm_status_t status;

	gtmkill_str.address = "gtmkill";
	gtmkill_str.length = sizeof("gtmkill")-1;
	gtmkill.rtn_name=gtmkill_str;
	gtmkill.handle = NULL;

	err[0] = '\0';

	CALLGTM( gtm_cip( &gtmkill, s_global, &err));

	if (0 != strlen( err )){
		snprintf(errmsg, maxmsglen, "cip_kill error: [%s]\n", err);
		fprintf( stderr, "error set: %s", err);
		return 100;
	}
	return 0;
} // end of cip_kill

int cip_zkill(char *s_global, char *errmsg, int maxmsglen) {
	gtm_char_t err[maxmsglen], msg[maxmsglen];
	gtm_string_t gtmzkill_str;
	ci_name_descriptor gtmzkill;
	gtm_status_t status;

	gtmzkill_str.address = "gtmzkill";
	gtmzkill_str.length = sizeof("gtmzkill")-1;
	gtmzkill.rtn_name=gtmzkill_str;
	gtmzkill.handle = NULL;

	err[0] = '\0';

	CALLGTM( gtm_cip( &gtmzkill, s_global, &err));

	if (0 != strlen( err )){
		snprintf(errmsg, maxmsglen, "cip_zkill error: [%s]\n", err);
		fprintf( stderr, "error set: %s", err);
		return 100;
	}
	return 0;
} // end of cip_zkill

int cip_xecute(char *s_global, char *errmsg, int maxmsglen) {
	gtm_char_t err[maxmsglen], msg[maxmsglen];
	gtm_string_t gtmxecute_str;
	ci_name_descriptor gtmxecute;
	gtm_status_t status;

	gtmxecute_str.address = "gtmxecute";
	gtmxecute_str.length = sizeof("gtmxecute")-1;
	gtmxecute.rtn_name=gtmxecute_str;
	gtmxecute.handle = NULL;

	err[0] = '\0';

	CALLGTM( gtm_cip( &gtmxecute, s_global, &err));

	if (0 != strlen( err )){
		snprintf(errmsg, maxmsglen, "cip_xecute error: [%s]\n", err);
		fprintf( stderr, "error set: %s", err);
		return 100;
	}
	return 0;
} // end of cip_xecute

int cip_order(char *s_global, char *s_ret, char *errmsg, int maxmsglen, int maxretlen, char *direction) {
        gtm_char_t err[maxmsglen], msg[maxmsglen];
        gtm_string_t gtmorder_str, p_opt;
        ci_name_descriptor gtmorder;
        gtm_status_t status;

        gtmorder_str.address = "gtmorder";
        gtmorder_str.length = sizeof("gtmorder")-1;
        gtmorder.rtn_name=gtmorder_str;
        gtmorder.handle = NULL;

        err[0] = '\0';

        CALLGTM( gtm_cip( &gtmorder, s_global, direction, s_ret, &err));

        if (0 != strlen( err )){
                snprintf(errmsg, maxmsglen, "cip_order error: [%s]\n", err);
                fprintf( stderr, "error set: %s", err);
                return 100;
        }
        return 0;
} // end of cip_order


int cip_query(char *s_global, char *s_ret, char *errmsg, int maxmsglen, int maxretlen) {
        gtm_char_t err[maxmsglen], msg[maxmsglen];
        gtm_string_t gtmquery_str, p_opt;
        ci_name_descriptor gtmquery;
        gtm_status_t status;

        gtmquery_str.address = "gtmquery";
        gtmquery_str.length = sizeof("gtmquery")-1;
        gtmquery.rtn_name=gtmquery_str;
        gtmquery.handle = NULL;

        err[0] = '\0';

        CALLGTM( gtm_cip( &gtmquery, s_global, s_ret, &err));

        if (0 != strlen( err )){
                snprintf(errmsg, maxmsglen, "cip_query error: [%s]\n", err);
                fprintf( stderr, "error set: %s", err);
                return 100;
        }
        return 0;
} // end of cip_query

int cip_data(char *s_global, unsigned int *i_ret, char *errmsg, int maxmsglen, int maxretlen) {
        gtm_char_t err[maxmsglen], msg[maxmsglen];
        gtm_string_t gtmdata_str, p_opt;
        ci_name_descriptor gtmdata;
        gtm_status_t status;

        gtmdata_str.address = "gtmdata";
        gtmdata_str.length = sizeof("gtmdata")-1;
        gtmdata.rtn_name=gtmdata_str;
        gtmdata.handle = NULL;

        err[0] = '\0';

        CALLGTM( gtm_cip( &gtmdata, s_global, i_ret, &err));

        if (0 != strlen( err )){
                snprintf(errmsg, maxmsglen, "cip_data error: [%s]\n", err);
                fprintf( stderr, "error set: %s", err);
                return 100;
        }
        return 0;
} // end of cip_data


*/
import "C"

import (
	"bytes"
	"errors"
	"os"
	"sync"
	"unsafe"

	"github.com/satori/go.uuid"

	"golang.org/x/crypto/ssh/terminal"
)

// maxmsglen maximum length of message from gt.m
const maxmsglen = 2048

// maxretlen maximum length of value retrieved from gt.m
const maxretlen = 1048576

// root working directory location
const workDir = "/tmp"

//global variables to store state of the terminal before gt.m init
var fd uintptr
var termAtStart *terminal.State

//global mutex to avoid parallel access to database (which is unsupported by used C API from gt.m)
var mu = &sync.Mutex{}

// generate goSessionId for multiple purposes
var u, err = uuid.NewV4()
var goSessionID = u.String()

//var goSessionID = uuid.NewV4().String()

//Set saves value to global in gt.m db
//Sample usage: gogtm.Set("^test","value")
func Set(global string, val string) error {

	if len(global) < 1 {
		return errors.New("Set failed - you must provide glvn")
	}

	_global := C.CString(global)
	_val := C.CString(val)
	errmsg := make([]byte, maxmsglen)

	defer C.free(unsafe.Pointer(_global))
	defer C.free(unsafe.Pointer(_val))
	mu.Lock()
	result := C.cip_set(_global, _val, (*C.char)(unsafe.Pointer(&errmsg[0])), C.int(len(errmsg)))
	mu.Unlock()
	if result != 0 {
		return errors.New("Set failed: " + string(result) + "Error message: " + string(errmsg))
	}
	return nil
} // end of Set

//Get returns the value of provided glvn if it has a value.
//If the variable has no value, the function returns a value specified by an optional second argument, and otherwise returns an empty string.
//Sample usage: gogtm.Get("test","nil")
func Get(global string, opt string) (string, error) {

	if len(global) < 1 {
		return "", errors.New("Get failed - you must provide glvn")
	}

	_global := C.CString(global)
	_opt := C.CString(opt)
	_ret := make([]byte, maxretlen)
	errmsg := make([]byte, maxmsglen)
	defer C.free(unsafe.Pointer(_global))
	defer C.free(unsafe.Pointer(_opt))

	p := C.malloc(C.size_t(maxmsglen))
	defer C.free(p)
	mu.Lock()
	result := C.cip_get(_global, _opt, (*C.char)(unsafe.Pointer(&_ret[0])), (*C.char)(unsafe.Pointer(&errmsg[0])), C.int(len(errmsg)), maxretlen)
	mu.Unlock()

	if result != 0 {
		return "", errors.New("Get failed: " + string(result) + "Error message: " + string(errmsg))
	}
	_ret = bytes.Trim(_ret, "\x00") // trim unused space
	return string(_ret), nil
} //end of Get

//GvStat returns gvstat for all regions
//Sample usage: gogtm.GvStat()
func GvStat() (string, error) {

	_ret := make([]byte, maxretlen)
	errmsg := make([]byte, maxmsglen)

	p := C.malloc(C.size_t(maxmsglen))
	defer C.free(p)
	mu.Lock()

	result := C.cip_gvstat((*C.char)(unsafe.Pointer(&errmsg[0])), (*C.char)(unsafe.Pointer(&_ret[0])), maxretlen)
	mu.Unlock()

	if result != 0 {
		return "", errors.New("Gvstat failed: " + string(result) + "Error message: " + string(errmsg))
	}
	_ret = bytes.Trim(_ret, "\x00") // trim unused space
	return string(_ret), nil
} //end of Gvstat

//Start should be used as the initiator of connection to gt.m
func Start() error {

	createRoutines(workDir)

	fd = os.Stdin.Fd()
	termAtStart, _ = terminal.GetState(int(fd))
	{
		result := C.gtm_init()
		if result != 0 {
			return errors.New("gtm_init failed: " + string(result))
		}
	}
	errmsg := C.CString("")
	defer C.free(unsafe.Pointer(errmsg))
	mu.Lock()
	result := C.cip_init(errmsg, maxmsglen)
	mu.Unlock()
	if result != 0 {
		return errors.New("CIP Init failed: " + string(result) + "Error MSG: " + C.GoString(errmsg))
	}
	return nil
} // end of Start

//Stop closes the connection gently.
func Stop() error {

	mu.Lock()
	result := C.gtm_exit()
	mu.Unlock()
	if result != 0 {
		return errors.New("gtm_exit failed: " + string(result))
	}
	cleanRoutines(workDir)
	terminal.Restore(int(fd), termAtStart)
	return nil
} // end of Stop

//Kill deletes global variable and its descendant nodes
func Kill(global string) error {

	if len(global) < 1 {
		return errors.New("Kill failed - you must provide [glvn | (glvn[,...]) | *lname | *lvn ]")
	}

	_global := C.CString(global)
	errmsg := make([]byte, maxmsglen)
	defer C.free(unsafe.Pointer(_global))
	mu.Lock()
	result := C.cip_kill(_global, (*C.char)(unsafe.Pointer(&errmsg[0])), C.int(len(errmsg)))
	mu.Unlock()
	if result != 0 {
		return errors.New("Kill failed: " + string(result) + "Error message: " + string(errmsg))
	}
	return nil
} // end of Kill

//ZKill deletes the data value for a variable name without affecting the nodes descended from that node.
func ZKill(global string) error {

	if len(global) < 1 {
		return errors.New("ZKill failed - you must provide glvn")
	}

	_global := C.CString(global)
	errmsg := make([]byte, maxmsglen)
	defer C.free(unsafe.Pointer(_global))
	mu.Lock()
	result := C.cip_zkill(_global, (*C.char)(unsafe.Pointer(&errmsg[0])), C.int(len(errmsg)))
	mu.Unlock()
	if result != 0 {
		return errors.New("ZKill failed: " + string(result) + "Error message: " + string(errmsg))
	}
	return nil
} // end of ZKill

//Xecute runs the M code
func Xecute(code string) error {

	if len(code) < 1 {
		return errors.New("Xecute failed - you must provide some code")
	}

	_code := C.CString(code)
	errmsg := make([]byte, maxmsglen)
	defer C.free(unsafe.Pointer(_code))
	mu.Lock()
	result := C.cip_xecute(_code, (*C.char)(unsafe.Pointer(&errmsg[0])), C.int(len(errmsg)))
	mu.Unlock()
	if result != 0 {
		return errors.New("Xecute failed: " + string(result) + "Error message: " + string(errmsg))
	}
	return nil
} // end of Xecute

//Order returns the next key or glvn
//You can also provide flag '-1' in dir variable. This way order will return previous key or glvn.
func Order(global string, dir string) (string, error) {

	if len(global) < 1 {
		return "", errors.New("Order failed - you must provide glvn")
	}

	if dir != "-1" {
		dir = "1"
	}

	_global := C.CString(global)
	_dir := C.CString(dir)
	_ret := make([]byte, maxretlen)
	errmsg := make([]byte, maxmsglen)
	defer C.free(unsafe.Pointer(_global))
	defer C.free(unsafe.Pointer(_dir))
	mu.Lock()
	result := C.cip_order(_global, (*C.char)(unsafe.Pointer(&_ret[0])), (*C.char)(unsafe.Pointer(&errmsg[0])), C.int(len(errmsg)), maxretlen, _dir)
	mu.Unlock()
	if result != 0 {
		return "", errors.New("Order failed: " + string(result) + "Error message: " + string(errmsg))
	}
	_ret = bytes.Trim(_ret, "\x00") // trim unused space
	return string(_ret), nil
} //end of Order

//Query returns the next subscripted local or global variable node name, independent of level, which follows
//the node specified by its argument in M collating sequence and has a data value.
func Query(global string) (string, error) {

	if len(global) < 1 {
		return "", errors.New("Query failed - you must provide glvn")
	}

	_global := C.CString(global)
	_ret := make([]byte, maxretlen)
	errmsg := make([]byte, maxmsglen)
	defer C.free(unsafe.Pointer(_global))
	mu.Lock()
	result := C.cip_query(_global, (*C.char)(unsafe.Pointer(&_ret[0])), (*C.char)(unsafe.Pointer(&errmsg[0])), C.int(len(errmsg)), maxretlen)
	mu.Unlock()
	if result != 0 {
		return "", errors.New("Query failed: " + string(result) + "Error message: " + string(errmsg))
	}
	_ret = bytes.Trim(_ret, "\x00") // trim unused space
	return string(_ret), nil
} //end of Query

//Data returns an 2 boolean values describing the value and descendent status of a local or global variable.
//		- If the variable is undefined Data returns false, false, nil.
//		- If the variable has a value but no descendants returns true, false, nil
//		- If the variable has descendants but no value returns false, true, nil
//		- If the variable has a value and descendants returns true, true, nil
//		in case of any error returns false, false, error
func Data(global string) (value, descendent bool, err error) {

	value = false
	descendent = false

	if len(global) < 1 {
		return value, descendent, errors.New("Data failed - you must provide glvn")
	}

	_global := C.CString(global)
	_ret := C.uint(1000)

	errmsg := make([]byte, maxmsglen)
	defer C.free(unsafe.Pointer(_global))
	mu.Lock()
	status := C.cip_data(_global, &_ret, (*C.char)(unsafe.Pointer(&errmsg[0])), C.int(len(errmsg)), maxretlen)
	mu.Unlock()
	if status != 0 {
		return value, descendent, errors.New("Data failed: " + string(status) + "Error message: " + string(errmsg))
	}

	switch _ret {
	case 1:
		value = true
	case 10:
		descendent = true
	case 11:
		value = true
		descendent = true
	case 1000:
		err = errors.New("Data failed: no code was returned by gt.m")
	}

	return
} //end of Data
