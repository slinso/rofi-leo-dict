package rld

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strings"
)

func ScrapeLeo(from, to, word string) ([]Section, error) {
	var v LeoXML
	ret := make([]Section, 0)

	url := "http://dict.leo.org/dictQuery/m-vocab/ende/query.xml?" +
		"&tolerMode=nof" +
		"&lp=" + from + to +
		"&lang=" + from +
		"&rmWords=off" +
		"&rmSearch=on" +
		"&search=" + word +
		"&resultOrder=basic" +
		"&multiwordShowSingle=on" +
		"&sectLenMax=16" +
		"&n=1"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(body, &v)
	if err != nil {
		return nil, err
	}

	/**********************************************************************
	 * translations
	**********************************************************************/
	for i := range v.Sectionlist.Section {
		s := Section{}
		s.Title = v.Sectionlist.Section[i].SctTitle

		for j := range v.Sectionlist.Section[i].Entry {
			e := v.Sectionlist.Section[i].Entry[j]
			entry := make(map[string]string)

			for k := range e.Side {
				entry[e.Side[k].Lang] = strings.Join(e.Side[k].Words.Word, ", ")
			}
			s.Entries = append(s.Entries, entry)
		}

		ret = append(ret, s)
	}

	/**********************************************************************
	 * simliar
	**********************************************************************/
	s := Section{}
	s.Title = "similar"
	entry := make(map[string]string)

	for k := range v.Similar.Side {
		entry[v.Similar.Side[k].Lang] = strings.Join(v.Similar.Side[k].Word, ", ")
	}

	s.Entries = append(s.Entries, entry)
	ret = append(ret, s)

	/**********************************************************************
	 * synonym
	**********************************************************************/
	s = Section{}
	s.Title = "synonyms"
	entry = make(map[string]string)

	for k := range v.Ffsynlist.Side {
		entry[v.Ffsynlist.Side[k].Lang] = strings.Join(v.Ffsynlist.Side[k].Word, ", ")
	}

	s.Entries = append(s.Entries, entry)
	ret = append(ret, s)

	return ret, nil
}
