package UIComponents

import (
	"fmt"
	"myproject/Config"
	"testing"
)

func TestVillagesDDL(t *testing.T) {

	Config.RunConfig()
	s := VillagesDDL()
	fmt.Println(s)
}

func TestGetVillagesDB(t *testing.T) {
	Config.RunConfig()
	m := GetVillagesDB()
	fmt.Println(m)
}
