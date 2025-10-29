package UIComponents

import (
	"fmt"
	"myproject/Config"
	"myproject/utils"
	"myproject/utilsDB"
	"strings"
)

func VillagesDDL() string {
	//get the local base directory and then load the file
	baseDir := Config.BaseDirPath
	villagesDir := baseDir + "UI\\UIComponents\\villages.html"
	control := utils.LoadFile(villagesDir)
	optionstring := OptionString()
	control = strings.ReplaceAll(control, "$villages", optionstring)
	fmt.Println(control)

	return control

}
func OptionString() string {
	s1 := `<option value="$villageID">$villageName</option>`
	villages := []string{}
	m := GetVillagesDB()
	for k, v := range m {
		s2 := s1
		s2 = strings.ReplaceAll(s2, "$villageID", utils.Int2string(k))
		s2 = strings.ReplaceAll(s2, "$villageName", v)
		villages = append(villages, s2)
	}
	villagesString := strings.Join(villages, "")
	return villagesString
}
func GetVillagesDB() map[int]string {
	m := make(map[int]string)
	fmt.Println(m)
	rows := utilsDB.GetRows("select villageId, villageName from Villages")
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		m[id] = name
	}
	return m

}
