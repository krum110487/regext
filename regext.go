package regext

import (
	"bytes"
	"fmt"
	"regexp"
)

//Regext is the main struct
type Regext struct {
	Dataset [][]byte
	source  []byte
	stack   [][]byte
}

//NewRegext gets the instance of Regext.
func NewRegext(data []byte) *Regext {
	var newSet [][]byte
	newSet = append(newSet, data)
	return &Regext{Dataset: newSet, source: data, stack: newSet}
}

/*FindLast TODO: Doesn't work just yet....*/
func (r *Regext) FindLast(regex string) *Regext {
	return r.Find(regex, -1, 1)
}

//FindFirst is a wrapper of Find to shorthand the skip and count
func (r *Regext) FindFirst(regex string) *Regext {
	return r.Find(regex, 0, 1)
}

//FindAll is a wrapper of Find to shorthand the skip and count
func (r *Regext) FindAll(regex string) *Regext {
	return r.Find(regex, 0, -1)
}

//Find or one of it's equivelant functions must be called to load the dataset
func (r *Regext) Find(regex string, skip int, count int) *Regext {
	re := regexp.MustCompile(regex)
	findAll := -1
	if skip <= 0 && count > 0 {
		findAll = count
	}

	var newDataset [][]byte
	for _, ds := range r.Dataset {
		res := re.FindAll(ds, findAll)

		//Check for bounds
		if count >= len(res) || count <= 0 {
			count = len(res)
		}
		if skip <= 0 {
			skip = 0
		}

		//Build new result set in order.
		for _, newAppend := range res[skip:count] {
			newDataset = append(newDataset, newAppend)
		}
		r.Dataset = newDataset
	}
	return r
}

//ReplaceAll will replace any match with the value supplied.
func (r *Regext) ReplaceAll(regex string, value string) *Regext {
	re := regexp.MustCompile(regex)
	var newDataset [][]byte
	for _, ds := range r.Dataset {
		res := re.ReplaceAll(ds, []byte(value))
		newDataset = append(newDataset, res)
	}
	r.Dataset = newDataset
	return r
}

//FilterOut will remove anything in the dataset that matches the regex
func (r *Regext) FilterOut(regexs ...string) *Regext {
	var newDataset [][]byte
	for _, ds := range r.Dataset {
		pass := false
		for _, regex := range regexs {
			re := regexp.MustCompile(regex)
			if !re.Match(ds) {
				pass = true
				break
			}
		}
		if pass {
			newDataset = append(newDataset, ds)
		}
	}
	r.Dataset = newDataset
	return r
}

//FilterOutAny - filters out anything that matches in the list
func (r *Regext) FilterOutAny(regexs ...string) *Regext {
	var newDataset [][]byte
	for _, ds := range r.Dataset {
		found := true
		for _, regex := range regexs {
			re := regexp.MustCompile(regex)
			if re.Match(ds) {
				found = false
				break
			}
		}
		if found {
			newDataset = append(newDataset, ds)
		}
	}
	r.Dataset = newDataset
	return r
}

//Filter will keep anything in the dataset that matches the regex
func (r *Regext) Filter(regexs ...string) *Regext {
	var newDataset [][]byte
	for _, ds := range r.Dataset {
		pass := true
		for _, regex := range regexs {
			re := regexp.MustCompile(regex)
			if !re.Match(ds) {
				pass = false
				break
			}
		}
		if pass {
			newDataset = append(newDataset, ds)
		}
	}
	r.Dataset = newDataset
	return r
}

//FilterAny will take all of the items and filter all of them
func (r *Regext) FilterAny(regexs ...string) *Regext {
	var newDataset [][]byte
	for _, ds := range r.Dataset {
		for _, regex := range regexs {
			re := regexp.MustCompile(regex)
			if re.Match(ds) {
				newDataset = append(newDataset, ds)
				break
			}
		}
	}
	r.Dataset = newDataset
	return r
}

//FilterByLen - will filter by the length of the byte
func (r *Regext) FilterByLen(min int, max int) *Regext {
	var newDataset [][]byte
	for _, ds := range r.Dataset {
		dslen := len(ds)
		if (dslen >= min || dslen < 0) && (dslen <= max || dslen < 0) {
			newDataset = append(newDataset, ds)
		}
	}
	r.Dataset = newDataset
	return r
}

//FilterOutByLen - will filter by the length of the byte
func (r *Regext) FilterOutByLen(min int, max int) *Regext {
	var newDataset [][]byte
	for _, ds := range r.Dataset {
		dslen := len(ds)
		if max < 0 {
			max = min
		}
		if (dslen < min) || (dslen > max) {
			newDataset = append(newDataset, ds)
		}
	}
	r.Dataset = newDataset
	return r
}

//Trim - trims whitespace
func (r *Regext) Trim() *Regext {
	return r.DeleteAny(`^[\s\r\n]+`, `[\s\r\n]+$`)
}

//DeleteAny - filters out anything that matches in the list
func (r *Regext) DeleteAny(regexs ...string) *Regext {
	var newDataset [][]byte
	for _, ds := range r.Dataset {
		for _, regex := range regexs {
			re := regexp.MustCompile(regex)
			ds = re.ReplaceAll(ds, []byte(""))
		}
		newDataset = append(newDataset, ds)
	}
	r.Dataset = newDataset
	return r
}

//Split - splits the results based upon a regex found.
func (r *Regext) Split(regexs ...string) *Regext {
	for _, regex := range regexs {
		var newDataset [][]byte
		for _, ds := range r.Dataset {
			re := regexp.MustCompile(regex)
			split := re.Split(string(ds[:len(ds)]), -1)
			for _, sp := range split {
				newDataset = append(newDataset, []byte(sp))
			}
		}
		r.Dataset = newDataset
	}
	return r
}

//SplitIgnoreBlanks - same as split, but ignores blanks.
func (r *Regext) SplitIgnoreBlanks(regexs ...string) *Regext {
	for _, regex := range regexs {
		var newDataset [][]byte
		for _, ds := range r.Dataset {
			re := regexp.MustCompile(regex)
			split := re.Split(string(ds[:len(ds)]), -1)
			for _, sp := range split {
				if sp != "" {
					newDataset = append(newDataset, []byte(sp))
				}
			}
		}
		r.Dataset = newDataset
	}
	return r
}

//String - Returns an array of strings
func (r *Regext) String() string {
	bytesJoined := bytes.Join(r.Dataset, []byte(""))
	return string(bytesJoined[:len(bytesJoined)])
}

//JoinStr - Joins the entries together using a string separator
func (r *Regext) JoinStr(sep string) *Regext {
	return r.Join([]byte(sep))
}

//Join - Joins the entries together
func (r *Regext) Join(sep []byte) *Regext {
	var newDataset [][]byte
	r.Dataset = append(newDataset, bytes.Join(r.Dataset, sep))
	return r
}

//ToLower - Sets all of the entries to lowercase
func (r *Regext) ToLower() *Regext {
	for i, ds := range r.Dataset {
		r.Dataset[i] = bytes.ToLower(ds)
	}
	return r
}

//ToUpper - Sets all of the entries to uppercase
func (r *Regext) ToUpper() *Regext {
	for i, ds := range r.Dataset {
		r.Dataset[i] = bytes.ToUpper(ds)
	}
	return r
}

//PrintRaw is a quick and dirty print to the console mostly for debug
func (r *Regext) PrintRaw() *Regext {
	r.Println("", "")
	return r
}

//Print prints the current dataset to the console without new line
func (r *Regext) Print(prefix string, postfix string) *Regext {
	fmt.Printf("%s%q%s", prefix, r.Dataset, postfix)
	return r
}

//Println prints the current dataset to the console
func (r *Regext) Println(prefix string, postfix string) *Regext {
	fmt.Printf("%s%q%s\n", prefix, r.Dataset, postfix)
	return r
}

//GetData returns the data, breaking the chain.
func (r *Regext) GetData() [][]byte {
	return r.Dataset
}
