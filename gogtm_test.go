package gogtm

import "testing"

func TestCheckParams(t *testing.T) {

	env := &gtmenv{gbldir: ""}

	if err := env.checkParams(); err == nil {
		t.Error("empty glbdir should be returned as an error")
	} else {
		t.Logf("returned: %s", err)
	}

	env = &gtmenv{
		gbldir: "gogtm_test.go",
		chset:  ""}
	if err := env.checkParams(); err == nil {
		t.Error("empty chset should be returned as an error")
	} else {
		t.Logf("returned: %s", err)
	}
	env = &gtmenv{
		gbldir: "gogtm_test.go",
		chset:  "UTF-8"}
	if err := env.checkParams(); err == nil {
		t.Error("utf with no icu set shoud return an error")
	} else {
		t.Logf("returned: %s", err)
	}

	env = &gtmenv{
		gbldir: "gogtm_test.go",
		chset:  "M",
		log:    "gogtm_test.go"}
	if err := env.checkParams(); err == nil {
		t.Error("file set as a log destination should return an error")
	} else {
		t.Logf("returned: %s", err)
	}

	env = &gtmenv{
		gbldir: "gogtm_test.go",
		chset:  "M",
		log:    "/tmp",
		tmp:    ""}
	if err := env.checkParams(); err == nil {
		t.Error("unknown dist destination should return an error")
	} else {
		t.Logf("returned: %s", err)
	}

	env = &gtmenv{
		gbldir: "gogtm_test.go",
		chset:  "M",
		log:    "/tmp",
		tmp:    "/tmp",
		dist:   ""}
	if err := env.checkParams(); err == nil {
		t.Error("unknown dist destination should return an error")
	} else {
		t.Logf("returned: %s", err)
	}
	env = &gtmenv{
		gbldir: "gogtm_test.go",
		chset:  "M",
		log:    "/tmp",
		tmp:    "/tmp",
		dist:   "/tmp"}
	if err := env.checkParams(); err != nil {
		t.Error("this one should be ok")
	} else {
		t.Logf("returned: %s", env)
	}
	env = &gtmenv{
		gbldir: "gogtm_test.go",
		chset:  "UTF-8",
		icu:    "5.0",
		log:    "/tmp",
		tmp:    "/tmp",
		dist:   "/tmp"}
	if err := env.checkParams(); err != nil {
		t.Error("this one should be ok")
	} else {
		t.Logf("returned: %s", env)
	}
}
