The golang encoding/json library is really quite ugly. It requires constant casts to deal with any structure that you don't have a guaranteed form for (and thusly a struct to match), and needs a bunch of boilerplate code.  This library aims to extend and address the existing structures, and clean up all of that boilerplate.

Usage:
    
    func TestJson(inputString string) map[string]string {
        jsonObject, err := jsonUtil.JsonObjectFromString(inputString)

        if err != nil {
            fmt.Println(err.Error())
            return nil
        }

        var jsonString string

        jsonString, err = jsonObject.String("keyForString")

        if err != nil {
            fmt.Println(err.Error())
            return nil
        }

        return {"keyForString":jsonString}
    }

Obviously, there's still the boilerplate of error handling, but it's a work in progress.