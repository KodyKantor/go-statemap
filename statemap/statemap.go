package statemap

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type State struct {
	Color interface{} `json:"color"`
	Value int         `json:"value"`
}

type Datum struct {
	Time   string      `json:"time"`
	Entity string      `json:"entity"`
	State  int         `json:"state"`
	Tag    interface{} `json:"tag"`
}

type description struct {
	Entity      string `json:"entity"`
	Description string `json:"description"`
}

type metadata struct {
	Start      [2]int64 `json:"start"`
	Title      string   `json:"title"`
	Host       string   `json:"host"`
	EntityKind string   `json:"entityKind,omitempty"`
	// name -> state mapping
	States map[string]State `json:"states"`
}

type Statemap struct {
	metadata   metadata
	stateData  map[string][]Datum
	firstState time.Time
}

func New(title string, host string, entityKind string) Statemap {
	mdata := metadata{
		Title:      title,
		Host:       host,
		EntityKind: entityKind,
		States:     map[string]State{},
	}

	return Statemap{
		metadata:  mdata,
		stateData: map[string][]Datum{},
	}
}

func (s *Statemap) SetState(entityName, stateName, tag, color string, time time.Time) {

	var col interface{} = nil
	if color != "" {
		col = color
	}

	/* Find the existing state info, or create it. */
	if _, ok := s.metadata.States[stateName]; !ok {
		s.metadata.States[stateName] = State{
			Color: col,
			Value: len(s.metadata.States),
		}
	}
	st := s.metadata.States[stateName]

	if s.firstState.IsZero() || s.firstState.After(time) {
		s.firstState = time
		s.metadata.Start[0] = time.Unix()
		s.metadata.Start[1] = time.UnixNano() - (time.Unix() * 1000000000)
	}

	var mtag interface{}
	mtag = tag
	if tag == "" {
		mtag = nil
	}

	d := Datum{
		Time:   strconv.FormatInt(time.UnixNano(), 10),
		Entity: entityName,
		State:  st.Value,
		Tag:    mtag,
	}

	/* Reminder: append() can operate on a nil slice. */
	s.stateData[entityName] = append(s.stateData[entityName], d)
}

func (s *Statemap) Dump() string {

	ret := ""
	if res, err := json.Marshal(s.metadata); err != nil {
		fmt.Println("error marshal json: ", err)
	} else {
		ret = fmt.Sprintf("%+v\n", string(res))
	}

	for _, v := range s.stateData {
		for _, d := range v {
			if t, err := strconv.ParseInt(d.Time, 10, 64); err == nil {
				d.Time = strconv.FormatInt(t-s.firstState.UnixNano(), 10)
			} else {
				fmt.Println("err converting time to int64:", err)
			}
			if res, err := json.Marshal(d); err != nil {
				fmt.Println("err marshal datum json")
			} else {
				ret = fmt.Sprintf("%s%s\n", ret, string(res))
			}
		}
	}

	return (ret)
}
