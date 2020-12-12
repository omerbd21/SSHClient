package windows

import (
	ole2 "github.com/go-ole/go-ole"
	"github.com/mattn/go-ole"
	"github.com/mattn/go-ole/oleutil"
)

func ViewPrograms() []string {
	// init COM, oh yeah
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	unknown, _ := oleutil.CreateObject("WbemScripting.SWbemLocator")
	defer unknown.Release()

	wmi, _ := unknown.QueryInterface((*ole2.GUID)(ole.IID_IDispatch))
	defer wmi.Release()

	// service is a SWbemServices
	serviceRaw, _ := oleutil.CallMethod(wmi, "ConnectServer")
	service := serviceRaw.ToIDispatch()
	defer service.Release()
	// result is a SWBemObjectSet
	resultRaw, _ := oleutil.CallMethod(service, "ExecQuery", "SELECT * FROM Win32_Product")
	result := resultRaw.ToIDispatch()
	defer result.Release()
	countVar, _ := oleutil.GetProperty(result, "Count")
	count := int(countVar.Val)
	var programs []string
	for i := 0; i < count; i++ {
		// item is a SWbemObject, but really a Win32_Process
		itemRaw, _ := oleutil.CallMethod(result, "ItemIndex", i)
		item := itemRaw.ToIDispatch()
		defer item.Release()
		asString, _ := oleutil.GetProperty(item, "Name")
		programs = append(programs,asString.ToString())
	}
	return programs
}
