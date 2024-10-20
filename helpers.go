package todo

import (
	"fmt"
	"reflect"
)

// MethodInfo contains information about a method
type MethodInfo struct {
	Name        string
	Method      reflect.Method
	InputTypes  []reflect.Type
	OutputTypes []reflect.Type
}

func getInterfaceMethods(i interface{}) []string {
	// Get the type of the interface
	interfaceType := reflect.TypeOf(i).Elem()

	// Check if it's actually an interface
	if interfaceType.Kind() != reflect.Interface {
		return nil
	}

	// Slice to store method names
	methods := make([]string, interfaceType.NumMethod())

	// Iterate through all methods
	for i := 0; i < interfaceType.NumMethod(); i++ {
		method := interfaceType.Method(i)
		methods[i] = method.Name
	}

	return methods
}

func mapMethodToFunction(i interface{}) map[string]MethodInfo {
	// Get the type of the interface
	interfaceType := reflect.TypeOf(i).Elem()

	// Check if it's actually an interface
	if interfaceType.Kind() != reflect.Interface {
		return nil
	}

	// Create map to store method information
	methodMap := make(map[string]MethodInfo)

	// Get methods from GetAllMethods
	methods := getInterfaceMethods(i)

	// Iterate through methods and create detailed information
	for _, methodName := range methods {
		method, _ := interfaceType.MethodByName(methodName)
		methodType := method.Type

		// Get input parameter types (excluding receiver)
		inputTypes := make([]reflect.Type, methodType.NumIn())
		for j := 0; j < methodType.NumIn(); j++ {
			inputTypes[j] = methodType.In(j)
		}

		// Get output parameter types
		outputTypes := make([]reflect.Type, methodType.NumOut())
		for j := 0; j < methodType.NumOut(); j++ {
			outputTypes[j] = methodType.Out(j)
		}

		methodMap[methodName] = MethodInfo{
			Name:        methodName,
			Method:      method,
			InputTypes:  inputTypes,
			OutputTypes: outputTypes,
		}
	}

	return methodMap
}

func callMethod(obj interface{}, methodName string, args ...interface{}) ([]interface{}, error) {
	objValue := reflect.ValueOf(obj)
	method := objValue.MethodByName(methodName)

	if !method.IsValid() {
		return nil, fmt.Errorf("method %s not found", methodName)
	}

	// Convert args to reflect.Value
	reflectArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		reflectArgs[i] = reflect.ValueOf(arg)
	}

	// Call the method
	results := method.Call(reflectArgs)

	// Convert results back to interface{}
	returnValues := make([]interface{}, len(results))
	for i, result := range results {
		returnValues[i] = result.Interface()
	}

	return returnValues, nil
}
