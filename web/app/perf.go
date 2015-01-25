package app

import (
	"bufio"
	"github.com/gophergala/GopherKombat/common/user"
	"net/http"
	"os"
	"os/exec"
)

func PerfHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := GetCurrentUser(r)
	data := make(map[string]interface{})
	data["loggedIn"] = ok
	if ok {
		data["user"] = user
	}
	render(w, "perf", data)

}

func PerfSubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		user, ok := GetCurrentUser(r)
		data := make(map[string]interface{})
		if ok {
			r.ParseForm()
			code1 := r.PostFormValue("code1")
			code2 := r.PostFormValue("code2")
			validated1, message1 := validate(user, code1, "code1")
			validated2, message2 := validate(user, code2, "code2")
			if validated1 && validated2 {
				// TODO entry point
			} else {
				remove(user)
			}

			data["success1"] = validated1
			data["success2"] = validated2
			data["message1"] = message1
			data["message2"] = message2
		} else {
			data["success"] = ok
			data["message"] = "You are not logged in."
		}
		renderJson(w, r, data)
	} else {
		http.Error(w, "post only", http.StatusMethodNotAllowed)
	}
}

func validate(user *user.User, code, name string) (bool, string) {
	err := save(user, code, name)
	if err != nil {
		return false, "Failed to save file."
	}
	file := os.Getenv("GOPATH") + "/src/blueprints/" + user.Name + "/" + name + ".go"
	out, err := exec.Command("go", "build", file).CombinedOutput()
	if err != nil {
		//panic(err)
	}
	message := string(out)
	if message == "" {
		message = "OK"
	}
	return true, message
}

func save(user *user.User, code, name string) error {
	dir := os.Getenv("GOPATH") + "/src/blueprints/" + user.Name

	err := os.Chdir(dir)
	if os.IsNotExist(err) {
		err = os.Mkdir(dir, 0777)
	}
	if err != nil {
		panic(err)
		return err
	}
	f, err := os.Create(name + ".go")
	if os.IsExist(err) {
		err = os.Remove(name + ".go")
		f, err = os.Create(name + ".go")
	}
	if err != nil {
		panic(err)
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.WriteString(code)

	if err != nil {
		panic(err)
		return err
	}
	w.Flush()
	return nil
}

func remove(user *user.User) {
	dir := os.Getenv("GOPATH") + "/src/blueprints/" + user.Name

	err := os.Chdir(dir)
	if os.IsNotExist(err) {
		err = os.Mkdir(dir, 0777)
	}
	if err != nil {
		panic(err)
	}
	err = os.Remove("code1.go")
	err = os.Remove("code2.go")
	if err != nil {
		panic(err)
	}
}
