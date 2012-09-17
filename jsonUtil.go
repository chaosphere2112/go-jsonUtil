package jsonUtil

import "encoding/json"
import "errors"

//JSON Object management– get arrays, objects, strings, numbers, booleans from a json object

type JsonObject map[string]interface{}

//Creates a JsonObject from a string, or returns the error that json.Unmarshal gives
func JsonFromString(value string) (obj JsonObject, errormsg error) {
	
	var blob interface{}

	err := json.Unmarshal([]byte(value), &blob)

	if err == nil {

		mapped := blob.(map[string]interface{})

		return mapped, nil
	}

	return nil, err
}

//JsonObject method: Returns the object stored at @param key, or an error
func (object JsonObject) Object(key string) (obj JsonObject, errormsg error){

	newObj, ok := object[key].(map[string]interface{})

	if ok {
		return newObj, nil
	}

	return nil, errors.New("Object["+key+"] is not an object.")

}

//JsonObject method: Returns the string stored at @param key, or an error
func (object JsonObject) String(key string) (obj string, errormsg error){

	str, ok := object[key].(string)

	if ok {
		return str, nil
	}

	return "", errors.New("Object["+key+"] is not a string.")

}

//JsonObject method: Returns the number stored at @param key, or an error
func (object JsonObject) Number(key string) (obj float64, errormsg error) {

	num, ok := object[key].(float64)

	if ok {
		return num, nil
	}

	return float64(0.0), errors.New("Object["+key+"] is not a number.")

}

//JsonObject method: Returns the bool stored at @param key, or an error
func (object JsonObject) Bool(key string) (obj bool, errormsg error) {

	boo, ok := object[key].(bool)

	if ok {
		return boo, nil
	}

	return false, errors.New("Object["+key+"] is not a boolean.")

}

//Type to allow methods to be added for easy access to typed subobjects
type JsonArray []interface{}

//JsonObject method: Returns the array stored at @param key, or an error
func (object JsonObject) Array(key string) (obj JsonArray, errormsg error) {

	array, ok := object[key].([]interface{})

	if ok {
		return array, nil
	}

	return nil, errors.New("Object["+key+"] is not an array.")

}

//JsonArray method: Returns the object stored at @param index, or an error
func (array JsonArray) Object(index int) (obj JsonObject, errormsg error) {

	object, ok := array[index].(map[string]interface{})

	if ok {
		return object, nil
	}

	return nil, errors.New("Object["+string(index)+"] is not an object.")
}

//JsonArray method: Returns the string stored at @param index, or an error
func (array JsonArray) String(index int) (obj string, errormsg error) {

	object, ok := array[index].(string)

	if ok {
		return object, nil
	}

	return "", errors.New("Object["+string(index)+"] is not a string.")
}

//JsonArray method: Returns the number stored at @param index, or an error
func (array JsonArray) Number(index int) (obj float64, errormsg error) {

	object, ok := array[index].(float64)

	if ok {
		return object, nil
	}

	return float64(0), errors.New("Object["+string(index)+"] is not a number.")
}

//JsonArray method: Returns the bool stored at @param index, or an error
func (array JsonArray) Bool(index int) (obj bool, errormsg error) {

	object, ok := array[index].(bool)

	if ok {
		return object, nil
	}

	return false, errors.New("Object["+string(index)+"] is not a bool.")
}

//JsonArray method: Returns the array stored at @param index, or an error
func (array JsonArray) Array(index int) (obj JsonArray, errormsg error) {

	object, ok := array[index].([]interface{})

	if ok {
		return object, nil
	}

	return nil, errors.New("Object["+string(index)+"] is not an array.")
}